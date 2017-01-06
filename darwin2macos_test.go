package darwin2macos

import (
	"testing"
)

func present(t *testing.T, darwin string, macos string) {
    str, present := ToMacOS(darwin)
    if present == false {
        t.Error(darwin, "is not present.")
    }
	if str != macos {
		t.Error(darwin, "is not equal to", macos)
	}
}

func unknwon(t *testing.T, darwin string) {
    str, present := ToMacOS(darwin)
    if present == true {
        t.Error(darwin, "is present.")
    }
	if str != "Unknown" {
		t.Error(darwin, "is not equal to \"Unknown\"")
	}
}

func TestToMacOS(t *testing.T) {
	present(t, "16.3.0", "10.12.2")
    present(t, "16.1.0", "10.12.1")
    present(t, "16.0.0", "10.12.0")

    unknwon(t, "0.0.0")
}

func TestDebugDump(t *testing.T) {
    DebugDumpVersions()
}
