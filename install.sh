#!/bin/bash

apt update
apt upgrade
wget -c https://github.com/wavesplatform/Waves/releases/download/v1.3.15/waves_1.3.15_all.deb
apt install -f ./waves_1.3.15_all.deb 
mkdir /var/lib/aint
wget https://raw.githubusercontent.com/aintdigital/aint-node/main/waves.conf
wget https://raw.githubusercontent.com/aintdigital/aint-node/main/application.ini
mv waves.conf /etc/waves/waves.conf
mv application.ini /etc/waves/application.ini