Example

User create

```
curl --location 'localhost:8080/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Username":"AlexUser",
    "Password":"SecurePassword",
    "Name":"Alex Mel",
    "Email":"alex@alex.com.ar"
}'
```
User Login
```
curl --location 'localhost:8080/login/' \
--header 'Content-Type: application/json' \
--data '{
    "Username":"Alex",
    "Password":"SecurePassword"
}'
```
UrlShort create
```
curl --location 'localhost:8080/shorter/' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTY4ODAwNDIzMSwidXNlcm5hbWUiOiJBbGV4In0.Ohp2R9q1CxgRln_082QyRejQcyV0V4ojqU9MUEl7Iig' \
--data '{
    "url": "https://www.google.com"
}'
```

Go to page using UrlShorted
```
curl --location 'localhost:8080/RPubNtMWUlOF'
```

Delete UrlShorted
```
curl --location --request DELETE 'localhost:8080/shorter/RPubNtMWUlOF' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTY4ODAwNDIzMSwidXNlcm5hbWUiOiJBbGV4In0.Ohp2R9q1CxgRln_082QyRejQcyV0V4ojqU9MUEl7Iig'
```