// +build go1.11 go1.12

package status

import (
	"context"

	"google.golang.org/grpc/codes"
)

func contextStatus(err error) (*Status, bool) {
	switch err {
	case context.DeadlineExceeded:
		return New(codes.DeadlineExceeded, err.Error()), true
	case context.Canceled:
		return New(codes.Canceled, err.Error()), true
	}
	return nil, false
}
