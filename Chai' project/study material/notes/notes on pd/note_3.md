## PD 代码阅读笔记——从宏观结构角度出发

/pd

	/.github:向作者提出自己的代码时使用
	
	/client:PD client的相关代码，其中包括变量的定义，子函数的定义

	/cmd:只有一个文件main.go是运行pd的入口

	/conf:配置文件.toml

	/docs:一个api.html+2个介绍使用方法以及工作流程的md文件

	/pkg:写包的地方

	/scripts:.sh文件 shell

	/server:关键文件[1]

	/table:有关数据库的table

	/tests:测试代码区域

	/tools:工具性代码
	
	其他一些文件[2]

#### server 代码阅读

/pd/server

	/api:api的设计（.raml文件）

	/checker

	/config

	/core

	/id

	/join

	/kv

	/member

	/namespace

	/placement

	/region_syncer

	/schedule

	/schedulers

	/statistics

	/tso

	其它文件[3]

#### server 之中

最重要的是```.../server/schedule```,```.../server/schedulers```


#### schedulers

**balance_region.go**



**balance_leader.go**

