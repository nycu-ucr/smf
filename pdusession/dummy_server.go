package pdusession

import (
	"log"

	"github.com/nycu-ucr/gonet/http"

	"github.com/free5gc/path_util"
	"github.com/nycu-ucr/http2_util"
	"github.com/nycu-ucr/logger_util"
	"github.com/nycu-ucr/smf/logger"
	"github.com/nycu-ucr/smf/pfcp"
	"github.com/nycu-ucr/smf/pfcp/udp"
)

func DummyServer() {
	router := logger_util.NewGinWithLogrus(logger.GinLog)

	AddService(router)

	go udp.Run(pfcp.Dispatch)

	smfKeyLogPath := path_util.Free5gcPath("free5gc/smfsslkey.log")
	smfPemPath := path_util.Free5gcPath("free5gc/support/TLS/smf.pem")
	smfkeyPath := path_util.Free5gcPath("free5gc/support/TLS/smf.key")

	var server *http.Server
	if srv, err := http2_util.NewServer(":29502", smfKeyLogPath, router); err != nil {
	} else {
		server = srv
	}

	if err := server.ListenAndServeTLS(smfPemPath, smfkeyPath); err != nil {
		log.Fatal(err)
	}
}
