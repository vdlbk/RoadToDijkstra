# Challenge 4: Daily temperatures

https://leetcode.com/problems/daily-temperatures/

## Description

Given a list of daily temperatures `T`, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put `0` instead.    

For example, given the list of temperatures `T = [73, 74, 75, 71, 69, 72, 76, 73]`, your output should be `[1, 1, 4, 2, 1, 1, 0, 0]`.

**Note:** The length of temperatures will be in the range `[1, 30000]`. Each temperature will be an integer in the range `[30, 100]`.

### First round

#### Solution

First, I choosed to trust directly the input so I won't check the length of my input and the validity of each temperature.    

I need to init first an array of the same length than my input.
A silly solution would be to find the next warmer temperature for each day. So for this first round, I'm going for it and let's see what happened.

##### Note post-solution
Even if it's a silly solution, it's quite a good solution that runs pretty well.


#### Complexity
`O(n log n)`

### Second round

The first solution works fine, but I definitely want to improve this.

#### Solution

The second solution wasn't obvious, but I got an intuition that we can do something by starting from the end.     
The idea is to memorize the index of the next warmer temperatures *(we've visited)* in an array of 101 temperatures.     
Then find the closest warmer temperature and compute the distance between the current position and his.

#### Complexity

For each item in the list we will, in the worst case, execute a loop of 71 item. We may interpret it as a negligible constant.
`O(71n) ==> O(n)`