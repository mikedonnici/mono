# Mono GRPC Server

## Configuration

- Configuration can be read from env vars or from a config file
- To specify a config file pass `-cfg /path/to/file`
- Sample config file:

```env
MYSQL_DSN=user:pass@tcp(host:port)/dbname # required
GRPC_PORT=50052 # optional, defaults to 50051
SERVICE_NAME="GRPC Service" # optional and pointless
```

- Config values from env have precedence and should have `MONO_` prefix:

```shell
export MONO_MYSQL_DSN=user:pass@tcp(host:port)/dbname 
export MONO_GRPC_PORT=50053
export MONO_SERVICE_NAME="My Service"
```

## Running

- Start with make (ARGS are used to pass cfg file path):

```shell
make run ARGS="-cfg sample.cfg"
```



