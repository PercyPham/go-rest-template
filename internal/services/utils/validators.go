package utils

import "regexp"

// IsEmail checks if the string is a valid email address
func IsEmail(s string) bool {
	pattern := `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

// IsMobileNo checks if the string is a valid mobile
func IsMobileNo(s string) bool {
	pattern := `^0\d{9}$`
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}
