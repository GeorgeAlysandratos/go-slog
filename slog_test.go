package slog

import "testing"

func Test_LogStuff(t* testing.T) {
	SetBasename("slog_test")

	Info("Heyyy")
	Note("point me out")
	Warn("eeee")
	Error("adasd")

	Infof("Heyyy")
	Notef("point me out")
	Warnf("eeee")
	Errorf("adasd")
}
