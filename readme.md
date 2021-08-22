# Go - Gin Boilerplate

This is the boilerplate for Go-Gin application.

## Running the application

**Prerequisite:**
 - JSON file with environment configuration in the application directory to run this application quickly, hopefully without any major headaches. A sample `config-dev.json`  file is provided.
 - Can have number of separate configuration files for different environments. For example  `config-stg.json` and `config.json` for staging and production environment configuration.

_**Now it's time to run...**_

From the project root director, run:

cd .\server\

**Production**

    go run main.go -env=prod

**Staging**

    go run main.go -env=stg

**Development**

    go run main.go

    To register the user 
    URL : http://localhost:4700/o/register
    Method: POST
    Body: generate the jwt token string using key(PrivateKey) from config file
    sample token : 
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdE5hbWUiOiJSYWphdCIsImxhc3ROYW1lIjoiWWFkYXYiLCJlbWFpbCI6InJhamF0eTY5M0BnbWFpbC5jb20iLCJhZGRyZXNzIjoiTmFncHVyLCBNSCIsIm1vYmlsZSI6ODk4MzM5Mjk4M30.sRQ_CRw6PXBs7QVsGUsQ_MYb6uR0fr3_v_J1MkwciLM
    
    sample json to generate token
    {
        "firstName": "Rajat",
        "lastName": "Yadav",
        "email": "rajaty693@gmail.com",
        "address": "Nagpur, MH",
        "mobile": 8983392983
    }

    To Login the user
    URL : http://localhost:4700/o/login
    Method: POST
    Body: generate the jwt token string using key(PrivateKey) from config file
    sample token:
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6InJhamF0eTg5OCIsInBhc3N3b3JkIjoieWFkYXZyODk4In0.qdRlhtfCoIvizS1tg7m6oVfV69egZWTmorgmeziEi4s
    
    sample json to generate token
    {
        "userName": "rajaty898",
        "password": "yadavr898"
    }

_**Happy Coding...**_
