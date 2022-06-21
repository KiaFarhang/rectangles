# Rectangles Problem

This repo contains solutions for the rectangles problems described in [the PDF in this directory](./Rectangles%20Programming%20Sample.pdf). There are a few implementation details to note:

- Rectangles are constructed/defined by their top-left and bottom-right corners
- The intersection solution/implementation does not count points on the border of both rectangles as points of intersection
- The adjacency solution/implementation does not count shared corners as points of adjacency

## Building and running tests

- [Install Go](https://go.dev/doc/install) on your machine if you haven't already done so
- You can run a basic demonstration of the solutions by running the main program: `go run main.go`
- The unit tests have a more robust suite of validations. To run those, you can use `go test ./...` from this root directory

## Notes

- The only third-party library used is [testify](https://github.com/stretchr/testify), a testing library that helps cut down on boilerplate in assertions