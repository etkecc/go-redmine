package redmine

import (
	"io"

	redmine "github.com/nixys/nxs-go-redmine/v5"
)

type mockAPI struct {
	projectID int64
	userID    int64
}

func (m *mockAPI) ProjectSingleGet(_ string, _ redmine.ProjectSingleGetRequest) (redmine.ProjectObject, redmine.StatusCode, error) {
	return redmine.ProjectObject{ID: m.projectID}, 200, nil
}

func (m *mockAPI) UserCurrentGet(_ redmine.UserCurrentGetRequest) (redmine.UserObject, redmine.StatusCode, error) {
	return redmine.UserObject{ID: m.userID}, 200, nil
}

func (m *mockAPI) AttachmentUpload(_ string) (redmine.AttachmentUploadObject, redmine.StatusCode, error) {
	return redmine.AttachmentUploadObject{Token: "123", Filename: "test.txt"}, 201, nil
}

func (m *mockAPI) AttachmentUploadStream(_ io.Reader, filename string) (redmine.AttachmentUploadObject, redmine.StatusCode, error) {
	return redmine.AttachmentUploadObject{Token: "123", Filename: filename}, 201, nil
}

//nolint:gocritic // This is a mock
func (m *mockAPI) IssueCreate(_ redmine.IssueCreate) (redmine.IssueObject, redmine.StatusCode, error) {
	return redmine.IssueObject{ID: 123}, 201, nil
}

//nolint:gocritic // This is a mock
func (m *mockAPI) IssueUpdate(_ int64, _ redmine.IssueUpdate) (redmine.StatusCode, error) {
	return 200, nil
}

func (m *mockAPI) IssueSingleGet(id int64, _ redmine.IssueSingleGetRequest) (redmine.IssueObject, redmine.StatusCode, error) {
	return redmine.IssueObject{ID: id, Status: redmine.IssueStatusObject{ID: 3}}, 200, nil
}
