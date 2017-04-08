package codebreaker

import (
  "testing"
    "strconv"
    "fmt"
)

func Test_codebreakerAciertaX(t *testing.T){
  expected := "X"
  actual := codebreaker("2584")
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }

}
