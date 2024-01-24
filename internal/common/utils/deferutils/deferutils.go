// Package deferutils provides functions for func() error call in defer. Supports silent call and call with logger.
package deferutils

import (
	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
)

// ExecWithLogError support method for defer functions call which should return error.
func ExecWithLogError(callback func() error, log logger.BaseLogger) {
	if callback == nil {
		return
	}

	if err := callback(); err != nil {
		log.Infof("callback execution finished with error: %v", err)
	}
}

// ExecSilent support method for defer functions call which should ignore error.
func ExecSilent(callback func() error) {
	_ = callback()
}
