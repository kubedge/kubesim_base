import bluetooth, time

search_time = 10
addr = None
# Device names configured on handset for bluetooth discovery
device_names = ['Siraj','Krishna','Ayush','User1', 'User2']
device_addrs = []

print("Scanning for bluetooth devices...")
while True:
    nearby_devices = bluetooth.discover_devices(duration=search_time, 
                                                flush_cache=True, 
                                                lookup_names=True)
    discovered_devices = [ name for addr, name in nearby_devices]
    if len(nearby_devices) > 0:
        print("Devices Detected {}".format(discovered_devices))
    else:
        print("No devices Detected!")

    match=False
    for addr, name in nearby_devices:
        if name in device_names:
            print("Got {} in Bluetooth coverage, " \
                  "will go sleep for sometime".format(name))
            match=True
            subprocess.call('/goclient network Bluetooth')
    if not match:
        print("Still waiting for someone from our team to " \
            "come in bluetooth coverage area")
    time.sleep(2)
    print("Re-Scanning...")

