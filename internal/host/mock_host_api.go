// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/host (interfaces: API)

// Package host is a generated GoMock package.
package host

import (
	context "context"
	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	logrus "github.com/sirupsen/logrus"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// AutoAssignRole mocks base method
func (m *MockAPI) AutoAssignRole(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoAssignRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoAssignRole indicates an expected call of AutoAssignRole
func (mr *MockAPIMockRecorder) AutoAssignRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoAssignRole", reflect.TypeOf((*MockAPI)(nil).AutoAssignRole), arg0, arg1, arg2)
}

// CancelInstallation mocks base method
func (m *MockAPI) CancelInstallation(arg0 context.Context, arg1 *models.Host, arg2 string, arg3 *gorm.DB) *common.ApiErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelInstallation", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*common.ApiErrorResponse)
	return ret0
}

// CancelInstallation indicates an expected call of CancelInstallation
func (mr *MockAPIMockRecorder) CancelInstallation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelInstallation", reflect.TypeOf((*MockAPI)(nil).CancelInstallation), arg0, arg1, arg2, arg3)
}

// DisableHost mocks base method
func (m *MockAPI) DisableHost(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableHost", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableHost indicates an expected call of DisableHost
func (mr *MockAPIMockRecorder) DisableHost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableHost", reflect.TypeOf((*MockAPI)(nil).DisableHost), arg0, arg1, arg2)
}

// EnableHost mocks base method
func (m *MockAPI) EnableHost(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableHost", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableHost indicates an expected call of EnableHost
func (mr *MockAPIMockRecorder) EnableHost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHost", reflect.TypeOf((*MockAPI)(nil).EnableHost), arg0, arg1, arg2)
}

// GetHostByKubeKey mocks base method
func (m *MockAPI) GetHostByKubeKey(arg0 types.NamespacedName) (*common.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostByKubeKey", arg0)
	ret0, _ := ret[0].(*common.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostByKubeKey indicates an expected call of GetHostByKubeKey
func (mr *MockAPIMockRecorder) GetHostByKubeKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostByKubeKey", reflect.TypeOf((*MockAPI)(nil).GetHostByKubeKey), arg0)
}

// GetHostValidDisks mocks base method
func (m *MockAPI) GetHostValidDisks(arg0 *models.Host) ([]*models.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostValidDisks", arg0)
	ret0, _ := ret[0].([]*models.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostValidDisks indicates an expected call of GetHostValidDisks
func (mr *MockAPIMockRecorder) GetHostValidDisks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostValidDisks", reflect.TypeOf((*MockAPI)(nil).GetHostValidDisks), arg0)
}

// GetNextSteps mocks base method
func (m *MockAPI) GetNextSteps(arg0 context.Context, arg1 *models.Host) (models.Steps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSteps", arg0, arg1)
	ret0, _ := ret[0].(models.Steps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextSteps indicates an expected call of GetNextSteps
func (mr *MockAPIMockRecorder) GetNextSteps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSteps", reflect.TypeOf((*MockAPI)(nil).GetNextSteps), arg0, arg1)
}

// GetStagesByRole mocks base method
func (m *MockAPI) GetStagesByRole(arg0 models.HostRole, arg1 bool) []models.HostStage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStagesByRole", arg0, arg1)
	ret0, _ := ret[0].([]models.HostStage)
	return ret0
}

// GetStagesByRole indicates an expected call of GetStagesByRole
func (mr *MockAPIMockRecorder) GetStagesByRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStagesByRole", reflect.TypeOf((*MockAPI)(nil).GetStagesByRole), arg0, arg1)
}

// HandleInstallationFailure mocks base method
func (m *MockAPI) HandleInstallationFailure(arg0 context.Context, arg1 *models.Host) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInstallationFailure", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleInstallationFailure indicates an expected call of HandleInstallationFailure
func (mr *MockAPIMockRecorder) HandleInstallationFailure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInstallationFailure", reflect.TypeOf((*MockAPI)(nil).HandleInstallationFailure), arg0, arg1)
}

// HostMonitoring mocks base method
func (m *MockAPI) HostMonitoring() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HostMonitoring")
}

// HostMonitoring indicates an expected call of HostMonitoring
func (mr *MockAPIMockRecorder) HostMonitoring() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostMonitoring", reflect.TypeOf((*MockAPI)(nil).HostMonitoring))
}

// IndexOfStage mocks base method
func (m *MockAPI) IndexOfStage(arg0 models.HostStage, arg1 []models.HostStage) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexOfStage", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// IndexOfStage indicates an expected call of IndexOfStage
func (mr *MockAPIMockRecorder) IndexOfStage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexOfStage", reflect.TypeOf((*MockAPI)(nil).IndexOfStage), arg0, arg1)
}

// Install mocks base method
func (m *MockAPI) Install(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockAPIMockRecorder) Install(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockAPI)(nil).Install), arg0, arg1, arg2)
}

// IsInstallable mocks base method
func (m *MockAPI) IsInstallable(arg0 *models.Host) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstallable", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInstallable indicates an expected call of IsInstallable
func (mr *MockAPIMockRecorder) IsInstallable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstallable", reflect.TypeOf((*MockAPI)(nil).IsInstallable), arg0)
}

// IsRequireUserActionReset mocks base method
func (m *MockAPI) IsRequireUserActionReset(arg0 *models.Host) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRequireUserActionReset", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRequireUserActionReset indicates an expected call of IsRequireUserActionReset
func (mr *MockAPIMockRecorder) IsRequireUserActionReset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRequireUserActionReset", reflect.TypeOf((*MockAPI)(nil).IsRequireUserActionReset), arg0)
}

// IsValidMasterCandidate mocks base method
func (m *MockAPI) IsValidMasterCandidate(arg0 *models.Host, arg1 *common.Cluster, arg2 *gorm.DB, arg3 logrus.FieldLogger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidMasterCandidate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidMasterCandidate indicates an expected call of IsValidMasterCandidate
func (mr *MockAPIMockRecorder) IsValidMasterCandidate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidMasterCandidate", reflect.TypeOf((*MockAPI)(nil).IsValidMasterCandidate), arg0, arg1, arg2, arg3)
}

// PermanentHostsDeletion mocks base method
func (m *MockAPI) PermanentHostsDeletion(arg0 strfmt.DateTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PermanentHostsDeletion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PermanentHostsDeletion indicates an expected call of PermanentHostsDeletion
func (mr *MockAPIMockRecorder) PermanentHostsDeletion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PermanentHostsDeletion", reflect.TypeOf((*MockAPI)(nil).PermanentHostsDeletion), arg0)
}

// RefreshInventory mocks base method
func (m *MockAPI) RefreshInventory(arg0 context.Context, arg1 *common.Cluster, arg2 *models.Host, arg3 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshInventory", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshInventory indicates an expected call of RefreshInventory
func (mr *MockAPIMockRecorder) RefreshInventory(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshInventory", reflect.TypeOf((*MockAPI)(nil).RefreshInventory), arg0, arg1, arg2, arg3)
}

// RefreshStatus mocks base method
func (m *MockAPI) RefreshStatus(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshStatus indicates an expected call of RefreshStatus
func (mr *MockAPIMockRecorder) RefreshStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshStatus", reflect.TypeOf((*MockAPI)(nil).RefreshStatus), arg0, arg1, arg2)
}

// RegisterHost mocks base method
func (m *MockAPI) RegisterHost(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterHost", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterHost indicates an expected call of RegisterHost
func (mr *MockAPIMockRecorder) RegisterHost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHost", reflect.TypeOf((*MockAPI)(nil).RegisterHost), arg0, arg1, arg2)
}

// RegisterInstalledOCPHost mocks base method
func (m *MockAPI) RegisterInstalledOCPHost(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInstalledOCPHost", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterInstalledOCPHost indicates an expected call of RegisterInstalledOCPHost
func (mr *MockAPIMockRecorder) RegisterInstalledOCPHost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInstalledOCPHost", reflect.TypeOf((*MockAPI)(nil).RegisterInstalledOCPHost), arg0, arg1, arg2)
}

// ReportValidationFailedMetrics mocks base method
func (m *MockAPI) ReportValidationFailedMetrics(arg0 context.Context, arg1 *models.Host, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportValidationFailedMetrics", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportValidationFailedMetrics indicates an expected call of ReportValidationFailedMetrics
func (mr *MockAPIMockRecorder) ReportValidationFailedMetrics(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportValidationFailedMetrics", reflect.TypeOf((*MockAPI)(nil).ReportValidationFailedMetrics), arg0, arg1, arg2, arg3)
}

// ResetHost mocks base method
func (m *MockAPI) ResetHost(arg0 context.Context, arg1 *models.Host, arg2 string, arg3 *gorm.DB) *common.ApiErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetHost", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*common.ApiErrorResponse)
	return ret0
}

// ResetHost indicates an expected call of ResetHost
func (mr *MockAPIMockRecorder) ResetHost(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetHost", reflect.TypeOf((*MockAPI)(nil).ResetHost), arg0, arg1, arg2, arg3)
}

// ResetHostValidation mocks base method
func (m *MockAPI) ResetHostValidation(arg0 context.Context, arg1, arg2 strfmt.UUID, arg3 string, arg4 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetHostValidation", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetHostValidation indicates an expected call of ResetHostValidation
func (mr *MockAPIMockRecorder) ResetHostValidation(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetHostValidation", reflect.TypeOf((*MockAPI)(nil).ResetHostValidation), arg0, arg1, arg2, arg3, arg4)
}

// ResetPendingUserAction mocks base method
func (m *MockAPI) ResetPendingUserAction(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPendingUserAction", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPendingUserAction indicates an expected call of ResetPendingUserAction
func (mr *MockAPIMockRecorder) ResetPendingUserAction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPendingUserAction", reflect.TypeOf((*MockAPI)(nil).ResetPendingUserAction), arg0, arg1, arg2)
}

// SetBootstrap mocks base method
func (m *MockAPI) SetBootstrap(arg0 context.Context, arg1 *models.Host, arg2 bool, arg3 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBootstrap", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBootstrap indicates an expected call of SetBootstrap
func (mr *MockAPIMockRecorder) SetBootstrap(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootstrap", reflect.TypeOf((*MockAPI)(nil).SetBootstrap), arg0, arg1, arg2, arg3)
}

// SetDiskSpeed mocks base method
func (m *MockAPI) SetDiskSpeed(arg0 context.Context, arg1 *models.Host, arg2 string, arg3, arg4 int64, arg5 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDiskSpeed", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDiskSpeed indicates an expected call of SetDiskSpeed
func (mr *MockAPIMockRecorder) SetDiskSpeed(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDiskSpeed", reflect.TypeOf((*MockAPI)(nil).SetDiskSpeed), arg0, arg1, arg2, arg3, arg4, arg5)
}

// SetUploadLogsAt mocks base method
func (m *MockAPI) SetUploadLogsAt(arg0 context.Context, arg1 *models.Host, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUploadLogsAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUploadLogsAt indicates an expected call of SetUploadLogsAt
func (mr *MockAPIMockRecorder) SetUploadLogsAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUploadLogsAt", reflect.TypeOf((*MockAPI)(nil).SetUploadLogsAt), arg0, arg1, arg2)
}

// UnRegisterHost mocks base method
func (m *MockAPI) UnRegisterHost(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnRegisterHost", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnRegisterHost indicates an expected call of UnRegisterHost
func (mr *MockAPIMockRecorder) UnRegisterHost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRegisterHost", reflect.TypeOf((*MockAPI)(nil).UnRegisterHost), arg0, arg1, arg2)
}

// UpdateApiVipConnectivityReport mocks base method
func (m *MockAPI) UpdateApiVipConnectivityReport(arg0 context.Context, arg1 *models.Host, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApiVipConnectivityReport", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApiVipConnectivityReport indicates an expected call of UpdateApiVipConnectivityReport
func (mr *MockAPIMockRecorder) UpdateApiVipConnectivityReport(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApiVipConnectivityReport", reflect.TypeOf((*MockAPI)(nil).UpdateApiVipConnectivityReport), arg0, arg1, arg2)
}

// UpdateConnectivityReport mocks base method
func (m *MockAPI) UpdateConnectivityReport(arg0 context.Context, arg1 *models.Host, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConnectivityReport", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConnectivityReport indicates an expected call of UpdateConnectivityReport
func (mr *MockAPIMockRecorder) UpdateConnectivityReport(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnectivityReport", reflect.TypeOf((*MockAPI)(nil).UpdateConnectivityReport), arg0, arg1, arg2)
}

// UpdateHostname mocks base method
func (m *MockAPI) UpdateHostname(arg0 context.Context, arg1 *models.Host, arg2 string, arg3 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostname", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHostname indicates an expected call of UpdateHostname
func (mr *MockAPIMockRecorder) UpdateHostname(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostname", reflect.TypeOf((*MockAPI)(nil).UpdateHostname), arg0, arg1, arg2, arg3)
}

// UpdateImageStatus mocks base method
func (m *MockAPI) UpdateImageStatus(arg0 context.Context, arg1 *models.Host, arg2 *models.ContainerImageAvailability, arg3 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImageStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImageStatus indicates an expected call of UpdateImageStatus
func (mr *MockAPIMockRecorder) UpdateImageStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImageStatus", reflect.TypeOf((*MockAPI)(nil).UpdateImageStatus), arg0, arg1, arg2, arg3)
}

// UpdateInstallProgress mocks base method
func (m *MockAPI) UpdateInstallProgress(arg0 context.Context, arg1 *models.Host, arg2 *models.HostProgress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstallProgress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstallProgress indicates an expected call of UpdateInstallProgress
func (mr *MockAPIMockRecorder) UpdateInstallProgress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstallProgress", reflect.TypeOf((*MockAPI)(nil).UpdateInstallProgress), arg0, arg1, arg2)
}

// UpdateInstallationDisk mocks base method
func (m *MockAPI) UpdateInstallationDisk(arg0 context.Context, arg1 *gorm.DB, arg2 *models.Host, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstallationDisk", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstallationDisk indicates an expected call of UpdateInstallationDisk
func (mr *MockAPIMockRecorder) UpdateInstallationDisk(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstallationDisk", reflect.TypeOf((*MockAPI)(nil).UpdateInstallationDisk), arg0, arg1, arg2, arg3)
}

// UpdateInventory mocks base method
func (m *MockAPI) UpdateInventory(arg0 context.Context, arg1 *models.Host, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInventory indicates an expected call of UpdateInventory
func (mr *MockAPIMockRecorder) UpdateInventory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockAPI)(nil).UpdateInventory), arg0, arg1, arg2)
}

// UpdateKubeKeyNS mocks base method
func (m *MockAPI) UpdateKubeKeyNS(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKubeKeyNS", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateKubeKeyNS indicates an expected call of UpdateKubeKeyNS
func (mr *MockAPIMockRecorder) UpdateKubeKeyNS(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKubeKeyNS", reflect.TypeOf((*MockAPI)(nil).UpdateKubeKeyNS), arg0, arg1, arg2)
}

// UpdateLogsProgress mocks base method
func (m *MockAPI) UpdateLogsProgress(arg0 context.Context, arg1 *models.Host, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLogsProgress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLogsProgress indicates an expected call of UpdateLogsProgress
func (mr *MockAPIMockRecorder) UpdateLogsProgress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLogsProgress", reflect.TypeOf((*MockAPI)(nil).UpdateLogsProgress), arg0, arg1, arg2)
}

// UpdateMachineConfigPoolName mocks base method
func (m *MockAPI) UpdateMachineConfigPoolName(arg0 context.Context, arg1 *gorm.DB, arg2 *models.Host, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMachineConfigPoolName", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMachineConfigPoolName indicates an expected call of UpdateMachineConfigPoolName
func (mr *MockAPIMockRecorder) UpdateMachineConfigPoolName(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMachineConfigPoolName", reflect.TypeOf((*MockAPI)(nil).UpdateMachineConfigPoolName), arg0, arg1, arg2, arg3)
}

// UpdateNTP mocks base method
func (m *MockAPI) UpdateNTP(arg0 context.Context, arg1 *models.Host, arg2 []*models.NtpSource, arg3 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNTP", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNTP indicates an expected call of UpdateNTP
func (mr *MockAPIMockRecorder) UpdateNTP(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNTP", reflect.TypeOf((*MockAPI)(nil).UpdateNTP), arg0, arg1, arg2, arg3)
}

// UpdateRole mocks base method
func (m *MockAPI) UpdateRole(arg0 context.Context, arg1 *models.Host, arg2 models.HostRole, arg3 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRole indicates an expected call of UpdateRole
func (mr *MockAPIMockRecorder) UpdateRole(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockAPI)(nil).UpdateRole), arg0, arg1, arg2, arg3)
}
