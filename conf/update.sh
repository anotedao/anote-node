#!/bin/bash

# Get the files
# wget https://github.com/anotedigital/anote-node/raw/main/conf/waves.conf
wget https://github.com/wavesplatform/Waves/releases/download/v1.4.20/waves_1.4.20_all.deb

# Stop waves
sudo service waves stop

# Remove old blockchain
sudo rm -rf /var/lib/anote/*

# Install new waves node
sudo dpkg -i waves_1.4.20_all.deb

# Configure new blockchain
# chmod +x anote-node
# mv seed secrets
# source ./secrets
# sed -i "s/D5u2FjJFcdit5di1fYy658ufnuzPULXRYG1YNVq68AH5/$ENCODED/g" waves.conf
# sed -i "s/DTMZNMkjDzCwxNE1QLomcp5sXEQ9A3Mdb2RziN41BrYA/$KEYENCODED/g" waves.conf
# sed -i "s/127.0.0.1:/$PUBLICIP:/g" waves.conf
# mv waves.conf /etc/waves/waves.conf
sed -i "s/17]/17, 18, 19, 20]/g" /etc/waves/waves.conf
sed -i "s/min-block-time = 5s/min-block-time = 5s\r\n        dao-address = \"3AVTze8bR1SqqMKv3uLedrnqCuWpdU7GZwX\"\r\n        xtn-buyback-address = \"3AVTze8bR1SqqMKv3uLedrnqCuWpdU7GZwX\"/g" /etc/waves/waves.conf
sed -i "s/desired = 100000000/desired = 100000000\r\n        term-after-capped-reward-feature = 100000/g" /etc/waves/waves.conf

# Stop waves
sudo service waves start