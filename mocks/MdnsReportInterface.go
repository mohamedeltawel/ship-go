// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"
)

// MdnsReportInterface is an autogenerated mock type for the MdnsReportInterface type
type MdnsReportInterface struct {
	mock.Mock
}

type MdnsReportInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MdnsReportInterface) EXPECT() *MdnsReportInterface_Expecter {
	return &MdnsReportInterface_Expecter{mock: &_m.Mock}
}

// ReportMdnsEntries provides a mock function with given fields: entries
func (_m *MdnsReportInterface) ReportMdnsEntries(entries map[string]*api.MdnsEntry) {
	_m.Called(entries)
}

// MdnsReportInterface_ReportMdnsEntries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReportMdnsEntries'
type MdnsReportInterface_ReportMdnsEntries_Call struct {
	*mock.Call
}

// ReportMdnsEntries is a helper method to define mock.On call
//   - entries map[string]*api.MdnsEntry
func (_e *MdnsReportInterface_Expecter) ReportMdnsEntries(entries interface{}) *MdnsReportInterface_ReportMdnsEntries_Call {
	return &MdnsReportInterface_ReportMdnsEntries_Call{Call: _e.mock.On("ReportMdnsEntries", entries)}
}

func (_c *MdnsReportInterface_ReportMdnsEntries_Call) Run(run func(entries map[string]*api.MdnsEntry)) *MdnsReportInterface_ReportMdnsEntries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]*api.MdnsEntry))
	})
	return _c
}

func (_c *MdnsReportInterface_ReportMdnsEntries_Call) Return() *MdnsReportInterface_ReportMdnsEntries_Call {
	_c.Call.Return()
	return _c
}

func (_c *MdnsReportInterface_ReportMdnsEntries_Call) RunAndReturn(run func(map[string]*api.MdnsEntry)) *MdnsReportInterface_ReportMdnsEntries_Call {
	_c.Call.Return(run)
	return _c
}

// NewMdnsReportInterface creates a new instance of MdnsReportInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMdnsReportInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MdnsReportInterface {
	mock := &MdnsReportInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
