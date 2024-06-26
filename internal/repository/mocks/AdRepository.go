// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	model "advertisement-rest-api-http-service/internal/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AdRepository is an autogenerated mock type for the AdRepository type
type AdRepository struct {
	mock.Mock
}

// CreateAd provides a mock function with given fields: ctx, ad
func (_m *AdRepository) CreateAd(ctx context.Context, ad *model.Ad) (string, error) {
	ret := _m.Called(ctx, ad)

	if len(ret) == 0 {
		panic("no return value specified for CreateAd")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Ad) (string, error)); ok {
		return rf(ctx, ad)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Ad) string); ok {
		r0 = rf(ctx, ad)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Ad) error); ok {
		r1 = rf(ctx, ad)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdByID provides a mock function with given fields: ctx, id, fields
func (_m *AdRepository) GetAdByID(ctx context.Context, id string, fields bool) (*model.Ad, error) {
	ret := _m.Called(ctx, id, fields)

	if len(ret) == 0 {
		panic("no return value specified for GetAdByID")
	}

	var r0 *model.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) (*model.Ad, error)); ok {
		return rf(ctx, id, fields)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *model.Ad); ok {
		r0 = rf(ctx, id, fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, id, fields)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAds provides a mock function with given fields: ctx, page, sort, order
func (_m *AdRepository) GetAds(ctx context.Context, page int, sort string, order string) ([]*model.Ad, error) {
	ret := _m.Called(ctx, page, sort, order)

	if len(ret) == 0 {
		panic("no return value specified for GetAds")
	}

	var r0 []*model.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) ([]*model.Ad, error)); ok {
		return rf(ctx, page, sort, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) []*model.Ad); ok {
		r0 = rf(ctx, page, sort, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, string) error); ok {
		r1 = rf(ctx, page, sort, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAdRepository creates a new instance of AdRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdRepository {
	mock := &AdRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
