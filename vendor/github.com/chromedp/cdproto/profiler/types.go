package profiler

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/chromedp/cdproto/runtime"
)

// ProfileNode profile node. Holds callsite information, execution statistics
// and child nodes.
type ProfileNode struct {
	ID            int64               `json:"id"`                      // Unique id of the node.
	CallFrame     *runtime.CallFrame  `json:"callFrame"`               // Function location.
	HitCount      int64               `json:"hitCount,omitempty"`      // Number of samples where this node was on top of the call stack.
	Children      []int64             `json:"children,omitempty"`      // Child node ids.
	DeoptReason   string              `json:"deoptReason,omitempty"`   // The reason of being not optimized. The function may be deoptimized or marked as don't optimize.
	PositionTicks []*PositionTickInfo `json:"positionTicks,omitempty"` // An array of source position ticks.
}

// Profile Profile.
type Profile struct {
	Nodes      []*ProfileNode `json:"nodes"`                // The list of profile nodes. First item is the root node.
	StartTime  float64        `json:"startTime"`            // Profiling start timestamp in microseconds.
	EndTime    float64        `json:"endTime"`              // Profiling end timestamp in microseconds.
	Samples    []int64        `json:"samples,omitempty"`    // Ids of samples top nodes.
	TimeDeltas []int64        `json:"timeDeltas,omitempty"` // Time intervals between adjacent samples in microseconds. The first delta is relative to the profile startTime.
}

// PositionTickInfo specifies a number of samples attributed to a certain
// source position.
type PositionTickInfo struct {
	Line  int64 `json:"line"`  // Source line number (1-based).
	Ticks int64 `json:"ticks"` // Number of samples attributed to the source line.
}

// CoverageRange coverage data for a source range.
type CoverageRange struct {
	StartOffset int64 `json:"startOffset"` // JavaScript script source offset for the range start.
	EndOffset   int64 `json:"endOffset"`   // JavaScript script source offset for the range end.
	Count       int64 `json:"count"`       // Collected execution count of the source range.
}

// FunctionCoverage coverage data for a JavaScript function.
type FunctionCoverage struct {
	FunctionName    string           `json:"functionName"`    // JavaScript function name.
	Ranges          []*CoverageRange `json:"ranges"`          // Source ranges inside the function with coverage data.
	IsBlockCoverage bool             `json:"isBlockCoverage"` // Whether coverage data for this function has block granularity.
}

// ScriptCoverage coverage data for a JavaScript script.
type ScriptCoverage struct {
	ScriptID  runtime.ScriptID    `json:"scriptId"`  // JavaScript script id.
	URL       string              `json:"url"`       // JavaScript script name or url.
	Functions []*FunctionCoverage `json:"functions"` // Functions contained in the script that has coverage data.
}

// TypeObject describes a type collected during runtime.
type TypeObject struct {
	Name string `json:"name"` // Name of a type collected with type profiling.
}

// TypeProfileEntry source offset and types for a parameter or return value.
type TypeProfileEntry struct {
	Offset int64         `json:"offset"` // Source offset of the parameter or end of function for return values.
	Types  []*TypeObject `json:"types"`  // The types for this parameter or return value.
}

// ScriptTypeProfile type profile data collected during runtime for a
// JavaScript script.
type ScriptTypeProfile struct {
	ScriptID runtime.ScriptID    `json:"scriptId"` // JavaScript script id.
	URL      string              `json:"url"`      // JavaScript script name or url.
	Entries  []*TypeProfileEntry `json:"entries"`  // Type profile entries for parameters and return values of the functions in the script.
}
