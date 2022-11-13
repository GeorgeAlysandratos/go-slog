package slog

import "testing"

func Test_LogStuff(t* testing.T) {
	Info("Heyyy")
	Warn("eeee")
	Error("adasd")

	Infof("Heyyy")
	Warnf("eeee")
	Errorf("adasd")
}
