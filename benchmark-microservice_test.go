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

import (
	"container/list"
	"testing"

	"github.com/ef-ds/benchmark"
	"github.com/ef-ds/deque"
	"github.com/ef-ds/stack"
	cookiejar "gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

func BenchmarkMicroserviceList(b *testing.B) {
	var l *list.List
	tests.Microservice(
		b,
		func() {
			l = list.New()
		},
		func(v interface{}) {
			l.PushBack(v)
		},
		func() (interface{}, bool) {
			return l.Remove(l.Back()), true
		},
		func() bool {
			return l.Front() == nil
		},
	)
}

func BenchmarkMicroserviceSlice(b *testing.B) {
	var q *CustomSliceStack
	tests.Microservice(
		b,
		func() {
			q = NewCustomSliceStack()
		},
		func(v interface{}) {
			q.Push(v.(*benchmark.TestValue))
		},
		func() (interface{}, bool) {
			return q.Pop()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkMicroserviceCookiejar(b *testing.B) {
	var q *cookiejar.Stack
	tests.Microservice(
		b,
		func() {
			q = cookiejar.New()
		},
		func(v interface{}) {
			q.Push(v)
		},
		func() (interface{}, bool) {
			return q.Pop(), true
		},
		func() bool {
			return q.Size() == 0
		},
	)
}

func BenchmarkMicroserviceDeque(b *testing.B) {
	var q *deque.Deque
	tests.Microservice(
		b,
		func() {
			q = deque.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkMicroserviceStack(b *testing.B) {
	var q *stack.Stack
	tests.Microservice(
		b,
		func() {
			q = stack.New()
		},
		func(v interface{}) {
			q.Push(v)
		},
		func() (interface{}, bool) {
			return q.Pop()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}
