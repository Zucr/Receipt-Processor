package db

import (
	"fmt"
	"github.com/google/uuid"
	"receipt-processor/model"
	
)

var receipts = make(map[string]model.Receipt)

func ProcessReceipt(receipt model.Receipt) string {
	myRandomReceiptID := fmt.Sprintf(uuid.New().String())
	receipts[myRandomReceiptID] = receipt
	return myRandomReceiptID
}

func GetReceiptByID(receiptID string) (model.Receipt, bool) {
	fmt.Println("Receipt: ", receiptID)//might help with testing
	receipt, exists := receipts[receiptID]
	return receipt, exists
}
