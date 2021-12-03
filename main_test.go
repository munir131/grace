package grace

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGrace(t *testing.T) {
	Init(10 * time.Second)
	go func() {
		time.Sleep(5 * time.Millisecond)
		SetTrue()
	}()
	go func() {
		check := CheckGraceSig()
		assert.Equal(t, check, false)
	}()

	time.Sleep(6 * time.Millisecond)
	check := CheckGraceSig()
	assert.Equal(t, check, true)

	assert.Equal(t, PeriodOver(), false)

}
