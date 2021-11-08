### 介绍

DFS(Depth First Search)意思是深度优先搜索，在寻找**所有可能性**类型的题目中经常可以用到。

解决回溯问题，关键是这三个点：
1. 已经做出的选择
2. 还可以做的选择
3. 回溯的终止条件

#### dfs框架

```python
result = []
def dfs(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```


### LeetCode相关题目

[46.全排列](https://leetcode-cn.com/problems/permutations/)

[51.N皇后](https://leetcode-cn.com/problems/n-queens/)

