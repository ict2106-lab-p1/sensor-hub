<p align="center"><img src="mod3.png" width="300"></p>
<p align="center">
<a href="https://github.com/ict2106-lab-p1/sensor-hub/actions/workflows/release.yml"><img src="https://github.com/ict2106-lab-p1/sensor-hub/actions/workflows/release.yml/badge.svg" alt="Build Status"></a>
</a>
</p>


## Usage
Grab the latest [release](https://github.com/ict2106-lab-p1/sensor-hub/releases) for your platform and unzip. No additional dependencies are required, just download and run from the command line.

On first launch, a config file (config.yaml) will be created. Adjust as required.

### Boot the server
`./sensor-hub serve`

By default, the server will listen on `localhost:8000`. You can change this by passing in a different interface+port to `-l/--listen`.

`./sensor-hub serve -l 0.0.0.0:3000`

### Building/Dev
If you just want to run the application, grab the latest release. You do not need to build/install anything to run this CLI app.

`go run main.go serve`
