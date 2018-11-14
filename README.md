# [TH3D Unified Firmware](https://github.com/houseofbugs) compiler

It's like Marlin with more stuff. Now automated.

# Usage

```sh
git clone --recurse-submodules -j8 https://github.com/mplewis/th3d-ufw-compiler
cd th3d-ufw-compiler
docker build . --tag th3d-ufw-compiler
docker run --mount type=bind,source=/path/to/your/build_config,target=/config th3d-ufw-compiler
```

Note: You can't use `,readonly` for the bind target because the container gives you the compiled firmware by copying it into your mounted directory.

## Build Config

The above `docker run` command mounts `/path/to/your/build_config` inside the Docker container. The compile script uses files in this directory to run your build:

```
build_config
├── Configuration.h
└── platformio.ini
```

## `Configuration.h`

A valid copy of [`Configuration.h`](https://github.com/houseofbugs/TH3D-Marlin/blob/23b773ad8d067462de155fc9aeee2936bc4e4624/TH3DUF/Configuration.h) from TH3D-Marlin. You'll need to uncomment the `#define` lines that enumerate your printer's model and features.

## `platformio.ini`

A valid [PlatformIO Project Configuration File](http://docs.platformio.org/en/latest/projectconf.html) that targets your printer's hardware. It must include a default configuration that `platformio run` will use.

Example valid `platformio.ini` for a Creality CR-10, which uses the "Sanguino" ATmega1284p:

```
[platformio]
env_default = printer

[env:printer]
platform = atmelavr
framework = arduino
board = sanguino_atmega1284p
```

# License

MIT
