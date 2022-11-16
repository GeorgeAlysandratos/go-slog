package slog

import "testing"

func Test_LogStuff(t* testing.T) {
	SetBasename("slog_test")

	Info("Heyyy", 15, 3.000565)
	Note("point me out", 15, 3.000565)
	Warn("eeee", 15, 3.000565)
	Error("adasd", 15, 3.000565, "adsad")

	Infof("Heyyy", 15, 3.000565)
	Notef("point me out", 15, 3.000565)
	Warnf("eeee", 15, 3.000565)
	Errorf("adasd", 15, 3.000565)
}
