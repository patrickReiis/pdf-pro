package model

import (
	"fmt"
	"time"
)

// Returns a slice with the length of 1
// The element is a timestamp of the time when the this function was called
// The format is ISO 8601
func getRequestTimeStamp() []string {
	return getRequestTimeStampImpl()
}

// Implementation of the `getRequestTimeStamp` function
func getRequestTimeStampImpl() []string {
	return []string{fmt.Sprintf(`%s`, time.Now().UTC().Format(time.RFC3339))}
}
