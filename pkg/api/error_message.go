package api

// errorMessage takes an error message and returns it in a more structured format.
func errorMessage(message string) map[string]string {
	return map[string]string{"error": message}
}

// errorsList takes a list of error messages and returns them in a more structured format.
func errorsList(errs []string) map[string][]string {
	return map[string][]string{"errors": errs}
}
