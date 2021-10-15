# example-web-service-gin

## HOW TO INSTALL
**Download the zipfile** \n

**run .exe file**

## HOW TO USE

** Open cmd or Powershell or git bash (Git bash is recommended) **

**GET - view all members** 
  --> curl http://localhost:8080/members

**GET - view single member** 
  --> curl http://localhost:8080/members/2
  
  
  ../members/id <- put the id to see

**POST**
  --> curl -X POST -i -H 'Content-Type: application/json' -d '{"id":"6", "name":"Himani","team":"Punjab Kings"}' http://localhost:8080/members
  
**PUT**
  --> curl -X PUT -i http://localhost:8080/members/6/Chennai+Super+Kings
  
  
  ../members/id <- put the id to be edited  
  ../members/id/team -<put the team name (use + for whitespace)
  
**DELTE**
  --> curl -X DELETE -i http://localhost:8080/members/4
  
      ../members/id <- put the id you want to delete
