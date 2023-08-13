STEP
1. Crud API ()
    - validation (request, entity model)
    - get, post, put, delete
    - add response for not method allowed & no route
2. Database ORM ()
    - table automigrate & populate data for testing
    - model represent table
    - add, getById, getAll, update, softDelete
    - validation in fields each model
3. Goroutine
4. Security such as JWT, auth (basicauth)
    - authentication using username & password
    - add response body with 401 status code
    - login, refresh token, logout, register
    - resource cant be access without login
5. Session management
    - JWT Based Session Management
    - Login, Register, Refresh Token, Logout
    - Credential stored in DB
6. Environtment Variable
7. Unit Test
    - anytest for test platground
8. Logging
    - logger for gorm, so far, logger specified for gorm
9. Message Broker as microservices
10. Best Practice
    - which one is better, create struct like class with method or struct with all function ???
11. Docker Containerization
12. Deployment with Docker of course


Plan
1. add password hashing
2. book db need to be more complicated, add relationship, test gorm, modify model etc (book & author separated)
3. add concurrentcy using goroutine

Project
1. Messaging Backend System
2. 