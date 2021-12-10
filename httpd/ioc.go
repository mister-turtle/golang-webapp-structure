package httpd

import (
	"context"

	"github.com/mister-turtle/golang-webapp-structure/evidence"
)

type iocService interface {
	Create(ctx context.Context, ioc evidence.IOC) error
	FindAll(ctx context.Context) ([]evidence.IOC, error)
}
