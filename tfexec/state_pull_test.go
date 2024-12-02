// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"testing"

	"github.com/chushi-io/tofu-exec/tfexec/internal/testutil"
)

func TestStatePull(t *testing.T) {
	tf, err := NewTofu(t.TempDir(), tfVersion(t, testutil.Alpha_v1_9))
	if err != nil {
		t.Fatal(err)
	}

	tf.SetEnv(map[string]string{})

	t.Run("tfstate", func(t *testing.T) {
		statePullCmd := tf.statePullCmd(context.Background(), nil)

		assertCmd(t, []string{
			"state",
			"pull",
		}, nil, statePullCmd)
	})
}
