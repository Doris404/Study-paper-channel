## 常用操作与基础知识

#### 在服务器上跑Pd模拟数据库 

1. 启动202.112.113.24服务器

2. 开启5个同样的页面

3. 使用下面的语句将它们全部转到202.112.113.25服务器
```
ssh root@202.112.113.25
```
4. 开启一个pd，具体方法是：先进入wyy/gopath/src/github.com/pingcap/pd,然后输入```make```编译pd,然后再运行sh -x run.sh运行一个pd


5. 开启三个TiKV,具体方法是：点开一个页面，进入```tikv```,输入```sh -x run.sh```,进入```tikv1```,输入```sh -x run.sh```,进入```tikv2```,输入```sh -x run.sh```

6. 开启go-ycsb,具体方法是：先进入wyy/gopath/src/github.com/pingcap/go-ycsb，然后输入```sh -x run.sh```

7. 为了看最终结果，可以打开日志文件看细节，可以用vim语句

#### 做测试

**Target**

1. 正常情况下tikv2上的leader size最大，其他都是0

2. 打印出来leadersize并查看

**遇到的问题**

1. swap 文件的产生是因为在这次vim操作之前，曾经发生过一次vim操作，该操作可能发生中断，所以在当前文件夹里产涩会给你了一个swap文件。这种swap文件不能由ls操作发现，想要发现它的存在可以通过ls  -a语句。通过删除swap文件可以使得这一报错得到解决。

2. "make"pd报错：cannot find package  


#### 常用技术

1. 在服务器上使用linux专用语句进行操作，参见：https://www.runoob.com/linux/linux-tutorial.html

