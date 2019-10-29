# MPK

Dead simple web app which shows the current position of Trams and Buses in Wroclaw.

## Usage

First step is to create new file in the repo: `.env` whit Google Maps Key.

```bash
echo 'GOOGLE_MAPS_KEY=<YOUR_KEY_GOES_HERE>' >> .env
```

Then to run application:

```bash
./start.sh tram 4
```

where `tram` points to type of vehicles which should be shown on the map and `4` is the number of particular type of vehicle.

When app is started you just need to open: [http://localhost:9000](http://localhost:9000) and show in the realtime, position of you bus/tram.

## Requirements

* Go > 1.6

## Why

Some time ago I used tram to go to home and usually I was late... So with this app my life was much more easier
