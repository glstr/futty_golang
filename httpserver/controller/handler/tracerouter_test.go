package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTraceRouterHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TraceRouterHandler(tt.args.c)
		})
	}
}

func TestGetRouterInfoHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetRouterInfoHandler(tt.args.c)
		})
	}
}
