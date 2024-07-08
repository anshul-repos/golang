# Rest API Server

## Types of REST API calls
- GET, POST, PUT, PATCH, DELETE
- PUT vs PATCH : 
   - PUT is method of modifying resources OR updates the entire resource. Similar to POST
   - PATCH  PATCH does partial update OR only that field is updated without modifying the other fields.

## Error Codes

- 1XX : Informational
- 2XX : Success      (ex 200-OK, 201-Created, 202-Accepted)
- 3XX : Redirection
- 4XX : Client Error (ex 400 - BadRequest, 401-Unauthorized, 403-Forbidden)
- 5XX : Server Error  (ex 500-Internal Server Error, 502-Bad Gateway)

## Gorilla Mux

- mux.NewRouter()
- mux.NewRouter.HandleFunc()
- mux.Vars()

## Json package

### Json.NewEncoder().Encode() :
- used in GET method
- writes the data to (w) response.Writer

### Json.NewDecoder().Decode() :
- used in POST method
- reads the data from (r) request
