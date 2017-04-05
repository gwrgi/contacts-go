package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gwrgi/contacts-go/types"
	_ "github.com/mattn/go-sqlite3"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetUser",
		"GET",
		"/contacts/v1/users",
		GetUser,
	},
	Route{
		"GetContacts",
		"GET",
		"/contacts/v1/contacts",
		GetContacts,
	},
	Route{
		"Budget_GetTransaction",
		"GET",
		"/budget/v1/transactions",
		GetTransactions,
	},
	Route{
		"Budget_PostTransaction",
		"POST",
		"/budget/v1/transactions",
		PostTransaction,
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunction)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	data, err := base64.StdEncoding.DecodeString(r.Header["Authorization"][0][6:])
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hello, %q", data)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	var expense types.Expense
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	checkErr(err)

	err = r.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &expense); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	fmt.Printf("Expense=(%+v)\n", expense)

	e := RecordExpense(expense)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}

func RecordExpense(e types.Expense) types.Expense {
	db, err := sql.Open("sqlite3", "./budget.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO expenses (merchant, amount, category_id, date) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(e.Merchant, e.Amount, e.CategoryId, e.Date.Unix())
	checkErr(err)

	_, err = res.LastInsertId()
	checkErr(err)

	return e
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
