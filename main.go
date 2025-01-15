package main

import (
	"net/http"
	"receipt-processor/handler"  
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /receipts/process", handler.ProcessReceiptHandler) 
	router.HandleFunc("GET /receipts/{id}/points", handler.GetReceiptPointsHandler)
	
	http.ListenAndServe(":8080", router)
}
