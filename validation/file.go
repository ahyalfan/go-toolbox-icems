package validation

import (
	"mime/multipart"
	"slices"
	"strings"
)

func IsValidFileExtension(filename string, validExtensions []string) bool {
	ext := strings.ToLower(getFileExtension(filename))

	// for _, validExt := range validExtensions {
	// 	if ext == validExt {
	// 		return true
	// 	}
	// }
	// return false

	return slices.Contains(validExtensions, ext)
}

func getFileExtension(filename string) string {
	ext := strings.ToLower(multipart.FileHeader{Filename: filename}.Filename)
	return ext[strings.LastIndex(ext, "."):]
}
