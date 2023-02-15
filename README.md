# Go-Bank
A simple bank api with golang, postgres with JWT authentication



## .env file
```
POSTGRES_CONN_STRING = "username=yourDbUsername dbname=yourDbName password=yourDbPassword sslmode=disable"

DB_NAME = ""
DB_USERNAME = ""
DB_PASSWORD = ""


JWT_SECRET = ""
```

## docker-compose for starting postgresql database
```
$ docker-compose up -d
```


## to build files 
```
$ go build main.go
$ ./main
```