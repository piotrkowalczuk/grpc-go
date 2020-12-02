// +build !go1.11 !go1.12

package errorutils

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ContextErr converts the error from context package into a status error.
// It returns false if given error is not coming from context package.
func ContextErr(err error) (error, bool) {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return status.Error(codes.DeadlineExceeded, err.Error()), true
	case errors.Is(err, context.Canceled):
		return status.Error(codes.Canceled, err.Error()), true
	}
	return nil, false
}
