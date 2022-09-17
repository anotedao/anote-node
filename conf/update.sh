#!/bin/bash

sudo supervisorctl stop anote
rm anote-node
apt remove supervisor

wget -c https://github.com/wavesplatform/Waves/releases/download/v1.4.8/waves_1.4.8_all.deb
wget https://github.com/anotedigital/anote-node/releases/download/v1.2.3/anote-node

dpkg -i waves_1.4.8_all.deb
chmod +x anote-node

./anote-node -update
