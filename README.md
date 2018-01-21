# rpi_notifier

## Build:
```
glide install
GOARCH=arm GOARM=7 go build
```

## Config file:
All configuration parameters can be configured using a single json file as below:
The default port is 7777

```json
{
    "Port": 7777,
    "ButtonPin": 3,
    "Leds": {
        "red": [
            12,
            13
        ], "green": [
            15,
            16 
        ]
    }
}
```

## Usage:
### start server
```
./rpi_notifier config.json

```

### red leds
```
curl http://pi:7777/led/red/{on|off|blink}
```

### green leds
```
curl http://pi:7777/led/green{on|off|blink}
```

### clear all leds
```
curl http://pi:7777/leds/clear
```
