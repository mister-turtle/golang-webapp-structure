package memory

import (
	"context"

	"github.com/mister-turtle/golang-webapp-structure/evidence"
)

type MemoryIOC []evidence.IOC

func (i *MemoryIOC) Create(_ context.Context, ip evidence.IOC) error {
	*i = append(*i, ip)
	return nil
}

func (i *MemoryIOC) FindAll(_ context.Context) ([]evidence.IOC, error) {
	return *i, nil
}
