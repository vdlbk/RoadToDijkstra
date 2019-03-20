# Challenge 1

https://www.youtube.com/watch?v=10WnvBk9sZc

## Description

Write a function that takes two strings, `s1` and `s2` and returns the longest common subqequence of `s1` and `s2`
It's case sensitive

| `s1` | `s2` | expected result |
| ---- | ---- | --------------- |
| "ABAZDC" | "BACBAD" | "ABAD" |
| "AGGTAB" | "GXTXAYB" | "GTAB" |
| "aaaa" | "aa" | "aa" | 
| "" | "..." | "" |
| "ABBA" | "ABCABA" | "ABBA" |

### First round

1.  Naive algorithm
2.  Performance does not matter

#### Solution decided

I've decided to try to use 2 cursors