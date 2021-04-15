package model

type RateLimter interface {
	CanUse() bool
}

type SingleRateLimter struct {
}

type GlobalRateLimter struct {
}
