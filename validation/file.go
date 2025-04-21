package validation

import (
	"mime/multipart"
	"slices"
	"strings"
)

// IsValidFileExtension checks if the given file's extension matches one of the valid extensions in the provided list.
// It returns true if the file extension is in the list of valid extensions, otherwise returns false.
// The comparison is case-insensitive.
//
// Example usage:
//
//	validExtensions := []string{"jpg", "png", "gif"}
//	isValid := IsValidFileExtension("image.JPG", validExtensions)
func IsValidFileExtension(filename string, validExtensions []string) bool {
	ext := strings.ToLower(getFileExtension(filename))

	return slices.Contains(validExtensions, ext)
}

// getFileExtension extracts the file extension from the given filename.
// The extension is returned in lowercase and includes the dot (e.g., ".jpg").
// If the filename has no extension, it returns an empty string.
//
// Example usage:
//
//	ext := getFileExtension("image.JPG")  // returns ".jpg"
func getFileExtension(filename string) string {
	ext := strings.ToLower(multipart.FileHeader{Filename: filename}.Filename)
	return ext[strings.LastIndex(ext, "."):]
}
