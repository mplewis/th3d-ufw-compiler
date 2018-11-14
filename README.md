# [TH3D Unified Firmware](https://github.com/houseofbugs) compiler

It's like Marlin with more stuff. Now automated.

# Usage

Start the server:

```sh
git clone --recurse-submodules -j8 https://github.com/mplewis/th3d-ufw-compiler
cd th3d-ufw-compiler
./build.sh
./run.sh
```

Make the REST request:

```sh
http --json POST 'http://localhost:8080/' \
    'Content-Type':'application/json' \
    config_header="...Configuration.h contents..." \
    pio_config="...platformio.ini contents..."
```

Receive a structured response:

```json
{ "compiled_hex": ":100000000C9431160C945B160C945B160C945B16D6..." }
```

# Build Config

Running './run.sh' exposes the Docker server on local port 8080. When you make a POST request, it must include a JSON body with two values that specify your desired firmware configuration:

## `config_header`

A valid copy of [`Configuration.h`](https://github.com/houseofbugs/TH3D-Marlin/blob/23b773ad8d067462de155fc9aeee2936bc4e4624/TH3DUF/Configuration.h) from TH3D-Marlin. You'll need to uncomment the `#define` lines that enumerate your printer's model and features.

## `pio_config`

A valid [PlatformIO Project Configuration File](http://docs.platformio.org/en/latest/projectconf.html) that targets your printer's hardware. It must include a default configuration that `platformio run` will use.

Example valid `pio_config` for a Creality CR-10, which uses the "Sanguino" ATmega1284p:

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
