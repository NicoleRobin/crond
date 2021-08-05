package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Debug("Debug log")
	Debugf("Debugf log, %s", time.Now().String())
	Info("Info log")
	Infof("Infof log, %s", time.Now().String())
	Warn("Warn log")
	Warnf("Warnf log, %s", time.Now().String())
	Error("Error log")
	Errorf("Errorf log, %s", time.Now().String())
}
