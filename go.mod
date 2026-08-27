module github.com/zoido/errassert

go 1.25.0

// 1.55.0 was the 1st version that supported go1.13 error wrapping.
// We want to use the lowest version possible to avoid pushing dependencies to
// clients.
// 1.56.3 is the lowest that is deemed safe by dependabot.
require google.golang.org/grpc v1.82.1

require (
	golang.org/x/sys v0.43.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260414002931-afd174a4e478 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
