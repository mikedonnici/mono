import grpc
from fastapi import FastAPI, HTTPException

import statusv1.status_pb2
import statusv1.status_pb2_grpc

app = FastAPI()


@app.get("/")
async def root():
    channel = grpc.insecure_channel('localhost:50051')
    stub = statusv1.status_pb2_grpc.StatusServiceStub(channel)
    req = statusv1.status_pb2.FetchStatusRequest()
    try:
        res = stub.FetchStatus(req)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"{e}")

    return {"msg": res.text}
