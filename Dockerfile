FROM ubuntu:16.04

WORKDIR /build

RUN apt-get update
RUN apt-get install -y python2.7 python-pip
RUN pip install -U platformio

RUN mkdir -p /build/lib
RUN mkdir -p /build/src
COPY platformio.ini /build
RUN platformio lib install 7
RUN platformio platform install atmelavr
# TODO: Figure out how to cache the installation of framework-arduinoavr

COPY TH3D-Marlin/TH3DUF /build/src
COPY Configuration.h /build/src
RUN platformio run
RUN cat .pioenvs/sanguino_atmega1284p/firmware.hex
