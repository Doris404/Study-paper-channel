### Reinforcement Learning

* Markov decision processes
* Bellman equations
* Value Interation alg

#### Markov decision processes---MDP <S,A,T,R>

* states : $s \in S$
* actions : $a \in A$ 
* transitions : T(s,a,s') = P(s'|s,a) 在s状态下采取a行动最总到达s'的概率
* reward : R(s) R(s,a,s')...

Utility function---be bounded in some way

* Final horizon: 只计算n步的奖励
* Discounted rewards: R(s0)+γR(s1)+γ^2R(s2)…

Policy


