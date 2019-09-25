## notes on pd 6 : PD 的功能是什么

PD是 TiDB 里面全局中心总控节点，它负责整个集群的调度，负责全局 ID 的生成，以及全局时间戳 TSO 的生成等。PD 还保存着整个集群 TiKV 的元信息，负责给 client 提供路由功能。

PD集成了etcd，现阶段PD有集群启动方式，```initial-cluster```的静态方式以及```join```的动态方式

动态初始化：先启动pd1，再启动pd2,加入到pd1的集群李

etcd的端口需要监听2379和2380两个端口，其中2379用来处理外部请求的，2380在etcd peer之间相互通信用

PD启动之后，我们需要选出一个leader提供对外服务（PD的leader?)


#### 心跳

PD所有关于集群的数据都是由Tikv主动心跳上报的，pd对Tikv的调度也是在心跳的时候完成的。通常PD会处理两种心跳：1.Tikv自身store的心跳2.store里面region的leader peer 上报的心跳

PD在handleStoreHeartbeat函数处理store的心跳，具体步骤为将心跳里面当前的store的一些状态缓存在cache里面，store的状态包括store有多少个region有多少个region的leader peer在这个store上等

PD在handleRegionHeartbeat函数处理region的心跳，只有leader peer才上报所属的region信息，收到信息后PD将信息放到cache里面，如果PD发现region的epoch有变化，就将这个region的信息放入cache里面。然后PD会对这个region进行具体的调度，譬如发现peer数目不够，添加新的peer等

region的epoch里面有```conf_ver```和```version```分别表示region的不同版本状态，如果region发生了membership changes，则conf_ver+1，如果region发生了split或者merge，则version+1
