# Logship
Simple CLI tool for fetching file data or log entries from remote host.

[![Build Status](https://travis-ci.org/iobestar/jbalancer.svg?branch=master)](https://travis-ci.org/iobestar/logship)

#### Build from source

    $ go get -u github.com/iobestar/logship/cmd/logship

### Server

Logship in server mode serves file data in form of log units. Log unit identifies  
group of files based on glob pattern. Log unit is defined with id and glob pattern.  
Log unit definition format: <log_unit_id_n>:<glob_pattern_n:<log_unit_id_n+1>:<glob_pattern_n+1


Run logship server:

    $ logship server --address=":11034" --logunits="myservice_error:/var/log/myservice/error.log*:myservice_output:/var/log/myservice/output.log*"

Run as logship server as Docker:

    $ docker run --name=logship -p 11034:11034 -v "/var/log/myservice":"/var/log/myservice" -e LOG_UNITS="myservice_error:/var/log/myservice/error.log*:myservice_output:/var/log/myservice/output.log*" -d iobestar/logship

### Client (CLI)

Default configuration (logship.yml):

    log_readers:
      - id: default
        log_pattern: "^(?P<datetime>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{3}).*"
        date_time_layout: "2006-01-02 15:04:05.000"

Log reader defines log pattern as regex and date time layout using Go date pattern notation.

#### Fetch log units

    $ logship --target=localhost:11034
    
target: logship server address

#### Fetch last lines by count

    $ logship --target=localhost:11034 nlines <unit> <number_of_lines>
    
target: logship server address  
unit: id of log unit    
number_of_lines: count of lines  
    
#### Fetch last logs by count

    $ logship --config=logship.yml --target=localhost:11034 nlogs <unit> <number_of_logs> <reader_id>
    
config: configuration path  
target: logship server address  
unit: id of log unit  
number_of_logs: count of logs  
reader_id: id of reader from configuration; first log reader definition will be used if not specified
    
#### Fetch last logs by time

    $ logship --config=logship.yml --target=localhost:11034 tlogs <unit> <duration> <reader_id>
    
config: configuration path  
target: logship server address  
unit: id of log unit  
duration: "1s", "10m", "5h"  
reader_id: id of reader from configuration; first log reader definition will be used if not specified
