package grace

import (
	"sync"
	"time"
)

var graceSig safeGraceSig

type safeGraceSig struct {
	sig        bool
	receivedAt time.Time
	lock       sync.Mutex
	period     time.Duration
}

// Init creates graceSig channel
func Init(period time.Duration) {
	graceSig = safeGraceSig{sig: false, period: period}
}

// SetTrue creates graceSig channel
func SetTrue() {
	graceSig.lock.Lock()
	graceSig.sig = true
	graceSig.receivedAt = time.Now()
	graceSig.lock.Unlock()
}

// PeriodOver returns true if grace period is over
func PeriodOver() bool {
	return time.Since(graceSig.receivedAt) > graceSig.period
}

// CheckGraceSig return signals
func CheckGraceSig() bool {
	graceSig.lock.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer graceSig.lock.Unlock()
	return graceSig.sig
}
