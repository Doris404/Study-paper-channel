## 常用操作与基础知识

#### 在服务器上跑Pd模拟数据库 

1.启动202.112.113.24服务器
2.开启5个同样的页面
3.使用下面的语句将它们全部转到202.112.113.25服务器
```
ssh root@202.112.113.25
```
4.开启一个pd，具体方法是：先进入wyy/gopath/src/github.com/pingcap/pd,然后输入```make```
5.开启三个TiKV,具体方法是：点开一个页面，进入```tikv```,输入```sh -x run.sh```,进入```tikv1```,输入```sh -x run.sh```,进入```tikv2```,输入```sh -x run.sh```
6.开启go-ycsb,具体方法是：先进入wyy/gopath/src/github.com/pingcap/go-ycsb，然后输入```sh -x run.sh```
7.为了看最终结果，可以打开日志文件看细节，可以用vim语句