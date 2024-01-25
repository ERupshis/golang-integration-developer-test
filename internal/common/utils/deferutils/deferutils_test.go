package deferutils

import (
	"fmt"
	"testing"

	"github.com/erupshis/golang-integration-developer-test/internal/common/logger"
)

func TestExecuteWithLogError(t *testing.T) {
	log := logger.CreateMock()

	type args struct {
		callback func() error
		log      logger.BaseLogger
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid def case",
			args: args{
				callback: func() error {
					return nil
				},
				log: log,
			},
		},
		{
			name: "error from callback",
			args: args{
				callback: func() error {
					return fmt.Errorf("test err")
				},
				log: log,
			},
		},
	}
	for _, ttTmp := range tests {
		tt := ttTmp
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ExecWithLogError(tt.args.callback, tt.args.log)
		})
	}
}

func TestExecSilent(t *testing.T) {
	type args struct {
		callback func() error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid def case",
			args: args{
				callback: func() error {
					return nil
				},
			},
		},
		{
			name: "error from callback",
			args: args{
				callback: func() error {
					return fmt.Errorf("test err")
				},
			},
		},
	}
	for _, ttTmp := range tests {
		tt := ttTmp
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ExecSilent(tt.args.callback)
		})
	}
}
