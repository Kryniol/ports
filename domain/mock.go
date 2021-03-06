// Code generated by mockery v2.9.4. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockPortRepository is an autogenerated mock type for the PortRepository type
type mockPortRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *mockPortRepository) Get(ctx context.Context, id PortID) (*Port, error) {
	ret := _m.Called(ctx, id)

	var r0 *Port
	if rf, ok := ret.Get(0).(func(context.Context, PortID) *Port); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Port)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, PortID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, port
func (_m *mockPortRepository) Save(ctx context.Context, port Port) error {
	ret := _m.Called(ctx, port)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Port) error); ok {
		r0 = rf(ctx, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockPortReader is an autogenerated mock type for the PortReader type
type mockPortReader struct {
	mock.Mock
}

// Read provides a mock function with given fields: ctx
func (_m *mockPortReader) Read(ctx context.Context) (<-chan Port, error) {
	ret := _m.Called(ctx)

	var r0 chan Port
	if rf, ok := ret.Get(0).(func(context.Context) chan Port); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan Port)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
