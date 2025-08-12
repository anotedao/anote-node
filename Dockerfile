FROM ubuntu:22.04

RUN echo 'APT::Install-Suggests "0";' >> /etc/apt/apt.conf.d/00-docker

RUN echo 'APT::Install-Recommends "0";' >> /etc/apt/apt.conf.d/00-docker

RUN DEBIAN_FRONTEND=noninteractive \
  && apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y wget curl apt-utils ca-certificates-java fontconfig-config fonts-dejavu-core java-common libavahi-client3 libavahi-common-data libavahi-common3 libcups2 libfontconfig1 libgraphite2-3 libharfbuzz0b libjpeg-turbo8 libjpeg8 liblcms2-2 libpcsclite1 openjdk-17-jre-headless \
  && rm -rf /var/lib/apt/lists/*

RUN mkdir /var/lib/anote

COPY data /var/lib/anote/data/

COPY conf/waves.conf waves.conf

COPY conf/run.sh run.sh

COPY waves-all-1.4.20.jar waves-all-1.4.20.jar

COPY anote-node anote-node

CMD [ "./run.sh" ]