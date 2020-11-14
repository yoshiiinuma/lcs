
package lcs

import (
  "fmt"
  "strings"
)

const DEBUG = true

func LCS(S string, T string) int {
  if DEBUG {
    fmt.Println("------------------------------------------")
    fmt.Println(S, T)
    fmt.Println("------------------------------------------")
  }
  s := strings.Split(S, "")
  t := strings.Split(T, "")
  return lcs(s, t)
}

func setupDP(s []string, t []string) ([][]int, int, int) {
  N := len(s) + 1
  M :=  len(t) + 1
  DP := make([][]int, N)
  for i := 0; i < N; i++ {
    DP[i] = make([]int, M)
  }
  return DP, N, M
}

func showDP(DP [][]int) {
  N := len(DP)
  for i := 0; i < N; i++ {
    fmt.Printf("%v\n", DP[i])
  }
}

func lcs(s []string, t []string) int {
  DP, N, M := setupDP(s, t)

  for i := 1; i < N; i++ {
    for j := 1; j < M; j++ {
      DP[i][j] = DP[i-1][j]
      if DP[i][j] < DP[i][j-1] {
        DP[i][j] = DP[i][j-1]
      }
      if s[i-1] == t[j-1] {
        if DP[i][j] < DP[i-1][j-1] + 1 {
          DP[i][j] = DP[i-1][j-1] + 1
        }
      }
    }
  }
  if DEBUG { showDP(DP) }
  return DP[N-1][M-1]
}
