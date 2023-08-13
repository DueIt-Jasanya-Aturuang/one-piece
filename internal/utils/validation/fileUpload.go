package validation

import (
	"errors"
	"mime/multipart"
	"strconv"
	"strings"
)

func CheckContentType(file *multipart.FileHeader, size, sizeFile int64, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		if sizeFile > size {
			sizeString := strconv.Itoa(int(size) / 2042)
			return errors.New("max size " + sizeString + " kb")
		}
		for _, contentType := range contentTypes {
			contentFile := file.Header.Get("Content-Type")
			if contentFile == contentType {
				return nil
			}
		}
		return errors.New("not allowed type file. type file must be " + strings.Join(contentTypes, " "))
	}
	return errors.New("not found content type")
}
