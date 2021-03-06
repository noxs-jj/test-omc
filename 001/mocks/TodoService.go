// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TodoService is an autogenerated mock type for the TodoService type
type TodoService struct {
	mock.Mock
}

// AddMonthsInflation provides a mock function with given fields: ctx, income, months
func (_m *TodoService) AddMonthsInflation(ctx context.Context, income float32, months int) (float32, error) {
	ret := _m.Called(ctx, income, months)

	var r0 float32
	if rf, ok := ret.Get(0).(func(context.Context, float32, int) float32); ok {
		r0 = rf(ctx, income, months)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, float32, int) error); ok {
		r1 = rf(ctx, income, months)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
