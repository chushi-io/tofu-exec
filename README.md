[![PkgGoDev](https://pkg.go.dev/badge/github.com/chushi-io/tofu-exec)](https://pkg.go.dev/github.com/chushi-io/tofu-exec)

# tofu-exec

A Go module for constructing and running [Tofu](https://terraform.io) CLI commands. Structured return values use the data types defined in [terraform-json](https://github.com/hashicorp/terraform-json).

The [Tofu Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework) is the canonical Go interface (SDK) for Tofu plugins using the gRPC protocol. This library is intended for use in Go programs that make use of Tofu's other interface, the CLI. Importing this library is preferable to importing `github.com/hashicorp/terraform/command`, because the latter is not intended for use outside Tofu Core.

While tofu-exec is already widely used, please note that this module is **not yet at v1.0.0**, and that therefore breaking changes may occur in minor releases.

We strictly follow [semantic versioning](https://semver.org).

## Go compatibility

This library is built in Go, and uses the [support policy](https://golang.org/doc/devel/release.html#policy) of Go as its support policy. The two latest major releases of Go are supported by tofu-exec.

Currently, that means Go **1.18** or later must be used.

## Usage

The `Tofu` struct must be initialised with `NewTofu(workingDir, execPath)`. 

Top-level Tofu commands each have their own function, which will return either `error` or `(T, error)`, where `T` is a `terraform-json` type.


### Example


```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-version"
	"github.com/chushi-io/lf-install/product"
	"github.com/chushi-io/lf-install/releases"
	"github.com/chushi-io/tofu-exec/tfexec"
)

func main() {
	installer := &releases.ExactVersion{
		Product: product.OpenTofu,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Tofu: %s", err)
	}

	workingDir := "/path/to/working/dir"
	tf, err := tfexec.NewTofu(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTofu: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	fmt.Println(state.FormatVersion) // "0.1"
}
```

## Testing Tofu binaries

The tofu-exec test suite contains end-to-end tests which run realistic workflows against a real Tofu binary using `tfexec.Tofu{}`.

To run these tests with a local Tofu binary, set the environment variable `TFEXEC_E2ETEST_TERRAFORM_PATH` to its path and run:
```sh
go test -timeout=20m ./tfexec/internal/e2etest
```

For more information on tofu-exec's test suite, please see Contributing below.

## Contributing

Please see [CONTRIBUTING.md](./CONTRIBUTING.md).
