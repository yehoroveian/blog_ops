package appmon

import (
	"runtime"
	"runtime/debug"

	"gitlab.com/blog/ops/pkg/metrics"
)

// Set metrics that are required for all go applications.
func init() {
	metrics.MustRegister(appBuildInfo, goDepsBuildInfo, metricsStandard)

	// Set actual metrics standard version.
	metricsStandard.WithLabelValues(metricsStandardVersion).Set(1)

	// Set goBuildInfo labels.
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range buildInfo.Deps {
			goDepsBuildInfo.WithLabelValues(dep.Path, dep.Version, dep.Sum).Set(1)
		}
	}
}

// Local constant that specifies a version of this package,
// that was used in your application.
// IMPORTANT:
//   - This version must follow semver semantics. Follow the link for more details: semver.org.
//   - This version must be updated each time you add or change any metrics that were defined before.
const metricsStandardVersion = "0.0.1"

// RegisterAppBuildInfo registers a special metric that collects version, revision (commit) and branch (git).
// Call this function somewhere in main().
func RegisterAppBuildInfo(buildEnv, gitCommit, gitBranch string) {
	appBuildInfo.WithLabelValues(buildEnv, gitCommit, gitBranch, runtime.Version()).Set(1)
}

var (
	// appBuildInfo collects information about each app build through these metric labels.
	appBuildInfo = metrics.NewGaugeVec(
		"app_build_info",
		"A required metric with a constant '1' value, labeled by app version, code revision, git branch and "+
			"version of Go, from which application was built.",
		[]string{"env", "git_branch", "git_commit", "go_version"},
	)

	// goDepsBuildInfo collects information about each go build through these metric labels.
	goDepsBuildInfo = metrics.NewGaugeVec(
		"go_build_info",
		"A required metric with a constant '1' value, labeled by path, version and checksum of each dependency.",
		[]string{"path", "version", "sum"},
	)

	// metricsStandard specifies the version of metrics standard.
	metricsStandard = metrics.NewGaugeVec(
		"app_metrics_standard",
		"A required metric with a constant '1' value, labeled by version of the metrics standard.",
		[]string{"version"},
	)
)
