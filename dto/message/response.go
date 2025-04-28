package msg

import "errors"

var (
	/* Error Message */
	ErrInternalServer = errors.New("Something went wrong, please try again later.")
	ErrUnauthorized   = errors.New("User is not authorized to perform this action.")
	ErrCacheMiss      = errors.New("Cache miss: Requested data not found.")
	ErrGetFile        = errors.New("Error getting form file")
	ErrFailedType     = errors.New("Type failed: failed type in parameter.")
)
