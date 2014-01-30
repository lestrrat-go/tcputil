package tcputil

import (
  "fmt"
  "log"
  "os/exec"
  "net"
  "syscall"
  "testing"
  "time"
)

func Example() {
  p, err := EmptyPort()
  if err != nil {
    log.Fatal(err)
  }

  cmd := exec.Command("memcached", "-p", fmt.Sprintf("%d", p))
  cmd.SysProcAttr = &syscall.SysProcAttr {
    Setpgid: true,
  }
  go cmd.Run()

  err = WaitLocalPort(p, 30 * time.Second)
  if err != nil {
    log.Fatal(err)
  }

  // Now do whatever with memcached!

  cmd.Process.Kill()
}

func TestBasic(t *testing.T) {
  p, err := EmptyPort()
  if err != nil {
    t.Fatalf("Could not find a port to listen to: %s", err)
  }

  go func() {
    t.Logf("Attempting to bind to port %d", p)
    l, err := net.Listen("tcp", fmt.Sprintf(":%d", p))
    if err != nil {
      t.Fatalf("Could not listen on port %d: %s", p, err)
    }

    t.Logf("Accepting new connection...")
    _, err = l.Accept()
    if err != nil {
      t.Fatalf("Could not accept connection: %s", err)
    }
  }()

  t.Logf("Waiting for port to be available...")
  err = WaitLocalPort(p, 1 * time.Minute)
  if err != nil {
    t.Fatalf("Waited for port to open, but could not connect to %d: %s", p, err)
  }
}