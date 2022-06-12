#!/bin/bash

while fuser /var/lib/apt/lists/lock >/dev/null 2>&1 ; do
echo "Waiting for other apt-get instances to exit"
# Sleep to avoid pegging a CPU core while polling this lock
sleep 10
done

apt update
apt upgrade

apt install -y supervisor ca-certificates-java fontconfig-config fonts-dejavu-core java-common libavahi-client3 libavahi-common-data libavahi-common3 libcups2 libfontconfig1 libgraphite2-3 libharfbuzz0b libjpeg-turbo8 libjpeg8 liblcms2-2 libpcsclite1
  openjdk-17-jre-headless

wget -c https://github.com/wavesplatform/Waves/releases/download/v1.4.6/waves_1.4.6_all.deb
dpkg -i waves_1.4.6_all.deb

mkdir /var/lib/anote
chown -R waves:waves /var/lib/anote/

wget https://raw.githubusercontent.com/anonutopia/anote-node/main/conf/waves.conf
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/conf/application.ini

wget https://github.com/anonutopia/anote-node/releases/download/v1.0.1/anote-node
chmod +x anote-node
./anote-node -init
source ./seed
sed -i "s/ENCODED/$ENCODED/g" waves.conf
sed -i "s/DTMZNMkjDzCwxNE1QLomcp5sXEQ9A3Mdb2RziN41BrYA/$KENCODED/g" waves.conf
mv waves.conf /etc/waves/waves.conf

mv application.ini /etc/waves/application.ini

wget https://raw.githubusercontent.com/anonutopia/anote-node/main/conf/anote.conf
mv anote.conf /etc/supervisor/conf.d/

wget https://raw.githubusercontent.com/anonutopia/anote-node/main/config.json
sed -i "s/ADDRESS/$ADDRESS/g" waves.conf
sed -i "s/KEY/$KEY/g" waves.conf

service supervisor restart
service waves restart