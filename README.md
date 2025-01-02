# CRUD Project In Go

This is my practice project that I worked on while learning. I got the concept for the project from Udemy. It's a CRUD project using Golang.

In this project, I used:

- Gin for handling HTTP requests
- bcrypt for hashing passwords
- JWT for authorization

In this CRUD project, I handle authorization using tokens. When creating, updating, or deleting an event, there are validation checks to ensure that only the authorized user (the one who created the event) can update or delete it.

Here are the available requests in this CRUD project along with their respective responses:

#### Create User 

request:

```
POST http://localhost:8080/signup
content-type: application/json

{
	"email" : "Rafid@pathao.com",        
	"password" :"test"    
	
}

```

response: 

```
HTTP/1.1 201 Created

{
  "event": {
    "Id": 8,
    "Email": "Rafid@pathao.com",
    "Password": "$2a$14$JqDZU2Nh6dHuLbLpnWF.dOmG9TKpMlNzFkggLNABIt9I8.U94sUK."
  },
  "message": "User Created"
}

```

#### User login 

request:

```
POST http://localhost:8080/login
content-type: application/json

{
	"email" : "Rafid@pathao.com",        
	"password" :"test" 
}

```

response: 

```
HTTP/1.1 201 Created

HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Thu, 02 Jan 2025 19:26:22 GMT
Content-Length: 197
Connection: close

{
  "message": "Login Successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJhZmlkQHBhdGhhby5jb20iLCJleHAiOjE3MzU4NTMxODIsInVzZXJJZCI6OH0.iOZ-ehYX1vLlRqbiGS7sXcsHFpkNqaBgw0OSaufC_TY"
}

```


#### Create Event 

request:

```
POST http://localhost:8080/events
content-type: application/json
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJhZmlkQHBhdGhhby5jb20iLCJleHAiOjE3MzU4NTMxODIsInVzZXJJZCI6OH0.iOZ-ehYX1vLlRqbiGS7sXcsHFpkNqaBgw0OSaufC_TY


{
	"name" : "Rafid",        
	"description" : "Welcome" ,
	"location" : "Banasree",    
	"dateTime" : "2025-02-01T15:30:00.000Z"    
	
}

```

response: 

```
HTTP/1.1 201 Created

{
  "event": {
    "Id": 2,
    "Name": "Rafid",
    "Description": "Welcome",
    "Location": "Banasree",
    "DateTime": "2025-02-01T15:30:00Z",
    "UserID": 8
  },
  "message": "Event Created"
}

```


#### Update Event 

request:

```
PUT  http://localhost:8080/events/3   (3 is event id)
Content-Type: application/json
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJhZmlkQHBhdGhhby5jb20iLCJleHAiOjE3MzU4NTM2MDUsInVzZXJJZCI6OH0._HiS78JGYQDaNPMV1a-q7SMYNWyASyOtxst6OVltEIc

{
	"name" : "Rafid5",        
	"description" : "Welcome2" ,
	"location" : "AB",    
	"dateTime" : "2025-02-01T15:30:00.000Z"    
	
}

```

response: 

```
HTTP/1.1 200 OK

{
  "event": {
    "Id": 3,
    "Name": "Rafid5",
    "Description": "Welcome2",
    "Location": "AB",
    "DateTime": "2025-02-01T15:30:00Z",
    "UserID": 0
  },
  "message": "Event Updated"
}

```

#### Delete Event 

request:

```
DELETE  http://localhost:8080/events/3

Content-Type: application/json
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJhZmlkQHBhdGhhby5jb20iLCJleHAiOjE3MzU4NTM2MDUsInVzZXJJZCI6OH0._HiS78JGYQDaNPMV1a-q7SMYNWyASyOtxst6OVltEIc


```

response: 

```
HTTP/1.1 200 OK

{
  "message": "Event Deleted"
}

```

#### Delete Event 

request:

```
DELETE  http://localhost:8080/events/3

Content-Type: application/json
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJhZmlkQHBhdGhhby5jb20iLCJleHAiOjE3MzU4NTM2MDUsInVzZXJJZCI6OH0._HiS78JGYQDaNPMV1a-q7SMYNWyASyOtxst6OVltEIc


```

response: 

```
HTTP/1.1 200 OK

{
  "message": "Event Deleted"
}

```

#### Collect event info by event id

request:

```
GET http://localhost:8080/events/2


```

response: 

```
HTTP/1.1 200 OK

{
  "Id": 2,
  "Name": "Rafid",
  "Description": "Welcome",
  "Location": "Banasree",
  "DateTime": "2025-02-01T15:30:00Z",
  "UserID": 8
}

```


#### Registration for a event 

request:

```
POST http://localhost:8080/events/1/register

Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJhZmlkQHBhdGhhby5jb20iLCJleHAiOjE3MzU4NTM2MDUsInVzZXJJZCI6OH0._HiS78JGYQDaNPMV1a-q7SMYNWyASyOtxst6OVltEIc



```

response: 

```
HTTP/1.1 201 Created

{
  "message": "Successfully register!!"
}

```

#### Delete Registration for a event 

request:

```
DELETE http://localhost:8080/events/1/register

Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IlJhZmlkQHBhdGhhby5jb20iLCJleHAiOjE3MzU4NTM2MDUsInVzZXJJZCI6OH0._HiS78JGYQDaNPMV1a-q7SMYNWyASyOtxst6OVltEIc




```

response: 

```
HTTP/1.1 201 Created

{
  "message": "Cancel registration successfully!!"
}

```

#### if Unauthorized/expaired token used

response: 

```
HTTP/1.1 401 Unauthorized

{
  "message": "Unauthorized Token"
}

```



















