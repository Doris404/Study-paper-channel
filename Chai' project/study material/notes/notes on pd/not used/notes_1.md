## /server/coordinator.go 
 
#### 变量

* coordinator 是一个结构体

* **scheduleCfg的定义有疑惑**
* coordinator.go中```func(c *coordinator) patrolRegions()```部分学习

#### 函数

* newCoordinator: creates a new coordinator
* patrolRegions: check if the regions need to do some operations
* drivePushOperator: push the unfinished operator to the excutor
* checkRegion: if PD has restarted, it need to check learners added before and promote them.Don't check isRaftLearnerEnabled cause it maybe disable learner feature but there are still some learners to promote.
* **run**:

## balance_leader.go

#### 变量

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

#### 函数

```
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
```

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

可以参考网站：https://tour.go-zh.org/methods/4的内容
* source store 和target store是两个关键，代码的逻辑是：transferLeaderOut是将leader从source store移出，transferLeaderIn是将leader从其它节点移到 target store，为了这两个操作产生了一些参数，例如：target score, source score
* 涉及到关键操作时才写简洁的注释，并非关键代码可以不写注释

## balance_region.go

#### 变量

* balaceRegionScheduler 是一个结构体

#### 函数

```
func (s *balanceRegionScheduler) Schedule(cluster schedule.Cluster, opInfluence schedule.OpInfluence) []*schedule.Operator {}
```
是一个关键的函数这个函数通过了解```schedule.Cluster```和```schedule.OpInfluence```得到一个使得region平衡的schedule.Operator

- 得到一个store，```stores := cluster.GetStores```，这个store就是一个tikv节点
- 得到一个source,```source := s.selector.SelectSource(cluster,stores)```，source是store（TiKV)里面有着最高分的region，这个region可以被选择为```balance source```
- 对应的错误处理，以及日志的更新
- 在```balanceRegionRetryLimit```的限制内进行尝试，尝试的具体操作是
	- ```region ：= cluster.RandFollowerRegion(source.Getld(),core.HealthRegion())```
	- 对应的错误处理
- 如果对region进行GetPeers操作后结果的长度不等于cluster.GetMaxRepilcas()，那么这个region是异常的region


