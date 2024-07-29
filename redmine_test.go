package redmine

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestNew(t *testing.T) {
	logger := zerolog.New(nil)
	api := &mockAPI{projectID: 123, userID: 456}

	r, err := New(
		withAPI(api),
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectIdentifier("my-project"),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if r == nil {
		t.Fatal("Expected Redmine client, but got nil")
	}

	if r.cfg.ProjectID != 123 {
		t.Errorf("Expected ProjectID to be 123, but got %d", r.cfg.ProjectID)
	}
	if r.cfg.UserID != 456 {
		t.Errorf("Expected UserID to be 456, but got %d", r.cfg.UserID)
	}
}

func TestNewIssue(t *testing.T) {
	logger := zerolog.New(nil)
	r, _ := New(
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectID(123),
		WithUserID(456),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)
	r.cfg.api = &mockAPI{}

	issueID, err := r.NewIssue("Test Issue", "email", "test@example.com", "This is a test issue.", 0)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if issueID != 123 {
		t.Errorf("Expected issue ID to be 123, but got %d", issueID)
	}
}

func TestUpdateIssue(t *testing.T) {
	logger := zerolog.New(nil)
	r, _ := New(
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectID(123),
		WithUserID(456),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)
	r.cfg.api = &mockAPI{}

	statusID := r.StatusToID(Done)
	err := r.UpdateIssue(123, statusID, "Closing issue.")
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
}

func TestIsClosed(t *testing.T) {
	logger := zerolog.New(nil)
	r, _ := New(
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectID(123),
		WithUserID(456),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)
	r.cfg.api = &mockAPI{}

	closed, err := r.IsClosed(123)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if !closed {
		t.Error("Expected issue to be closed, but it was not")
	}
}

func TestGetNotes(t *testing.T) {
	logger := zerolog.New(nil)
	r, _ := New(
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectID(123),
		WithUserID(456),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)
	r.cfg.api = &mockAPI{}

	notes, err := r.GetNotes(123)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if notes != nil {
		t.Errorf("Expected no notes, but got %v", notes)
	}
}

func TestRedmineShutdown(_ *testing.T) {
	logger := zerolog.New(nil)
	r, _ := New(
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectID(123),
		WithUserID(456),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)

	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
	}()

	r.Shutdown()
}
