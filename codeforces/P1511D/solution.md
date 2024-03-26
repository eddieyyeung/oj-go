## Solution

Problem: [D. Min Cost String](https://codeforces.com/problemset/problem/1511/D)

Tutorial: https://codeforces.com/blog/entry/89634

###  Note
The main thought of the solution is to find Eulerian Cycle in the Eulerian Circuit (The graph structure of problem has already matched the Eulerian Circuit). 

So we can use [Hierholzer’s Algorithm](https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/) to find the Eulerian Path in $O(E)$

### References

- [Hierholzer’s Algorithm for directed graph](https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/)
- [Euler Circuit in a Directed Graph](https://www.geeksforgeeks.org/euler-circuit-directed-graph/)