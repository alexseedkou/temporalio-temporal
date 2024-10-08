// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Proto-type adapters for the mock library
package protomock

import (
	"fmt"

	"go.temporal.io/api/temporalproto"
	"go.uber.org/mock/gomock"
)

type protoEq struct {
	want any
}

type protoSliceEq struct {
	want []any
}

// Eq returns a gomock.Matcher that uses proto.Equal to check equality
func Eq(want any) gomock.Matcher {
	return protoEq{want}
}

// Matches returns true when the provided value equals the expectation
func (m protoEq) Matches(x any) bool {
	if m.want == nil && x == nil {
		return true
	}

	return temporalproto.DeepEqual(m.want, x)
}

func (m protoEq) String() string {
	return fmt.Sprintf("is equal to %v (%T)", m.want, m.want)
}

func SliceEq(want []any) gomock.Matcher {
	return protoSliceEq{want}
}

func (m protoSliceEq) Matches(x any) bool {
	xs, ok := x.([]any)
	if !ok || len(m.want) != len(xs) {
		return false
	}

	for i := range m.want {
		if !temporalproto.DeepEqual(m.want[i], xs[i]) {
			return false
		}
	}
	return true
}

func (m protoSliceEq) String() string {
	return fmt.Sprintf("is equal to %v", m.want)
}
