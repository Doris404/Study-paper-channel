## TIKV ѧϰ�ʼ�

* Storage Engin
* Raft��һ����Э��
* gRPC��ͨѶ���
* Prometheus: ���ϵͳ
* Fail:����ע��Ŀ�
* TiKV
* PD

#### raft-rs
> һ����ҳ�������raft�㷨��http://thesecretlivesofdata.com/raft/

#### Prometheus
> rust-prometheus�Ǽ��ϵͳPrometheus��Rust�ͻ��˿⣬Prometheus֧�ֵ�4��ָ�꣺Counter,Gauge
,Histogram,Summary����֧��ǰ����,����ϸ�������https://pingcap.com/blog-cn/tikv-source-code-reading-4/
��https://pingcap.com/blog-cn/tikv-source-code-reading-3/

**Counter**

Counter����򵥣����õ�ָ�ꡣ���������֣�Counter��һ���ʺ��ڼ������ۼƵ�ָ�꣬����
����

**Gauge**

Gauge���������²�����ָ�꣬�ṩ��```inc()```,```dec()```,```add()```,```sub()```��```set()```,����
����ָ��

**Histogram**

Histogram��ֱ��ͼ�����˻��������⻹���Լ����λ��

**Summary**

SummaryĿǰ��δ��rust-promtheus��ʵ��

**�����÷�**

* ������Ҫ�ռ���ָ��
* �ڴ����ض�λ�õ���ָ���ṩ�Ľӿ��ռ���¼ָ��ֵ
* ʵ��HTTP Pull Service ʹ��Prometheus ���Զ��ڷ����ռ�����ָ�꣬����ʹ��
rust-prometheus�ṩ��Push ���ܶ��ڽ��ռ�����ָ���ϴ���Pushgateway

#### fail-rs 

> fail-rs ����TIKV��ͨ��������ʹ��fail point�������쳣������Ӷ����д�����ԣ������
https://pingcap.com/blog-cn/tikv-source-code-reading-5/

�ڷֲ�ʽϵͳ��ʱ��Ĺ�ϵ�Ƿǳ��ؼ��ģ���������������ִ��˳���෴���͵����˽�Ȼ��ͬ��
����������Ƕ������ݿ���˵����֤���ݵ�һ���Էǳ���Ҫ�������Ҫ��һЩ���Թ�������
����fail point���ڵ��������ڡ�

**�����÷�**

```
#[macro_use]
extern crate fail;

fn say_hello() {
    fail_point!(��before_print��);
    println!(��Hello World~��);
}

fn main() {
    say_hello();
    fail::cfg("before_print", "panic");
    say_hello();
}
```
���н�����£�
```
Hello World~
thread 'main' panicked at 'failpoint before_print panic' ...
```

**Fail point��Ϊ**

```
[<pct>%][<cnt>*]<type>[(args...)][-><more terms>]
```

* pct: ��Ϊ��ִ��ʱ�аٷ�֮pct�ļ��ʴ�
* cnt: ��Ϊ�ܹ��ܱ������Ĵ���
* type:��Ϊ����
* arts:��Ϊ����

#### raft-rs��־���ƹ���

> https://pingcap.com/blog-cn/tikv-source-code-reading-6/

**MsgAppend & MsgAppendResponse**

#### gRPC Server�ĳ�ʼ������������

> https://pingcap.com/blog-cn/tikv-source-code-reading-7/

* ����һ���������ڲ���Ϊÿһ����ɶ�������һ���߳�
* ���Ŵ���Server���󣬰󶨶˿ڣ�����һ����������ע�������Serverr��
* ������Server��start������������ľ���ʵ�ֹ���������Call�ϣ����������е���ɶ����С�

#### gRPC-rs�ķ�װ��ʵ��

> https://pingcap.com/blog-cn/tikv-source-code-reading-8/

gRPC C Core �ṩ��һ�������ṩgRPC�Ļ���ʵ�֣�������������Ҫ���``` grpc_channel```,```grpc_completion_queue```,```grpc_call```

#### Service�㴦�����̽���

> https://pingcap.com/blog-cn/tikv-source-code-reading-9/

Service �������```src/server```

#### Snapshot�ķ��ͺͽ���

>https://pingcap.com/blog-cn/tikv-source-code-reading-10/

ʲô��snapshot:

* snapshot��ĳ��ʱ��ϵͳ״̬�Ŀ��գ�������Ǵ˿�ϵͳ״̬���ݣ��Ա����û����Իָ���ϵͳ����ʱ�̵�״̬
* ������˵����ȫ���Խ�Snapshot������ͨ��```RaftMessage```�����ͣ��������������һЩʵ������
    * Snapshot��ʱ������������������������׵�������ӵ����������������Region����Raftѡ�ٳ�ʱ
    * �����ȴ�����Snapshot���ڴ�
    * ������Ϣ���ܵ���gRPC��message size��������
    
Ϊʲô��Ҫsnapshot:

* ��������£�leader��follower֮��ͨ��append log����ͬ����leader�ᶨ�ڴ�����ϵ�log�����follower����崻����ָ������ȱʧ��log�Ѿ���leader�ڵ�������ˣ�ֻ��ͨ��Snapshot����ͬ��
* Raft�����½ڵ㣬�����½ڵ�ûͬ�����κ���־��ֻ��ͨ������Snapshot��ͬ��
* ���ݡ��ָ�������Ӧ�ò���Ҫdumpһ��State Machine����������

���ʵ��snapshot:

* snap-worker��snapshot���շ�������������
    * **���ͣ�**��װ```RaftMessage```��һ��```SnapTask����Send```���񣬽���```snap-worker```����
    * **���գ�**��װ```RaftMessage```��һ��```SnapTask����Recv```���񣬽���```snap-worker```����


#### Storage-������Ʋ�

> https://pingcap.com/blog-cn/tikv-source-code-reading-11/

#### �ֲ�ʽ����















