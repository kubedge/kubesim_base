"""The Python implementation of the GRPC helloworld.Greeter server."""

from concurrent import futures
import time

import grpc

import kubedge_pb2
import kubedge_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class Kubedge(kubedge_pb2_grpc.KubedgeServicer):

    def FiveGDemo(self, request, context):
        return kubedge_pb2.EPCReply(message='%s:: Recieved data from Enode' % request.protocol)

    def DetectNW(self, request, context):
        return kubedge_pb2.EPCReply(message='Got connected to a Device through : %s' % request.network)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    kubedge_pb2_grpc.add_KubedgeServicer_to_server(Kubedge(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()

