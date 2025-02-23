// Code generated by MockGen. DO NOT EDIT.
// Source: types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/capability/types"
	types1 "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	types2 "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	exported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	gomock "github.com/golang/mock/gomock"
)

// MockRelationshipsKeeper is a mock of RelationshipsKeeper interface.
type MockRelationshipsKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockRelationshipsKeeperMockRecorder
}

// MockRelationshipsKeeperMockRecorder is the mock recorder for MockRelationshipsKeeper.
type MockRelationshipsKeeperMockRecorder struct {
	mock *MockRelationshipsKeeper
}

// NewMockRelationshipsKeeper creates a new mock instance.
func NewMockRelationshipsKeeper(ctrl *gomock.Controller) *MockRelationshipsKeeper {
	mock := &MockRelationshipsKeeper{ctrl: ctrl}
	mock.recorder = &MockRelationshipsKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelationshipsKeeper) EXPECT() *MockRelationshipsKeeperMockRecorder {
	return m.recorder
}

// HasUserBlocked mocks base method.
func (m *MockRelationshipsKeeper) HasUserBlocked(ctx types.Context, user, blocker string, subspaceID uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUserBlocked", ctx, user, blocker, subspaceID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUserBlocked indicates an expected call of HasUserBlocked.
func (mr *MockRelationshipsKeeperMockRecorder) HasUserBlocked(ctx, user, blocker, subspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUserBlocked", reflect.TypeOf((*MockRelationshipsKeeper)(nil).HasUserBlocked), ctx, user, blocker, subspaceID)
}

// MockChannelKeeper is a mock of ChannelKeeper interface.
type MockChannelKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockChannelKeeperMockRecorder
}

// MockChannelKeeperMockRecorder is the mock recorder for MockChannelKeeper.
type MockChannelKeeperMockRecorder struct {
	mock *MockChannelKeeper
}

// NewMockChannelKeeper creates a new mock instance.
func NewMockChannelKeeper(ctrl *gomock.Controller) *MockChannelKeeper {
	mock := &MockChannelKeeper{ctrl: ctrl}
	mock.recorder = &MockChannelKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelKeeper) EXPECT() *MockChannelKeeperMockRecorder {
	return m.recorder
}

// ChanCloseInit mocks base method.
func (m *MockChannelKeeper) ChanCloseInit(ctx types.Context, portID, channelID string, chanCap *types0.Capability) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChanCloseInit", ctx, portID, channelID, chanCap)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChanCloseInit indicates an expected call of ChanCloseInit.
func (mr *MockChannelKeeperMockRecorder) ChanCloseInit(ctx, portID, channelID, chanCap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChanCloseInit", reflect.TypeOf((*MockChannelKeeper)(nil).ChanCloseInit), ctx, portID, channelID, chanCap)
}

// GetChannel mocks base method.
func (m *MockChannelKeeper) GetChannel(ctx types.Context, srcPort, srcChan string) (types2.Channel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", ctx, srcPort, srcChan)
	ret0, _ := ret[0].(types2.Channel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockChannelKeeperMockRecorder) GetChannel(ctx, srcPort, srcChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannel), ctx, srcPort, srcChan)
}

// GetNextSequenceSend mocks base method.
func (m *MockChannelKeeper) GetNextSequenceSend(ctx types.Context, portID, channelID string) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSequenceSend", ctx, portID, channelID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNextSequenceSend indicates an expected call of GetNextSequenceSend.
func (mr *MockChannelKeeperMockRecorder) GetNextSequenceSend(ctx, portID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSequenceSend", reflect.TypeOf((*MockChannelKeeper)(nil).GetNextSequenceSend), ctx, portID, channelID)
}

// SendPacket mocks base method.
func (m *MockChannelKeeper) SendPacket(ctx types.Context, channelCap *types0.Capability, packet exported.PacketI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPacket", ctx, channelCap, packet)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPacket indicates an expected call of SendPacket.
func (mr *MockChannelKeeperMockRecorder) SendPacket(ctx, channelCap, packet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPacket", reflect.TypeOf((*MockChannelKeeper)(nil).SendPacket), ctx, channelCap, packet)
}

// MockClientKeeper is a mock of ClientKeeper interface.
type MockClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeeperMockRecorder
}

// MockClientKeeperMockRecorder is the mock recorder for MockClientKeeper.
type MockClientKeeperMockRecorder struct {
	mock *MockClientKeeper
}

// NewMockClientKeeper creates a new mock instance.
func NewMockClientKeeper(ctrl *gomock.Controller) *MockClientKeeper {
	mock := &MockClientKeeper{ctrl: ctrl}
	mock.recorder = &MockClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientKeeper) EXPECT() *MockClientKeeperMockRecorder {
	return m.recorder
}

// GetClientConsensusState mocks base method.
func (m *MockClientKeeper) GetClientConsensusState(ctx types.Context, clientID string) (exported.ConsensusState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientConsensusState", ctx, clientID)
	ret0, _ := ret[0].(exported.ConsensusState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetClientConsensusState indicates an expected call of GetClientConsensusState.
func (mr *MockClientKeeperMockRecorder) GetClientConsensusState(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientConsensusState", reflect.TypeOf((*MockClientKeeper)(nil).GetClientConsensusState), ctx, clientID)
}

// MockConnectionKeeper is a mock of ConnectionKeeper interface.
type MockConnectionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionKeeperMockRecorder
}

// MockConnectionKeeperMockRecorder is the mock recorder for MockConnectionKeeper.
type MockConnectionKeeperMockRecorder struct {
	mock *MockConnectionKeeper
}

// NewMockConnectionKeeper creates a new mock instance.
func NewMockConnectionKeeper(ctrl *gomock.Controller) *MockConnectionKeeper {
	mock := &MockConnectionKeeper{ctrl: ctrl}
	mock.recorder = &MockConnectionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionKeeper) EXPECT() *MockConnectionKeeperMockRecorder {
	return m.recorder
}

// GetConnection mocks base method.
func (m *MockConnectionKeeper) GetConnection(ctx types.Context, connectionID string) (types1.ConnectionEnd, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", ctx, connectionID)
	ret0, _ := ret[0].(types1.ConnectionEnd)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection.
func (mr *MockConnectionKeeperMockRecorder) GetConnection(ctx, connectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockConnectionKeeper)(nil).GetConnection), ctx, connectionID)
}

// MockPortKeeper is a mock of PortKeeper interface.
type MockPortKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPortKeeperMockRecorder
}

// MockPortKeeperMockRecorder is the mock recorder for MockPortKeeper.
type MockPortKeeperMockRecorder struct {
	mock *MockPortKeeper
}

// NewMockPortKeeper creates a new mock instance.
func NewMockPortKeeper(ctrl *gomock.Controller) *MockPortKeeper {
	mock := &MockPortKeeper{ctrl: ctrl}
	mock.recorder = &MockPortKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortKeeper) EXPECT() *MockPortKeeperMockRecorder {
	return m.recorder
}

// BindPort mocks base method.
func (m *MockPortKeeper) BindPort(ctx types.Context, portID string) *types0.Capability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindPort", ctx, portID)
	ret0, _ := ret[0].(*types0.Capability)
	return ret0
}

// BindPort indicates an expected call of BindPort.
func (mr *MockPortKeeperMockRecorder) BindPort(ctx, portID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindPort", reflect.TypeOf((*MockPortKeeper)(nil).BindPort), ctx, portID)
}

// MockScopedKeeper is a mock of ScopedKeeper interface.
type MockScopedKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockScopedKeeperMockRecorder
}

// MockScopedKeeperMockRecorder is the mock recorder for MockScopedKeeper.
type MockScopedKeeperMockRecorder struct {
	mock *MockScopedKeeper
}

// NewMockScopedKeeper creates a new mock instance.
func NewMockScopedKeeper(ctrl *gomock.Controller) *MockScopedKeeper {
	mock := &MockScopedKeeper{ctrl: ctrl}
	mock.recorder = &MockScopedKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScopedKeeper) EXPECT() *MockScopedKeeperMockRecorder {
	return m.recorder
}

// AuthenticateCapability mocks base method.
func (m *MockScopedKeeper) AuthenticateCapability(ctx types.Context, cap *types0.Capability, name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateCapability", ctx, cap, name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthenticateCapability indicates an expected call of AuthenticateCapability.
func (mr *MockScopedKeeperMockRecorder) AuthenticateCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateCapability", reflect.TypeOf((*MockScopedKeeper)(nil).AuthenticateCapability), ctx, cap, name)
}

// ClaimCapability mocks base method.
func (m *MockScopedKeeper) ClaimCapability(ctx types.Context, cap *types0.Capability, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimCapability", ctx, cap, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimCapability indicates an expected call of ClaimCapability.
func (mr *MockScopedKeeperMockRecorder) ClaimCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimCapability", reflect.TypeOf((*MockScopedKeeper)(nil).ClaimCapability), ctx, cap, name)
}

// GetCapability mocks base method.
func (m *MockScopedKeeper) GetCapability(ctx types.Context, name string) (*types0.Capability, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapability", ctx, name)
	ret0, _ := ret[0].(*types0.Capability)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetCapability indicates an expected call of GetCapability.
func (mr *MockScopedKeeperMockRecorder) GetCapability(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapability", reflect.TypeOf((*MockScopedKeeper)(nil).GetCapability), ctx, name)
}
