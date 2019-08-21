## TIKV 学习笔记

* Storage Engin
* Raft：一致性协议
* gRPC：通讯框架
* Prometheus: 监控系统
* Fail:错误注入的库
* TiKV
* PD

#### raft-rs
> 一个网页帮助理解raft算法：http://thesecretlivesofdata.com/raft/

#### Prometheus
> rust-prometheus是监控系统Prometheus的Rust客户端库，Prometheus支持的4种指标：Counter,Gauge
,Histogram,Summary，它支持前三种,具体细节请见：https://pingcap.com/blog-cn/tikv-source-code-reading-4/
和https://pingcap.com/blog-cn/tikv-source-code-reading-3/

**Counter**

Counter是最简单，常用的指标。正如其名字，Counter是一种适合于计数、累计的指标，单调
递增

**Gauge**

Gauge适用于上下波动的指标，提供的```inc()```,```dec()```,```add()```,```sub()```与```set()```,用于
更新指标

**Histogram**

Histogram即直方图，除了基本计数外还可以计算分位数

**Summary**

Summary目前还未在rust-promtheus中实现

**基本用法**

* 定义想要收集的指标
* 在代码特定位置调用指标提供的接口收集记录指标值
* 实现HTTP Pull Service 使得Prometheus 可以定期访问收集到的指标，或者使用
rust-prometheus提供的Push 功能定期将收集到的指标上传到Pushgateway

#### fail-rs 

> fail-rs 帮助TIKV在通常测试中使用fail point来构建异常情况，从而进行代码测试，详见：
https://pingcap.com/blog-cn/tikv-source-code-reading-5/

在分布式系统中时序的关系是非常关键的，可能两个操作的执行顺序相反，就导致了截然不同的
结果，尤其是对于数据库来说，保证数据的一致性非常重要，因此需要做一些测试工作，这
便是fail point存在的意义所在。

**基本用法**

```
#[macro_use]
extern crate fail;

fn say_hello() {
    fail_point!(“before_print”);
    println!(“Hello World~”);
}

fn main() {
    say_hello();
    fail::cfg("before_print", "panic");
    say_hello();
}
```
运行结果如下：
```
Hello World~
thread 'main' panicked at 'failpoint before_print panic' ...
```

**Fail point行为**

```
[<pct>%][<cnt>*]<type>[(args...)][-><more terms>]
```

* pct: 行为被执行时有百分之pct的几率触
* cnt: 行为总共能被触发的次数
* type:行为类型
* arts:行为参数

#### raft-rs日志复制过程

> https://pingcap.com/blog-cn/tikv-source-code-reading-6/

**MsgAppend & MsgAppendResponse**

#### gRPC Server的初始化和启动流程

> https://pingcap.com/blog-cn/tikv-source-code-reading-7/

* 创建一个环境，内部会为每一个完成队列启动一个线程
* 接着创建Server对象，绑定端口，并将一个或多个服务注册在这个Serverr上
* 最后调用Server的start方法，将服务的具体实现关联到若干Call上，并塞进所有的完成队列 
中。

#### gRPC-rs的封装与实现

> https://pingcap.com/blog-cn/tikv-source-code-reading-8/


















