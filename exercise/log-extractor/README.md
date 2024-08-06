# Nix package build

### Build the Package
To build the package, run the following command in the logparser directory:

```sh
nix-build
```
This command will create a result directory containing the built package.


### Modify default.nix with GOCACHE
Go build process is trying to use a default cache location (/homeless-shelter/Library/Caches/go-build) that it cannot write to because the file system is read-only. This is a common issue when using Go in restricted or sandboxed environments like Nix.

To resolve this, you need to configure Go to use a different cache directory. You can set the GOCACHE environment variable to a writable directory within the Nix build environment.

### Run the Executable
To run the built executable, use the following command:

```s
./result/bin/logparser
```
