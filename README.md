# golang_RestEndpoints_foods

This project displays foods in a Restaurant.

For this project I used postman a browser extention.

Use `go build food.go && ./food` to run the project

## Display all

http://localhost:6921/foods/ displays all foods and their respective ID's. 
Make sure to be in the "Get" method.

![Alt text](/Users/Parth/Documents/GooglecodeIn/go/src/RestEndpoints.jpg?raw=true "Optional Title")

## Display by ID

http://localhost:6921/foods/(id) displays the food associated with a particular ID. 
Make sure to be in the "Get" method.

## Create

http://localhost:6921/foods/create/(id) Creates a food with the id entered. 
Make sure to be in the "Post" method.
Add the new value in the body section in the header.

## Delete

http://localhost:6921/foods/delete/(id) deletes the item with the ID entered. 
Make sure to be in the "Delete" method.

## Update

http://localhost:6921/foods/update/(id) allows you to edit a food. 
Make sure to be in the "Put" method.

Add the new value in the body section in the header.
