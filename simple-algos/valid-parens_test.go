package algos

import "testing"

func TestValidParens(t *testing.T) {
	res := ValidParens("{[()]}")
	if res != true {
		t.Errorf("ValidParens({[()]}) = %t; want true", res)
	}
}