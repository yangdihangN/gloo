#!/usr/bin/env bash

set -ex

# install docker
sudo apt-get update
sudo apt-get remove docker docker-engine docker.io -y
sudo apt install docker.io -y
sudo systemctl start docker
sudo systemctl enable docker
docker --version
sudo groupadd docker
sudo usermod -aG docker $USER

# install unzip
sudo apt install unzip -y

# install nomad
wget https://releases.hashicorp.com/nomad/0.9.4/nomad_0.9.4_linux_amd64.zip
unzip nomad_0.9.4_linux_amd64.zip
rm -f nomad_0.9.4_linux_amd64.zip
chmod +x nomad
sudo mv nomad /usr/local/bin/

# install consul
wget https://releases.hashicorp.com/consul/0.9.4/consul_0.9.4_linux_amd64.zip
unzip consul_0.9.4_linux_amd64.zip
rm -f consul_0.9.4_linux_amd64.zip
chmod +x consul
sudo mv consul /usr/local/bin/

# install vault
wget https://releases.hashicorp.com/vault/0.9.4/vault_0.9.4_linux_amd64.zip
unzip vault_0.9.4_linux_amd64.zip
rm -f vault_0.9.4_linux_amd64.zip
chmod +x vault
sudo mv vault /usr/local/bin/
