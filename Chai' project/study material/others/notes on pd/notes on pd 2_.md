## Notes on PD 2 

Placement Driver(PD)��TiDB ����ȫ�������ܿؽڵ㣬������������Ⱥ�ĵ��ȣ�
����ȫ��ID�����ɣ��Լ�ȫ��ʱ���TSO�����ɵȡ�PD������������ȺTiKV��Ԫ��Ϣ
�����client�ṩ·�ɹ���

> https://pingcap.com/blog-cn/placement-driver/
> https://zhuanlan.zhihu.com/p/24809131?refer=newsql

PD���赣�ĵ���������⣬��etcd�����������⣬ͬʱPDͨ��etcd��raft��֤��
���ݵ�ǿһ���ԣ����õ������ݶ�ʧ����


#### ��ʼ��

PD ������ etcd������ͨ����������Ҫ���������������������ܱ�֤���ݵİ�ȫ��
�ֽ׶� PD ��**��Ⱥ������ʽ**��**initial-cluster** �ľ�̬��ʽ�Լ�** join**�Ķ�
̬��ʽ���ַ���

��etcd���棬������Ҫ����2379��2380�����˿ڣ�2379�����ⲿ����2380����
etcd peer֮���໥ͨ���õ�


#### ѡ��

��PD���������Ǿ���Ҫѡ��һ��leader�����ṩ������Ȼetcd������raft-leader
���������leader��PD��leader�ǲ�һ����

