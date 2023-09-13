package validation

import (
	"fmt"
)

var required = "field ini tidak boleh dikosongkan"
var minString = "field ini tidak boleh kurang dari %d"
var maxString = "field ini tidak boleh lebih dari %d"
var fileSize = "maximal size harus %d kb atau %d mb"
var fileContent = "file content harus %s"

const image = "image"

func contentType(typeContent string) []string {
	switch typeContent {
	case image:
		return []string{
			"image/png", "image/jpeg", "image/jpg",
		}
	}

	return []string{
		"image/png", "image/jpeg", "image/jpg",
	}
}
func checkContentType(headerContentType string, typeContent string) bool {
	if headerContentType == "" {
		return false
	}

	var status bool
	for _, v := range contentType(typeContent) {
		if headerContentType == v {
			return true
		}
		status = false
	}
	return status
}

func maxMinString(s string, min, max int) string {
	switch {
	case len(s) < min:
		return fmt.Sprintf(minString, 3)
	case len(s) > max:
		return fmt.Sprintf(maxString, 32)
	}

	return ""
}
