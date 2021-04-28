# Teamserver
## Summary
The teamserver portion of project is the brains. It controls all the implants and will dynamically implement functions the operators can use as implants tell them. With the teamserver you can do scheduled grouping, and easy plugin implementation. 

## Disclaimer
THIS TOOK IS MEANT FOR COMPETITION USE ONLY. THE DEVELOPERS AND CONTRIBUTERS ARE NOT RESPONSIBLE FOR ANY MISUSE OF THIS CODE.

## Teamserver Setup
**NOTE:** Use the deployment repo if you want to easily deploy the whole project at once, this should usually only be used for development.

There is a `docker-compose` file that will spin up all the services needed to run the teamserver in this repo. If you want to not run the Go code in the container you will need to comment out that container and change the config file. Then you can run the go code with `go run cmd/hivemind/server.go config/config.yaml`. 

## Go PKGs 
The teamserver has Go project has serveral packages it uses in order to run. The packages are:

* comms
    * Used for anything to do with handling the communications of the platform.
    * The models for the network commication are located here.
* conf
    * Used to read the config file. 
* crud
    * Used for any interaction with the database
* listeners
    * Place listeners here which will call the `comms` package for data handling.
    * Ex: tcp / http / udp
* model
    * Holds all the models that the teamserver uses to pull and save data to the database and the REST API. 
* rest
    * This is the logic for the REST API. You would add routes here.