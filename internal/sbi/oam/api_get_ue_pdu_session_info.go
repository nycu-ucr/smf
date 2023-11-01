package oam

import (
	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/smf/internal/sbi/producer"
	"github.com/nycu-ucr/util/httpwrapper"
)

func HTTPGetUEPDUSessionInfo(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["smContextRef"] = c.Params.ByName("smContextRef")

	smContextRef := req.Params["smContextRef"]
	HTTPResponse := producer.HandleOAMGetUEPDUSessionInfo(smContextRef)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}