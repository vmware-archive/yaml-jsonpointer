// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package atomicfile

import (
	"io"
	"os"

	"github.com/vmware-labs/go-yaml-edit/atomicfile"
	"golang.org/x/text/transform"
)

// Writer returns a io.WriteCloser that writes data to a temporary file
// which gets renamed atomically as filename upon Commit.
func Writer(filename string, perm os.FileMode) (*AtomicWriter, error) {
	return atomicfile.Writer(filename, perm)
}

// An AtomicWriter is a writer that atomically writes into a file once the Commit method is called.
type AtomicWriter = atomicfile.AtomicWriter

// WriteFrom copies data from a read into a destination file identified by filename.
// If the file already exists, it's replaced atomically with the new content and the
// original file permissions are preserved.
func WriteFrom(filename string, r io.Reader, perm os.FileMode) error {
	return atomicfile.WriteFrom(filename, r, perm)
}

// WriteFile is a drop-in replacement for ioutil.WriteFile that writes the file atomically.
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	return atomicfile.WriteFile(filename, data, perm)
}

// Transform reads the content of an existing file, passes it through a transformer and writes it back atomically.
func Transform(t transform.Transformer, filename string) error {
	return atomicfile.Transform(t, filename)
}
