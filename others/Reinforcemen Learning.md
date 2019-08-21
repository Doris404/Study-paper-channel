## Reinforcement Learning
> һЩ���׶������ӣ�https://www.zhihu.com/question/41775291
* Markov decision processes
* Bellman equations
* Value Interation alg

#### Markov decision processes---MDP <S,A,T,R>

* states : $s \in S$
* actions : $a \in A$ 
* transitions : T(s,a,s') = P(s'|s,a) ��s״̬�²�ȡa�ж����ܵ���s'�ĸ���
* reward : R(s) R(s,a,s')...

**MDP**

<S,A,T,R>����ÿһ��״̬��Ѱ����ѵ�a $\pi(s)$��
ʹ�����ջ����������$\pi^*(s) = \argmax_aQ^*(s,a)$

**RL**

<S,A>

$T(s,a,s') = \frac^{count(s,a,s')}_{\sum_s'count(s,a,s')}

$R(s,a,s') = \frac^{\sum_tR(s,a,s')}_{count(s,a,s')}$


**Utility function---be bounded in some way**

* Final horizon: ֻ����n���Ľ���
* Discounted rewards: $R(s0)+��R(s1)+��^2R(s2)��$

**Policy**
ÿһ���������д���

#### Bellman Equations

���������̣�Bellman Equation����Ҳ����Ϊ��̬�滮���̡����������̽�һ�����ӵĶ�̬
�滮�������С���⣬�����������

> һ������BE��������Դ��https://zhuanlan.zhihu.com/p/35261164

$$V^*(s) = \max_{a}Q^*(s,a)$$
$$Q^*(s,a) = \sum_{s'}T(s,a,s')(R(s,a,s') + /gammaV^*(s'))$$

#### Value Interation alg

> https://blog.csdn.net/songrotek/article/details/51378582

$V^*_k(s)$: k��֮��״̬s�����value,��k���������ʱ��$V^*_k(s)$����$V(s)$

**VI alg**

![](1.png "VI alg")

**Q-value Iteration**







