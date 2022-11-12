# DEV 

Adding things to Go:

- Create proto defs
- Create the internal pkg folder: `go/internal/[pkg-name]`
- Generate pkg code: `make pb-go`

Then:

- Add base file: `internal/[pkg-name]/pkg-name.go`
  - Internal representation of proto messages as structs
  - Message methods for each that maps the internal struct to its protobuf equivalent  

- Add manager file: `internal/[pkg-name]/manager.go` that attaches the appropriate data store


- Add pkg data manager type to server type 
- Add manage value to server value in main and register with grpc server
- Add a file in root of grpc-server named after package - this contains implementations 

