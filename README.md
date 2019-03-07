# Static IPAM Driver
### IP address non-management for libnetwork

Static overlapping subnets are not a supported feature of the standard libnetwork IPAM driver. This
driver is designed to do as little management as possible, giving the user full control over the IP
addresses assigned to their containers. This is crucial when created a number of docker networks
seperate on layer 2 and you want overlapping layer 3 (IP) domains.

## Installation as a service with SysV (Debian/Ubuntu)
```bash
# Download the service script and install it to init.d
sudo curl -L https://raw.githubusercontent.com/nategraf/static-ipam-driver/master/sysv.sh -o /etc/init.d/static-ipam
sudo chmod +x /etc/init.d/static-ipam

# Download the driver to usr/local/bin
sudo curl -L https://github.com/nategraf/static-ipam-driver/releases/latest/download/static-ipam-driver.linux.amd64 -o /usr/local/bin/static-ipam
sudo chmod +x /usr/local/bin/static-ipam

# Activate the service
sudo update-rc.d static-ipam defaults
sudo service static-ipam start

# Verify that it is running
sudo stat /run/docker/plugins/static.sock
#  File: /run/docker/plugins/static.sock
#  Size: 0               Blocks: 0          IO Block: 4096   socket
#  ...
```
