package env

import "os"

func GetTempDir() string {
	temp := os.Getenv("TEMP")
	if temp == "" {
		temp = "/public"
	}
	return temp
}
