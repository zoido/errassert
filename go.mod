module github.com/zoido/errassert

go 1.24.0

// 1.55.0 was the 1st version that supported go1.13 error wrapping.
// We want to use the lowest version possible to avoid pushing dependencies to
// clients.
// 1.56.3 is the lowest that is deemed safe by dependabot.
require google.golang.org/grpc v1.79.3

require (
	golang.org/x/sys v0.39.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
