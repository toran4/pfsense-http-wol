# HTTP to WakeOnLan thru Pfsense
REST API to trigger Wake-on-Lan thru pfSense. 

You need a [pfSense](https://www.pfsense.org/) router apable of sending Wake-On-Lan requests to devices in the network.

## Usage
Can be used to wake a device capable of WoL by a REST request. Can be used in situations when wake-on-lan can be hard to achive like when on VPN or between different VLAN or subnets.

Send WoL request by do a GET to http://{server adress}/wol?mac={mac adress}&if={interface}



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