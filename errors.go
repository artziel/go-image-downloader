package ImageDownloader

import "errors"

var ErrNotFound = errors.New("resource not found")
var ErrServerError = errors.New("remote server error")
var ErrInvalidDataType = errors.New("invalid data type")
