# Parking Lot

A simple CLI based automated ticketing system to manage day to day work in a multi-storey parking lot.

## Setup

```bash
$ bin/setup
```

Running above script will ensure 
* you have go installed on your system. 
* you have a GOPATH setup
* your source code is added to $GOPATH/src
* your source code is build and binary is added to bin/parking_lot

## Execution

### Interactive mode

```bash
$ bin/parking_lot
```
### File input mode<sup>[2](#Appendix)</sup>

```bash
$ bin/parking_lot <relative or absolute path of the file>
```

## Features

Following commands are supported: 
* `create_parking_lot <number of slots>` : creates the parking lot with specified slots
* `park <registration number> <colour>` : parks the car at the nearest empty slot in the parking lot or display an error message for parking lot full
* `leave <slot number>` : marks the specified slot as empty or display error message if already empty
* `status` : shows the current status of the parking lot
* `registration_numbers_for_cars_with_colour <color1> [, ...<colour n>]` : shows the registration number of all the vehicles having specified color(s)
* `slot_numbers_for_cars_with_colour <color1>[, ...<colour n>]` :  shows the slot number of all the vehicles having specified color(s)
* `colours_for_cars_with_registration_number <reg num 1>[,...<reg num n>]` : shows the colour of vehicles having specified registration number(s)
* `slot_numbers_for_cars_with_registration_number <reg num 1>[,...<reg num n>]` : shows the slot numbers of vehicles having specified registration number(s)
* `colours_for_slot_number <slot num 1>[,...<slot num n>]` : shows the colour of vehicles parked at specified slot number(s)
* `registration_numbers_for_slot_number <slot num 1>[,...<slot num n>]` : shows the registration numbers of vehicles parked at specified slot number(s)
* `exit` : exits the application

## Development
To setup the system for development execute the setup script.

### Running tests
* Unit
    ```bash
    $ bin/setup go test ./...
    ```
* Functional
    ```bash
    $ bin/run_functional_tests
    ```
