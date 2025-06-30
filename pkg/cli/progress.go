// file: pkg/cli/progress.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174001

package cli

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// ProgressBar displays a simple terminal progress bar
type ProgressBar struct {
	mu       sync.Mutex
	current  int
	total    int
	prefix   string
	width    int
	lastDraw time.Time
}

// NewProgressBar creates a new progress bar with the given total and prefix
func NewProgressBar(total int, prefix string) *ProgressBar {
	return &ProgressBar{
		current: 0,
		total:   total,
		prefix:  prefix,
		width:   40,
	}
}

// Update increments the progress and redraws the bar
func (p *ProgressBar) Update(file string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.current++
	
	// Throttle updates to avoid excessive drawing
	if time.Since(p.lastDraw) < 100*time.Millisecond && p.current < p.total {
		return
	}
	p.lastDraw = time.Now()
	
	p.draw(file)
}

// Finish completes the progress bar
func (p *ProgressBar) Finish() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.current = p.total
	p.draw("")
	fmt.Println() // Final newline
}

// draw renders the progress bar
func (p *ProgressBar) draw(file string) {
	if p.total == 0 {
		return
	}
	
	percent := float64(p.current) / float64(p.total)
	filled := int(percent * float64(p.width))
	
	bar := strings.Repeat("█", filled) + strings.Repeat("░", p.width-filled)
	
	// Truncate filename if too long
	displayFile := file
	maxFileLen := 30
	if len(displayFile) > maxFileLen {
		displayFile = "..." + displayFile[len(displayFile)-maxFileLen+3:]
	}
	
	// Clear line and redraw
	fmt.Fprintf(os.Stderr, "\r%s [%s] %d/%d (%.1f%%) %s", 
		p.prefix, bar, p.current, p.total, percent*100, displayFile)
}

// FileCounter tracks files processed for progress display
type FileCounter struct {
	mu      sync.Mutex
	files   []string
	current int
}

// NewFileCounter creates a new file counter
func NewFileCounter() *FileCounter {
	return &FileCounter{
		files: make([]string, 0),
	}
}

// Add adds a file to be processed
func (f *FileCounter) Add(file string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.files = append(f.files, file)
}

// Total returns the total number of files
func (f *FileCounter) Total() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.files)
}

// SetTotal sets the total number of files (for unknown totals)
func (f *FileCounter) SetTotal(total int) {
	f.mu.Lock()
	defer f.mu.Unlock()
	// Expand or shrink the slice as needed
	if len(f.files) < total {
		for len(f.files) < total {
			f.files = append(f.files, "")
		}
	} else if len(f.files) > total {
		f.files = f.files[:total]
	}
}