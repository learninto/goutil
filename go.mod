module github.com/learninto/goutil

go 1.14

require (
	bou.ke/monkey v1.0.2
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.2.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dlmiddlecote/sqlstats v1.0.1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-kiss/memcache v0.0.0-20210719092635-467cdb8e19df
	github.com/go-kiss/net/pool v0.0.0-20210719091328-f4192f07e5b8
	github.com/go-kiss/redis v0.0.0-20210719094043-637dbcd540c2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.1
	github.com/jarcoal/httpmock v1.0.6
	github.com/json-iterator/go v1.1.10
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-sqlite3 v1.14.3 // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.3.0
	github.com/robfig/cron v1.2.0
	github.com/shopspring/decimal v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/gjson v1.8.1
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.uber.org/automaxprocs v1.3.0
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f
	google.golang.org/protobuf v1.25.0
)

replace github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.9.0
