package consumer

import (
	"context"
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/openapi"
	"github.com/nycu-ucr/openapi/Nsmf_PDUSession"
	"github.com/nycu-ucr/openapi/models"
	"github.com/nycu-ucr/smf/internal/logger"
)

func SendSMContextStatusNotification(uri string) (*models.ProblemDetails, error) {
	if uri != "" {
		request := models.SmContextStatusNotification{}
		request.StatusInfo = &models.StatusInfo{
			ResourceStatus: models.ResourceStatus_RELEASED,
		}
		configuration := Nsmf_PDUSession.NewConfiguration()
		client := Nsmf_PDUSession.NewAPIClient(configuration)

		logger.CtxLog.Infoln("[SMF] Send SMContext Status Notification")
		httpResp, localErr := client.
			IndividualSMContextNotificationApi.
			SMContextNotification(context.Background(), uri, request)

		if localErr == nil {
			if httpResp.StatusCode != http.StatusNoContent {
				return nil, openapi.ReportError("Send SMContextStatus Notification Failed")
			}

			logger.PduSessLog.Tracef("Send SMContextStatus Notification Success")
		} else if httpResp != nil {
			defer func() {
				if resCloseErr := httpResp.Body.Close(); resCloseErr != nil {
					logger.ConsumerLog.Errorf("SMContextNotification response body cannot close: %+v", resCloseErr)
				}
			}()
			logger.PduSessLog.Warnf("Send SMContextStatus Notification Error[%s]", httpResp.Status)
			if httpResp.Status != localErr.Error() {
				return nil, localErr
			}
			problem := localErr.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
			return &problem, nil
		} else {
			logger.PduSessLog.Warnln("Http Response is nil in comsumer API SMContextNotification")
			return nil, openapi.ReportError("Send SMContextStatus Notification Failed[%s]", localErr.Error())
		}
	}
	return nil, nil
}