package evidence

import "context"

type IOCRepository interface {
	Create(context.Context, IOC) error
	FindAll(context.Context) ([]IOC, error)
}
