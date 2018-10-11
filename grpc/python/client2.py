"""The Python implementation of the GRPC kubedge client."""

from __future__ import print_function

import grpc
import sys

import kubedge_pb2
import kubedge_pb2_grpc


def run(conn_mode):
    # IP need to be changed when running from a container
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = kubedge_pb2_grpc.KubedgeStub(channel)
        response = stub.DetectNW(kubedge_pb2.UERequest(network=conn_mode))

    print(response.message)


if __name__ == '__main__':
    conn_mode = 'BLUETOOTH'
    if len(sys.argv) > 1:
        conn_mode = sys.argv[1]
    run(conn_mode)

