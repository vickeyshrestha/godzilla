// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HealthHandler is an autogenerated mock type for the HealthHandler type
type HealthHandler struct {
	mock.Mock
}

// HealthCheck provides a mock function with given fields: _a0, _a1
func (_m *HealthHandler) HealthCheck(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}
