# Terraform Provider Commons

Collection of common function for developing Terraform provider.

## RecordingHttpPayloadTransport

PayloadTransport with response recording capability.

## utility package

This package provides a set of helper functions for working with strings and diagnostic information in your Terraform
provider.

### Functions

### NullString

The `NullString` function receives a string and returns a `types.String` that is either null or non-null based on
whether the input string is empty.

### TestError

`TestError` function checks for an error in your code, adds it to the diagnostic, and returns a boolean indicating
whether an error has occurred.

### TestDiagnostic

The `TestDiagnostic` function is used to append additional diagnostic information and check if there are any errors.

### TestDiagnostics

`TestDiagnostics` function allows you to append multiple diagnostics at once and then checks if there are any errors.
