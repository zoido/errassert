module github.com/zoido/errassert/grpcerrassert

go 1.25.0

require (
	github.com/zoido/errassert v1.0.0
	google.golang.org/grpc v1.83.2
)

require (
	golang.org/x/sys v0.47.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

// The root module is always developed and tested together with this module.
// Consumers ignore this directive and use the required version instead.
replace github.com/zoido/errassert => ../
