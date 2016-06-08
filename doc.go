// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// See https://github.com/golang/go/blob/master/CONTRIBUTORS
// Licensed under the same terms as Go itself:
// https://github.com/golang/go/blob/master/LICENSE

// Package subtest provides a backwards-compatible way to run subtests.
//
// For Go 1.7 and higher, it uses testing.T's run method. For lower versions
// it mimics subtests by logging additional information.
package subtest
