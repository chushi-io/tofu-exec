// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"testing"

	"github.com/chushi-io/tofu-exec/tfexec/internal/testutil"
)

func TestOutputCmd(t *testing.T) {
	td := t.TempDir()

	tf, err := NewTofu(td, tfVersion(t, testutil.Alpha_v1_9))
	if err != nil {
		t.Fatal(err)
	}

	// empty env, to avoid environ mismatch in testing
	tf.SetEnv(map[string]string{})

	t.Run("defaults", func(t *testing.T) {
		outputCmd := tf.outputCmd(context.Background())

		assertCmd(t, []string{
			"output",
			"-no-color",
			"-json",
		}, nil, outputCmd)
	})

	t.Run("override all defaults", func(t *testing.T) {
		outputCmd := tf.outputCmd(context.Background(),
			State("teststate"))

		assertCmd(t, []string{
			"output",
			"-no-color",
			"-json",
			"-state=teststate",
		}, nil, outputCmd)
	})
}
