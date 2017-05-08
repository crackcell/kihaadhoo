/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Tests.
 *
 * @file signal_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 11 13:38:15 2014
 *
 **/

package signal

import (
	"fmt"
	"os"
	"testing"
	"time"
	"syscall"
)

//===================================================================
// Public APIs
//===================================================================

func TestEverything(t *testing.T) {
	sset := NewSignalHandlerSet()

	handler := func(s os.Signal, arg interface{}) {
		fmt.Printf("handle signal: %v\n", s)
		if s == syscall.SIGTERM {
			fmt.Printf("signal terminate received, exited normally\n")
			os.Exit(0)
		}
	}
	sset.Register(syscall.SIGINT, handler)
	sset.Register(syscall.SIGUSR1, handler)
	sset.Register(syscall.SIGUSR2, handler)
	sset.Register(syscall.SIGTERM, handler)

	sset.Start()

	time.Sleep(time.Hour)
}

//===================================================================
// Private
//===================================================================
