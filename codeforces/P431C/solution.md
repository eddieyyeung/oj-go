## Solution

Problem: [C. k-Tree](https://codeforces.com/problemset/problem/431/C)

It is a **dp** problem, so we can figure out a formula below:

$$
dp(n, state) = \begin{cases}
   \sum_{i=1}^{d-1}{dp(i, 0)} & \text{if $state=0$} \\
   \sum_{i=d}^{k}{dp(i, 0)} + \sum_{i=1}^{k}{dp(i,1)} & \text{if $state=1$}
\end{cases}
$$

where is,
- $n$, $k$, $d$ is the number where is mentioned in the problem.
- $state$ represents whether the current $dp$ has a value at least $d$
