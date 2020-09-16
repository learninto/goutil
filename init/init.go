package init

import (
	_ "net/http/pprof" // 注册 pprof 接口

	_ "go.uber.org/automaxprocs" // 根据容器配额设置 maxprocs
)
