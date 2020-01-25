//Package stringparse implements additional sanitization against strings with spaces
//Only Dashes(-), Underscores(_), and AlphaNumeric(a-zA-Z0-9) should return True
package stringparse

import (
"regexp"
)

//Parse function validations the string against regex, ensuring only dashes, underscores and alphanumeric
func Parse(s string) bool{
	var status bool
	status = false
	pattern := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)

	if pattern.MatchString(s) {
		status = true
	}
	return status
}
