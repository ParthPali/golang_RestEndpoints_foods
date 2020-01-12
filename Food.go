package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Food struct {
    ID        string   `json:"id,omitempty"`
    Name      string   `json:"name,omitempty"`
}

var foods []Food

func GetFoodEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range foods {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Food{})
}

func GetFoodsEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(foods)
}

func CreateFoodEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var newfood Food
    _ = json.NewDecoder(req.Body).Decode(&newfood)
    newfood.ID = params["id"]
    foods = append(foods, newfood)
    json.NewEncoder(w).Encode(foods)
}

func UpdateFoodEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var newfood Food
    _ = json.NewDecoder(req.Body).Decode(&newfood)
    newfood.ID = params["id"]

    for index, item := range foods {
        if item.ID == params["id"] {
            //foods = append(foods[:index], newfood)
            foods = append(foods[:index], foods[index+1:]...)
            foods = append(foods, newfood)
            //json.NewEncoder(w).Encode(newfood)
            break
        }
    }
    json.NewEncoder(w).Encode(foods)
}

func DeleteFoodEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range foods {
        if item.ID == params["id"] {
            foods = append(foods[:index], foods[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(foods)
}

func main() {
    router := mux.NewRouter()
    foods = append(foods, Food{ID: "1", Name: "Pizza"})
    foods = append(foods, Food{ID: "2", Name: "Burger"})
    foods = append(foods, Food{ID: "3", Name: "Taco"})
    foods = append(foods, Food{ID: "4", Name: "French Fries"})
    foods = append(foods, Food{ID: "5", Name: "Smoothie"})
    foods = append(foods, Food{ID: "6", Name: "Noodles"})
    router.HandleFunc("/foods", GetFoodsEndpoint).Methods("GET")
    router.HandleFunc("/foods/{id}", GetFoodEndpoint).Methods("GET")
    router.HandleFunc("/foods/create/{id}", CreateFoodEndpoint).Methods("POST")
    router.HandleFunc("/foods/update/{id}", UpdateFoodEndpoint).Methods("PUT")
    router.HandleFunc("/foods/delete/{id}", DeleteFoodEndpoint).Methods("DELETE")
    err := http.ListenAndServe(":6921", router)
    if err != nil { log.Fatal("ListenAndServe: ", err) }
}