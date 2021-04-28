# What To Know
Here is the base layout again from the Teamserver overview page. This will be important depending on what you want to develop for the teamserver. 

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

# What To Do When Developing
When developing to actually get something merged, you must develop with the style of the server. Also, you will have to auto regenerate the documentation to place it in the docs folder. There is a section on how to do this. 
