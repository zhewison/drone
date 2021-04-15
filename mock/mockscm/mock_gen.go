// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/drone/go-scm/scm (interfaces: ContentService,GitService,OrganizationService,PullRequestService,RepositoryService,UserService)

// Package mockscm is a generated GoMock package.
package mockscm

import (
	context "context"
	scm "github.com/drone/go-scm/scm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContentService is a mock of ContentService interface
type MockContentService struct {
	ctrl     *gomock.Controller
	recorder *MockContentServiceMockRecorder
}

// MockContentServiceMockRecorder is the mock recorder for MockContentService
type MockContentServiceMockRecorder struct {
	mock *MockContentService
}

// NewMockContentService creates a new mock instance
func NewMockContentService(ctrl *gomock.Controller) *MockContentService {
	mock := &MockContentService{ctrl: ctrl}
	mock.recorder = &MockContentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContentService) EXPECT() *MockContentServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockContentService) Create(arg0 context.Context, arg1, arg2 string, arg3 *scm.ContentParams) (*scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockContentServiceMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContentService)(nil).Create), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockContentService) Delete(arg0 context.Context, arg1, arg2, arg3 string) (*scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockContentServiceMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockContentService)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Find mocks base method
func (m *MockContentService) Find(arg0 context.Context, arg1, arg2, arg3 string) (*scm.Content, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Content)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockContentServiceMockRecorder) Find(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockContentService)(nil).Find), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockContentService) List(arg0 context.Context, arg1, arg2, arg3 string, arg4 scm.ListOptions) ([]*scm.ContentInfo, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*scm.ContentInfo)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockContentServiceMockRecorder) List(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContentService)(nil).List), arg0, arg1, arg2, arg3, arg4)
}

// Update mocks base method
func (m *MockContentService) Update(arg0 context.Context, arg1, arg2 string, arg3 *scm.ContentParams) (*scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockContentServiceMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockContentService)(nil).Update), arg0, arg1, arg2, arg3)
}

// MockGitService is a mock of GitService interface
type MockGitService struct {
	ctrl     *gomock.Controller
	recorder *MockGitServiceMockRecorder
}

// MockGitServiceMockRecorder is the mock recorder for MockGitService
type MockGitServiceMockRecorder struct {
	mock *MockGitService
}

// NewMockGitService creates a new mock instance
func NewMockGitService(ctrl *gomock.Controller) *MockGitService {
	mock := &MockGitService{ctrl: ctrl}
	mock.recorder = &MockGitServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitService) EXPECT() *MockGitServiceMockRecorder {
	return m.recorder
}

// CompareChanges mocks base method
func (m *MockGitService) CompareChanges(arg0 context.Context, arg1, arg2, arg3 string, arg4 scm.ListOptions) ([]*scm.Change, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareChanges", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*scm.Change)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CompareChanges indicates an expected call of CompareChanges
func (mr *MockGitServiceMockRecorder) CompareChanges(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareChanges", reflect.TypeOf((*MockGitService)(nil).CompareChanges), arg0, arg1, arg2, arg3, arg4)
}

// FindBranch mocks base method
func (m *MockGitService) FindBranch(arg0 context.Context, arg1, arg2 string) (*scm.Reference, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBranch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Reference)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindBranch indicates an expected call of FindBranch
func (mr *MockGitServiceMockRecorder) FindBranch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBranch", reflect.TypeOf((*MockGitService)(nil).FindBranch), arg0, arg1, arg2)
}

// FindCommit mocks base method
func (m *MockGitService) FindCommit(arg0 context.Context, arg1, arg2 string) (*scm.Commit, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCommit", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Commit)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindCommit indicates an expected call of FindCommit
func (mr *MockGitServiceMockRecorder) FindCommit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCommit", reflect.TypeOf((*MockGitService)(nil).FindCommit), arg0, arg1, arg2)
}

// FindTag mocks base method
func (m *MockGitService) FindTag(arg0 context.Context, arg1, arg2 string) (*scm.Reference, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTag", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Reference)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindTag indicates an expected call of FindTag
func (mr *MockGitServiceMockRecorder) FindTag(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTag", reflect.TypeOf((*MockGitService)(nil).FindTag), arg0, arg1, arg2)
}

// ListBranches mocks base method
func (m *MockGitService) ListBranches(arg0 context.Context, arg1 string, arg2 scm.ListOptions) ([]*scm.Reference, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranches", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*scm.Reference)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBranches indicates an expected call of ListBranches
func (mr *MockGitServiceMockRecorder) ListBranches(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockGitService)(nil).ListBranches), arg0, arg1, arg2)
}

// ListChanges mocks base method
func (m *MockGitService) ListChanges(arg0 context.Context, arg1, arg2 string, arg3 scm.ListOptions) ([]*scm.Change, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChanges", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*scm.Change)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListChanges indicates an expected call of ListChanges
func (mr *MockGitServiceMockRecorder) ListChanges(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChanges", reflect.TypeOf((*MockGitService)(nil).ListChanges), arg0, arg1, arg2, arg3)
}

// ListCommits mocks base method
func (m *MockGitService) ListCommits(arg0 context.Context, arg1 string, arg2 scm.CommitListOptions) ([]*scm.Commit, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*scm.Commit)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCommits indicates an expected call of ListCommits
func (mr *MockGitServiceMockRecorder) ListCommits(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockGitService)(nil).ListCommits), arg0, arg1, arg2)
}

// ListTags mocks base method
func (m *MockGitService) ListTags(arg0 context.Context, arg1 string, arg2 scm.ListOptions) ([]*scm.Reference, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*scm.Reference)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTags indicates an expected call of ListTags
func (mr *MockGitServiceMockRecorder) ListTags(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockGitService)(nil).ListTags), arg0, arg1, arg2)
}

// MockOrganizationService is a mock of OrganizationService interface
type MockOrganizationService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationServiceMockRecorder
}

// MockOrganizationServiceMockRecorder is the mock recorder for MockOrganizationService
type MockOrganizationServiceMockRecorder struct {
	mock *MockOrganizationService
}

// NewMockOrganizationService creates a new mock instance
func NewMockOrganizationService(ctrl *gomock.Controller) *MockOrganizationService {
	mock := &MockOrganizationService{ctrl: ctrl}
	mock.recorder = &MockOrganizationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationService) EXPECT() *MockOrganizationServiceMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockOrganizationService) Find(arg0 context.Context, arg1 string) (*scm.Organization, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*scm.Organization)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockOrganizationServiceMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockOrganizationService)(nil).Find), arg0, arg1)
}

// FindMembership mocks base method
func (m *MockOrganizationService) FindMembership(arg0 context.Context, arg1, arg2 string) (*scm.Membership, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMembership", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Membership)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindMembership indicates an expected call of FindMembership
func (mr *MockOrganizationServiceMockRecorder) FindMembership(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMembership", reflect.TypeOf((*MockOrganizationService)(nil).FindMembership), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockOrganizationService) List(arg0 context.Context, arg1 scm.ListOptions) ([]*scm.Organization, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*scm.Organization)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockOrganizationServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrganizationService)(nil).List), arg0, arg1)
}

// MockPullRequestService is a mock of PullRequestService interface
type MockPullRequestService struct {
	ctrl     *gomock.Controller
	recorder *MockPullRequestServiceMockRecorder
}

// MockPullRequestServiceMockRecorder is the mock recorder for MockPullRequestService
type MockPullRequestServiceMockRecorder struct {
	mock *MockPullRequestService
}

// NewMockPullRequestService creates a new mock instance
func NewMockPullRequestService(ctrl *gomock.Controller) *MockPullRequestService {
	mock := &MockPullRequestService{ctrl: ctrl}
	mock.recorder = &MockPullRequestServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPullRequestService) EXPECT() *MockPullRequestServiceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPullRequestService) Close(arg0 context.Context, arg1 string, arg2 int) (*scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close
func (mr *MockPullRequestServiceMockRecorder) Close(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPullRequestService)(nil).Close), arg0, arg1, arg2)
}

// Create mocks base method
func (m *MockPullRequestService) Create(arg0 context.Context, arg1 string, arg2 *scm.PullRequestInput) (*scm.PullRequest, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.PullRequest)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create
func (mr *MockPullRequestServiceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPullRequestService)(nil).Create), arg0, arg1, arg2)
}

// CreateComment mocks base method
func (m *MockPullRequestService) CreateComment(arg0 context.Context, arg1 string, arg2 int, arg3 *scm.CommentInput) (*scm.Comment, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Comment)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateComment indicates an expected call of CreateComment
func (mr *MockPullRequestServiceMockRecorder) CreateComment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockPullRequestService)(nil).CreateComment), arg0, arg1, arg2, arg3)
}

// DeleteComment mocks base method
func (m *MockPullRequestService) DeleteComment(arg0 context.Context, arg1 string, arg2, arg3 int) (*scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockPullRequestServiceMockRecorder) DeleteComment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockPullRequestService)(nil).DeleteComment), arg0, arg1, arg2, arg3)
}

// Find mocks base method
func (m *MockPullRequestService) Find(arg0 context.Context, arg1 string, arg2 int) (*scm.PullRequest, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.PullRequest)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockPullRequestServiceMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPullRequestService)(nil).Find), arg0, arg1, arg2)
}

// FindComment mocks base method
func (m *MockPullRequestService) FindComment(arg0 context.Context, arg1 string, arg2, arg3 int) (*scm.Comment, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Comment)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindComment indicates an expected call of FindComment
func (mr *MockPullRequestServiceMockRecorder) FindComment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindComment", reflect.TypeOf((*MockPullRequestService)(nil).FindComment), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockPullRequestService) List(arg0 context.Context, arg1 string, arg2 scm.PullRequestListOptions) ([]*scm.PullRequest, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*scm.PullRequest)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockPullRequestServiceMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPullRequestService)(nil).List), arg0, arg1, arg2)
}

// ListChanges mocks base method
func (m *MockPullRequestService) ListChanges(arg0 context.Context, arg1 string, arg2 int, arg3 scm.ListOptions) ([]*scm.Change, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChanges", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*scm.Change)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListChanges indicates an expected call of ListChanges
func (mr *MockPullRequestServiceMockRecorder) ListChanges(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChanges", reflect.TypeOf((*MockPullRequestService)(nil).ListChanges), arg0, arg1, arg2, arg3)
}

// ListComments mocks base method
func (m *MockPullRequestService) ListComments(arg0 context.Context, arg1 string, arg2 int, arg3 scm.ListOptions) ([]*scm.Comment, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*scm.Comment)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListComments indicates an expected call of ListComments
func (mr *MockPullRequestServiceMockRecorder) ListComments(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockPullRequestService)(nil).ListComments), arg0, arg1, arg2, arg3)
}

// Merge mocks base method
func (m *MockPullRequestService) Merge(arg0 context.Context, arg1 string, arg2 int) (*scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Merge", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Merge indicates an expected call of Merge
func (mr *MockPullRequestServiceMockRecorder) Merge(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockPullRequestService)(nil).Merge), arg0, arg1, arg2)
}

// MockRepositoryService is a mock of RepositoryService interface
type MockRepositoryService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryServiceMockRecorder
}

// MockRepositoryServiceMockRecorder is the mock recorder for MockRepositoryService
type MockRepositoryServiceMockRecorder struct {
	mock *MockRepositoryService
}

// NewMockRepositoryService creates a new mock instance
func NewMockRepositoryService(ctrl *gomock.Controller) *MockRepositoryService {
	mock := &MockRepositoryService{ctrl: ctrl}
	mock.recorder = &MockRepositoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryService) EXPECT() *MockRepositoryServiceMockRecorder {
	return m.recorder
}

// CreateHook mocks base method
func (m *MockRepositoryService) CreateHook(arg0 context.Context, arg1 string, arg2 *scm.HookInput) (*scm.Hook, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHook", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Hook)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateHook indicates an expected call of CreateHook
func (mr *MockRepositoryServiceMockRecorder) CreateHook(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHook", reflect.TypeOf((*MockRepositoryService)(nil).CreateHook), arg0, arg1, arg2)
}

// CreateStatus mocks base method
func (m *MockRepositoryService) CreateStatus(arg0 context.Context, arg1, arg2 string, arg3 *scm.StatusInput) (*scm.Status, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Status)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateStatus indicates an expected call of CreateStatus
func (mr *MockRepositoryServiceMockRecorder) CreateStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStatus", reflect.TypeOf((*MockRepositoryService)(nil).CreateStatus), arg0, arg1, arg2, arg3)
}

// DeleteHook mocks base method
func (m *MockRepositoryService) DeleteHook(arg0 context.Context, arg1, arg2 string) (*scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHook", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHook indicates an expected call of DeleteHook
func (mr *MockRepositoryServiceMockRecorder) DeleteHook(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHook", reflect.TypeOf((*MockRepositoryService)(nil).DeleteHook), arg0, arg1, arg2)
}

// Find mocks base method
func (m *MockRepositoryService) Find(arg0 context.Context, arg1 string) (*scm.Repository, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*scm.Repository)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockRepositoryServiceMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepositoryService)(nil).Find), arg0, arg1)
}

// FindHook mocks base method
func (m *MockRepositoryService) FindHook(arg0 context.Context, arg1, arg2 string) (*scm.Hook, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindHook", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scm.Hook)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindHook indicates an expected call of FindHook
func (mr *MockRepositoryServiceMockRecorder) FindHook(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindHook", reflect.TypeOf((*MockRepositoryService)(nil).FindHook), arg0, arg1, arg2)
}

// FindPerms mocks base method
func (m *MockRepositoryService) FindPerms(arg0 context.Context, arg1 string) (*scm.Perm, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPerms", arg0, arg1)
	ret0, _ := ret[0].(*scm.Perm)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindPerms indicates an expected call of FindPerms
func (mr *MockRepositoryServiceMockRecorder) FindPerms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPerms", reflect.TypeOf((*MockRepositoryService)(nil).FindPerms), arg0, arg1)
}

// List mocks base method
func (m *MockRepositoryService) List(arg0 context.Context, arg1 scm.ListOptions) ([]*scm.Repository, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*scm.Repository)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockRepositoryServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepositoryService)(nil).List), arg0, arg1)
}

// ListHooks mocks base method
func (m *MockRepositoryService) ListHooks(arg0 context.Context, arg1 string, arg2 scm.ListOptions) ([]*scm.Hook, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHooks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*scm.Hook)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListHooks indicates an expected call of ListHooks
func (mr *MockRepositoryServiceMockRecorder) ListHooks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHooks", reflect.TypeOf((*MockRepositoryService)(nil).ListHooks), arg0, arg1, arg2)
}

// ListStatus mocks base method
func (m *MockRepositoryService) ListStatus(arg0 context.Context, arg1, arg2 string, arg3 scm.ListOptions) ([]*scm.Status, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*scm.Status)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListStatus indicates an expected call of ListStatus
func (mr *MockRepositoryServiceMockRecorder) ListStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStatus", reflect.TypeOf((*MockRepositoryService)(nil).ListStatus), arg0, arg1, arg2, arg3)
}

// UpdateHook mocks base method
func (m *MockRepositoryService) UpdateHook(arg0 context.Context, arg1, arg2 string, arg3 *scm.HookInput) (*scm.Hook, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHook", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scm.Hook)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateHook indicates an expected call of UpdateHook
func (mr *MockRepositoryServiceMockRecorder) UpdateHook(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHook", reflect.TypeOf((*MockRepositoryService)(nil).UpdateHook), arg0, arg1, arg2, arg3)
}

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockUserService) Find(arg0 context.Context) (*scm.User, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*scm.User)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockUserServiceMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserService)(nil).Find), arg0)
}

// FindEmail mocks base method
func (m *MockUserService) FindEmail(arg0 context.Context) (string, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEmail", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindEmail indicates an expected call of FindEmail
func (mr *MockUserServiceMockRecorder) FindEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEmail", reflect.TypeOf((*MockUserService)(nil).FindEmail), arg0)
}

// FindLogin mocks base method
func (m *MockUserService) FindLogin(arg0 context.Context, arg1 string) (*scm.User, *scm.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLogin", arg0, arg1)
	ret0, _ := ret[0].(*scm.User)
	ret1, _ := ret[1].(*scm.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindLogin indicates an expected call of FindLogin
func (mr *MockUserServiceMockRecorder) FindLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLogin", reflect.TypeOf((*MockUserService)(nil).FindLogin), arg0, arg1)
}
