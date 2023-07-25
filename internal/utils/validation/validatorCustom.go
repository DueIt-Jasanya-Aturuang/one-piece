package validation

import "fmt"

func MsgForTag(tag, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("This Field Min Character %s", param)
	case "max":
		return fmt.Sprintf("This Field Max Character %s", param)
	}
	return ""
}
