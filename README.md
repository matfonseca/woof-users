# User service

## Build service

- ```make build```

## Start service

- ```make Start```

## Endpoints

### /users

####Methods

- POST

example body: 
```
{
    "mail":"mail@mail.com",
    "password":"password",
    "name": "Pepe",
    "lastname": "Diaz"
}
```

### /users/{id}

- GET
