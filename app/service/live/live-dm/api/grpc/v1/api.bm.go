// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api/grpc/v1/api.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api/grpc/v1/api.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

// ============
// DM Interface
// ============

type DM interface {
	SendMsg(ctx context.Context, req *SendMsgReq) (resp *SendMsgResp, err error)

	GetHistory(ctx context.Context, req *HistoryReq) (resp *HistoryResp, err error)
}

var v1DMSvc DM

// @params SendMsgReq
// @router GET /xlive/live-dm/v1/dM/SendMsg
// @response SendMsgResp
func dMSendMsg(c *bm.Context) {
	p := new(SendMsgReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DMSvc.SendMsg(c, p)
	c.JSON(resp, err)
}

// @params HistoryReq
// @router GET /xlive/live-dm/v1/dM/GetHistory
// @response HistoryResp
func dMGetHistory(c *bm.Context) {
	p := new(HistoryReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DMSvc.GetHistory(c, p)
	c.JSON(resp, err)
}

// RegisterV1DMService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1DMService(e *bm.Engine, svc DM, midMap map[string]bm.HandlerFunc) {
	v1DMSvc = svc
	e.GET("/xlive/live-dm/v1/dM/SendMsg", dMSendMsg)
	e.GET("/xlive/live-dm/v1/dM/GetHistory", dMGetHistory)
}
