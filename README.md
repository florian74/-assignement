# Assignement

This repo contains the code of an assignment we can give to candidates for technical positions.

# Context

Before flying, a pilot needs to file a "Flight Plan" with some information about what the pilot intends to do, such as:

* the type of aircraft
* the departure airport
* the destination airport
* the estimated take-off time
* the route the pilot intends to follow.

The Flight Plan is sent to Eurocontrol and forwarded to all "ANSP", Air Navigation Service Provider (Skyguide in Switzerland, DSNA in France, DFS in Germany).

Upon reception, ANSPs store and process the flight plan to compute different information such as the time and position the flight is supposed to enter their airspace.

The flight plan is one of the most important domain object in Air Traffic Control business.

# Assignment

In this assignment, the candidate is asked to implement the reception and management of flight plans.

There is a flight plan generator provided [here](https://github.com/SkySoft-ATM/assignement/blob/master/cmd/fplgen/main.go). The generator sends Flight Plan in JSON over UDP on localhost:6669.

The Flight Plan object is defined as `IFPL` [here](https://github.com/SkySoft-ATM/assignement/blob/master/adexp/ifpl.go)

Here is a sample message emitted by the flight plan generator :

```json
{
    "Title": "IFPL",
    "Addr": [
        "EGZYTTTE",
        "EGLLZTZP",
        "EHAAZQZX",
        "EGLLZTZP",
        "EHAAZQZX",
        "EGLLZTZR"
    ],
    "Adep": "EGLL",
    "Ades": "LSZH",
    "Arcid": "KVN979",
    "Arctyp": "A320",
    "Ceqpt": "SRGWY",
    "Eobd": "20191106",
    "Eobt": "2111",
    "Filtim": "20191105",
    "IfplId": "OU52187947",
    "Origin": "-NETWORKTYPE AFTN -FAC EGLLZTZR",
    "Seqpt": "C",
    "Wktrc": "M",
    "Opr": "ABC",
    "Pnb": "B2",
    "Reg": "GAAPO",
    "Rmk": "HKKXQCBZ",
    "Rvr": 200,
    "Sel": "DSGL",
    "Src": "FPL",
    "Ttleet": "1377403",
    "Rfl": "F582",
    "Speed": "N0525",
    "Fltrul": "I",
    "Fltyp": "S",
    "Route": [
        "LIFFY",
        "LIFFY5A",
        "UL620",
        "REDFA/N0390F230"
    ],
    "Altrnt1": "",
    "Eetfir": null,
    "Rtepts": null,
    "Atsrt": null
}
```


## General rules

The candidate can implement the solution in Go or another language he feels more comfortable with.

## Step 1

The candidate must explain & implement an application that listens to Flight Plan objects received by UDP, stores them in a persistant storage of their choice.

## Step 2

The candidate must explain & implement a REST API to retrieve flight plans.

The API should support Flight Plan filters on: 
* ADEP (departure Airport)
* ADES (destination Airport)
* ARCID (Aircraft Identifier)
* IFPL ID

## Step 3

The candidate must package his application using Docker.
With the tooling of his choice the candidate must demonstrate that his server works as intended in steps 2.



USAGE

docker-compose build
docker-compose up
go run cmd/fplgen/main.go

Check search on: 
http://localhost:4201/flight/IfplId/R
http://localhost:4201/flight/Adep/