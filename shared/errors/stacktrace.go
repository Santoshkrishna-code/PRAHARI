package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// Frame represents a single program caller frame.
type Frame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

func (f Frame) String() string {
	return fmt.Sprintf("%s (%s:%d)", f.Function, f.File, f.Line)
}

// StackTrace is a slice of call frames.
type StackTrace []Frame

func (s StackTrace) String() string {
	var sb strings.Builder
	for i, frame := range s {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(frame.String())
	}
	return sb.String()
}

// CaptureStackTrace records program call frames skipping designated lead frames.
func CaptureStackTrace(skip int) StackTrace {
	const maxDepth = 32
	pcs := make([]uintptr, maxDepth)
	
	// Skip runtime.Callers, CaptureStackTrace, plus whatever caller requests
	n := runtime.Callers(skip+2, pcs)
	if n == 0 {
		return nil
	}
	
	pcs = pcs[:n]
	frames := runtime.CallersFrames(pcs)
	
	var trace StackTrace
	for {
		frame, more := frames.Next()
		// Filter out standard go runtime bootstrap frames
		if !strings.Contains(frame.Function, "runtime.") {
			trace = append(trace, Frame{
				Function: frame.Function,
				File:     filterFilePath(frame.File),
				Line:     frame.Line,
			})
		}
		if !more {
			break
		}
	}
	
	return trace
}

func filterFilePath(path string) string {
	idx := strings.LastIndex(path, "/shared/")
	if idx != -1 {
		return path[idx+1:]
	}
	idx = strings.LastIndex(path, "/services/")
	if idx != -1 {
		return path[idx+1:]
	}
	
	// Fallback to simple basename
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		return strings.Join(parts[len(parts)-2:], "/")
	}
	return path
}
