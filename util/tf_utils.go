package util

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NullString converts a string to a null
// or non-null types.String depending on if it's empty.
func NullString(value string) types.String {
	if len(value) == 0 {
		// If input string is empty, then return a null types.String.
		return types.StringNull()
	} else {
		// If input string is non-empty, then return a non-null types.String with the same value.
		return types.StringValue(value)
	}
}

// TestError checks if there's error and appends the error to the diagnostic.
func TestError(diagnostic *diag.Diagnostics, err error, message string) bool {
	if err != nil {
		// If error exists,
		// add custom message and error message to the diagnostic, and return true.
		diagnostic.AddError(message, err.Error())
		return true
	} else {
		// If no error, return false.
		return false
	}
}

// TestDiagnostic appends additional diagnostic to given diagnostic and checks if there's any error.
func TestDiagnostic(diagnostic *diag.Diagnostics, additional diag.Diagnostics) bool {
	// Append additional diagnostic to existing one.
	diagnostic.Append(additional...)
	// If there's any error in diagnostic, return true.
	if diagnostic.HasError() {
		return true
	} else {
		// If not, return false.
		return false
	}
}

// TestDiagnostics appends multiple additional diagnostics to given diagnostic and checks if there's any error.
func TestDiagnostics(diagnostic *diag.Diagnostics, additional ...diag.Diagnostics) bool {
	// Loop over each additional diagnostic and append it to existing one.
	for _, diags := range additional {
		diagnostic.Append(diags...)
	}
	// If there's any error in diagnostic, return true.
	if diagnostic.HasError() {
		return true
	} else {
		// If not, return false.
		return false
	}
}
