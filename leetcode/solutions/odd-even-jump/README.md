[https://leetcode-cn.com/problems/odd-even-jump/](https://leetcode-cn.com/problems/odd-even-jump/)

### 1
暴力解法，超时

### 2

单调栈

### 3

基于红黑树的TreeMap，后续研究下红黑树

### 4

基本的原理就是使用了BST。

一开始打算用两棵BST，一棵来计算floor，一棵来计算ceiling。进过改进后，只要将重复的元素更新一下就可以共用一棵BST了。

采用了两个数组odds和evens来记录可以通过奇/偶次跳跃到当前元素。

后面随着元素插入到BST，再计算当前元素的floor和ceiling，再通过odds和evens记录的跳跃情况，来动态计算可到达次数count