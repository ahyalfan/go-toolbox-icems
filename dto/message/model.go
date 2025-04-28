package msg

import "errors"

var (
	/* Error Message */
	ErrCreate       = errors.New("Failed to create the record. Please try again later.")
	ErrRead         = errors.New("Failed to retrieve data. Please try again later.")
	ErrMultipleRead = errors.New("Failed to retrieved multiple data. Please try again later")
	ErrUpdate       = errors.New("Failed to update the record. Please try again later.")
	ErrDelete       = errors.New("Failed to delete the record. Please try again later.")
	ErrRestore      = errors.New("Failed to restore the record. Please try again later.")

	/* Sucecss Message */
	SuccessCreate  = "The record has been successfully created."
	SuccessRead    = "Data retrieved successfully."
	SuccessUpdate  = "The record has been successfully updated."
	SuccessDelete  = "The record has been successfully delete."
	SuccessRestore = "The record has been successfully restore."
)
