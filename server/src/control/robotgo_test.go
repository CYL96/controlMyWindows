package control

import (
	"github.com/go-vgo/robotgo"
	"testing"
)

func Test_robotGo_MouseMove(t *testing.T) {
	r := &RobotGo{}
	robotgo.Scale = true
	r.MouseMove(1936, 370)
}
