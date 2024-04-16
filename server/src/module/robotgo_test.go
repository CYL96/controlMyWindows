package module

import (
	"testing"

	"github.com/go-vgo/robotgo"
)

func Test_robotGo_MouseMove(t *testing.T) {
	r := &RobotGo{}
	robotgo.Scale = true
	r.MouseMove(1936, 370)
}
