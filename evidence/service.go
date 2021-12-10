package evidence

import (
	"context"
)

type IOCService struct {
	repository IOCRepository
}

func (i IOCService) Create(ctx context.Context, ioc IOC) error {
	return i.repository.Create(ctx, ioc)
}

func (i IOCService) FindAll(ctx context.Context) ([]IOC, error) {
	return i.repository.FindAll(ctx)
}

func NewIOCService(repository IOCRepository) IOCService {
	return IOCService{repository: repository}
}
