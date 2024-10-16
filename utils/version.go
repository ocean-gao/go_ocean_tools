package utils

import "fmt"

var Version = "1.0.0"

func GetVersion() string {
	return fmt.Sprintf("v%v", Version)
}
