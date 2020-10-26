package ethernet

import (
	"github.com/pkg/errors"
	"math/rand"
)

func CheckChannelStatus() error {
	if rand.Intn(100) >= 30 {
		return nil
	}
	return errors.New("channel busy")
}
