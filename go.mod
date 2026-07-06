module github.com/nycu-ucr/smf

go 1.21

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/davecgh/go-spew v1.1.1
	github.com/free5gc/aper v1.0.6-0.20250102035630-3ddc831eed6a
	github.com/free5gc/ngap v1.0.9
	github.com/google/uuid v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	github.com/smartystreets/goconvey v1.8.1
	github.com/stretchr/testify v1.9.0
	github.com/urfave/cli v1.22.5
	go.uber.org/mock v0.4.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/nycu-ucr/gin v0.0.0-20230307121200-573befe2dfbe
	github.com/nycu-ucr/gock v0.0.0-20230928062355-60066588379a
	github.com/nycu-ucr/gonet v0.0.0-20250423025127-ecbd471be7ee
	github.com/nycu-ucr/nas v0.0.0-20250219150919-dbf6340e0fb2
	github.com/nycu-ucr/openapi v0.0.0-20250423030004-ca6aa49065b0
	github.com/nycu-ucr/pfcp v0.0.0-20231101064624-0deadf0e6271
	github.com/nycu-ucr/util v0.0.0-20230928120650-6ed674c090b3
)

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/free5gc/logger_conf v1.0.0 // indirect
	github.com/free5gc/logger_util v1.0.0 // indirect
	github.com/free5gc/openapi v1.0.9-0.20250102040443-6ef25ba56c88 // indirect
	github.com/free5gc/path_util v1.0.0 // indirect
	github.com/free5gc/pfcp v1.0.6 // indirect
	github.com/free5gc/tlv v1.0.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.3 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nycu-ucr/net v0.0.0-20250421092837-e154d11df6a2 // indirect
	github.com/nycu-ucr/oauth2 v0.0.0-20250423025658-33ac33ae87b5 // indirect
	github.com/nycu-ucr/onvmpoller v0.0.0-20230807070551-64ddd3797912 // indirect
	github.com/nycu-ucr/sse v0.0.0-20221108140034-8e09fddc7347 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smarty/assertions v1.15.0 // indirect
	github.com/tim-ywliu/nested-logrus-formatter v1.3.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.49.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/nycu-ucr/onvmpoller => ../xio
replace github.com/nycu-ucr/openapi => ../../third_party/openapi
replace github.com/nycu-ucr/util => ../../third_party/util-20230928
