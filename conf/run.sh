#!/bin/bash

if [ -f secrets ]; then
    java -jar waves-all-1.4.20.jar /var/lib/anote/waves.conf
else 
    ./anote-node -init
    source ./secrets

    mv data /var/lib/anote/
    mv waves.conf /var/lib/anote/waves.conf

    sed -i "s/D5u2FjJFcdit5di1fYy658ufnuzPULXRYG1YNVq68AH5/$ENCODED/g" /var/lib/anote/waves.conf
    sed -i "s/DTMZNMkjDzCwxNE1QLomcp5sXEQ9A3Mdb2RziN41BrYA/$KEYENCODED/g" /var/lib/anote/waves.conf
    sed -i "s/Anote Node/$ADDRESS/g" /var/lib/anote/waves.conf

    java -jar waves-all-1.4.20.jar /var/lib/anote/waves.conf
fi