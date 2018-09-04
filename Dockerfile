FROM ubuntu:16.04
WORKDIR /build

# Install PlatformIO
RUN apt-get update
RUN apt-get install -y python2.7 python-pip
RUN pip install -U platformio

# Install AtmelAVR and U8glib
RUN mkdir -p /build/lib
RUN mkdir -p /build/src
COPY cache_helpers/platformio.ini /build
RUN platformio platform install atmelavr
RUN platformio lib install 7

# Add TH3D Marlin source
COPY TH3D-Marlin/TH3DUF /build/src
RUN rm /build/src/Configuration.h

# First run: installing Arduino framework, caching most built files
COPY cache_helpers/Configuration.h /build/src
RUN platformio run

CMD cat .pioenvs/sanguino_atmega1284p/firmware.hex
