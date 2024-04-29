# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import cluster_pb2 as v0_dot_cluster__pb2
from . import core_pb2 as v0_dot_core__pb2


class NodesManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListNodes = channel.unary_stream(
                '/mindxv.v0.NodesManager/ListNodes',
                request_serializer=v0_dot_core__pb2.EmptyMessage.SerializeToString,
                response_deserializer=v0_dot_cluster__pb2.Node.FromString,
                )
        self.AddNode = channel.unary_stream(
                '/mindxv.v0.NodesManager/AddNode',
                request_serializer=v0_dot_cluster__pb2.Node.SerializeToString,
                response_deserializer=v0_dot_cluster__pb2.Node.FromString,
                )
        self.RemoveNode = channel.unary_unary(
                '/mindxv.v0.NodesManager/RemoveNode',
                request_serializer=v0_dot_cluster__pb2.Node.SerializeToString,
                response_deserializer=v0_dot_core__pb2.EmptyMessage.FromString,
                )
        self.LoadInfo = channel.unary_unary(
                '/mindxv.v0.NodesManager/LoadInfo',
                request_serializer=v0_dot_core__pb2.EmptyMessage.SerializeToString,
                response_deserializer=v0_dot_cluster__pb2.NodeLoadInfo.FromString,
                )


class NodesManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListNodes(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddNode(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveNode(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LoadInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_NodesManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListNodes': grpc.unary_stream_rpc_method_handler(
                    servicer.ListNodes,
                    request_deserializer=v0_dot_core__pb2.EmptyMessage.FromString,
                    response_serializer=v0_dot_cluster__pb2.Node.SerializeToString,
            ),
            'AddNode': grpc.unary_stream_rpc_method_handler(
                    servicer.AddNode,
                    request_deserializer=v0_dot_cluster__pb2.Node.FromString,
                    response_serializer=v0_dot_cluster__pb2.Node.SerializeToString,
            ),
            'RemoveNode': grpc.unary_unary_rpc_method_handler(
                    servicer.RemoveNode,
                    request_deserializer=v0_dot_cluster__pb2.Node.FromString,
                    response_serializer=v0_dot_core__pb2.EmptyMessage.SerializeToString,
            ),
            'LoadInfo': grpc.unary_unary_rpc_method_handler(
                    servicer.LoadInfo,
                    request_deserializer=v0_dot_core__pb2.EmptyMessage.FromString,
                    response_serializer=v0_dot_cluster__pb2.NodeLoadInfo.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mindxv.v0.NodesManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class NodesManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListNodes(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/mindxv.v0.NodesManager/ListNodes',
            v0_dot_core__pb2.EmptyMessage.SerializeToString,
            v0_dot_cluster__pb2.Node.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddNode(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/mindxv.v0.NodesManager/AddNode',
            v0_dot_cluster__pb2.Node.SerializeToString,
            v0_dot_cluster__pb2.Node.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveNode(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.NodesManager/RemoveNode',
            v0_dot_cluster__pb2.Node.SerializeToString,
            v0_dot_core__pb2.EmptyMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def LoadInfo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.NodesManager/LoadInfo',
            v0_dot_core__pb2.EmptyMessage.SerializeToString,
            v0_dot_cluster__pb2.NodeLoadInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)