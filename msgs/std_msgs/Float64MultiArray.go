package std_msgs

import (
	"github.com/aler9/goroslib/msg"
)

type Float64MultiArray struct {
	msg.Package `ros:"std_msgs"`
	Layout      MultiArrayLayout
	Data        []float64
}
