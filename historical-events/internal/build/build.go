package build // import github.com/aliaksandrrachko/historian/historical-events/internal/build

import "runtime"

var (
	version   = "v0.0.1-SNAPSHOT"
	gitCommit = ""
	gitAuthor = ""
	time      = ""
)

type BuildInfo struct {
	Version   string
	GitCommit string
	GitAuthor string
	BuildTime string
	GoVersion string
}

func Get() BuildInfo {
	return BuildInfo{
		Version:   version,
		GitCommit: gitCommit,
		GitAuthor: gitAuthor,
		BuildTime: time,
		GoVersion: runtime.Version(),
	}
}
