package main

import "math"

const (
	controlPI  = 0
	controlPID = 1
)

type Tuner struct {
	input, output, outputStart, NoiseBand, OStep, lastTime, refVal, absMin, absMax, kp, ki, kd, ku, pu float64
	ControlType, lookbackSec, nLookback, sampleTime int
	running bool
}

func (t *Tuner) Cancel() {
	t.running = false
}

func (t *Tuner) Finish() {
	t.output = t.outputStart
	_ = 4 * (2 * t.OStep)/((t.absMax - t.absMin) * math.Pi)

}

func (t *Tuner) GetKp() float64 {
	if t.ControlType == controlPID {
		return 0.6 * t.ku
	}
	return 0.4 * t.ku
}

func (t *Tuner) GetKi() float64 {
	if t.ControlType == controlPID {
		return 1.2 * t.ku / t.pu
	}
	return 0.48 * t.ku / t.pu
}

func (t *Tuner) GetKd() float64 {
	if t.ControlType == controlPID {
		return 0.075 * t.ku * t.pu
	}
	return 0
}

func (t *Tuner) SetLookbackSec(value int) {
	if value < 1 {
		value = 1
	}
	if value < 25 {
		t.nLookback = 4 * value
		t.sampleTime = 250
	} else {
		t.nLookback = 100
		t.sampleTime = 10 * value
	}
}

func (t *Tuner) GetLoockBackSec() int {
	return t.nLookback * t.sampleTime / 1000
}