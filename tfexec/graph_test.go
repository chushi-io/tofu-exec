// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"testing"

	"github.com/chushi-io/tofu-exec/tfexec/internal/testutil"
)

func TestGraphCmd_v1(t *testing.T) {
	td := t.TempDir()

	tf, err := NewTofu(td, tfVersion(t, testutil.Alpha_v1_9))
	if err != nil {
		t.Fatal(err)
	}

	// empty env, to avoid environ mismatch in testing
	tf.SetEnv(map[string]string{})

	t.Run("defaults", func(t *testing.T) {
		graphCmd, _ := tf.graphCmd(context.Background())

		assertCmd(t, []string{
			"graph",
		}, nil, graphCmd)
	})

	t.Run("override all defaults", func(t *testing.T) {
		graphCmd, _ := tf.graphCmd(context.Background(),
			GraphPlan("teststate"),
			DrawCycles(true),
			GraphType("output"))

		assertCmd(t, []string{
			"graph",
			"-plan=teststate",
			"-draw-cycles",
			"-type=output",
		}, nil, graphCmd)
	})
}
