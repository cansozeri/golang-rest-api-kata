# golang-rest-api-kata
REST-API in GoLang with Clean Architecture

# Crucial Points

- Delivering a Working RESTful API.
- Clean and Production Ready Code
- Error Handling
- Comments and Documentation
- Unit and/or Integration Tests
- Avoid Over Engineering

# Clean Architecture

### What is Clean Architecture?
In his book “Clean Architecture: A Craftsman’s Guide to Software Structure and Design” famous author Robert “Uncle Bob” Martin presents an architecture with some important points like testability and independence of frameworks, databases and interfaces.

### The constraints in the Clean Architecture are :

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

# How to use?
### Mongo Endpoint
You should send a request to https://golang-clean-rest-api.herokuapp.com/api/v1/records/search with a valid JSON payload.
A valid payload consists of four parameters. `startDate`, `endDate`, `minCount` and `maxCount`.
If you fail to deliver any of these parameters error code 1 will greet you as the following.

```
{
    "code": 1,
    "msg": "{parameter} is a required field"
}
```

If you provide a valid payload it will deliver the correct result as intended.
It will check between the startDate and endDate, after selecting the correct interval it will find summation of the `counts` array in the DB and project it as `totalCount`.
Another match process will happen to ensure `totalCount` is between `minCount` and `maxCount`.
Final product will be delivered as a JSON which consists of `code`, `msg` and `records` fields.
`code` can be either `0`, `1` or `2`. Defaults to `0`.
* `0` indicates `success`.
* `1` indicates validation error.
  `msg` can be anything defined to explain the `code`. Defaults to `Success`.
  `records` is the filtered results given as an object.
* `2` indicates other generic errors.

## Sample inputs and outputs
All of the below were performed by sending POST requests to https://golang-clean-rest-api.herokuapp.com/api/v1/records/search
```
Request payload:
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```

```
Response payload:
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "key":"TAKwGc6Jr4i8Z487",
            "createdAt":"2017-01-28T01:22:14.398Z",
            "totalCount":2800
        },
        {
            "key":"NAeQ8eX7e5TEg7oH",
            "createdAt":"2017-01-27T08:19:14.135Z",
            "totalCount":2900
        }
    ]
}
```

```
Request payload: (Notice that minCount is greater than maxCount)
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2900,
    "maxCount": 2800
}
```

```
Response payload:
{
    "code": 1,
    "msg": "maxCount must be greater than or equal to MinCount"
}
```

```
Request payload: (Notice that endDate is less than startDate)
{
    "startDate": "2016-01-26",
    "endDate": "2015-01-01",
    "minCount": 2700,
    "maxCount": 3000
}
```

```
Response payload:
{
    "code": 1,
    "msg": "endDate must be greater than or equal to StartDate"
}
```

```
Request payload: (Notice that minCount parameter is missing, randomly chosen)
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "maxCount": 3000
}
```

```
Response payload:
{
    "code": 1,
    "msg": "minCount is a required field"
}
```
### In-Memory Endpoint
#### Create
You should send a request to https://golang-clean-rest-api.herokuapp.com/api/v1/in-memory with a valid JSON payload.
A valid payload consists of two parameters. `key` and `value`.
If you fail to deliver any of these parameters error code 1 will greet you as the following.

```
{
    "code": 1,
    "msg": "{parameter} is a required field"
}
```

If you provide a valid payload it will deliver the echo of the request as intended.

#### Get
You should send a request to https://golang-clean-rest-api.herokuapp.com/api/v1/in-memory?key=active-tabs with a query parameter called key.
If you fail to deliver key parameter error code 1 will greet you as the following.

```
{
    "code": 1,
    "msg": "key is a required field"
}
```
## Sample inputs and outputs for In-Memory Endpoint
All of the below were performed by sending POST requests to https://golang-clean-rest-api.herokuapp.com/api/v1/in-memory
```
Request payload:
{
    "key": "active-tabs",
    "value": "getir"
}
```

```
Response payload:
{
    "key": "active-tabs",
    "value": "getir"
}
```

```
Request payload: (Notice that key is missing)
{
    "key": "",
    "value": "getir"
}
```

```
Response payload:
{
    "code": 1,
    "msg": "key is a required field"
}
```

```
Request payload: (Notice that value is missing)
{
    "key": "active-tabs",
    "value": ""
}
```

```
Response payload:
{
    "code": 1,
    "msg": "value is a required field"
}
```

All of the below were performed by sending GET requests to https://golang-clean-rest-api.herokuapp.com/api/v1/in-memory
```
Request:
https://golang-clean-rest-api.herokuapp.com/api/v1/in-memory?key=active-tabs
```

```
Response payload:
{
    "key": "active-tabs",
    "value": "getir"
}
```

```
Request: (the key has not been created before)
https://golang-clean-rest-api.herokuapp.com/api/v1/in-memory?key=not-in-memory
```

```
Response payload:
{
    "code": 2,
    "msg": "entity not found on database"
}
```

### Local Deployment
If you want to run this project on your local machine use "make docker-run" command.

You can use localhost:8080 as a host to reach the endpoints.

