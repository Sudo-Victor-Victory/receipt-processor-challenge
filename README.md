# Receipt Processor

My Implementation of [Fetch's Payment Processor challenge](https://github.com/fetch-rewards/receipt-processor-challenge). Here you can find the prompt, and API specificaiton.

My implementation was made using Golang, Gin, Google's UUID, and Testify. The app and tests are hosted within Docker.

# How it works
Gin is the center of this project. Gin allowed for quick development of a Middleware - functions that can define routes, and process requests & responses. Similarly, it quickly allowed me to enforce JSON validation rules. 

From there I use Google's UUID to generate unique user ID's on the fly, fufilling a requirement in the spec.

Testify was a given, if there is middleware being written then we have to have some tests verifying API calls (System test) and function truthyness (unit tests).

Lastly, Docker was chosen not only for it meeting a requirement of the spec, but for ease of building, testing, and environment isolation, but for its use in testing and deployment.


Note:
I elected to use Gin instead of implementing this all by hand using Go's builting web handler functions (anything in net/http like http.ResponseWriter) for speed in prototyping API functionality and server interactions. 

### The Endpoints
- **POST** /receipts/process     Inserts a receipt into our program.
- **GET** /receipts/:id/points    Retrieves point total from a receipt in our program.

## How to run the API server using Docker

Straightforward.

1. Download the repo
2. Enter directory of receipt-processor-challenge
3. Open a command window (git bash, bash, linux terminal, etc)
4. Run `docker-compose build receipt-processor`
5. Run `docker-compose run receipt-processor`
You'll know you succeeded when you see `[GIN-debug] Listening and serving HTTP on localhost:1313`
6. Open a new command window (git bash, bash, linux terminal, etc) 

The rest of these commands are optional if you know what you are doing. Here's what I did
1. Run docker ps in the new command window
2. Look at the **contianer ID** for the image named **receipt-processor:latest**
3. Run ```docker exec -it 3cc4f55b270b sh``` 

This puts you into the container where the Server and API are running. From here you can run all of the following commands in this document no problem!


## How to Use the API / example requests and responses.
Also fairly straightforward. Once your server is up (serving on localhost:1313) do the following

### Sample Post cURL request command
This is the **POST** Request for the **/receipts/process endpoint**.
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
      },{
        "shortDescription": "Emils Cheese Pizza",
        "price": "12.25"
      },{
        "shortDescription": "Knorr Creamy Chicken",
        "price": "1.26"
      },{
        "shortDescription": "Doritos Nacho Cheese",
        "price": "3.35"
      },{
        "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
        "price": "12.00"
      }
    ],
    "total": "36.53"
  }' \
  http://localhost:1313/receipts/process
  ```

### Sample Post response
```{"id":"c67d2ed2-479f-4bce-95eb-0a493a104465"}```

### Sample Get cURL request command
This is the **GET** Request for the **/receipts/:id/points**  endpoint.
```curl http://localhost:1313/receipts/{ID_RESPONSE_FROM_POST_REQUEST}/points```

### Sample GET response
```{"points":"28"}```


# Running the tests
Within this Repo there are tests that can be ran to verify the accuracy of my functions. 
Assuming you did the optional steps to be inside of the docker container, here's the command you can run.
```
go test -v
```
Which will run all of the tests, and you will know that the tests are successful if the last line looks like
```
ok      workspace       0.007s
```
