package util

import (
	"fmt"
)

var Required = "field ini tidak boleh dikosongkan"
var minString = "field ini tidak boleh kurang dari %d"
var maxString = "field ini tidak boleh lebih dari %d"
var minNumeric = "field ini tidak boleh kurang dari %d"
var maxNumeric = "field ini tidak boleh lebih dari %d"
var FileSize = "maximal size harus %d kb atau %d mb"
var FileContent = "file content harus %s"

const Image = "image"

func ContentType(typeContent string) []string {
	switch typeContent {
	case Image:
		return []string{
			"image/png", "image/jpeg", "image/jpg",
		}
	}

	return []string{
		"image/png", "image/jpeg", "image/jpg",
	}
}
func CheckContentType(headerContentType string, typeContent string) bool {
	if headerContentType == "" {
		return false
	}

	var status bool
	for _, v := range ContentType(typeContent) {
		if headerContentType == v {
			return true
		}
		status = false
	}
	return status
}

func MaxMinString(s string, min, max int) string {
	switch {
	case len(s) < min:
		return fmt.Sprintf(minString, min)
	case len(s) > max:
		return fmt.Sprintf(maxString, max)
	}

	return ""
}

func MaxMinNumeric(i int, min, max int) string {
	switch {
	case i < min:
		return fmt.Sprintf(minNumeric, min)
	case i > max:
		return fmt.Sprintf(maxNumeric, max)
	}

	return ""
}
