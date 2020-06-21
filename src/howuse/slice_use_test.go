package howuse

import "testing"

func TestRemoveFront(t *testing.T) {
	example := []int{1}
	ret, err := RemoveFront(&example)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("ret:%d, list:%v", ret, example)
}
