// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"testing"

	"github.com/chushi-io/tofu-exec/tfexec/internal/testutil"
)

func TestMetadataFunctionsCmd(t *testing.T) {
	td := t.TempDir()

	tf, err := NewTofu(td, tfVersion(t, testutil.Alpha_v1_9))
	// tf, err := NewTofu(td, tfVersion(t, testutil.Latest_v1_4)) // TODO! enable after 1.4 release
	if err != nil {
		t.Fatal(err)
	}

	// empty env, to avoid environ mismatch in testing
	tf.SetEnv(map[string]string{})

	functionsCmd := tf.metadataFunctionsCmd(context.Background())

	assertCmd(t, []string{
		"metadata",
		"functions",
		"-json",
	}, nil, functionsCmd)
}
