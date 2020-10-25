### How to run Webservice
1. Clone this repository:
    `git clone https://github.com/danielwetan/backend-99`
2. Start the server:
    `cd backend-99/soal-3-4/go run main.go`
3. Run app in the browser on the port http://localhost:3000


**Request**:
- Method: `POST`
- Endpoint: `/api/palindrome`
- Header:
    - Content-Type: `application/json`
    - Accept: `application/json`
- Body:
```json
{
    "start": "int",
    "end": "int",
}
```

**Response**:
```json
{
    "status": "string",
    "body": {
      "totalPalindrome": "int",
    }
}
```
