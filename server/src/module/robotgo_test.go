package module

import (
	"testing"
	"time"
)

func Test_robotGo_MouseMove(t *testing.T) {
	r := &RobotGo{}
	// robotgo.Scale = true
	r.MouseMoveRelative(10, 10)
	time.Sleep(1 * time.Second)
	r.MouseMoveRelative(10, 10)
	time.Sleep(1 * time.Second)
	r.MouseMoveRelative(10, 10)
	time.Sleep(1 * time.Second)
	r.MouseMoveRelative(10, 10)
}
