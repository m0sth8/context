package context

import (
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
)

const testMsg = "Test message"

func TestLogger(t *testing.T) {

	memLog := &log.Logger{
		Handler: memory.New(),
	}
	memLog.Debug(testMsg)
	ctx := WithLogger(Background(), memLog)
	rmemLog := GetLogger(ctx)

	if rmemLog != memLog {
		t.Fatalf("Expected %+v logger, but got %+v", memLog, rmemLog)
	}

	actualHandler, ok := memLog.Handler.(*memory.Handler)

	if !ok {
		t.Fatalf("Unexpected handler %+v", memLog.Handler)
	}

	if len(actualHandler.Entries) != 1 {
		t.Fatalf("Expected only one entry, but got %d", len(actualHandler.Entries))
	}

	if actualHandler.Entries[0].Message != testMsg {
		t.Fatalf(`Wrong message in entry, expected "%s" got "%s"`, testMsg, actualHandler.Entries[0].Message)
	}
}
