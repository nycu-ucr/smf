package context

import (
	"fmt"
	"net"
	"os"
	"sync/atomic"

	"github.com/google/uuid"

	"github.com/free5gc/pfcp/pfcpType"
	"github.com/free5gc/pfcp/pfcpUdp"
	"github.com/nycu-ucr/openapi/Nnrf_NFDiscovery"
	"github.com/nycu-ucr/openapi/Nnrf_NFManagement"
	"github.com/nycu-ucr/openapi/Nudm_SubscriberDataManagement"
	"github.com/nycu-ucr/openapi/models"
	"github.com/nycu-ucr/smf/factory"
	"github.com/nycu-ucr/smf/logger"
)

func init() {
	smfContext.NfInstanceID = uuid.New().String()
}

var smfContext SMFContext

type SMFContext struct {
	Name         string
	NfInstanceID string

	URIScheme    models.UriScheme
	BindingIPv4  string
	RegisterIPv4 string
	SBIPort      int
	CPNodeID     pfcpType.NodeID

	UDMProfile models.NfProfile

	UPNodeIDs []pfcpType.NodeID
	Key       string
	PEM       string
	KeyLog    string

	SnssaiInfos []SnssaiSmfInfo

	NrfUri                         string
	NFManagementClient             *Nnrf_NFManagement.APIClient
	NFDiscoveryClient              *Nnrf_NFDiscovery.APIClient
	SubscriberDataManagementClient *Nudm_SubscriberDataManagement.APIClient

	UserPlaneInformation *UserPlaneInformation

	// Now only "IPv4" supported
	// TODO: support "IPv6", "IPv4v6", "Ethernet"
	SupportedPDUSessionType string

	//*** For ULCL ** //
	ULCLSupport         bool
	ULCLGroups          map[string][]string
	UEPreConfigPathPool map[string]*UEPreConfigPaths
	UEDefaultPathPool   map[string]*UEDefaultPaths
	LocalSEIDCount      uint64
}

// RetrieveDnnInformation gets the corresponding dnn info from S-NSSAI and DNN
func RetrieveDnnInformation(Snssai models.Snssai, dnn string) *SnssaiSmfDnnInfo {
	for _, snssaiInfo := range SMF_Self().SnssaiInfos {
		if snssaiInfo.Snssai.Sst == Snssai.Sst && snssaiInfo.Snssai.Sd == Snssai.Sd {
			return snssaiInfo.DnnInfos[dnn]
		}
	}
	return nil
}

func AllocateLocalSEID() uint64 {
	atomic.AddUint64(&smfContext.LocalSEIDCount, 1)
	return smfContext.LocalSEIDCount
}

func InitSmfContext(config *factory.Config) {
	if config == nil {
		logger.CtxLog.Error("Config is nil")
		return
	}

	logger.CtxLog.Infof("smfconfig Info: Version[%s] Description[%s]", config.Info.Version, config.Info.Description)
	configuration := config.Configuration
	if configuration.SmfName != "" {
		smfContext.Name = configuration.SmfName
	}

	sbi := configuration.Sbi
	if sbi == nil {
		logger.CtxLog.Errorln("Configuration needs \"sbi\" value")
		return
	} else {
		smfContext.URIScheme = models.UriScheme(sbi.Scheme)
		smfContext.RegisterIPv4 = factory.SMF_DEFAULT_IPV4 // default localhost
		smfContext.SBIPort = factory.SMF_DEFAULT_PORT_INT  // default port
		if sbi.RegisterIPv4 != "" {
			smfContext.RegisterIPv4 = sbi.RegisterIPv4
		}
		if sbi.Port != 0 {
			smfContext.SBIPort = sbi.Port
		}

		if tls := sbi.TLS; tls != nil {
			smfContext.Key = tls.Key
			smfContext.PEM = tls.PEM
		}

		smfContext.BindingIPv4 = os.Getenv(sbi.BindingIPv4)
		if smfContext.BindingIPv4 != "" {
			logger.CtxLog.Info("Parsing ServerIPv4 address from ENV Variable.")
		} else {
			smfContext.BindingIPv4 = sbi.BindingIPv4
			if smfContext.BindingIPv4 == "" {
				logger.CtxLog.Warn("Error parsing ServerIPv4 address as string. Using the 0.0.0.0 address as default.")
				smfContext.BindingIPv4 = "0.0.0.0"
			}
		}
	}

	if configuration.NrfUri != "" {
		smfContext.NrfUri = configuration.NrfUri
	} else {
		logger.CtxLog.Warn("NRF Uri is empty! Using localhost as NRF IPv4 address.")
		smfContext.NrfUri = fmt.Sprintf("%s://%s:%d", smfContext.URIScheme, "127.0.0.1", 29510)
	}

	if pfcp := configuration.PFCP; pfcp != nil {
		if pfcp.Port == 0 {
			pfcp.Port = pfcpUdp.PFCP_PORT
		}
		pfcpAddrEnv := os.Getenv(pfcp.Addr)
		if pfcpAddrEnv != "" {
			logger.CtxLog.Info("Parsing PFCP IPv4 address from ENV variable found.")
			pfcp.Addr = pfcpAddrEnv
		}
		if pfcp.Addr == "" {
			logger.CtxLog.Warn("Error parsing PFCP IPv4 address as string. Using the 0.0.0.0 address as default.")
			pfcp.Addr = "0.0.0.0"
		}
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", pfcp.Addr, pfcp.Port))
		if err != nil {
			logger.CtxLog.Warnf("PFCP Parse Addr Fail: %v", err)
		}

		smfContext.CPNodeID.NodeIdType = 0
		smfContext.CPNodeID.NodeIdValue = addr.IP.To4()
	}

	smfContext.SnssaiInfos = make([]SnssaiSmfInfo, 0, len(configuration.SNssaiInfo))

	for _, snssaiInfoConfig := range configuration.SNssaiInfo {
		snssaiInfo := SnssaiSmfInfo{}
		snssaiInfo.Snssai = SNssai{
			Sst: snssaiInfoConfig.SNssai.Sst,
			Sd:  snssaiInfoConfig.SNssai.Sd,
		}

		snssaiInfo.DnnInfos = make(map[string]*SnssaiSmfDnnInfo)

		for _, dnnInfoConfig := range snssaiInfoConfig.DnnInfos {
			dnnInfo := SnssaiSmfDnnInfo{}
			if dnnInfoConfig.DNS != nil {
				dnnInfo.DNS.IPv4Addr = net.ParseIP(dnnInfoConfig.DNS.IPv4Addr).To4()
				dnnInfo.DNS.IPv6Addr = net.ParseIP(dnnInfoConfig.DNS.IPv6Addr).To16()
			}
			if dnnInfoConfig.PCSCF != nil {
				dnnInfo.PCSCF.IPv4Addr = net.ParseIP(dnnInfoConfig.PCSCF.IPv4Addr).To4()
			}
			snssaiInfo.DnnInfos[dnnInfoConfig.Dnn] = &dnnInfo
		}
		smfContext.SnssaiInfos = append(smfContext.SnssaiInfos, snssaiInfo)
	}

	// Set client and set url
	ManagementConfig := Nnrf_NFManagement.NewConfiguration()
	ManagementConfig.SetBasePath(SMF_Self().NrfUri)
	smfContext.NFManagementClient = Nnrf_NFManagement.NewAPIClient(ManagementConfig)

	NFDiscovryConfig := Nnrf_NFDiscovery.NewConfiguration()
	NFDiscovryConfig.SetBasePath(SMF_Self().NrfUri)
	smfContext.NFDiscoveryClient = Nnrf_NFDiscovery.NewAPIClient(NFDiscovryConfig)

	smfContext.ULCLSupport = configuration.ULCL

	smfContext.SupportedPDUSessionType = "IPv4"

	smfContext.UserPlaneInformation = NewUserPlaneInformation(&configuration.UserPlaneInformation)

	SetupNFProfile(config)
}

func InitSMFUERouting(routingConfig *factory.RoutingConfig) {
	if !smfContext.ULCLSupport {
		return
	}

	if routingConfig == nil {
		logger.CtxLog.Error("configuration needs the routing config")
		return
	}

	logger.CtxLog.Infof("ue routing config Info: Version[%s] Description[%s]",
		routingConfig.Info.Version, routingConfig.Info.Description)

	UERoutingInfo := routingConfig.UERoutingInfo
	smfContext.UEPreConfigPathPool = make(map[string]*UEPreConfigPaths)
	smfContext.UEDefaultPathPool = make(map[string]*UEDefaultPaths)
	smfContext.ULCLGroups = make(map[string][]string)

	for groupName, routingInfo := range UERoutingInfo {
		logger.CtxLog.Debugln("Set context for ULCL group: ", groupName)
		smfContext.ULCLGroups[groupName] = routingInfo.Members
		uePreConfigPaths, err := NewUEPreConfigPaths(routingInfo.SpecificPaths)
		if err != nil {
			logger.CtxLog.Warnln(err)
		} else {
			smfContext.UEPreConfigPathPool[groupName] = uePreConfigPaths
		}
		ueDefaultPaths, err := NewUEDefaultPaths(smfContext.UserPlaneInformation, routingInfo.Topology)
		if err != nil {
			logger.CtxLog.Warnln(err)
		} else {
			smfContext.UEDefaultPathPool[groupName] = ueDefaultPaths
		}
	}
}

func SMF_Self() *SMFContext {
	return &smfContext
}

func GetUserPlaneInformation() *UserPlaneInformation {
	return smfContext.UserPlaneInformation
}

func GetUEDefaultPathPool(groupName string) *UEDefaultPaths {
	return smfContext.UEDefaultPathPool[groupName]
}
