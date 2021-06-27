# Example Go Build, Test & Depoly

Example repository to show GitHub actions to build test deploy a Go application.

This example is designed to show how GitHub Actions can be used to sutomatically build and publish a Go application to a external locations (for example a S3 bucket) where it can be referenced.

The main use case for this was to allow me to create basic Go apps on an iPad (which lacks the Go tool chain) but have pre-build binaries available for my use when I arrive at work.

The `golang-build` actions themselves are held in the [actions](https://github.com/rikwatson/actions) repository.
