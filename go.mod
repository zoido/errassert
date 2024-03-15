module github.com/zoido/errassert

go 1.13

// 1.55.0 was the 1st version that supported go1.13 error wrapping.
// We want to use the lowest version possible to avoid pushing dependencies to
// clients.
require google.golang.org/grpc v1.55.1
