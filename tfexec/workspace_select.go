// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import "context"

// WorkspaceSelect represents the workspace select subcommand to the Tofu CLI.
func (tf *Tofu) WorkspaceSelect(ctx context.Context, workspace string) error {
	// TODO: [DIR] param option

	return tf.runTofuCmd(ctx, tf.buildTofuCmd(ctx, nil, "workspace", "select", "-no-color", workspace))
}
