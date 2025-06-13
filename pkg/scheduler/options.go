// file: pkg/scheduler/options.go
package scheduler

import "time"

// Options defines settings for scheduled execution.
type Options struct {
	// Interval specifies time between runs.
	Interval time.Duration
	// Jitter adds a random delay up to this duration before each run.
	Jitter time.Duration
	// SkipFirst prevents the task from running immediately when true.
	SkipFirst bool
	// MaxRuns stops the scheduler after this many executions. 0 means unlimited.
	MaxRuns int
}
