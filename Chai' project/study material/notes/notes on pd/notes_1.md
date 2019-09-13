## /server/coordinator.go 
 
#### 学习心得

* 在go中Channel是一个核心类型，你可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯(communication)。 timer和ticker是关于时间的两个channel
* go func(){}()以并发方式调用匿名函数func
* channel的操作符是```<-```
	- ch <- v // 发送值v到Channel ch中
	- v := <-ch // 从Channel ch 中接收数据，并将数据赋值给v变量
* coordinator的定义
```
type coordinator struct {
	sync.RWMutex

	wg     sync.WaitGroup
	ctx    context.Context//ctx.Done()
	cancel context.CancelFunc

	cluster          *RaftCluster//cluster.isPrepared(),cluster.opt.Load().Clone(),cluster.GetPatrolRegion,cluster.GetPatrolRegionInterval()
	learnerChecker   *checker.LearnerChecker
	replicaChecker   *checker.ReplicaChecker
	namespaceChecker *checker.NamespaceChecker
	mergeChecker     *checker.MergeChecker
	regionScatterer  *schedule.RegionScatterer
	schedulers       map[string]*scheduleController
	opController     *schedule.OperatorController
	classifier       namespace.Classifier
	hbStreams        *heartbeatStreams
}
```
* Package zap provides fast, structured, leveled logging. 详见网址：https://godoc.org/go.uber.org/zap
* patrolScanRegionLimit = 128 // It takes about 14 minutes to iterate 1 million regions.
* **scheduleCfg的定义有疑惑**
* coordinator.go中func(c *coordinator) patrolRegions()部分学习

## balance_leader.go

#### 定义变量

* balanceLeaderRetryLimit = 10 重试次数最多10次
* balanceLeaderScheduler 一个结构体，代表平衡leader的调度器，定义如下

```
type balanceLeaderScheduler struct {
	*baseScheduler
	name         string
	selector     *selector.BalanceSelector
	taintStores  *cache.TTLUint64
	opController *schedule.OperatorController
	counter      *prometheus.CounterVec
}
```

* BalanceLeaderCreateOption 一类函数，这类函数用于创造一个有选项的调度

#### 定义函数

* func init
* func newBalanceLeaderScheduler
* func WithBalanceLeaderCounter
* func WithBalanceLeaderName
* func (l *balanceLeaderScheduler) GetName() string
* func (l *balanceLeaderScheduler) GetType() string
* func (l *balanceLeaderScheduler) IsScheduleAllowed
* func (l *balanceLeaderScheduler) Schedule
* func (l *balanceLeaderScheduler) transferLeaderOut
* func (l *balanceLeaderScheduler) transferLeaderIn
* func (l *balanceLeaderScheduler) createOperator

##### func newBalanceLeaderScheduler

> newBalanceLeaderScheduler creates a scheduler that tends to keep leaders on each store balanced.

> 创造一个调度使得leaders在各个store上平衡分布（我们的想法与此不同，我们希望让性能最好的做leader）

##### func WithBalanceLeaderCounter

> WithBalanceLeaderCounter sets the counter for the scheduler.

##### func WithBalanceLeaderName

> WithBalanceLeaderName sets the name for the scheduler.

##### func WithBalanceLeaderName

> WithBalanceLeaderName sets the name for the scheduler.

##### func GetName()

> get the name of a balanceLeaderScheduler

##### func GetType()

> get the type of a balanceLeaderScheduler

##### func IsScheduleAllowed()

> return a bool 

##### func Schedule

> return a pointer to an operation

##### createOperator

> createOperator creates the operator according to the source and target store.If the region is hot or the difference between the two stores is tolerable, then no new operator neet to be created,otherwise create an operator that transfers the leader from the source store to the target store for the region.
 
 
##### 疑惑与发现
* go 语言创建子函数的方法```
func (l *balanceLeaderScheduler) Schedule(cluster schedule.Cluster) []*operator.Operator```是什么意思？
* source store 和target store是两个关键，代码的逻辑是：transferLeaderOut是将leader从source store移出，transferLeaderIn是将leader从其它节点移到 target store，为了这两个操作产生了一些参数，例如：target score, source score
* 涉及到关键操作时才写剪短的注释，并非关键代码可以不写注释
* 

