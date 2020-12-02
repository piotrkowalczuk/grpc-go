// +build !go1.11 !go1.12

package status

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
)

func contextStatus(err error) (*Status, bool) {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return New(codes.DeadlineExceeded, err.Error()), true
	case errors.Is(err, context.Canceled):
		return New(codes.Canceled, err.Error()), true
	}
	return nil, false
}
