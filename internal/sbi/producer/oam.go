package producer

import (
	"github.com/nycu-ucr/gonet/http"
	"strconv"

	"github.com/nycu-ucr/openapi/models"
	"github.com/nycu-ucr/smf/internal/context"
	"github.com/nycu-ucr/util/httpwrapper"
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

func HandleOAMGetUEPDUSessionInfo(smContextRef string) *httpwrapper.Response {
	smContext := context.GetSMContextByRef(smContextRef)
	if smContext == nil {
		httpResponse := &httpwrapper.Response{
			Header: nil,
			Status: http.StatusNotFound,
			Body:   nil,
		}

		return httpResponse
	}

	httpResponse := &httpwrapper.Response{
		Header: nil,
		Status: http.StatusOK,
		Body: PDUSessionInfo{
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
		},
	}
	return httpResponse
}
