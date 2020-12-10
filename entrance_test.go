package nblog

import (
	"fmt"
	"nblog/nbutil"
	"reflect"
	"sync"
	"testing"
	"go.uber.org/zap"
)

func TestDefaultLogger(t *testing.T) {
	logger := DefaultLogger()
	l2 := DefaultLogger()
	var loggerPool = sync.Pool{
		New: func() interface{} {
			return logger
		},
	}
	loggerPool.Put(l2)
	l3 := loggerPool.Get().(*zap.SugaredLogger)
	l := loggerPool.Get().(*zap.SugaredLogger)
	fmt.Println(l3,l,777)
	fmt.Println(nbutil.Uuid(),888)
	l.Info("0000000000")
}

func TestNewLogger(t *testing.T) {
	type args struct {
		serverName string
	}
	tests := []struct {
		name    string
		args    args
		want    *zap.SugaredLogger
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLogger(tt.args.serverName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() got = %v, want %v", got, tt.want)
			}
		})
	}
}
