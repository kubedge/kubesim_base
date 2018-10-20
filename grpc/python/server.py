# Copyright 2018 Kubedge.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

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

