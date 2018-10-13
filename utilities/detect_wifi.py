import os
import time
import subprocess

device_names = ['Siraj','Krishna','Ayush','RICOH', 'User2']
device_addrs = []
command = "sudo arp-scan -l | tail -n +3 | head -n -3 > /tmp/wifi_connected_devices"

print("Scanning for devices connected to WIFI...")
while True:
    #I know it is a security issue to use shell=True :)
    #getting the Popen with complex was tricky. will fix this later
    subprocess.call(command, shell=True)
    detected_devices = open('/tmp/wifi_connected_devices', "r")

    with open("/tmp/wifi_connected_devices") as tmpfile:
         data=tmpfile.readlines()

    match = False
    for name in device_names:
        for line in data:
            if name in line:
                print("Got {} in Wifi coverage".format(name))
                subprocess.call('/goclient network WIFI')
                match=True

    if not match:
        print("Still waiting for someone from our team to " \
            "come in wifi coverage area")
    time.sleep(5)
    print("Re-Scanning...")
