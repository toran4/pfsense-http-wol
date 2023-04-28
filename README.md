# HTTP to WakeOnLan thru pfSense

This project provides a REST API that allows you to trigger Wake-on-Lan (WoL) requests through pfSense, a popular open-source router/firewall platform. With this API, you can wake up devices on your network remotely, even if they are on a different VLAN or subnet. 

## Background

Wake-on-Lan is a network protocol that allows you to wake up a computer or other device from a sleep or hibernate state by sending a special "magic packet" over the network. While Wake-on-Lan is a useful feature, it can be hard to achieve in certain network configurations, such as when you are on a VPN or when devices are on different subnets. This project makes it easy to wake up devices on your network, regardless of their location.

You need a [pfSense](https://www.pfsense.org/) router capable of sending Wake-On-Lan requests to devices in the network.

## Usage
Can be used to wake a device capable of WoL by a REST request. Can be used in situations when wake-on-lan can be hard to achive like when on VPN or between different VLAN or subnets.

Send WoL request by do a GET to http://{server address}/wol?mac={mac address}&if={interface} where mac is the mac address of the computer to wake and if s the interface in pfSense that the cmputers network is on.

Example to wake a computer with mac 34:31:4d:69:39:03 on interface opt4  when pfSense is runnon on ip 192.168.0.1: ```https://192.168.0.1/wol?mac=34:31:4d:69:39:03&if=opt4```



## Installation
Run in docker by pulling [toran4/pfsense-http-wol](https://hub.docker.com/r/toran4/pfsense-http-wol). 

Example when running using docker-compose:
```
---
version: '3'
services:
  http_wol:
    image: toran4/pfsense-http-wol:latest
    ports:
      - '8080:8080'
    restart: always
    environment:
      - CONFIG_DIR=/config/
    volumes:
      - /opt/http_wol:/config
```

### Configuration  
Configure by placing an file named ```.env``` in the ```CONFIG_DIR```. In the config file you set Pfsense URL, username and password

Example:
```
PFSENSE_URL=https://192.168.0.1
PFSENSE_USER=wol_user
PFSENSE_PASSWORD=wol_password
```
