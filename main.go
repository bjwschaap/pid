package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, ys []float64
	input := 24.0594
	setpoint := 100.0
	interval := 1
	noise := 0.01
	pid := NewPID(interval, 3.0, 5.0, 1.0)
	//pid.SetOutputLimits(-200.0, 200.0)
	i := 0.0
	for math.Abs(setpoint - input) > 0.01 {
		i++
		output := pid.Compute(setpoint, input)
		input +=  noise * output
		x = append(x, i)
		y = append(y, input)
		ys = append(ys, output)
		//if i == 125 {
		//	setpoint = 0
		//}
		//if i == 200 {
		//	pid.SetTunings(8.0, 0.2, 0.1)
		//}
	}
	if err := makeChart(x, y, ys); err != nil {
		panic(err)
	}
	fmt.Println("done")
}
