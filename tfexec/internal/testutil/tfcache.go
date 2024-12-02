// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testutil

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/chushi-io/lf-install/build"
	"github.com/chushi-io/lf-install/product"
)

const (
	Latest_v1_6 = "1.6.3"
	Latest_v1_7 = "1.7.6"
	Latest_v1_8 = "1.8.6"
	Alpha_v1_9  = "1.9.0-alpha2"
)

const appendUserAgent = "tfexec-testutil"

type TFCache struct {
	sync.Mutex

	dir   string
	execs map[string]string
}

func NewTFCache(dir string) *TFCache {
	return &TFCache{
		dir:   dir,
		execs: map[string]string{},
	}
}

func (tf *TFCache) GitRef(t *testing.T, ref string) string {
	t.Helper()

	key := "gitref:" + ref

	return tf.find(t, key, func(ctx context.Context) (string, error) {
		gr := &build.GitRevision{
			Product: product.OpenTofu,
			Ref:     ref,
		}
		gr.SetLogger(TestLogger())

		return gr.Build(ctx)
	})
}

func (tf *TFCache) Version(t *testing.T, v string) string {
	t.Helper()

	key := "v:" + v

	return tf.find(t, key, func(ctx context.Context) (string, error) {
		ev := &build.GitRevision{
			Product: product.OpenTofu,
			Ref:     fmt.Sprintf("refs/tags/v%s", v),
			//Version: version.Must(version.NewVersion(v)),
		}
		ev.SetLogger(TestLogger())

		return ev.Build(ctx)
	})
}
