# Project Desgin
This project is designed to be extremely modular and adaptive to however you want to use it. The server uses modules to be able to control either an implant or command and control server (c2). To use the server you must tell it what you modules are so that it knows how to generate and create actions for you.

## Terms
* Framework - Infrastructure with defined integration requirements 
* API - Application Programming Interface
* Teamserver - Server acting as operator interface for command and control
* Command and Control (C2) - manages implants
* Agent - A type of implant
* Implant - Code running on a victim machine
* Transport - Protocol for communicating between C2 and Implant

## Registration
An implant or C2 will initialize its connection to the teamserver and go through a registration process. To be properly registered the entity being registered needs to inform the teamserver what modules it has to use. 

*NOTE:* The hivemind developers suggest naming your modules in an easy to recongize format. If there is a module that you have for a specific implant we would suggest you name it `<implant_name>_<module_name>`. For a C2 we suggest `c2_<c2_name>_<module_name>`. This will help keep everything organized.

## After Registration
After an implant or C2 is registered it can always repull its information down if connection is lost. Once registration is complete the implants or C2s should send `ActionRequests` to the teamserver. If there is anything for them to do they will then recieve an `Action` that they complete that has an ID. Once the `Action` is completed they should respond to the teamserver again with the response and the idea in an `ActionResponse` structure.

## Diagrams
TODO.