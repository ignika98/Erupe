# Contributing an ABI config

CC only supports a set of known OS/arch pairs (defined in [abi_platforms.go](../abi_platforms.go)).

You are welcome (and highly encouraged) to contribute your ABI config to CC.

To do this, run a C helper to generate size/alignment information for basic types (for GCC):

```bash
gcc ./cmd/cabi/main.c && ./a.out
```

Copy this information to [abi_platforms.go](../abi_platforms.go), under the key generated by:

```bash
echo "{\"$(go env GOOS)\", \"$(go env GOARCH)\"}"
```

Also, please add a comment with the version of the compiler used to generate the mapping.