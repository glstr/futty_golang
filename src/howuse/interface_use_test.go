package howuse

import "testing"

func TestCrossTheRiver(t *testing.T) {
	ti := &Tiger{}
	CrossTheRiver(ti)
	d := &Duck{}
	CrossTheRiver(d)
}
