import grpc

from flask import Flask
from authlib.integrations.flask_oauth2 import ResourceProtector
from starlette.exceptions import HTTPException

from .auth import Auth0JWTBearerTokenValidator

from google.protobuf.json_format import MessageToJson

from status.v1 import status_pb2
from status.v1 import status_pb2_grpc

require_auth = ResourceProtector()
validator = Auth0JWTBearerTokenValidator(
    "mono-dev.au.auth0.com",
    "https://mono-http.md70.net"
)
require_auth.register_token_validator(validator)

app = Flask(__name__)


@app.get("/")
def index():
    channel = grpc.insecure_channel('localhost:50051')
    stub = status_pb2_grpc.StatusServiceStub(channel)
    req = status_pb2.FetchStatusRequest()
    try:
        status = stub.FetchStatus(req)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"{e}")

    return MessageToJson(status)


@app.get("/api/public")
def api_public():
    """No access token required."""
    response = (
        "Hello from a public endpoint! You don't need to be"
        " authenticated to see this."
    )
    return {"msg": response}


@app.get("/api/private")
@require_auth(None)  # this is the auth decorator, None refers to scope
def api_private():
    """A valid access token is required."""
    response = (
        "Hello from a private endpoint! You need to be"
        " authenticated to see this."
    )
    return {"msg": response}


@app.get("/api/scoped")
@require_auth("read:messages")
def api_scoped():
    """A valid access token and scope are required."""
    response = (
        "Hello from a private endpoint! You need to be"
        " authenticated and have a scope of read:messages to see"
        " this."
    )
    return {"msg": response}
