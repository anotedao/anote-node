#!/bin/bash

# Wait for apt to finish
while fuser /var/lib/apt/lists/lock >/dev/null 2>&1 ; do
echo "Waiting for other apt-get instances to exit"
# Sleep to avoid pegging a CPU core while polling this lock
sleep 10
done

# Update, upgrade and install dependencies
export DEBIAN_FRONTEND=noninteractive
apt update
apt upgrade -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" --force-yes
apt install -y supervisor ca-certificates-java fontconfig-config fonts-dejavu-core java-common libavahi-client3 libavahi-common-data libavahi-common3 libcups2 libfontconfig1 libgraphite2-3 libharfbuzz0b libjpeg-turbo8 libjpeg8 liblcms2-2 libpcsclite1
  openjdk-17-jre-headless

# Get files
wget -c https://github.com/wavesplatform/Waves/releases/download/v1.4.7/waves_1.4.7_all.deb
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/conf/waves.conf
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/conf/application.ini
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/config.sample.json
wget https://raw.githubusercontent.com/anonutopia/anote-node/main/conf/anote.conf
wget https://github.com/anonutopia/anote-node/releases/download/v1.1.3/anote-node

# Install Waves node
dpkg -i waves_1.4.7_all.deb
mkdir /var/lib/anote
chown -R waves:waves /var/lib/anote/
cp waves.conf /etc/waves/waves.conf

# Install Anote node
mv config.sample.json config.json
chmod +x anote-node
./anote-node -init
source ./seed
sed -i "s/D5u2FjJFcdit5di1fYy658ufnuzPULXRYG1YNVq68AH5/$ENCODED/g" waves.conf
sed -i "s/DTMZNMkjDzCwxNE1QLomcp5sXEQ9A3Mdb2RziN41BrYA/$KENCODED/g" waves.conf
sed -i "s/127.0.0.1:/$PUBLICIP:/g" waves.conf
mv waves.conf /etc/waves/waves.conf
mv application.ini /etc/waves/application.ini
mv anote.conf /etc/supervisor/conf.d/

# Prepare node config file
sed -i "s/ADDRESS/$ADDRESS/g" config.json
sed -i "s/PUBLICKEY/$PUBLICKEY/g" config.json
sed -i "s/PRIVATEKEY/$PRIVATEKEY/g" config.json
sed -i "s/KEY/$KEY/g" config.json

# Remove extra files and folders
rm -rf /var/lib/waves
rm -rf /var/lib/anote/wallet

# Secure the node
adduser --quiet --disabled-password --gecos "" anon
chpasswd <<<"anon:$KEY"
sed -i "s/sudo:x:27:/sudo:x:27:anon/g" /etc/group
sed -i "s/PasswordAuthentication no/PasswordAuthentication yes/g" /etc/ssh/sshd_config
echo "anonnode" > /etc/hostname

reboot