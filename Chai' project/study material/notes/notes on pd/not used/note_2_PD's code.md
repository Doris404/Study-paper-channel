## PD 代码阅读笔记 

> from blog：https://pingcap.com/blog-cn/pd-scheduler/

当数据量增量超过当前机器的物理存储极限时，我们需要将一部分数据迁移到其它机器上面去。TiKV是通过range的方式进行数据切分的，我们使用Region来表示一个数据range，每个Region有多个副本peer，通常为了保证安全我们会使用至少3个副本。

最开始系统只有一个region，当数据量增大到超过region最大规格（64MB）时，region就会分裂成2个region。region是PD调度TiKV的基本单位。当增加一个TiKV，PD会将原来TiKV里面分布的一些Region调度到这个新增的TiKV上面，以保证整个数据均衡地分布在多个TiKV上面。

以上，我们仅仅考虑了数据的均衡，我们另外应该也考虑计算的均衡。试想如果在有3个TiKV的情况下，所有leader都在某一个TiKV上面会造成这个TiKV的性能瓶颈，最好的方法就是leader能够均衡在不同的TiKV上面（我们的想法恰恰与此不同，我们并不需要最终leader在TiKV上的分布是平衡的）。

PD主要会对两种资源进行分配与调度，存储storage以及计算leader。

#### Scheduler

scheduler的定义可以在/server/schedulers.go中找到，具体定义如下：

```
type Scheduler interface {
	GetName() string
	// GetType should in accordance with the name passing to schedule.RegisterScheduler()
	GetType() string
	GetMinInterval() time.Duration
	GetNextInterval(interval time.Duration) time.Duration
	Prepare(cluster Cluster) error
	Cleanup(cluster Cluster)
	Schedule(cluster Cluster) []*operator.Operator
	IsScheduleAllowed(cluster Cluster) bool
}
```

**然而博客上写的好像和这个不一样??**

博客上写的是：

```
// Scheduler is an interface to schedule resources.
type Scheduler interface {
	GetName() string
	GetResourceKind() ResourceKind
	Schedule(cluster *clusterInfo) Operator
}
```

Scheduler是用来调度资源的接口，```GetName```返回Scheduler的名字，此名字唯一。```GetResourceKind```返回Scheduler要处理的资源类型（现阶段只有leader和storage两种）。Scheduler则是进行实际的调度，它需要的参数就是整个集群的信息，会产生实际调度操作Operator。

#### Operator

```
// Operator is an interface to schedule region.
type Operator interface {
	GetRegionID() uint64
	GetResourceKind() ResourceKind
	Do(region *regionInfo) (*pdpb.RegionHeartbeatResponse, bool)
}
```
from blog：https://pingcap.com/blog-cn/pd-scheduler/

#### Selector

```
// Selector is an interface to select source and target store to schedule.
type Selector interface {
	SelectSource(stores []*storeInfo, filters ...Filter) *storeInfo
	SelectTarget(stores []*storeInfo, filters ...Filter) *storeInfo
}
```
from blog：https://pingcap.com/blog-cn/pd-scheduler/

Selector根据传入的storeInfo列表，已经一批Fileter选择合适的source和
target，供Scheduler实际调度。

#### Filter

```
// Filter is an interface to filter source and target store.
type Filter interface {
	// Return true if the store should not be used as a source store.
	FilterSource(store *storeInfo) bool
	// Return true if the store should not be used as a target store.
	FilterTarget(store *storeInfo) bool
}
```
from blog：https://pingcap.com/blog-cn/pd-scheduler/

Filter返回true，不能选择这个store。

#### Controller

虽然我们希望调度越快越好，但是我们同时也应该保证调度不能影响现有的系统，不能对现有系统造成太大的波动。


例如，在做storage调度时，Pd需要将Region的某一个副本从一个TiKV迁移到另一个TiKV，该Region的leader peer会首先将在目标TiKV上面添加一个新的peer，leader生成当前region的snapshot然后发送给follower。follower收到snapshot将其apply到自己的状态机中。leader会给要迁移出去的peer发送删除命令。

```
// Controller is an interface to control the speed of different schedulers.
type Controller interface {
	Ctx() context.Context
	Stop()
	GetInterval() time.Duration
	AllowSchedule() bool
}
```
from blog：https://pingcap.com/blog-cn/pd-scheduler/

controller负责控制整个调度的速度，```GetInterval```返回调度间隔时间，即两次调度之间的间隔时长。```AllowSchedule```表示是否允许调度。

#### Coordinator

PD使用coordinator来管理所有Scheduler以及Controller。

```
// ScheduleController combines Scheduler with Controller.
type ScheduleController struct {
	Scheduler
	Controller
}
```
from blog：https://pingcap.com/blog-cn/pd-scheduler/

## pd 代码阅读 2

> https://pingcap.com/blog-cn/placement-driver/

> PD如何与TiKV,TiDB协作交互

#### 初始化

PD集成了etcd(一个分布式的，一致的key-value存储），通常我们需要至少三个副本，才能保证数据的安全。PD目前有两中集群启动方式，```initial-cluster```的静态方法与```join```的动态方法。

我们需要了解etcd的端口，默认监听2379和2380两个端口。2379处理外部请求，2380是etcd peer之间通信用的。```initial-cluster```应用2380端口，```join```应用2379端口。二者互斥，我们只能用一种方式初始化集群。

#### 选举

PD启动后会选出一个leader对外服务，这个leader与etcd中raft的leader不一样。PD的leader选举如下：

1. 检查集群中是否有leader，有则watch这个leader，leader掉了就重新开始1

2. 没有leader则开始campaign创建一个Lessor，通过etcd事务机制写入相关信息，如果leader key的CreateRevision为0，表明其他PD还没写入，则将自己的leader相关信息写入，同时带上一个lease（租期）。如果事务执行失败，表明其他的PD已经成为 leader，则重新回到1。

3. 成为leader，定期保活。PD崩溃，原先写入的leader key因为lease到期自动删除，其他的PD可以watch到，重新选举。

4. 初始化raft cluster，从etcd中重新载入集群的元信息，拿到最新的TSO信息。

5. 定期更新TSO，监听lessor是否过期，以及外面是否主动退出。

#### TSO

TSO是一个全局时间戳，它是TiDB实现分布式事务的基石。因此我们需要PD可以快速大量地为事务分配TSO，同时也需要保证分配的TSO是单调递增的。

TSO是一个int64的整形，它由physical time+logical time两个部分组成。physical time是当前unix time的毫秒时间，而logical time是一个最大```1～`8```的计数器。这意味着1ms，PD最多可以分配262144个TSO。

PD保存与分配TSO的策略如下：

1. 当PD成为leader，会从etcd获取上一次保存的时间，如果发现本地的时间比这个小，则继续等待直到当前时间大于这个值。

2. 但PD能分配 TSO后，首先会向etcd申请一个最大时间，例如：当前时间为t1,每次最多申请3s的时间窗口，PD会向etcd保存t1+3s的时间值，然后PD就能在内存中直接使用这一段时间窗口。当前时间t2大于t1+3s后，PD就会再向etcd继续更新为t2+3s

3. 因为PD在内存中保存了一个可分配的时间窗口，所以外面请求TSO的时候，PD能直接在内存里计算TSO并返回

4. client批量向PD获取TSO。

#### 心跳

PD所有关于集群的数据都是由TiKV主动心跳上报的，PD对于TiKV的调度也是在心跳的时候完成的。PD会处理两种心跳，TiKV自身store的心跳，store里面region的leaderpeer上报的心跳。

	**store的心跳**
	
	在```handleStoreHeartbeat```函数中处理，主要将心跳里当前的store的一些状态缓存到cache里面。store的状态包括该store有多少个region，有多少个region的leader peer在该store上面等。后续调度会应用这些信息。

	**region的心跳**

	在```handleRegionHeartbeat```函数中处理，只有leader peer才会上报所属Region的信息，follower peer不会。

region的epoch中有```conf_ver```和```version```，分别表示region的不同版本状态。无论在PD还是在TiKV我们都通过epoch判断region是否发生变化，从而拒绝一些危险操作。

#### Split/Merge

PD在Region的heartbeat里对Region进行调度，接着在heartbeat的返回值中判断调度是否成功。

	**split**

	1. leader peer定期检查Region所占空间是否超过一个阈值，超过了就分裂

	2. leader peer 向PD发送请求分裂的指令，PD在handleAskSplit里处理。一个Region因而分裂成两个，这两个分裂形成的Region，一个会继承之前Region的所有元信息，另一个是有PD重新生成的，并且返回给leader。

	3. leader peer写入一个split raft log，在apply的时候执行，这样region就分裂成两个

	4. 分裂成功后，TiKV告诉PD，PD就在handleReportSplit里面处理，更新cache相关信息，并持久化到etcd。

#### 路由

当client要对key写入一个值：

1. client先从PD获取key属于哪一个region，PD将这个region相关元信息返回。

2. client自己cache，然后直接给region的leader peer 发送命令。

3. 当Region的leader转移时，TiKV会返回```NotLead```的错误，并附新leader的地址，client在cache里随即更新

4. 当Region的version变化（例如split），则key可能已经落入新的Region，client此时会收到```StateCommand```错误，于是重新从PD获取，回到状态1。

## pd代码阅读3

#### PD: Placement Driver

> https://pingcap.com/blog-cn/the-design-and-implementation-of-multi-raft/

PD 是 TiKV 的全局中央控制器，存储整个 TiKV 集群的元数据信息，负责整个 TiKV 集群的调度，全局 ID 的生成，以及全局 TSO 授时等。

在 TiKV 里面，跟 PD 的交互是放在源码的 pd 目录下,我的工作应该在这个目录下展开。现在跟 PD 的交互都是通过自己定义的 RPC 实现，协议非常简单，在 pd/mod.rs 里面我们直接提供了用于跟 PD 进行交互的 Client trait，以及实现了 RPC Client。

**bootstrap_cluster**：当我们启动一个 TiKV 服务的时候，首先需要通过 is_cluster_bootstrapped 来判断整个 TiKV 集群是否已经初始化，如果还没有初始化，我们就会在该 TiKV 服务上面创建第一个 region。

**region_heartbeat**：定期 Region 向 PD 汇报自己的相关信息，供 PD 做后续的调度。譬如，如果一个 Region 给 PD 上报的 peers 的数量小于预设的副本数，那么 PD 就会给这个 Region 添加一个新的副本 Peer。

**store_heartbeat**：定期 store 向 PD 汇报自己的相关信息，供 PD 做后续调度。譬如，Store 会告诉 PD 当前的磁盘大小，以及剩余空间，如果 PD 发现空间不够了，就不会考虑将其他的 Peer 迁移到这个 Store 上面。

**ask_split/report_split**：当 Region 发现自己需要 split 的时候，就 ask_split 告诉 PD，PD 会生成新分裂 Region 的 ID ，当 Region 分裂成功之后，会 report_split 通知 PD。

#### Raftstore

在 TiKV 里面，Multi Raft 的实现是在 Raftstore 完成的，代码在 raftstore/store 目录。

#### Region

通常的数据分片算法就是 Hash 和 Range，TiKV 使用的 Range 来对数据进行数据分片。

**Region的protobuf协议**:

```
message Region {
    optional uint64 id                  = 1 [(gogoproto.nullable) = false];
    optional bytes  start_key           = 2;
    optional bytes  end_key             = 3;
    optional RegionEpoch region_epoch   = 4;
    repeated Peer   peers               = 5;
}
```

region_epoch：当一个 Region 添加或者删除 Peer，或者 split 等，我们就会认为这个 Region 的 epoch 发生的变化，RegionEpoch 的 conf_ver 会在每次做 ConfChange 的时候递增，而 version 则是会在每次做 split/merge 的时候递增。

#### RocksDB / Keys Prefix

#### Peer storage

在 TiKV 里面，一个 Peer 的创建有如下几种方式：

    1. 主动创建，通常对于第一个 Region 的第一个副本 Peer，我们采用这样的创建方式，初始化的时候，我们会将它的 Log Term 和 Index 设置为 5。

    2. 被动创建，当一个 Region 添加一个副本 Peer 的时候，当这个ConfChange 命令被 applied 之后， Leader 会给这个新增 Peer 所在的 Store 发送Message，Store 收到这个 Message 之后，发现并没有相应的 Peer 存在，并且确定这个Message 是合法的，就会创建一个对应的 Peer，但此时这个 Peer 是一个未初始化的 Peer，不知道所在的 Region 任何的信息，我们使用 0 来初始化它的 Log Term 和 Index。Leader 就能知道这个 Follower 并没有数据（0 到 5 之间存在 Log 缺口），Leader 就会给这个 Follower 直接发送 snapshot。

    3.Split 创建，当一个 Region 分裂成两个 Region，其中一个 Region 会继承分裂之前 Region 的元信息，只是会将自己的 Range 范围修改。而另一个 Region 相关的元信息，则会新建，新建的这个 Region 对应的 Peer，初始的 Log Term 和 Index 也是 5，因为这时候 Leader 和 Follower 都有最新的数据，不需要 snapshot。（注意：实际 Split 的情况非常的复杂，有可能也会出现发送 snapshot 的情况，但这里不做过多说明）。

#### Peer

对 Raft 的 Propose，ready 的处理都是在 Peer 里面完成的。

**Propose函数**

Peer 的 propose 是外部 Client command 的入口。Peer 会判断这个 command 的类型：



    * 如果是只读操作，并且 Leader 仍然是在 lease 有效期内，Leader 就能直接提供 local read，不需要走 Raft 流程。

    * 如果是 Transfer Leader 操作，Peer 首先会判断自己还是不是 Leader，同时判断需要变成新 Leader 的 Follower 是不是有足够新的 Log，如果条件都满足，Peer 就会调用 RawNode 的 transfer_leader 命令。

    * 如果是 Change Peer 操作，Peer 就会调用 RawNode propose_conf_change。

    * 剩下的，Peer 会直接调用 RawNode 的 propose。

在 propose 之前，Peer 也会将这个 command 对应的 callback 存到 PendingCmd 里面，当对应的 log 被 applied 之后，会通过 command 里面唯一的 uuid 找到对应的 callback 调用，并给 Client 返回相应的结果。

为了保证数据的一致性，Peer 在 execute 的时候，都只会将修改的数据保存到 RocksDB 的 WriteBatch 里面，然后在最后原子的写入到 RocksDB，写入成功之后，才修改对应的内存元信息。如果写入失败，我们会直接 panic，保证数据的完整性。

**Ready函数**

在 Peer 处理 ready 的时候，我们还会传入一个 Transport 对象，用来让 Peer 发送 message。它就只有一个函数 send，TiKV 实现的 Transport 会将需要 send 的 message 发到 Server 层，由 Server 层发给其他的节点。

#### Multi Raft

#### Server

Server 层就是 TiKV 的网络层





