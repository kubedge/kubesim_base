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
    filename ='/etc/kubedge/connected_ue.yaml'
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
        update_yaml(data={'connected': discovered_devices})
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

