// Code generated by mockery v2.14.0. DO NOT EDIT.

package mockery

import mock "github.com/stretchr/testify/mock"

// mockCalculation is an autogenerated mock type for the calculation type
type mockCalculation struct {
	mock.Mock
}

type mockCalculation_Expecter struct {
	mock *mock.Mock
}

func (_m *mockCalculation) EXPECT() *mockCalculation_Expecter {
	return &mockCalculation_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: a, b
func (_m *mockCalculation) Add(a int, b int) int {
	ret := _m.Called(a, b)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(a, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// mockCalculation_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type mockCalculation_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - a int
//   - b int
func (_e *mockCalculation_Expecter) Add(a interface{}, b interface{}) *mockCalculation_Add_Call {
	return &mockCalculation_Add_Call{Call: _e.mock.On("Add", a, b)}
}

func (_c *mockCalculation_Add_Call) Run(run func(a int, b int)) *mockCalculation_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *mockCalculation_Add_Call) Return(_a0 int) *mockCalculation_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTnewMockCalculation interface {
	mock.TestingT
	Cleanup(func())
}

// newMockCalculation creates a new instance of mockCalculation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockCalculation(t mockConstructorTestingTnewMockCalculation) *mockCalculation {
	mock := &mockCalculation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
