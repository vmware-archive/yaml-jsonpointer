// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

/*
Package splice allows to perform simple edits on a string, byte buffer or a file.

It allows to delete, insert or replace strings in a text buffer.

The core operation is: replace the current content of a given selection with a new string.
Deletion is just replacement with an empty string.
Insertion is just replacement at a zero length selection.

Selections are addressed by unicode character offsets, not byte offsets!.

The edit operation involves one single pass through the input.
*/
package splice

import (
	"io"

	"github.com/vmware-labs/go-yaml-edit/splice"
)

// T constructs a splice transformer given one or more operations.
// A splice transformer implements golang.org/x/text/transform.Transformer;
// that package contains many useful functions to apply the transformation.
func T(ops ...Op) *Transformer { return splice.T(ops...) }

type Transformer = splice.Transformer

// A Op captures a request to replace a selection with a replacement string.
// An idiomatic way to construct an Op instance is to call With or WithFunc on a Selection.
type Op = splice.Op

// A Selection selects a range of characters in the input string buffer.
// It's defined to be the range that starts at Start end ends before the End position.
// Positions are  unicode codepoint offsets, not byte offsets.
type Selection = splice.Selection

// Span constructs a Selection.
func Span(start, end int) Selection { return splice.Span(start, end) }

// Peek returns a slice of strings for each extent of the input reader.
// The order of the resulting slice matches the order of the provided selection slice
// (which can be in any order; slice provides the necessary sorting to guarantee a single
// scan pass on the reader).
func Peek(r io.Reader, sels ...Selection) ([]string, error) {
	return splice.Peek(r, sels...)
}
