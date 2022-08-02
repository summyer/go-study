package util2_test

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestIO(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.Println("hello world")
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.") // Calls os.Exit(1) after logginglog.Fatal("Bye.")// Calls panic() after logginglog.Panic("I'm bailing.")

}
