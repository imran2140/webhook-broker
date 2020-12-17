// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	data "github.com/imyousuf/webhook-broker/storage/data"
	mock "github.com/stretchr/testify/mock"
)

// ConsumerRepository is an autogenerated mock type for the ConsumerRepository type
type ConsumerRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: consumer
func (_m *ConsumerRepository) Delete(consumer *data.Consumer) error {
	ret := _m.Called(consumer)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.Consumer) error); ok {
		r0 = rf(consumer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: channelID, consumerID
func (_m *ConsumerRepository) Get(channelID string, consumerID string) (*data.Consumer, error) {
	ret := _m.Called(channelID, consumerID)

	var r0 *data.Consumer
	if rf, ok := ret.Get(0).(func(string, string) *data.Consumer); ok {
		r0 = rf(channelID, consumerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Consumer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(channelID, consumerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *ConsumerRepository) GetByID(id string) (*data.Consumer, error) {
	ret := _m.Called(id)

	var r0 *data.Consumer
	if rf, ok := ret.Get(0).(func(string) *data.Consumer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Consumer)
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

// GetList provides a mock function with given fields: channelID, page
func (_m *ConsumerRepository) GetList(channelID string, page *data.Pagination) ([]*data.Consumer, *data.Pagination, error) {
	ret := _m.Called(channelID, page)

	var r0 []*data.Consumer
	if rf, ok := ret.Get(0).(func(string, *data.Pagination) []*data.Consumer); ok {
		r0 = rf(channelID, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.Consumer)
		}
	}

	var r1 *data.Pagination
	if rf, ok := ret.Get(1).(func(string, *data.Pagination) *data.Pagination); ok {
		r1 = rf(channelID, page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*data.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *data.Pagination) error); ok {
		r2 = rf(channelID, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Store provides a mock function with given fields: consumer
func (_m *ConsumerRepository) Store(consumer *data.Consumer) (*data.Consumer, error) {
	ret := _m.Called(consumer)

	var r0 *data.Consumer
	if rf, ok := ret.Get(0).(func(*data.Consumer) *data.Consumer); ok {
		r0 = rf(consumer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Consumer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data.Consumer) error); ok {
		r1 = rf(consumer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
