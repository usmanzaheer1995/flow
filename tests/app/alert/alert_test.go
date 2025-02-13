package app

import (
	"errors"
	"testing"

	"github.com/ibilalkayy/flow/db/alert_db"
	internal_alert "github.com/ibilalkayy/flow/internal/app/alert"
	"github.com/ibilalkayy/flow/internal/common/structs"
)

func TestCreateAlert(t *testing.T) {
	testCases := []struct {
		name        string
		input       *structs.AlertVariables
		expectedMsg string
	}{
		{
			name: "ValidInput",
			input: &structs.AlertVariables{
				Category:  "first",
				Frequency: "hourly",
				Method:    "email",
			},
			expectedMsg: "Created alert for category: first, frequency: hourly, method: email\n",
		},
		{
			name: "NoCategory",
			input: &structs.AlertVariables{
				Category:  "",
				Frequency: "hourly",
				Method:    "email",
			},
			expectedMsg: "Created alert for category: , frequency: hourly, method: email\n",
		},
		{
			name: "NoFrequency",
			input: &structs.AlertVariables{
				Category:  "first",
				Frequency: "",
				Method:    "email",
			},
			expectedMsg: "Created alert for category: first, frequency: , method: email\n",
		},
		{
			name: "NoFrequency",
			input: &structs.AlertVariables{
				Category:  "first",
				Frequency: "hourly",
				Method:    "",
			},
			expectedMsg: "Created alert for category: first, frequency: hourly, method: \n",
		},
	}

	basePath := "../../../db/budget_db/migrations/"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := alert_db.CreateAlert(tc.input, basePath)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestAlertSetup(t *testing.T) {
	tests := []struct {
		name        string
		alertVars   *structs.AlertVariables
		expectedErr error
	}{
		// Test cases for valid inputs
		{
			name: "Valid inputs - Category amount alert",
			alertVars: &structs.AlertVariables{
				Frequency: "weekly",
				Method:    "cli",
				Category:  "food",
			},
			expectedErr: nil,
		},
		// Test cases for invalid inputs
		{
			name: "Invalid frequency",
			alertVars: &structs.AlertVariables{
				Frequency: "invalid",
				Method:    "cli",
				Category:  "food",
			},
			expectedErr: errors.New("invalid alert frequency"),
		},
		{
			name: "Invalid category",
			alertVars: &structs.AlertVariables{
				Frequency: "monthly",
				Method:    "email",
				Category:  "invalid",
			},
			expectedErr: errors.New("category is not given"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := internal_alert.AlertSetup(test.alertVars)
			if err != nil {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}
		})
	}
}

// Mock internal_budget package for testing
// type mockBudget struct{}

// func (m mockBudget) TotalBudgetAmount() (int, error) {
// 	return 1000, nil
// }

// // Mock transaction package for testing
// type mockTransaction struct{}

// func (m mockTransaction) TransactionAmount() int {
// 	return 800
// }

// func TestAlertMessage(t *testing.T) {
// 	// Mock dependencies
// 	_ = mockBudget{}
// 	_ = mockTransaction{}

// 	// Test case where transaction amount is less than total budget
// 	err := internal_alert.AlertMessage()
// 	if err != nil {
// 		t.Errorf("AlertMessage() returned an error, expected nil: %v", err)
// 	}

// 	// Test case where transaction amount is greater than or equal to total budget
// 	_ = mockTransaction{}
// 	err = internal_alert.AlertMessage()
// 	expectedError := "You can't spend more because your budget is set to 1000"
// 	if err == nil || err.Error() != expectedError {
// 		t.Errorf("AlertMessage() returned unexpected error: got %v, want %v", err, expectedError)
// 	}
// }
