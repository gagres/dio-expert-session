package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gagres/dio-expert-session-finance/model/transaction"
	"github.com/gagres/dio-expert-session-finance/util"
)

// GetTransactions get transactions function
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2020-04-05T11:45:26"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateTransaction create a new transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res) // o body vai ser convertido (se possível), para dentro do res

	fmt.Print(res)

	// w.Header().Set("Content-Type", "application/json")
}
