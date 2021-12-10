package memory

type MemoryRepository struct {
	IOC MemoryIOC
}

func NewRepository() *MemoryRepository {
	return &MemoryRepository{
		IOC: MemoryIOC{},
	}
}
