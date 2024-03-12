# Problem A2: Cheeseburger Corollary 2 Solution

### step 1
we all know according to the problem:
- $A$ is the cost of one single cheesburger
- $B$ is the cost of one double cheesburger
- $C$ is the money we have

let's make some definitions:
- $K_{max}$ is the answer we want K-decker cheeseburger.
- $N_S$ is the number of single cheeseburgers
- $N_D$ is the number of double cheeseburgers
- $N_{bun}$ is the number of buns of $K_{max}$-decker cheeseburger
- $N_{patty}$ is the number of buns of $K_{max}$-decker cheeseburger

we can find,

- $N_{buns} = 2N_S + 2N_D$
- $N_{patty} = N_S + 2N_D$

and

- $K_{max}$-decker cheeseburger has $K_{max} + 1$ buns
- $K_{max}$-decker cheeseburger has $K_{max}$ patties

so we get the derivation formula: 

- $2N_S + 2N_D = K_{max} + 1$
- $N_S + 2N_D = K_{max}$

and 

$K_{max} = min(2N_S + 2N_D-1, N_S + 2N_D)$

let's continue to list some cases:

### case 1
we buy 0 single, means $N_S = 0$, then 

$K_{1} = min(2N_D-1, 2N_D) = 2N_D - 1$

### case 2
we buy 1 single, means $N_S = 1$, then 

$K_{2} = min(2 + 2N_D-1, 1 + 2N_D) = 1+2N_D$

### case 3
we buy 2 and more than 2 singles, means $N_S \geq 2$, then 

$$
K_{max} = min(2N_S + 2N_D-1, N_S + 2N_D) \\
= N_S + 2N_D \\
= N_S + 2\lfloor (C-A*N_S)/B \rfloor   
$$

let's see this function: $y = x + 2(C-A*x)/B$, we should find the $y_{max}$ by according to $x$:

$y = x + 2(C-A*x)/B = (1-2A/B)x + 2C/B$

$$
y_{max} = 
  \begin{cases}
    (1-2A/B) * x_{max} + 2C/B = (1-2A/B) * \lfloor C/A \rfloor + 2C/B & \text{if 1-2A/B > 0} \\
    (1-2A/B) * x_{min} + 2C/B = (1-2A/B) * 2 + 2C/B & \text{if 1-2A/B < 0} 
  \end{cases}
$$

now we use $x_{max}$ to get $K_3$ means, we buy $\lfloor C/A \rfloor$ singles, no double.

and we use $x_{min}$ to get $K_4$ means, we buy 2 singles and as many doubles 

so, the final two situation is:

$K_3 = \lfloor C/A \rfloor$

$K_4 = N_S + 2N_D = 2 + 2 * (C - \lfloor C/A \rfloor * A) * B$

### conclusion

now we have $K_1$, $K_2$, $K_3$, $K_4$, finally we can get the answer $K_{max} = max(K_1, K_2, K_3, K_4)$


