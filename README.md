# Longest Common Subsequence Problem

Given two strings S { s1, s2, ..., sm } and T { t1, t2, ..., tn }, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.

If there is no common subsequence, return 0.

## Constraints

```
0 < m, n <= 1000
The input strings consist of lowercase English characters only.
```

## Examples

```
INPUT:

  S = "abcde"
  T = "ace" 

OUTPUT: 3  

The longest common subsequence is "ace" and its length is 3.
```

```
INPUT:

  S = "abc"
  T = "abc"

OUTPUT: 3

The longest common subsequence is "abc" and its length is 3.
```

```
INPUT:

  S = "abc"
  T = "def"

OUTPUT: 0

There is no such common subsequence, so the result is 0.
```

## Approach

This problem can be solved with Dynamic Programming.

Suppose DP[i][j] represents the longuest common sequence S[1..i] and T[1..j], then: 

```
i) When S[i+1] !=  T[j+1]

  DP[i+1][j+1] = MAX (
    DP[i][j+1],
    DP[i+1][j]
  )

ii) When S[i+1] ==  T[j+1]

  DP[i+1][j+1] = MAX (
    DP[i][j+1],
    DP[i+1][j],
    DP[i][j] + 1
  )
```

## Complexity Analysis

* Time Complexity: O(N*M)
* Space Complexity: O(N*M)

