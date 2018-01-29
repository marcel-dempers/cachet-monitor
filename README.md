# cachet-monitor
Solution that probes endpoints and updates cachet if there is an outage or recovery

## Prerequisites

You will need Cachet up and running. </br>
You will also need a URL (I.E `http://cachet:8000/api/v1`) of that Cachet server and authentication (I.E `zlvNVV0VxKkuRL0WM5ww`) token

## Configuration

You will need to provide a configuration file in `yaml` format.
Supported configuration sample :

```
monitors:
  - endpoint:
    type: probe
    name: GoogleHello
    url: http://google-hello:8080
    method: GET
    frequency: 5
    expectedResponseCode: 200
    timeoutInSec: 10
    maxfailures: 5
    cachet:
      componentid: 1
```

## Start something to monitor

You will need some endpoint to monitor:

```
docker run -it --rm --name google-hello -p 8080:8080 google/golang-hello
```

## Start the monitor

Check out `start-monitor.sh` to see how its started.