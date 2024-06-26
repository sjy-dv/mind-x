# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import core_pb2 as v0_dot_core__pb2
from . import dataset_pb2 as v0_dot_dataset__pb2


class DatasetManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.List = channel.unary_stream(
                '/mindxv.v0.DatasetManager/List',
                request_serializer=v0_dot_dataset__pb2.ListDatasetsRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.Dataset.FromString,
                )
        self.Get = channel.unary_unary(
                '/mindxv.v0.DatasetManager/Get',
                request_serializer=v0_dot_dataset__pb2.GetDatasetRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.Dataset.FromString,
                )
        self.Create = channel.unary_unary(
                '/mindxv.v0.DatasetManager/Create',
                request_serializer=v0_dot_dataset__pb2.Dataset.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.Dataset.FromString,
                )
        self.Delete = channel.unary_unary(
                '/mindxv.v0.DatasetManager/Delete',
                request_serializer=v0_dot_core__pb2.UUIDRequest.SerializeToString,
                response_deserializer=v0_dot_core__pb2.EmptyMessage.FromString,
                )
        self.GetDatasetSize = channel.unary_unary(
                '/mindxv.v0.DatasetManager/GetDatasetSize',
                request_serializer=v0_dot_dataset__pb2.GetDatasetRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.DatasetSize.FromString,
                )


class DatasetManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def List(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDatasetSize(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DatasetManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'List': grpc.unary_stream_rpc_method_handler(
                    servicer.List,
                    request_deserializer=v0_dot_dataset__pb2.ListDatasetsRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.Dataset.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=v0_dot_dataset__pb2.GetDatasetRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.Dataset.SerializeToString,
            ),
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=v0_dot_dataset__pb2.Dataset.FromString,
                    response_serializer=v0_dot_dataset__pb2.Dataset.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=v0_dot_core__pb2.UUIDRequest.FromString,
                    response_serializer=v0_dot_core__pb2.EmptyMessage.SerializeToString,
            ),
            'GetDatasetSize': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDatasetSize,
                    request_deserializer=v0_dot_dataset__pb2.GetDatasetRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.DatasetSize.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mindxv.v0.DatasetManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DatasetManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def List(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/mindxv.v0.DatasetManager/List',
            v0_dot_dataset__pb2.ListDatasetsRequest.SerializeToString,
            v0_dot_dataset__pb2.Dataset.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DatasetManager/Get',
            v0_dot_dataset__pb2.GetDatasetRequest.SerializeToString,
            v0_dot_dataset__pb2.Dataset.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DatasetManager/Create',
            v0_dot_dataset__pb2.Dataset.SerializeToString,
            v0_dot_dataset__pb2.Dataset.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DatasetManager/Delete',
            v0_dot_core__pb2.UUIDRequest.SerializeToString,
            v0_dot_core__pb2.EmptyMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetDatasetSize(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DatasetManager/GetDatasetSize',
            v0_dot_dataset__pb2.GetDatasetRequest.SerializeToString,
            v0_dot_dataset__pb2.DatasetSize.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class DataManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Insert = channel.unary_unary(
                '/mindxv.v0.DataManager/Insert',
                request_serializer=v0_dot_dataset__pb2.InsertRequest.SerializeToString,
                response_deserializer=v0_dot_core__pb2.EmptyMessage.FromString,
                )
        self.Update = channel.unary_unary(
                '/mindxv.v0.DataManager/Update',
                request_serializer=v0_dot_dataset__pb2.UpdateRequest.SerializeToString,
                response_deserializer=v0_dot_core__pb2.EmptyMessage.FromString,
                )
        self.Remove = channel.unary_unary(
                '/mindxv.v0.DataManager/Remove',
                request_serializer=v0_dot_dataset__pb2.RemoveRequest.SerializeToString,
                response_deserializer=v0_dot_core__pb2.EmptyMessage.FromString,
                )
        self.BatchInsert = channel.unary_unary(
                '/mindxv.v0.DataManager/BatchInsert',
                request_serializer=v0_dot_dataset__pb2.BatchRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.BatchResponse.FromString,
                )
        self.BatchUpdate = channel.unary_unary(
                '/mindxv.v0.DataManager/BatchUpdate',
                request_serializer=v0_dot_dataset__pb2.BatchRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.BatchResponse.FromString,
                )
        self.BatchRemove = channel.unary_unary(
                '/mindxv.v0.DataManager/BatchRemove',
                request_serializer=v0_dot_dataset__pb2.BatchRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.BatchResponse.FromString,
                )
        self.PartitionBatchInsert = channel.unary_unary(
                '/mindxv.v0.DataManager/PartitionBatchInsert',
                request_serializer=v0_dot_dataset__pb2.PartitionBatchRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.BatchResponse.FromString,
                )
        self.PartitionBatchUpdate = channel.unary_unary(
                '/mindxv.v0.DataManager/PartitionBatchUpdate',
                request_serializer=v0_dot_dataset__pb2.PartitionBatchRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.BatchResponse.FromString,
                )
        self.PartitionBatchRemove = channel.unary_unary(
                '/mindxv.v0.DataManager/PartitionBatchRemove',
                request_serializer=v0_dot_dataset__pb2.PartitionBatchRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.BatchResponse.FromString,
                )
        self.PartitionInfo = channel.unary_unary(
                '/mindxv.v0.DataManager/PartitionInfo',
                request_serializer=v0_dot_dataset__pb2.PartitionInfoRequest.SerializeToString,
                response_deserializer=v0_dot_dataset__pb2.PartitionInfoResponse.FromString,
                )


class DataManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Insert(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Remove(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchInsert(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchUpdate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BatchRemove(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PartitionBatchInsert(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PartitionBatchUpdate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PartitionBatchRemove(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PartitionInfo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DataManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Insert': grpc.unary_unary_rpc_method_handler(
                    servicer.Insert,
                    request_deserializer=v0_dot_dataset__pb2.InsertRequest.FromString,
                    response_serializer=v0_dot_core__pb2.EmptyMessage.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=v0_dot_dataset__pb2.UpdateRequest.FromString,
                    response_serializer=v0_dot_core__pb2.EmptyMessage.SerializeToString,
            ),
            'Remove': grpc.unary_unary_rpc_method_handler(
                    servicer.Remove,
                    request_deserializer=v0_dot_dataset__pb2.RemoveRequest.FromString,
                    response_serializer=v0_dot_core__pb2.EmptyMessage.SerializeToString,
            ),
            'BatchInsert': grpc.unary_unary_rpc_method_handler(
                    servicer.BatchInsert,
                    request_deserializer=v0_dot_dataset__pb2.BatchRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.BatchResponse.SerializeToString,
            ),
            'BatchUpdate': grpc.unary_unary_rpc_method_handler(
                    servicer.BatchUpdate,
                    request_deserializer=v0_dot_dataset__pb2.BatchRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.BatchResponse.SerializeToString,
            ),
            'BatchRemove': grpc.unary_unary_rpc_method_handler(
                    servicer.BatchRemove,
                    request_deserializer=v0_dot_dataset__pb2.BatchRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.BatchResponse.SerializeToString,
            ),
            'PartitionBatchInsert': grpc.unary_unary_rpc_method_handler(
                    servicer.PartitionBatchInsert,
                    request_deserializer=v0_dot_dataset__pb2.PartitionBatchRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.BatchResponse.SerializeToString,
            ),
            'PartitionBatchUpdate': grpc.unary_unary_rpc_method_handler(
                    servicer.PartitionBatchUpdate,
                    request_deserializer=v0_dot_dataset__pb2.PartitionBatchRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.BatchResponse.SerializeToString,
            ),
            'PartitionBatchRemove': grpc.unary_unary_rpc_method_handler(
                    servicer.PartitionBatchRemove,
                    request_deserializer=v0_dot_dataset__pb2.PartitionBatchRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.BatchResponse.SerializeToString,
            ),
            'PartitionInfo': grpc.unary_unary_rpc_method_handler(
                    servicer.PartitionInfo,
                    request_deserializer=v0_dot_dataset__pb2.PartitionInfoRequest.FromString,
                    response_serializer=v0_dot_dataset__pb2.PartitionInfoResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mindxv.v0.DataManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DataManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Insert(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/Insert',
            v0_dot_dataset__pb2.InsertRequest.SerializeToString,
            v0_dot_core__pb2.EmptyMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/Update',
            v0_dot_dataset__pb2.UpdateRequest.SerializeToString,
            v0_dot_core__pb2.EmptyMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Remove(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/Remove',
            v0_dot_dataset__pb2.RemoveRequest.SerializeToString,
            v0_dot_core__pb2.EmptyMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchInsert(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/BatchInsert',
            v0_dot_dataset__pb2.BatchRequest.SerializeToString,
            v0_dot_dataset__pb2.BatchResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchUpdate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/BatchUpdate',
            v0_dot_dataset__pb2.BatchRequest.SerializeToString,
            v0_dot_dataset__pb2.BatchResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def BatchRemove(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/BatchRemove',
            v0_dot_dataset__pb2.BatchRequest.SerializeToString,
            v0_dot_dataset__pb2.BatchResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PartitionBatchInsert(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/PartitionBatchInsert',
            v0_dot_dataset__pb2.PartitionBatchRequest.SerializeToString,
            v0_dot_dataset__pb2.BatchResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PartitionBatchUpdate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/PartitionBatchUpdate',
            v0_dot_dataset__pb2.PartitionBatchRequest.SerializeToString,
            v0_dot_dataset__pb2.BatchResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PartitionBatchRemove(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/PartitionBatchRemove',
            v0_dot_dataset__pb2.PartitionBatchRequest.SerializeToString,
            v0_dot_dataset__pb2.BatchResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PartitionInfo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mindxv.v0.DataManager/PartitionInfo',
            v0_dot_dataset__pb2.PartitionInfoRequest.SerializeToString,
            v0_dot_dataset__pb2.PartitionInfoResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
