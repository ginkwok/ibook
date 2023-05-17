// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ginkwok/ibook/dal (interfaces: Dal)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/ginkwok/ibook/model"
	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDal is a mock of Dal interface.
type MockDal struct {
	ctrl     *gomock.Controller
	recorder *MockDalMockRecorder
}

// MockDalMockRecorder is the mock recorder for MockDal.
type MockDalMockRecorder struct {
	mock *MockDal
}

// NewMockDal creates a new mock instance.
func NewMockDal(ctrl *gomock.Controller) *MockDal {
	mock := &MockDal{ctrl: ctrl}
	mock.recorder = &MockDalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDal) EXPECT() *MockDalMockRecorder {
	return m.recorder
}

// CheckUser mocks base method.
func (m *MockDal) CheckUser(arg0 *gorm.DB, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockDalMockRecorder) CheckUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockDal)(nil).CheckUser), arg0, arg1, arg2)
}

// CreateResv mocks base method.
func (m *MockDal) CreateResv(arg0 *gorm.DB, arg1 *model.Reservation) (*model.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResv", arg0, arg1)
	ret0, _ := ret[0].(*model.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResv indicates an expected call of CreateResv.
func (mr *MockDalMockRecorder) CreateResv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResv", reflect.TypeOf((*MockDal)(nil).CreateResv), arg0, arg1)
}

// CreateRoom mocks base method.
func (m *MockDal) CreateRoom(arg0 *gorm.DB, arg1 *model.Room) (*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", arg0, arg1)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockDalMockRecorder) CreateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockDal)(nil).CreateRoom), arg0, arg1)
}

// CreateSeat mocks base method.
func (m *MockDal) CreateSeat(arg0 *gorm.DB, arg1 *model.Seat) (*model.Seat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSeat", arg0, arg1)
	ret0, _ := ret[0].(*model.Seat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSeat indicates an expected call of CreateSeat.
func (mr *MockDalMockRecorder) CreateSeat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSeat", reflect.TypeOf((*MockDal)(nil).CreateSeat), arg0, arg1)
}

// CreateSeats mocks base method.
func (m *MockDal) CreateSeats(arg0 *gorm.DB, arg1 []*model.Seat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSeats", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSeats indicates an expected call of CreateSeats.
func (mr *MockDalMockRecorder) CreateSeats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSeats", reflect.TypeOf((*MockDal)(nil).CreateSeats), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockDal) CreateUser(arg0 *gorm.DB, arg1 *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDalMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDal)(nil).CreateUser), arg0, arg1)
}

// DeleteRoom mocks base method.
func (m *MockDal) DeleteRoom(arg0 *gorm.DB, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockDalMockRecorder) DeleteRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockDal)(nil).DeleteRoom), arg0, arg1)
}

// DeleteSeat mocks base method.
func (m *MockDal) DeleteSeat(arg0 *gorm.DB, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSeat", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSeat indicates an expected call of DeleteSeat.
func (mr *MockDalMockRecorder) DeleteSeat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSeat", reflect.TypeOf((*MockDal)(nil).DeleteSeat), arg0, arg1)
}

// DeleteSeatsOfRoom mocks base method.
func (m *MockDal) DeleteSeatsOfRoom(arg0 *gorm.DB, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSeatsOfRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSeatsOfRoom indicates an expected call of DeleteSeatsOfRoom.
func (mr *MockDalMockRecorder) DeleteSeatsOfRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSeatsOfRoom", reflect.TypeOf((*MockDal)(nil).DeleteSeatsOfRoom), arg0, arg1)
}

// GetAllRooms mocks base method.
func (m *MockDal) GetAllRooms(arg0 *gorm.DB) ([]*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRooms", arg0)
	ret0, _ := ret[0].([]*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRooms indicates an expected call of GetAllRooms.
func (mr *MockDalMockRecorder) GetAllRooms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRooms", reflect.TypeOf((*MockDal)(nil).GetAllRooms), arg0)
}

// GetAllSeatsOfRoom mocks base method.
func (m *MockDal) GetAllSeatsOfRoom(arg0 *gorm.DB, arg1 int64) ([]*model.Seat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSeatsOfRoom", arg0, arg1)
	ret0, _ := ret[0].([]*model.Seat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSeatsOfRoom indicates an expected call of GetAllSeatsOfRoom.
func (mr *MockDalMockRecorder) GetAllSeatsOfRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSeatsOfRoom", reflect.TypeOf((*MockDal)(nil).GetAllSeatsOfRoom), arg0, arg1)
}

// GetAvailableRooms mocks base method.
func (m *MockDal) GetAvailableRooms(arg0 *gorm.DB) ([]*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableRooms", arg0)
	ret0, _ := ret[0].([]*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableRooms indicates an expected call of GetAvailableRooms.
func (mr *MockDalMockRecorder) GetAvailableRooms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableRooms", reflect.TypeOf((*MockDal)(nil).GetAvailableRooms), arg0)
}

// GetResvByID mocks base method.
func (m *MockDal) GetResvByID(arg0 *gorm.DB, arg1 int64) (*model.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResvByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResvByID indicates an expected call of GetResvByID.
func (mr *MockDalMockRecorder) GetResvByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResvByID", reflect.TypeOf((*MockDal)(nil).GetResvByID), arg0, arg1)
}

// GetResvsBySeat mocks base method.
func (m *MockDal) GetResvsBySeat(arg0 *gorm.DB, arg1 int64) ([]*model.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResvsBySeat", arg0, arg1)
	ret0, _ := ret[0].([]*model.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResvsBySeat indicates an expected call of GetResvsBySeat.
func (mr *MockDalMockRecorder) GetResvsBySeat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResvsBySeat", reflect.TypeOf((*MockDal)(nil).GetResvsBySeat), arg0, arg1)
}

// GetResvsByUser mocks base method.
func (m *MockDal) GetResvsByUser(arg0 *gorm.DB, arg1 string) ([]*model.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResvsByUser", arg0, arg1)
	ret0, _ := ret[0].([]*model.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResvsByUser indicates an expected call of GetResvsByUser.
func (mr *MockDalMockRecorder) GetResvsByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResvsByUser", reflect.TypeOf((*MockDal)(nil).GetResvsByUser), arg0, arg1)
}

// GetRoomByID mocks base method.
func (m *MockDal) GetRoomByID(arg0 *gorm.DB, arg1 int64) (*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomByID indicates an expected call of GetRoomByID.
func (mr *MockDalMockRecorder) GetRoomByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomByID", reflect.TypeOf((*MockDal)(nil).GetRoomByID), arg0, arg1)
}

// GetSeatByID mocks base method.
func (m *MockDal) GetSeatByID(arg0 *gorm.DB, arg1 int64) (*model.Seat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeatByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Seat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeatByID indicates an expected call of GetSeatByID.
func (mr *MockDalMockRecorder) GetSeatByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeatByID", reflect.TypeOf((*MockDal)(nil).GetSeatByID), arg0, arg1)
}

// GetUserByName mocks base method.
func (m *MockDal) GetUserByName(arg0 *gorm.DB, arg1 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockDalMockRecorder) GetUserByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockDal)(nil).GetUserByName), arg0, arg1)
}

// UpdateResv mocks base method.
func (m *MockDal) UpdateResv(arg0 *gorm.DB, arg1 *model.Reservation) (*model.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResv", arg0, arg1)
	ret0, _ := ret[0].(*model.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResv indicates an expected call of UpdateResv.
func (mr *MockDalMockRecorder) UpdateResv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResv", reflect.TypeOf((*MockDal)(nil).UpdateResv), arg0, arg1)
}

// UpdateRoom mocks base method.
func (m *MockDal) UpdateRoom(arg0 *gorm.DB, arg1 *model.Room) (*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", arg0, arg1)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockDalMockRecorder) UpdateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockDal)(nil).UpdateRoom), arg0, arg1)
}

// UpdateSeat mocks base method.
func (m *MockDal) UpdateSeat(arg0 *gorm.DB, arg1 *model.Seat) (*model.Seat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSeat", arg0, arg1)
	ret0, _ := ret[0].(*model.Seat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSeat indicates an expected call of UpdateSeat.
func (mr *MockDalMockRecorder) UpdateSeat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSeat", reflect.TypeOf((*MockDal)(nil).UpdateSeat), arg0, arg1)
}
