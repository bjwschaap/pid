package main

import (
	"errors"
)

type PID struct {
	iTerm, lastInput, outMin, outMax, kp, ki, kd float64
	sampleTime int
	limitsActive bool
}

func NewPID(interval int, kp, ki, kd float64) *PID {
	return &PID {
		sampleTime: interval,
		kp: kp,
		ki: ki,
		kd: kd,
	}
}

func (pid *PID) SetOutputLimits(min, max float64) error {
	if min >= max {
		return errors.New("minimum should be smaller than maximum")
	}
	pid.limitsActive = true
	pid.outMin = min
	pid.outMax = max
	return nil
}

func (pid *PID) SetTunings(kp, ki, kd float64) error {
	if kp < 0 || ki < 0 || kd < 0 {
		return errors.New("tuning paramaters must be > 0")
	}
	pid.kp = kp
	pid.ki = ki * float64(pid.sampleTime)
	pid.kd = kd / float64(pid.sampleTime)
	return nil
}

func (pid *PID) SetSampleTime(seconds int) error {
	if seconds < 1 {
		return errors.New("the interval must be > 0 seconds")
	}
	var ratio float64
	if pid.sampleTime > 0 {
		ratio = float64(seconds) / float64(pid.sampleTime)
	} else {
		ratio = 1
	}
	pid.ki = pid.ki * ratio
	pid.kd = pid.kd / ratio
	pid.sampleTime = seconds
	return nil
}

func (pid *PID) Compute(setpoint, input float64) float64 {
	err := setpoint - input
	pid.iTerm += pid.ki * err
	if pid.limitsActive {
		if pid.iTerm > pid.outMax {
			pid.iTerm = pid.outMax
		} else if pid.iTerm < pid.outMin {
			pid.iTerm = pid.outMin
		}
	}
	dInput := input - pid.lastInput
	output := pid.kp * err + pid.ki * pid.iTerm + pid.kd * dInput
	if pid.limitsActive {
		if output > pid.outMax {
			output = pid.outMax
		} else if output < pid.outMin {
			output = pid.outMin
		}
	}
	return output
}
