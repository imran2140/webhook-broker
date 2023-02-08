// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	data "github.com/newscred/webhook-broker/storage/data"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// DeliveryJobRepository is an autogenerated mock type for the DeliveryJobRepository type
type DeliveryJobRepository struct {
	mock.Mock
}

// DispatchMessage provides a mock function with given fields: message, deliveryJobs
func (_m *DeliveryJobRepository) DispatchMessage(message *data.Message, deliveryJobs ...*data.DeliveryJob) error {
	_va := make([]interface{}, len(deliveryJobs))
	for _i := range deliveryJobs {
		_va[_i] = deliveryJobs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.Message, ...*data.DeliveryJob) error); ok {
		r0 = rf(message, deliveryJobs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *DeliveryJobRepository) GetByID(id string) (*data.DeliveryJob, error) {
	ret := _m.Called(id)

	var r0 *data.DeliveryJob
	if rf, ok := ret.Get(0).(func(string) *data.DeliveryJob); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.DeliveryJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobsForConsumer provides a mock function with given fields: consumer, jobStatus, page
func (_m *DeliveryJobRepository) GetJobsForConsumer(consumer *data.Consumer, jobStatus data.JobStatus, page *data.Pagination) ([]*data.DeliveryJob, *data.Pagination, error) {
	ret := _m.Called(consumer, jobStatus, page)

	var r0 []*data.DeliveryJob
	if rf, ok := ret.Get(0).(func(*data.Consumer, data.JobStatus, *data.Pagination) []*data.DeliveryJob); ok {
		r0 = rf(consumer, jobStatus, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.DeliveryJob)
		}
	}

	var r1 *data.Pagination
	if rf, ok := ret.Get(1).(func(*data.Consumer, data.JobStatus, *data.Pagination) *data.Pagination); ok {
		r1 = rf(consumer, jobStatus, page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*data.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*data.Consumer, data.JobStatus, *data.Pagination) error); ok {
		r2 = rf(consumer, jobStatus, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetJobsForMessage provides a mock function with given fields: message, page
func (_m *DeliveryJobRepository) GetJobsForMessage(message *data.Message, page *data.Pagination) ([]*data.DeliveryJob, *data.Pagination, error) {
	ret := _m.Called(message, page)

	var r0 []*data.DeliveryJob
	if rf, ok := ret.Get(0).(func(*data.Message, *data.Pagination) []*data.DeliveryJob); ok {
		r0 = rf(message, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.DeliveryJob)
		}
	}

	var r1 *data.Pagination
	if rf, ok := ret.Get(1).(func(*data.Message, *data.Pagination) *data.Pagination); ok {
		r1 = rf(message, page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*data.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*data.Message, *data.Pagination) error); ok {
		r2 = rf(message, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetJobsInflightSince provides a mock function with given fields: delta
func (_m *DeliveryJobRepository) GetJobsInflightSince(delta time.Duration) []*data.DeliveryJob {
	ret := _m.Called(delta)

	var r0 []*data.DeliveryJob
	if rf, ok := ret.Get(0).(func(time.Duration) []*data.DeliveryJob); ok {
		r0 = rf(delta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.DeliveryJob)
		}
	}

	return r0
}

// GetJobsReadyForInflightSince provides a mock function with given fields: delta
func (_m *DeliveryJobRepository) GetJobsReadyForInflightSince(delta time.Duration) []*data.DeliveryJob {
	ret := _m.Called(delta)

	var r0 []*data.DeliveryJob
	if rf, ok := ret.Get(0).(func(time.Duration) []*data.DeliveryJob); ok {
		r0 = rf(delta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.DeliveryJob)
		}
	}

	return r0
}

// GetPrioritizedJobsForConsumer provides a mock function with given fields: consumer, jobStatus, pageSize
func (_m *DeliveryJobRepository) GetPrioritizedJobsForConsumer(consumer *data.Consumer, jobStatus data.JobStatus, pageSize int) ([]*data.DeliveryJob, error) {
	ret := _m.Called(consumer, jobStatus, pageSize)

	var r0 []*data.DeliveryJob
	if rf, ok := ret.Get(0).(func(*data.Consumer, data.JobStatus, int) []*data.DeliveryJob); ok {
		r0 = rf(consumer, jobStatus, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.DeliveryJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data.Consumer, data.JobStatus, int) error); ok {
		r1 = rf(consumer, jobStatus, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkJobDead provides a mock function with given fields: deliveryJob
func (_m *DeliveryJobRepository) MarkJobDead(deliveryJob *data.DeliveryJob) error {
	ret := _m.Called(deliveryJob)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.DeliveryJob) error); ok {
		r0 = rf(deliveryJob)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkJobDelivered provides a mock function with given fields: deliveryJob
func (_m *DeliveryJobRepository) MarkJobDelivered(deliveryJob *data.DeliveryJob) error {
	ret := _m.Called(deliveryJob)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.DeliveryJob) error); ok {
		r0 = rf(deliveryJob)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkJobInflight provides a mock function with given fields: deliveryJob
func (_m *DeliveryJobRepository) MarkJobInflight(deliveryJob *data.DeliveryJob) error {
	ret := _m.Called(deliveryJob)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.DeliveryJob) error); ok {
		r0 = rf(deliveryJob)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkJobRetry provides a mock function with given fields: deliveryJob, earliestDelta
func (_m *DeliveryJobRepository) MarkJobRetry(deliveryJob *data.DeliveryJob, earliestDelta time.Duration) error {
	ret := _m.Called(deliveryJob, earliestDelta)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.DeliveryJob, time.Duration) error); ok {
		r0 = rf(deliveryJob, earliestDelta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequeueDeadJobsForConsumer provides a mock function with given fields: consumer
func (_m *DeliveryJobRepository) RequeueDeadJobsForConsumer(consumer *data.Consumer) error {
	ret := _m.Called(consumer)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.Consumer) error); ok {
		r0 = rf(consumer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDeliveryJobRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeliveryJobRepository creates a new instance of DeliveryJobRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeliveryJobRepository(t mockConstructorTestingTNewDeliveryJobRepository) *DeliveryJobRepository {
	mock := &DeliveryJobRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
