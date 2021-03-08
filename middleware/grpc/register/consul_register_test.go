package register

import (
	"testing"
	"time"
)

func TestNewConsulRegister(t *testing.T) {
	reg := NewConsulRegister("47.104.27.123:8500",5)
	reg.Register(RegisterInfo{
		Host:"127.0.0.1",Port:8080,ServiceName:"test",UpdateInterval:5*time.Second,
	})
	time.Sleep(60*time.Second)
}
