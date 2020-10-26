package ethernet

import (
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

var (
	MaxCollisionNumber = 3
)

func SendSymbol(symbol int32) {
	time.Sleep(100 * time.Millisecond)
}

func CollisionWindow() error {
	time.Sleep(500 * time.Millisecond)
	if rand.Intn(100) >= 70 {
		return errors.New("collision detected")
	}
	return nil
}
