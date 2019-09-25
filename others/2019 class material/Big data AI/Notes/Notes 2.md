## Notes 2

> DYNAMO：亚马逊数据库技术
> redis: 可持久化（备份到硬盘） 高性能kv Nosql 内存数据库
> Memcached: 高性能分布式数据库  
> Clustrix 集群数据库系统 Clustered Database System

#### 频繁模式 Frequent Pattern(FM)

> Principles of datamining: 6.7 pattern structure
> Data mining concepts and technology: chapter 7

**频繁项集模式（关联规则）**

**指标:Support & Confidence** p13

	A -> B
	support(支持度） = A,B同时出现次数/总共样例数
	confidence(可信度)  =A B 同时出现次数/ A独自出现的次数 [条件概率]

**Apriori 算法**
- 找到所有频繁项集
	- 由Lk生成Ck+1:自连接+检查是否违反公理
	- 由Ck生成Lk：计数，得到大于用户给出的最小频度的项形成Lk
- 生成关联规则:找到所有大于用户最小可信度的规则

对Apriori算法的改进P32

生成L1的同时生成哈希表，二者一起生成C2,减小C2的规模

P37:不在频繁桶中说明模式出现次数小于预期，但是出现在频繁桶中也不一定代表其模式出现频繁，因为有可能发生哈希冲突的情况

#### 序列模式挖掘 

**指标：Support & Confidence**
	
	support = 包含所给序列的序列数/总样例数
	confidence = 
