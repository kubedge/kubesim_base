sudo apt-get install python3
sudo python3 -m pip install grpcio
sudo python3 -m pip install grpcio-tools googleapis-common-protos

From python directory run this command to generate the srevere and client grpc code
python3 -m grpc_tools.protoc -I../proto --python_out=. --grpc_python_out=. ../proto/kubedge.proto

Start Server:
python3 server.py

Start Client:
python3 client1.py => to test protocol
python3 client2.py => to detect UE network

