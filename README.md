
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
