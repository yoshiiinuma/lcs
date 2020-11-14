
package lcs

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  var S string
  var T string

  S = "abc"
  T = "abc"
  Equal(t, 3, LCS(S, T))

  S = "abc"
  T = "def"
  Equal(t, 0, LCS(S, T))

  S = "abcde"
  T = "ace"
  Equal(t, 3, LCS(S, T))

  S = "ace"
  T = "abcde"
  Equal(t, 3, LCS(S, T))

  S = "ayyayyayya"
  T = "axaxa"
  Equal(t, 3, LCS(S, T))
}
