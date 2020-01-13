# golang_RestEndpoints_foods

This project displays foods in a Restaurant.

For this project I used postman a browser extention.

Use `go build food.go && ./food` to run the project

## Display all

http://localhost:6921/foods/ displays all foods and their respective ID's. 
Make sure to be in the "Get" method.

result:
[{"id":"1","name":"Pizza"},{"id":"2","name":"Burger"},{"id":"3","name":"Taco"},{"id":"4","name":"French Fries"},{"id":"5","name":"Smoothie"},{"id":"6","name":"Noodles"}]

## Display by ID

http://localhost:6921/foods/(id) displays the food associated with a particular ID. 
Make sure to be in the "Get" method.

## Create

http://localhost:6921/foods/create/(id) Creates a food with the id entered. 
Make sure to be in the "Post" method.
Add the new value in the body section in the header.

e.g 
{
"name":"French Fries"
}

## Delete

http://localhost:6921/foods/delete/(id) deletes the item with the ID entered. 
Make sure to be in the "Delete" method.

## Update

http://localhost:6921/foods/update/(id) allows you to edit a food. 
Make sure to be in the "Put" method.

Add the new value in the body section in the header.
e.g
{
"name":"French Fries"
}
