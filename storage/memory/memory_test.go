package memory

import (
	"context"
	"testing"
	"time"

	"github.com/mister-turtle/golang-webapp-structure/evidence"
)

func TestMemoryRepository(t *testing.T) {
	repo := NewRepository()
	serviceIOC := evidence.NewIOCService(&repo.IOC)
	newIOCs := []evidence.IOC{
		evidence.NewIOC("ipv4", "127.0.0.1", time.Now(), "test-service"),
		evidence.NewIOC("ipv4", "8.8.8.8", time.Now(), "test-service"),
		evidence.NewIOC("domain", "example.org", time.Now(), "test-service"),
		evidence.NewIOC("file-md5", "ace3ea1ab3ae555ccf3125c134b6ab2f", time.Now(), "test-service"),
	}

	for _, ioc := range newIOCs {
		err := serviceIOC.Create(context.TODO(), ioc)
		if err != nil {
			t.Errorf("could not create iocs: %v", err)
			return
		}
	}

	iocs, err := repo.IOC.FindAll(context.TODO())
	if err != nil {
		t.Errorf("could not query iocs: %v", err)
		return
	}

	if l := len(iocs); l != 4 {
		t.Errorf("expected 4 ios, instead got %d", l)
	}
	for i, ioc := range iocs {
		t.Logf("%02d: [%s] - %s (%s)\n", i, ioc.Date, ioc.Type, ioc.Value)
	}
}
