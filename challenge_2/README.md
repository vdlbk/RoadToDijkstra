# Challenge 2

https://www.youtube.com/watch?v=cdCeU8DJvPM

## First Round

Create a function that reverses a string


## Second Round

Optimize this function

## Third round

Find a missing item between two arrays.   
The second array contains all first items expect one, the function has to find it.    
For the challenge purposes, arrays are just arrays of int.

| a1 | a2 | expected result |
| -- | -- | --------------- |
| `[4, 12, 9, 5, 6]` | `[4, 9, 12, 6]` | 5 |
| `[4]` | `[]` | 4 |
| `[1, 2, 3]` | `[1, 2]` | 3 |
| `[7, 0, 23, 14, 8, 2, 4]` | `[0, 23, 14, 8, 2, 4]` | 7 |
| `[3, 5, 7, 9]` | `[2, 4, 6, 8]` | -1 |
| `[3, 5, 7, 9]` | `[3, 5, 7, 9]` | -1 |

### Solution

I choosed to sort the second array because If the current integer I try to find is less than the one i'm looging in the second array it's mean that it is missing in the second one.

*Note that there were no information about what's happen if there is the same number twice in a array. So I decided to ignore that*

### Complexity

The sort part takes `O(n log n)`   
The research part takes `O(n^2)` in the worth case I guess


