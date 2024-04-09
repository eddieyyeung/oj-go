## Solution

Problem: [C. Table Decorations](https://codeforces.com/problemset/problem/478/C)

###  Note
There is three situations:

First, if $max(r,g,b) - min(r,g,b) <= 1$, then the answer is $(r+g+b)/3$.

Second, if $max(r,g,b) \geq 2(r+g+b - max(r,g,b))$, then the answer is $r+g+b - max(r,g,b)$.

Third, if $max(r,g,b) \lt 2(r+g+b - max(r,g,b))$, and not meet the condition $max(r,g,b) - min(r,g,b) <= 1$, then we can split to two components:

- $max(r_1, g_1, b_1) = 2(r_1+g_1+b_1-max(r_1,g_1,b_1))$
- $max(r_2,g_2,b_2) - min(r_2,g_2,b_2) <= 1$

where is $(r,b,g) = (r_1,g_1,b_1) + (r_2,g_2,b_2)$, and the answers of both two components are:

- $r_1+g_1+b_1 - max(r_1,g_1,b_1)=(r_1+g_1+b_1)/3$
- $(r_2+g_2+b_2)/3$

and we can conclude the final answer is $(r_1+g_1+b_1)/3+(r_2+g_2+b_2)/3=(r+g+b)/3$. It is the same as first situation.