package task

import (
	generated "RestCrud/openapi"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Helper function to wrap const, because go doesn't allow to take address of a literal or const directly
func ptr[T any](v T) *T {
	return &v
}

func formatDate(t time.Time) string {
	return t.Format("02/01/2006")
}

func TestTaskRequestValidation(t *testing.T) {

	type testCase struct {
		// Input params
		task generated.TaskRequest
		// Expected values
		err error
	}

	t.Run("valid date", func(t *testing.T) {
		currentTime := time.Now()
		tests := []testCase{
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(1, 0, 0))}, err: nil},
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(0, 1, 0))}, err: nil},
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(0, 0, 1))}, err: nil},
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(0, 1, 1))}, err: nil},
		}

		for _, test := range tests {
			err := ValidateDate(&test.task)
			assert.NoError(t, err)
		}
	})

	t.Run("date in the past", func(t *testing.T) {
		currentTime := time.Now()
		tests := []testCase{
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(-1, 0, 0))}, err: ErrDueDateInPast},  // 1 year ago
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(0, -6, 0))}, err: ErrDueDateInPast},  // 6 months ago
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(-2, 0, 0))}, err: ErrDueDateInPast},  // 2 years ago
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(0, 0, -10))}, err: ErrDueDateInPast}, // 10 days ago
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(0, 0, -1))}, err: ErrDueDateInPast},  // 10 days ago
			{task: generated.TaskRequest{DueDate: formatDate(currentTime.AddDate(0, 0, 0))}, err: ErrDueDateInPast},   // current date
		}

		for _, test := range tests {
			err := ValidateDate(&test.task)
			assert.ErrorIs(t, err, ErrDueDateInPast)
		}
	})

	t.Run("invalid date format", func(t *testing.T) {
		tests := []testCase{
			{task: generated.TaskRequest{DueDate: "1/12/2043"}},
			{task: generated.TaskRequest{DueDate: "01/1/2043"}},
			{task: generated.TaskRequest{DueDate: "01/12/202243"}},
			{task: generated.TaskRequest{DueDate: "1/123/2043"}},
			{task: generated.TaskRequest{DueDate: "112/12/2043"}},
		}

		for _, test := range tests {
			err := ValidateDate(&test.task)
			assert.ErrorIs(t, err, ErrInvalidDateFormat)
		}
	})

	t.Run("valid status", func(t *testing.T) {
		tests := []testCase{
			{task: generated.TaskRequest{Status: ptr(generated.PENDING)}},
			{task: generated.TaskRequest{Status: ptr(generated.INPROGRESS)}},
			{task: generated.TaskRequest{Status: ptr(generated.COMPLETED)}},
			{task: generated.TaskRequest{Status: ptr(generated.CANCELLED)}},
		}

		for _, test := range tests {
			err := ValidateStatus(&test.task)
			assert.NoError(t, err)
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		tests := []testCase{
			{task: generated.TaskRequest{Status: (*generated.Status)(ptr("invalid"))}},
			{task: generated.TaskRequest{Status: (*generated.Status)(ptr("AWESOME"))}},
			{task: generated.TaskRequest{Status: (*generated.Status)(ptr("qwewq"))}},
			{task: generated.TaskRequest{Status: (*generated.Status)(ptr("eqwee1"))}},
		}

		for _, test := range tests {
			err := ValidateStatus(&test.task)
			assert.ErrorIs(t, err, ErrInvalidStatus)
		}
	})
}
