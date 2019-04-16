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

#### Solution

We have to push every element of the `pushed` in a stack, and each time, try to pop every element from this stack 
if the last element is equal to the current element of `popped`.    
At the end, if we succeed, the number of item popped must be equal to `|pushed|`

Exemple:    
```
// ... push `pushed` elements to `stack`
stack = [1, 2, 3, 4] 
stack[last] == popped[0] ? pop() 
stack = [1, 2, 3]
stack[last] == popped[1] ? push()
stack = [1, 2, 3, 5]
stack[last] == popped[1] ? pop ()
stack = [1, 2, 3]
stack[last] == popped[2] ? pop ()
stack = [1, 2]
stack[last] == popped[3] ? pop ()
stack = [1]
stack[last] == popped[4] ? pop ()
```


#### Complexity
`O(n)`

#### Reality

I totally failed this one, this solution is not mine. It became counter productive to continue on this one. 
I was stuck in a bad idea that never could have worked.
I wasn't able to think differently and I was quite disappointed by my thoughts on this one.     
The solution looks really easy and yet, I didn't succeed. I think it's time to forget it and try another one.