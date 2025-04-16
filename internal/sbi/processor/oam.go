package processor

import (
	"strconv"

	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/openapi/models"
	"github.com/nycu-ucr/smf/internal/context"
	"github.com/nycu-ucr/smf/pkg/factory"
)

type PDUSessionInfo struct {
	Supi         string
	PDUSessionID string
	Dnn          string
	Sst          string
	Sd           string
	AnType       models.AccessType
	PDUAddress   string
	SessionRule  models.SessionRule
	UpCnxState   models.UpCnxState
	Tunnel       context.UPTunnel
}

func (p *Processor) HandleOAMGetUEPDUSessionInfo(c *gin.Context, smContextRef string) {
	smContext := context.GetSMContextByRef(smContextRef)
	if smContext == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	pduSessionInfo := &PDUSessionInfo{
		Supi:         smContext.Supi,
		PDUSessionID: strconv.Itoa(int(smContext.PDUSessionID)),
		Dnn:          smContext.Dnn,
		Sst:          strconv.Itoa(int(smContext.SNssai.Sst)),
		Sd:           smContext.SNssai.Sd,
		AnType:       smContext.AnType,
		PDUAddress:   smContext.PDUAddress.String(),
		UpCnxState:   smContext.UpCnxState,
		// Tunnel: context.UPTunnel{
		// 	//UpfRoot:  smContext.Tunnel.UpfRoot,
		// 	ULCLRoot: smContext.Tunnel.UpfRoot,
		// },
	}
	c.JSON(http.StatusOK, pduSessionInfo)
}

func (p *Processor) HandleGetSMFUserPlaneInfo(c *gin.Context) {
	c.JSON(http.StatusOK, factory.SmfConfig.Configuration.UserPlaneInformation)
}
