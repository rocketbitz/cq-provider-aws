// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: ElastiCache)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	elasticache "github.com/aws/aws-sdk-go-v2/service/elasticache"
	gomock "github.com/golang/mock/gomock"
)

// MockElastiCache is a mock of ElastiCache interface.
type MockElastiCache struct {
	ctrl     *gomock.Controller
	recorder *MockElastiCacheMockRecorder
}

// MockElastiCacheMockRecorder is the mock recorder for MockElastiCache.
type MockElastiCacheMockRecorder struct {
	mock *MockElastiCache
}

// NewMockElastiCache creates a new mock instance.
func NewMockElastiCache(ctrl *gomock.Controller) *MockElastiCache {
	mock := &MockElastiCache{ctrl: ctrl}
	mock.recorder = &MockElastiCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElastiCache) EXPECT() *MockElastiCacheMockRecorder {
	return m.recorder
}

// DescribeCacheClusters mocks base method.
func (m *MockElastiCache) DescribeCacheClusters(arg0 context.Context, arg1 *elasticache.DescribeCacheClustersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheClusters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheClusters indicates an expected call of DescribeCacheClusters.
func (mr *MockElastiCacheMockRecorder) DescribeCacheClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheClusters", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheClusters), varargs...)
}
