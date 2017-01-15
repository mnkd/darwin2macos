package darwin2macos

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Darwin_(operating_system)
var version map[string]string

func prepare() {
	if len(version) != 0 {
		return
	}

	version = map[string]string{
		"16.3.0": "10.12.2",
		"16.1.0": "10.12.1",
		"16.0.0": "10.12.0",
		"15.6.0": "10.11.6",
		"15.0.0": "10.11.0",
	}
}

func DebugDumpVersions() {
	prepare()
	fmt.Printf("%#v\n", version)
}

func ToMacOS(darwin string) (string, bool) {
	prepare()
	macos, present := version[darwin]
	if present == false {
		macos = "Unknown"
	}
	return macos, present
}
