# Test if your router supports UPnP
import requests
import socket

def discover_upnp():
    # Send UPnP discovery request
    msg = \
        'M-SEARCH * HTTP/1.1\r\n' \
        'HOST:239.255.255.250:1900\r\n' \
        'ST:upnp:rootdevice\r\n' \
        'MX:2\r\n' \
        'MAN:"ssdp:discover"\r\n' \
        '\r\n'
    
    # Send multicast request
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.settimeout(5)
    sock.sendto(msg.encode(), ('239.255.255.250', 1900))
    
    try:
        data, addr = sock.recvfrom(1024)
        print(f"UPnP device found at {addr}: {data.decode()}")
    except socket.timeout:
        print("No UPnP devices found")

if __name__ =="__main__":
    discover_upnp()