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

# Add TH3D Marlin source and runner
COPY TH3D-Unified-U1.R2/TH3DUF_R2 /build/src
COPY runtime/compile.sh /build

# First run: installing Arduino framework, caching most built files
COPY cache_helpers/Configuration.h /build/src
RUN platformio run

# Cleanup files used as params for the real run
RUN rm /build/src/Configuration.h /build/platformio.ini

CMD /build/compile.sh
