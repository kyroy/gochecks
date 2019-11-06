package gotest

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"
)

type Result struct {
	Package string `json:"package"`
	Test    string `json:"test,omitempty"`
	// Result may be one of:
	//     - fail
	//     - skip
	//     - pass
	Result   string        `json:"result,omitempty"`
	Output   []string      `json:"output"`
	Elapsed  time.Duration `json:"elapsed"`
	Line     string        `json:"line,omitempty"`
	Coverage float64       `json:"coverage"`
}

func (r Result) String() string {
	b, err := json.Marshal(&r)
	if err != nil {
		return fmt.Sprintf("Package: %s, Test: %s, Result: %s, ...", r.Package, r.Test, r.Result)
	}
	return string(b)
}

type Results []*Result

func (r Results) Find(pkg, test string) (*Result, bool) {
	for _, res := range r {
		if res.Package == pkg && res.Test == test {
			return res, true
		}
	}
	return nil, false
}

func (r *Results) addOrUpdate(evt TestEvent) {
	if evt.Package == "" {
		return
	}
	res, ok := r.Find(evt.Package, evt.Test)
	if !ok {
		res = &Result{
			Package: evt.Package,
			Test:    evt.Test,
			Output:  []string{},
		}
		*r = append(*r, res)
	}
	switch evt.Action {
	case "output":
		res.Output = append(res.Output, evt.Output)
	case "fail":
		res.Line = findLine(res)
		fallthrough
	case "skip", "pass":
		res.Elapsed = time.Duration(evt.Elapsed*1000.) * time.Millisecond
		res.Result = evt.Action
		res.Coverage = findCoverage(res)
	}
}

var findRegex = regexp.MustCompile(`^\s*(.+_test.go:\d+):`)

func findLine(res *Result) string {
	for _, o := range res.Output {
		matches := findRegex.FindStringSubmatch(o)
		if len(matches) < 2 {
			continue
		}
		return matches[1]
	}
	return ""
}

var coverageRegex = regexp.MustCompile(`^coverage: (\d{1,3}\.\d+)\% of statements\s*$`)

func findCoverage(res *Result) float64 {
	for _, o := range res.Output {
		matches := coverageRegex.FindStringSubmatch(o)
		if len(matches) < 2 {
			continue
		}
		f, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			continue
		}
		return f / 100
	}
	return 0
}

func Parse(input io.Reader) (Results, error) {
	x := Results([]*Result{})
	results := &x
	dec := json.NewDecoder(input)
	for dec.More() {
		var evt TestEvent
		if err := dec.Decode(&evt); err != nil {
			return *results, err
		}
		results.addOrUpdate(evt)
	}
	return *results, nil
}
