# Challenge 3: Permutations

https://leetcode.com/problems/permutations/

## Description
Given a collection of distinct integers, return all possible permutations.

Example:
```
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

###### Draft
```

1 => 1 (1*1)
2 => 2 (2*1) 1 other fixes possible position for each indexes
3 => 6 (3*2) 2 others fixed possible position for each indexes
4 => 24 (4*6) 6 others fixed possible position for each other indexes
5 =>  120 ??? (5*24 ???) 24 others fixed possible position for each indexes

1, 2, 3, 4
1, 2, 4, 3
1, 3, 2, 4
1, 4, 2, 3
1, 3, 4, 2
1, 4, 3, 2

2, 1, 3, 4
2, 1, 4, 3
2, 3, 1, 4
2, 4, 1, 3
2, 3, 4, 1
2, 4, 3, 1

3, 1, 2, 4
3, 1, 4, 2
3, 2, 1, 4
3, 4, 1, 2
3, 2, 4, 1
3, 4, 2, 1

4, 1, 2, 3
4, 1, 3, 2
4, 2, 1, 3
4, 3, 1, 2
4, 2, 3, 1
4, 3, 2, 1
```

### First round

#### Solution

2 solutions came in my mind:
* By recursion: for each item in the list, add it with the permuation of all other items. Example: for `1, 2, 3`, Add `1` in every permutation possible for `2, 3`. Which is `1, 2, 3` and `1, 3, 2`. This could work perfectly fine and in terms of complexity I don't know yet. Let's say that it's doesn't matter :smile:
* Also by recursion, I can calculate the number of possible position for each item and allocate in advance a matrix, and add item to it. I guess that the complexity should be better, but for now I can't assure, it's just an idea. For exemple, I know that the number of possible position is the current lenght of input multiplied by the result of the previous input size. For an array of 4 items, it should be something like `4 * (3 * (2 * (1 * 1))` 

Finally I decided to code just the first solution. The second looks great but it's a bit more complicated.


#### Complexity
To be determined