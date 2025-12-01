package day7

import (
	"fmt"
	"time"
)

func Time() {
	now := time.Now()
	rfc3339string := now.Format(time.RFC3339)
	fmt.Println("RFC339 Format:", rfc3339string)

	formatted := now.Format("02-01-2006")
	fmt.Println("Formatted time:", formatted)
}
