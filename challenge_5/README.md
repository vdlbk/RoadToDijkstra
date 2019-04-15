# Challenge 5: Validate Stack Sequences

https://leetcode.com/problems/validate-stack-sequences/

## Description

Given two sequences `pushed` and `popped` **with distinct values**, return `true` if and only if this could have been the result of a sequence of push and pop operations on an initially empty stack.

### Notes

Note:

+  0 <= `|pushed|`, `|popped|`  <= 1000
+  0 <= `pushed[i]`, `popped[i]` < 1000
+  `pushed` is a permutation of `popped`.
+  `pushed` and `popped` have distinct values.

### Exemples

#### Example 1
**Input:** pushed = [1,2,3,4,5], popped = [4,5,3,2,1]     
**Output:** true    
**Explanation:** We might do the following sequence:    
push(1), push(2), push(3), push(4), pop() -> 4,    
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1    

#### Example 2
**Input:** pushed = [1,2,3,4,5], popped = [4,3,5,1,2]    
**Output:** false    
**Explanation:** 1 cannot be popped before 2.    

### First round

#### Draft
```javascript
// pushed = [1, 2]
push(1)
push(2)
pop() // popped = [2]
pop() // popped = [2, 1]
// --------
push(1)
pop() // popped = [1]
push(2)
pop() // popped = [1, 2]
```

#### Solution

TODO

#### Complexity
`O(n)`
