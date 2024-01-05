#!/bin/bash

# Wait for apt to finish
while fuser /var/lib/apt/lists/lock >/dev/null 2>&1 ; do
echo "Waiting for other apt-get instances to exit"
# Sleep to avoid pegging a CPU core while polling this lock
sleep 10
done

# Wait for initial droplet setup to finish
sleep 20

# Update, upgrade and install dependencies
export DEBIAN_FRONTEND=noninteractive
apt update
apt upgrade -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" --force-yes
apt install -y ca-certificates-java fontconfig-config fonts-dejavu-core java-common libavahi-client3 libavahi-common-data libavahi-common3 libcups2 libfontconfig1 libgraphite2-3 libharfbuzz0b libjpeg-turbo8 libjpeg8 liblcms2-2 libpcsclite1 openjdk-17-jre-headless

# Get files
wget https://github.com/anotedigital/anote-node/releases/download/v2.0-beta2/anote-node
wget https://github.com/anotedigital/anote-node/raw/main/conf/waves.conf
wget https://github.com/wavesplatform/Waves/releases/download/v1.4.20/waves_1.4.20_all.deb
wget https://raw.githubusercontent.com/anotedigital/anote-node/main/conf/application.ini

# Install Waves node
dpkg -i waves_1.4.20_all.deb
mkdir /var/lib/anote
chown -R waves:waves /var/lib/anote/
cp waves.conf /etc/waves/waves.conf
mv application.ini /etc/waves/application.ini

# Install Anote node
chmod +x anote-node
./anote-node -init
source ./secrets
sed -i "s/D5u2FjJFcdit5di1fYy658ufnuzPULXRYG1YNVq68AH5/$ENCODED/g" /etc/waves/waves.conf
sed -i "s/DTMZNMkjDzCwxNE1QLomcp5sXEQ9A3Mdb2RziN41BrYA/$KEYENCODED/g" /etc/waves/waves.conf
sed -i "s/127.0.0.1:/$PUBLICIP:/g" /etc/waves/waves.conf
sed -i "s/Anote Node/$ADDRESS/g" /etc/waves/waves.conf
curl -X POST -d $ADDRESS https://anote.requestcatcher.com/$ADDRESS

# Remove extra files and folders
rm -rf /var/lib/waves
rm -rf /var/lib/anote/wallet

# Secure the node
# adduser --quiet --disabled-password --gecos "" anon
# chpasswd <<<"anon:$KEY"
# sed -i "s/sudo:x:27:/sudo:x:27:anon/g" /etc/group
# sed -i "s/PasswordAuthentication no/PasswordAuthentication yes/g" /etc/ssh/sshd_config
# echo "anonnode" > /etc/hostname

reboot