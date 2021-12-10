package memory

import (
	"context"

	"github.com/mister-turtle/golang-webapp-structure/domain/evidence"
)

type MemoryIOC struct {
	IOCs []evidence.IOC
}

func (i *MemoryIOC) Create(_ context.Context, ip evidence.IOC) error {
	i.IOCs = append(i.IOCs, ip)
	return nil
}

func (i MemoryIOC) FindAll(_ context.Context) ([]evidence.IOC, error) {
	return i.IOCs, nil
}
