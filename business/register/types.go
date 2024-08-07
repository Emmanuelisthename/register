package register

import (
	"context"
	"dev/register.git/business/i"
)

// Constants
const ()

// Service encapsulates core register functionality
type Service struct {
	Log   i.Logger
	Store Store
}

// Store encapsulates third-party dependencies
type Store interface {
	Create(ctx context.Context) error
}
