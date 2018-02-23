tcputil
==========

Some Utilities To Help Your TCP-Related testing

[![Build Status](https://travis-ci.org/lestrrat-go/tcputil.png?branch=master)](https://travis-ci.org/lestrrat-go/tcputil)
[![GoDoc](https://godoc.org/github.com/lestrrat-go/tcputil?status.svg)](https://godoc.org/github.com/lestrrat-go/tcputil)

# SYNOPSIS

```go
  p, err := EmptyPort()
  if err != nil {
    log.Fatal(err)
  }

  cmd := exec.Command("memcached", "-p", fmt.Sprintf("%d", p))
  go cmd.Run()

  err = WaitLocalPort(p, 30 * time.Second)
  if err != nil {
    log.Fatal(err)
  }

  // Now do whatever with memcached!

  cmd.Process.Kill()
```

For a higher level wrapper for testing network servers on random local ports,
see [go-tcptest](https://github.com/lestrrat-go/tcptest)

API docs: http://godoc.org/github.com/lestrrat-go/tcputil
