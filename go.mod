module github.com/free5gc/smf

go 1.14

require (
	github.com/antihax/optional v1.0.0
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/free5gc/aper v1.0.0
	github.com/free5gc/flowdesc v1.0.0
	github.com/free5gc/http2_util v1.0.0
	github.com/free5gc/http_wrapper v1.0.0
	github.com/free5gc/idgenerator v1.0.0
	github.com/free5gc/logger_conf v1.0.0
	github.com/free5gc/logger_util v1.0.0
	github.com/free5gc/nas v1.0.0
	github.com/free5gc/ngap v1.0.1
	github.com/free5gc/openapi v1.0.0
	github.com/free5gc/path_util v1.0.0
	github.com/free5gc/util_3gpp v1.0.0
	github.com/free5gc/version v1.0.0
	github.com/gin-gonic/gin v1.7.2
	github.com/google/uuid v1.1.2
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/nycu-ucr/pfcp v0.0.0-20220602163725-3105ebc0c901
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.4
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/free5gc/pfcp => /home/public/l25gc/onvm-pfcp3.0.5/

replace github.com/free5gc/logger_util => /home/public/l25gc/logger_util/
