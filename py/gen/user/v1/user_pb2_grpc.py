# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from user.v1 import user_pb2 as user_dot_v1_dot_user__pb2


class UserServiceStub(object):
    """UserService defines a set of services for managing user data.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FetchUser = channel.unary_unary(
                '/user.v1.UserService/FetchUser',
                request_serializer=user_dot_v1_dot_user__pb2.FetchUserRequest.SerializeToString,
                response_deserializer=user_dot_v1_dot_user__pb2.FetchUserResponse.FromString,
                )


class UserServiceServicer(object):
    """UserService defines a set of services for managing user data.
    """

    def FetchUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FetchUser': grpc.unary_unary_rpc_method_handler(
                    servicer.FetchUser,
                    request_deserializer=user_dot_v1_dot_user__pb2.FetchUserRequest.FromString,
                    response_serializer=user_dot_v1_dot_user__pb2.FetchUserResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'user.v1.UserService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class UserService(object):
    """UserService defines a set of services for managing user data.
    """

    @staticmethod
    def FetchUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user.v1.UserService/FetchUser',
            user_dot_v1_dot_user__pb2.FetchUserRequest.SerializeToString,
            user_dot_v1_dot_user__pb2.FetchUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
