package hello

import "testing"

func TestHello(t *testing.T){
  expected := "Hello Go!"
  actual := hello()
  if actual != expected{
    t.Errorf("Test failed, Expected: '%s' , Actual: '%s'",expected,actual)
  }
}
