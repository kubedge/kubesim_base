FROM scratch
ADD grpc/go/kubedge_server/goserv /goserv
ADD grpc/go/kubedge_client/goclient /goclient
EXPOSE 50051
CMD ["/goserv"]
