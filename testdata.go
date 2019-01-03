// Copyright (c) 2018 ef-ds
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package stackbenchtests

import "github.com/ef-ds/benchmark"

var (
	tests benchmark.Benchmark
)

// Pure slice based test stack implementation-------------------------------------------------------

// CustomSliceStack represents an unbounded, dynamically growing deque customized
// to operate on testVale struct.
type CustomSliceStack struct {
	// The queue values.
	v []*benchmark.TestValue
}

// NewCustomSliceStack returns an initialized instance of CustomSliceStack.
func NewCustomSliceStack() *CustomSliceStack {
	return new(CustomSliceStack).Init()
}

// Init initializes or clears stack s.
func (s *CustomSliceStack) Init() *CustomSliceStack {
	s.v = make([]*benchmark.TestValue, 0)
	return s
}

// Len returns the length of stack s.
func (s *CustomSliceStack) Len() int { return len(s.v) }

// Back returns element at the back of stack s.
func (s *CustomSliceStack) Back() (*benchmark.TestValue, bool) {
	if len(s.v) == 0 {
		return nil, false
	}
	return s.v[len(s.v)-1], true
}

// Push adds element v to the back of stack s.
func (s *CustomSliceStack) Push(v *benchmark.TestValue) {
	s.v = append(s.v, v)
}

// Pop removes and returns the element at the back of stack s.
func (s *CustomSliceStack) Pop() (*benchmark.TestValue, bool) {
	if len(s.v) == 0 {
		return nil, false
	}

	tp := len(s.v) - 1
	v := s.v[tp]
	s.v[tp] = nil // Avoid memory leaks
	s.v = s.v[:tp]
	return v, true
}
