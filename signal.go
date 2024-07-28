/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package events

import (
	"os"
	"os/signal"
	"syscall"
)

// OnStopSignal calling a function if you send a system event stop
func OnStopSignal(callFunc func()) {
	quit := make(chan os.Signal, 4)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL) //nolint:staticcheck
	<-quit
	signal.Stop(quit)
	callFunc()
}

// OnHubSignal calling function if OS sends SIGHUP to re-read configuration file
func OnHubSignal(callFunc func()) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP)
	<-quit
	signal.Stop(quit)
	callFunc()
}

// OnCustomSignal calling a function if you send a system custom event
func OnCustomSignal(callFunc func(), sig ...os.Signal) {
	quit := make(chan os.Signal, len(sig))
	signal.Notify(quit, sig...)
	<-quit
	signal.Stop(quit)
	callFunc()
}
