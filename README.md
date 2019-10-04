# Logship
Simple CLI tool for fetching file data or log entries from remote host.

[![Build Status](https://travis-ci.org/iobestar/logship.svg?branch=master)](https://travis-ci.org/iobestar/logship)

#### Build from source

    $ go get github.com/iobestar/logship/cmd/logship
    
_Note: It is recommended to have `GO111MODULE=on` to ensure the correct
dependencies are used._

### Server

Logship in server mode serves file data in form of log units. Log unit identifies  
group of files based on glob pattern. Log unit is defined with id and glob pattern.

Configuration file (logship.yml):

    log_units:
      - id: service_error
        glob: /var/log/myservice/error.log*
      - id: service_output
        glob: /var/log/myservice/outout.log*
    

Run logship server:

    $ logship server --address=":11034" --logunits="myservice_error:/var/log/myservice/error.log*:myservice_output:/var/log/myservice/output.log*"

Run as logship server as Docker:

    $ docker run --name=logship -p 11034:11034 -v "/var/log/myservice":"/var/log/myservice" -e LOGSHIP_CONFIG="<logship.yml>" -d iobestar/logship


### Client (CLI)

    logship client --help

#### Fetch log units

    $ logship --target=localhost:11034
    
target: logship server address

#### Fetch last lines by count

    $ logship --target=localhost:11034 nlines <unit> <number_of_lines>
    
target: logship server address  
unit: id of log unit    
number_of_lines: count of lines  
    
#### Fetch last logs by count

    $ logship --target=localhost:11034 nlogs <unit> <number_of_logs> <reader_id>
    
config: configuration path  
target: logship server address  
unit: id of log unit  
number_of_logs: count of logs  
    
#### Fetch last logs by time

    $ logship --target=localhost:11034 tlogs <unit> <duration> <reader_id>
    
config: configuration path  
target: logship server address  
unit: id of log unit  
duration: "1s", "10m", "5h"  
