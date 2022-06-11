#!/bin/bash

apt update
apt upgrade
apt install supervisor
wget -c https://github.com/wavesplatform/Waves/releases/download/v1.4.6/waves_1.4.6_all.deb
apt install -f ./waves_1.4.6_all.deb
mkdir /var/lib/anote
chown -R waves:waves /var/lib/anote/
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/waves.conf
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/application.ini
mv waves.conf /etc/waves/waves.conf
mv application.ini /etc/waves/application.ini
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/anote.conf
mv anote.conf /etc/supervisor/conf.d/