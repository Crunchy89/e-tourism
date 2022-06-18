package slug

import "strings"

func Slug(s string) string {
	data := strings.ToLower(s)
	data = strings.Replace(data, " ", "-", -1)
	return data
}
