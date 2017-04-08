package calculadora

import "testing"

func Test_add_1_1(t *testing.T){
  expected := 2
  actual := add(1,1)
  if actual != expected{
    t.Errorf("Test failed, Expected: '%d' , Actual: '%d'",expected,actual)
  }
}

func Test_add(t *testing.T){
  expected := 4
  actual := add(1,3)
  if actual != expected{
    t.Errorf("Test failed, Expected: '%d' , Actual: '%d'",expected,actual)
  }
}
