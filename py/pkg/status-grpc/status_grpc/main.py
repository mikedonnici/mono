import grpc

from status_pb2_grpc import StatusServiceStub
from status_pb2 import StatusRequest

channel = grpc.insecure_channel('localhost:50051')
stub = StatusServiceStub(channel)
req = StatusRequest()
try:
    res = stub.FetchStatus(req)
except Exception as e:
    raise ValueError(e)

print(res)
