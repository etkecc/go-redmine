package redmine

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestNewConfig(t *testing.T) {
	logger := zerolog.New(nil)

	cfg := NewConfig(
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectIdentifier("my-project"),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)

	if cfg.Log == nil {
		t.Error("Expected cfg.Log to be set, but it was nil")
	}
	if cfg.Host != "https://redmine.example.com" {
		t.Errorf("Expected Host to be 'https://redmine.example.com', but got '%s'", cfg.Host)
	}
	if cfg.APIKey != "your_api_key" {
		t.Errorf("Expected APIKey to be 'your_api_key', but got '%s'", cfg.APIKey)
	}
	if cfg.ProjectIdentifier != "my-project" {
		t.Errorf("Expected ProjectIdentifier to be 'my-project', but got '%s'", cfg.ProjectIdentifier)
	}
	if cfg.TrackerID != 1 {
		t.Errorf("Expected TrackerID to be 1, but got %d", cfg.TrackerID)
	}
	if cfg.WaitingForOperatorStatusID != 1 {
		t.Errorf("Expected WaitingForOperatorStatusID to be 1, but got %d", cfg.WaitingForOperatorStatusID)
	}
	if cfg.WaitingForCustomerStatusID != 2 {
		t.Errorf("Expected WaitingForCustomerStatusID to be 2, but got %d", cfg.WaitingForCustomerStatusID)
	}
	if cfg.DoneStatusID != 3 {
		t.Errorf("Expected DoneStatusID to be 3, but got %d", cfg.DoneStatusID)
	}
}

func TestConfigEnabled(t *testing.T) {
	logger := zerolog.New(nil)

	cfg := NewConfig(
		WithLog(&logger),
		WithHost("https://redmine.example.com"),
		WithAPIKey("your_api_key"),
		WithProjectIdentifier("my-project"),
		WithTrackerID(1),
		WithWaitingForOperatorStatusID(1),
		WithWaitingForCustomerStatusID(2),
		WithDoneStatusID(3),
	)

	if !cfg.Enabled() {
		t.Error("Expected cfg.Enabled() to return true, but it returned false")
	}

	cfg.Host = ""
	if cfg.Enabled() {
		t.Error("Expected cfg.Enabled() to return false, but it returned true")
	}
}

func TestConfigDefaults(t *testing.T) {
	cfg := NewConfig()

	if cfg.Log == nil {
		t.Error("Expected cfg.Log to be set, but it was nil")
	}
	if cfg.Host != "" {
		t.Errorf("Expected Host to be '', but got '%s'", cfg.Host)
	}
	if cfg.APIKey != "" {
		t.Errorf("Expected APIKey to be '', but got '%s'", cfg.APIKey)
	}
	if cfg.ProjectIdentifier != "" {
		t.Errorf("Expected ProjectIdentifier to be '', but got '%s'", cfg.ProjectIdentifier)
	}
	if cfg.ProjectID != 0 {
		t.Errorf("Expected ProjectID to be 0, but got %d", cfg.ProjectID)
	}
	if cfg.UserID != 0 {
		t.Errorf("Expected UserID to be 0, but got %d", cfg.UserID)
	}
	if cfg.TrackerID != 0 {
		t.Errorf("Expected TrackerID to be 0, but got %d", cfg.TrackerID)
	}
	if cfg.WaitingForOperatorStatusID != 0 {
		t.Errorf("Expected WaitingForOperatorStatusID to be 0, but got %d", cfg.WaitingForOperatorStatusID)
	}
	if cfg.WaitingForCustomerStatusID != 0 {
		t.Errorf("Expected WaitingForCustomerStatusID to be 0, but got %d", cfg.WaitingForCustomerStatusID)
	}
	if cfg.DoneStatusID != 0 {
		t.Errorf("Expected DoneStatusID to be 0, but got %d", cfg.DoneStatusID)
	}
}

func TestConfigOptions(t *testing.T) {
	logger := zerolog.New(nil)

	cfg := &Config{}
	WithLog(&logger)(cfg)
	WithHost("https://redmine.example.com")(cfg)
	WithAPIKey("your_api_key")(cfg)
	WithProjectIdentifier("my-project")(cfg)
	WithProjectID(123)(cfg)
	WithUserID(456)(cfg)
	WithTrackerID(1)(cfg)
	WithWaitingForOperatorStatusID(1)(cfg)
	WithWaitingForCustomerStatusID(2)(cfg)
	WithDoneStatusID(3)(cfg)

	if cfg.Log != &logger {
		t.Error("Expected cfg.Log to be set to the provided logger, but it was not")
	}
	if cfg.Host != "https://redmine.example.com" {
		t.Errorf("Expected Host to be 'https://redmine.example.com', but got '%s'", cfg.Host)
	}
	if cfg.APIKey != "your_api_key" {
		t.Errorf("Expected APIKey to be 'your_api_key', but got '%s'", cfg.APIKey)
	}
	if cfg.ProjectIdentifier != "my-project" {
		t.Errorf("Expected ProjectIdentifier to be 'my-project', but got '%s'", cfg.ProjectIdentifier)
	}
	if cfg.ProjectID != 123 {
		t.Errorf("Expected ProjectID to be 123, but got %d", cfg.ProjectID)
	}
	if cfg.UserID != 456 {
		t.Errorf("Expected UserID to be 456, but got %d", cfg.UserID)
	}
	if cfg.TrackerID != 1 {
		t.Errorf("Expected TrackerID to be 1, but got %d", cfg.TrackerID)
	}
	if cfg.WaitingForOperatorStatusID != 1 {
		t.Errorf("Expected WaitingForOperatorStatusID to be 1, but got %d", cfg.WaitingForOperatorStatusID)
	}
	if cfg.WaitingForCustomerStatusID != 2 {
		t.Errorf("Expected WaitingForCustomerStatusID to be 2, but got %d", cfg.WaitingForCustomerStatusID)
	}
	if cfg.DoneStatusID != 3 {
		t.Errorf("Expected DoneStatusID to be 3, but got %d", cfg.DoneStatusID)
	}
}
