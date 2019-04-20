# Challenge 6: Number of Digit One

https://leetcode.com/problems/number-of-digit-one/

## Description

Given an integer n, count the total number of digit 1 appearing in all non-negative integers less than or equal to n.

### Exemple

**Input:** 13
**Output:** 6    
**Explanation:** Digit 1 occurred in the following numbers: 1, 10, 11, 12, 13.   


### First round
#### Solution

A silly solution would be to cast each integers in a string and count the number of 1. Just for my curiosity, 
I want to go for it in order to see how fast it can be and how it's possible to improve this solution.

It was indeed a really bad solution as expected, because the conversion to string take time.


#### Complexity
`O(n*log10(n))`



### Secound round
#### Solution

The first solution works but failed in terms of complexity and time. The main bottleneck for now is obviously the string conversion,
 so I want to get rid of it. This solution won't introduce any further string conversion which is definitively not what we want.    
My idea for this one, is to get the last digit of the current number, and do the same for the result of the division.    
I will increment the result by one each time the last digit is equals to `1`

##### Draft
```
n = 2112

2112 % 10 = 2;  2112 / 10 = 211;    r = 0
211  % 10 = 1;  211  / 10 = 21;     r = 1
21   % 10 = 1;  21   / 10 = 2;      r = 2
2    % 10 = 2;  1    / 10 = 0;      r = 2   
```

#### Complexity
`O(n*log10(n))`

#### Notes
Even if the complexity is the same *(at least I think)*, the algorithm runs quite faster without the string conversion, 
approximately 3 times faster.    
It works well but it's still slow and this solution is not acceptable for `leetcode`.    


### Solution

You will find the solution and benchmarks which show how fast it can be.

#### Complexity
`O(log10(n))`