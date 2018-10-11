"""The Python implementation of the GRPC kubedge client for protocol testing."""

from __future__ import print_function

import grpc

import kubedge_pb2
import kubedge_pb2_grpc


def run():
    # IP need to be changed when running from a container
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = kubedge_pb2_grpc.KubedgeStub(channel)
        response = stub.FiveGDemo(kubedge_pb2.EnodeRequest(protocol='X10000001'))
    print(response.message)


if __name__ == '__main__':
    run()

