# README.md

# Arithmetic Expression Calculator API

This project provides a simple HTTP web service that evaluates arithmetic expressions. Users can send an expression via a POST request, and the service will return the result of the evaluation or an error message.

## API Endpoint

### POST /api/v1/calculate

#### Request Body

The request should be in JSON format with the following structure:

```json
{
    "expression": "expression entered by the user"
}
Response
Success (HTTP 200)
If the expression is valid and successfully evaluated, the response will be:

 
{
    "result": "result of the expression"
}
Error (HTTP 422)
If the expression is invalid (contains unsupported characters), the response will be:

 
{
    "error": "Expression is not valid"
}
Internal Server Error (HTTP 500)
If there is any other error (e.g., during processing), the response will be:

 
{
    "error": "Internal server error"
}
Usage Examples
Successful Calculation
To evaluate an expression:

 
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2 * 2"
}'
Response:

 
{
    "result": "6"
}
Invalid Expression (422 Error)
If the expression contains invalid characters:

 
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2 * a"
}'
Response:

 
{
    "error": "Expression is not valid"
}
Internal Server Error (500 Error)
If an unexpected error occurs:

 
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 / 0"
}'
Response:

 
{
    "error": "Internal server error"
}
Running the Project
To run the project, use the following command:

 
Copy:
go run ./cmd/calc_service/...
