package util

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullString(t *testing.T) {
	assert.Equal(t, types.StringNull(), NullString(""), "they should be equal")
	assert.Equal(t, types.StringValue("test"), NullString("test"), "they should be equal")
}

func TestTestError(t *testing.T) {
	diags := &diag.Diagnostics{}
	assert.False(t, TestError(diags, nil, "no error"), "expected false for no error")
	assert.True(t, TestError(diags, errors.New("some error"), "error occurred"), "expected true for error")
}

func TestTestDiagnostic(t *testing.T) {
	diags := &diag.Diagnostics{}
	addDiag := []diag.Diagnostic{}
	assert.False(t, TestDiagnostic(diags, addDiag), "expected false for no additional diagnostics")

	addDiag = diag.Diagnostics{diag.NewErrorDiagnostic("Error", "error occurred")}
	assert.True(t, TestDiagnostic(diags, addDiag), "expected true for additional diagnostics")
}

func TestTestDiagnostics(t *testing.T) {
	diags := &diag.Diagnostics{}
	addDiags := []diag.Diagnostics{
		*diags,
	}
	assert.False(t, TestDiagnostics(diags, addDiags...), "expected false for no additional diagnostics")

	addDiags = []diag.Diagnostics{
		{diag.NewErrorDiagnostic("Error", "error occurred")},
	}
	assert.True(t, TestDiagnostics(diags, addDiags...), "expected true for additional diagnostics")
}
