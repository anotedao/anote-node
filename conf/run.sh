#!/bin/bash

if [ -f secrets ]; then
    java -jar waves-all-1.4.20.jar waves.conf
else 
    ./anote-node -init
    source ./secrets
    sed -i "s/D5u2FjJFcdit5di1fYy658ufnuzPULXRYG1YNVq68AH5/$ENCODED/g" waves.conf
    sed -i "s/DTMZNMkjDzCwxNE1QLomcp5sXEQ9A3Mdb2RziN41BrYA/$KEYENCODED/g" waves.conf
    sed -i "s/Anote Node/$ADDRESS/g" waves.conf

    java -jar waves-all-1.4.20.jar waves.conf
fi