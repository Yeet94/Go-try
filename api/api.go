package api
//This file contains the business logic and data structures for your API. 
// It defines the request and response formats, as well as any helper functions
// or handlers that process API requests.
import(
	"encoding/json"
	"net/http"
)

//Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

//Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	Balance int64
}

//Error Response
type Error struct{
	//Error Code
	Code int

	//Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWrite){
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)

