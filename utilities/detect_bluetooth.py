import bluetooth, time
import subprocess
import os
import yaml

search_time = 5
addr = None
# Device names configured on handset for bluetooth discovery
device_names = ['Siraj','Krishna','Ayush','User1', 'User2']
device_addrs = []

print("Scanning for bluetooth devices...")
def update_yaml(**kwargs):
    filename ='/etc/kubedge/connected.yaml'
    directory = os.path.dirname(filename)
    if not os.path.exists(directory):
         os.makedirs(directory)
    with open(filename, 'w+') as f:
        yaml_dict = yaml.load(f) or {}
    print(yaml_dict)
    yaml_dict.update(kwargs)
    with open(filename, 'w') as f:
        yaml.dump(yaml_dict, f)

while True:
    nearby_devices = bluetooth.discover_devices(duration=search_time,
                                                flush_cache=True,
                                                lookup_names=True)
    discovered_devices = [ name for addr, name in nearby_devices]
    if len(nearby_devices) > 0:
        print("Devices Detected {}".format(discovered_devices))
        update_yaml(data={'connected': matched_users})
    else:
        print("No devices Detected!")

    match=False
    matched_users = []
    for addr, name in nearby_devices:
        if name in device_names:
            print("Got {} in Bluetooth coverage, " \
                  "will go sleep for sometime".format(name))
            match=True
            matched_users.append(name)
    #update_yaml(data={'connected': matched_users})
    if not match:
        print("Still waiting for someone from our team to " \
            "come in bluetooth coverage area")
    time.sleep(2)
    print("Re-Scanning...")

