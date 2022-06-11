#!/bin/bash

apt update
apt upgrade
apt install supervisor
wget -c https://github.com/wavesplatform/Waves/releases/download/v1.3.15/waves_1.3.15_all.deb
apt install -f ./waves_1.3.15_all.deb 
mkdir /var/lib/anote
chown -R waves:waves /var/lib/anote/
wget https://raw.githubusercontent.com/anotedigital/anote-node/main/waves.conf
wget https://raw.githubusercontent.com/anotedigital/anote-node/main/application.ini
mv waves.conf /etc/waves/waves.conf
mv application.ini /etc/waves/application.ini