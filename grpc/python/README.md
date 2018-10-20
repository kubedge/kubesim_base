# Kubesim base python code

## Python-GRPC Connection

```bash
sudo apt-get install python3
sudo python3 -m pip install grpcio
sudo python3 -m pip install grpcio-tools googleapis-common-protos
```

From python directory run this command to generate the srevere and client grpc code

```bash
python3 -m grpc_tools.protoc -I../proto --python_out=. --grpc_python_out=. ../proto/kubedge.proto
```

## Start the Server:

```bash
python3 server.py
```

## Start the Client

```bash
python3 client1.py
```

## To detect UE network

```bash
python3 client2.py
```
