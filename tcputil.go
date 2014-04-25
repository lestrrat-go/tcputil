package tcputil

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// EmptyPort looks for an empty port to listen on local interface.
func EmptyPort() (int, error) {
	for p := 50000 + rand.Intn(1000); p < 60000; p++ {
		l, e := net.Listen("tcp", fmt.Sprintf(":%d", p))
		if e == nil {
			// yey!
			l.Close()
			return p, nil
		}
	}

	return 0, errors.New("error: Could not find an available port")
}

// WaitPort waits until you can connect to `addr`, up to `dur` amount of time
func WaitPort(addr string, dur time.Duration) error {
	timeout := time.Now().Add(dur)
	for time.Now().Before(timeout) {
		c, e := net.Dial("tcp", addr)
		if e == nil {
			c.Close()
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("error: Could not connect to '%s'", addr)
}

// WaitLocalPort until you can connect to `port` on localhost, up to `dur` amount of time
func WaitLocalPort(port int, dur time.Duration) error {
	return WaitPort(fmt.Sprintf(":%d", port), dur)
}
