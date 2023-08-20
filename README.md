STEP
1. Crud API ()
    - validation (request, entity model)
    - get, post, put, delete
    - add response for not method allowed & no route
2. Database ORM ()
    - table automigrate & populate data for testing
    - model represent table, login account
    - add, getById, getAll, update, softDelete
    - validation in fields each model
3. Goroutine
4. Caching Mechanism using redis
5. Authorization (Basic, JWT)
    - using username & password on basic auth
    - login, refresh token, logout, register
    - resource cant be access without login
6. Authentication
    - authentication using username & password
    - add response body with 401 status code
    - add hashing password
6. Session management (JWT)
    - JWT Based Session Management
    - Login, Register, Refresh Token, Logout
    - Credential stored in DB
7. Environtment Variable
    - using viper to read from env variable
    - using yaml file
    - create reader & env model in go struct, load & replace some config with the env variable one
8. Unit Test
    - anytest for test platground
9. Logging
    - logger for gorm, so far, logger specified for gorm
10. Message Broker as microservices
11. Best Practice
    - which one is better, create struct like class with method or struct with all function ???
12. Docker Containerization
13. Deployment with Docker of course


Plan
1. book db need to be more complicated, add relationship, test gorm, modify model etc (book & author separated)
2. add concurrentcy using goroutine
3. Covered with unit test
4. add http connection to other api
5. send email confirmation when register, add media storage system

Project
1. Messaging Backend System
2. 