## app Structure
- cmd > main file
- internal > all internal development, such as handlers, types, db, etc
    - data > interaction with database stuff and db models (basically all related with db)
    - handlers > routes handlers (functions that handle route req&res)
    - routes > routes of the api
    - types > typing stuff like structs or interfaces
    - config > setup config as DB connect
    - utils > functions that will help u through the app
