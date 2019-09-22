## Notes 3

> DYNAMO：亚马逊数据库技术
> redis: 可持久化（备份到硬盘） 高性能kv Nosql 内存数据库  

#### 频繁模式 Frequent Pattern(FM)

> Principles of datamining: 6.7 pattern structure
> Data mining concepts and technology: chapter 7

**频繁项集模式（关联规则）**

格结构中存每一个模式，遍历格结构，找到最频繁的模式

有三种方法可供选择：
* 基于读
* 基于写
* 基于指针

**指标:Support & Confidence** p13

A -> B

support(支持度） = A,B同时出现次数/总共样例数

confidence(可信度)  =A B 同时出现次数/ A独自出现的次数 [条件概率]

Apriori 算法 p13






