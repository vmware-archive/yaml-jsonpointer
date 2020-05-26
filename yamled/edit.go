// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package yamled

import (
	yamled "github.com/vmware-labs/go-yaml-edit"
	"github.com/vmware-labs/go-yaml-edit/splice"
	"gopkg.in/yaml.v3"
)

// Node returns a selection that spans over a YAML node.
func Node(n *yaml.Node) splice.Selection { return yamled.Node(n) }

// T creates a transformer that performs YAML-aware edit operations.
func T(ops ...splice.Op) *Transformer { return yamled.T(ops...) }

// A Transformer implements golang.org/x/text/transform.Transformer and can be used to perform
// precise in-place edits of yaml nodes in an byte stream.
type Transformer = yamled.Transformer
