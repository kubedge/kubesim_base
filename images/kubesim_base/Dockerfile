#FROM arm32v7/python 
FROM don41382/rpi-python3-with-bluetooth
RUN apt-get update && apt-get install -y \
    bluetooth \
    bluez \
    bluez-tools \
    python-bluez
ADD grpc/go/kubedge_server/goserv /goserv
ADD grpc/go/kubedge_client/goclient /goclient
ADD utilities/detect_bluetooth.py /detect_bluetooth.py
EXPOSE 50051
CMD ["python","detect_bluetooth.py"]
