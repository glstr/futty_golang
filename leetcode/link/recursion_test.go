package link

import "testing"

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	get := combinationSum(candidates, target)
	t.Errorf("get:%v", get)
}
