/*
Expose the REST endpoints for Products

base url :  http://localhost:8080/products

product = { id, name, description, cost }

Possible operations : GET, PUT, PATCH, DELETE, POST
//Now
GET - http://localhost:8080/products => return all the products
GET - http://localhost:8080/products/1 => return the product with id = 1 OR 404 if not found
POST - http://localhost:8080/products => Save the given product and return with 201 status

//Later
PUT - http://localhost:8080/products/1 => Update the product with id = 1 OR 404 if not found
DELETE - http://localhost:8080/products/1 => Delete the product with id = 1 OR 404 if not found
*/

/*
json.Marshal(product) & json.Unmarshal(product)

OR

json.NewEncoder(w).Encode(product) => to json
json.NewDecoder(r.Body).Decode(product) => from json
*/