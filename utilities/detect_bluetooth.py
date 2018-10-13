import bluetooth, time

print("Bluetooth Detection Demo!")

search_time = 10
addr = None
device_names = ['Siraj','Krishna','Aysuh','User1', 'User2']

while True:
    print("Searching for devices...")
    nearby_devices = bluetooth.discover_devices(duration=search_time, 
                                                flush_cache=True, 
                                                lookup_names=True)
    discovered_devices = [ name for addr, name in nearby_devices]
    if len(nearby_devices) > 0:
        print("Devices Detected {}".format(discovered_devices))
    else:
        print("No devices Detected!")

    for addr, name in nearby_devices:
        if name in device_names:
            print("Got {} in Boluetooth coverage, " \
                  "will go sleep for sometime".format(name))
            time.sleep(20)
    print("Still waiting for someone from our team to " \
          "come in coverage area")
    time.sleep(1)

