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
* ������Server��start������������ľ���ʵ�ֹ���������Call�ϣ����������е���ɶ��� 
�С�

#### gRPC-rs�ķ�װ��ʵ��

> https://pingcap.com/blog-cn/tikv-source-code-reading-8/


















