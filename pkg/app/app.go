package app

import (
	smf_context "github.com/nycu-ucr/smf/internal/context"
	"github.com/nycu-ucr/smf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *smf_context.SMFContext
	Config() *factory.Config
}
