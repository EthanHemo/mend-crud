# Superhero crud!
In this application we will be able to store and load superheros.  

# Run application  
To run the application, wither build it locally with the command:  
```
go run .  
```

Or build a container with the following steps:  
1. Copy repo to the url:  $GOPATH/src/mend-crud/  
2. run `docker build . -t crud` to build the image
3. run `docker run -p 8090:8090 crud:latest` to run the container locally.

# Database interface  
In this application you can choose the following databases:  
* Array  
* File  

To store superheros informations in different database sources, do the following:  
1. Create a file under "databases" folder with the type of database as the source.  
2. Implement the interface `crud` that is in the file `api.go`  
3. Add to the enum and the function `GetDatabase` in the file `api.go`  
4. Set the `main.go` file to use the enum you used.  

# APIs  
## /add  
Accept super hero parameters:  
```
{
    "Name": "Superman" ,
    "Age": 900, 
    "Power":"Super strong", 
    "Weakness":"Cheese" 
}  
```

## /update  
Accept super hero name and parameters parameters:  
```
{
    "Name": "Superman" ,
    "Superhero": {
        "Name": "Superman" ,
        "Age": 900, 
        "Power":"Super strong", 
        "Weakness":"Cheese" 
    }
}  
```

## /remove  
Accept super hero name only:  
```
{
    "Name": "Superman" 
}  
```

## /Get  
Accept super hero name only:  
```
{
    "Name": "Superman" 
}  
```
