package utils

import (
	"fmt"
	"time"
)

const (
	DefaultTimeout = 5 * time.Second
)

func GetFileNameWithTime() string {
	timestamp := time.Now().Unix()
	// to string
	return fmt.Sprintf("%d", timestamp)
}
