package cache_test

import (
	"testing"
	"time"

	"github.com/mgjules/minion/pkg/cache"
	"github.com/stretchr/testify/assert"
)

type test struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	Embedded    embedded
}

type embedded struct {
	ID   string
	Name string
}

func TestCacheCoster(t *testing.T) {
	t.Parallel()

	d := test{
		ID:   "36434a92-69dd-549c-9387-b9df7777ac35",
		Name: "Hettie Shelton",
		Description: `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
		Duis nec finibus dolor, eu egestas metus. 
		Proin cursus velit leo, sit amet eleifend turpis iaculis sollicitudin. 
		Nulla eu mollis lacus. Vivamus hendrerit rutrum velit non congue. 
		In a pellentesque urna. Interdum et malesuada fames ac ante ipsum primis in faucibus. 
		In felis felis, convallis accumsan dignissim non, tristique non justo. 
		Duis interdum pharetra massa et cursus. Etiam nec dignissim nibh. 
		Donec vitae arcu nec dolor dictum consectetur et ut nibh. 
		Suspendisse elementum dui eget dignissim mattis. Vivamus sed erat ac mi porta commodo. 
		Nunc ultrices ut nisl ut ullamcorper.
		`,
		CreatedAt: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC), //nolint:revive
		Embedded: embedded{
			ID:   "183ebee4-938f-5c77-9950-6f2afef2e3f0",
			Name: "Shane Ellis",
		},
	}

	assert.Equalf(t, int64(875), cache.Coster(d), "Size of test object does not match expected size") //nolint:revive
}

func BenchmarkCacheCoster(b *testing.B) {
	d := test{
		ID:   "36434a92-69dd-549c-9387-b9df7777ac35",
		Name: "Hettie Shelton",
		Description: `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
		Duis nec finibus dolor, eu egestas metus. 
		Proin cursus velit leo, sit amet eleifend turpis iaculis sollicitudin. 
		Nulla eu mollis lacus. Vivamus hendrerit rutrum velit non congue. 
		In a pellentesque urna. Interdum et malesuada fames ac ante ipsum primis in faucibus. 
		In felis felis, convallis accumsan dignissim non, tristique non justo. 
		Duis interdum pharetra massa et cursus. Etiam nec dignissim nibh. 
		Donec vitae arcu nec dolor dictum consectetur et ut nibh. 
		Suspendisse elementum dui eget dignissim mattis. Vivamus sed erat ac mi porta commodo. 
		Nunc ultrices ut nisl ut ullamcorper.
		`,
		CreatedAt: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC), //nolint:revive
		Embedded: embedded{
			ID:   "183ebee4-938f-5c77-9950-6f2afef2e3f0",
			Name: "Shane Ellis",
		},
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = cache.Coster(d)
	}
}
