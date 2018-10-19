FROM amd64/python 
RUN apt-get update && apt-get install -y \
    bluetooth \
    bluez \
    bluez-tools \
    python-bluez
ADD utilities/detect_bluetooth.py /detect_bluetooth.py
EXPOSE 50051
CMD ["python","detect_bluetooth.py"]
