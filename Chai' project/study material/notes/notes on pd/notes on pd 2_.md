## Notes on PD 2 

Placement Driver(PD)是TiDB 里面全局中心总控节点，它负责整个集群的调度，
负责全局ID的生成，以及全局时间戳TSO的生成等。PD保存着整个集群TiKV的元信息
负责给client提供路由功能

> https://pingcap.com/blog-cn/placement-driver/
> https://zhuanlan.zhihu.com/p/24809131?refer=newsql

PD无需担心单点故障问题，由etcd解决了这个问题，同时PD通过etcd的raft保证了
数据的强一致性，不用担心数据丢失问题


#### 初始化

PD 集成了 etcd，所以通常，我们需要启动至少三个副本，才能保证数据的安全。
现阶段 PD 有**集群启动方式**，**initial-cluster** 的静态方式以及** join**的动
态方式三种方法

在etcd里面，我们需要监听2379和2380两个端口，2379处理外部请求，2380则是
etcd peer之间相互通信用的


#### 选举

当PD启动后，我们就需要选出一个leader对外提供服务，虽然etcd自身有raft-leader
，但是这个leader和PD的leader是不一样的

