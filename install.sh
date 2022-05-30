#!/bin/bash

apt update
apt upgrade
wget -c https://github.com/wavesplatform/Waves/releases/download/v1.3.15/waves_1.3.15_all.deb
apt install -f ./waves_1.3.15_all.deb 
mkdir /var/lib/aint