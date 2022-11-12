# Mono GRPC Server

## Configuration

- `config.go` should specify all configuration values required to run the application
- `config_test.go` can be used to test the configuration itself (redundant?) 

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

- To start with `make` use `ARGS=...` to pass cfg file flag, eg:

```shell
make run ARGS="-cfg testdata/sample.cfg"
```



