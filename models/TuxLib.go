package models

import (
	"DATA_FWD_TAP/config"
	"fmt"

	"log"

	"gorm.io/gorm"
)

const (
	TRAN_TIMEOUT   = 300 // placeholder for transaction timeout
	REMOTE_TRNSCTN = 2
	LOCAL_TRNSCTN  = 1
)

type TransactionManager struct {
	Tran *gorm.DB
}

func (tm *TransactionManager) FnBeginTran() int {

	// Check if a transaction is already active (GORM does not provide tpgetlev equivalent, manage this externally if needed)

	var TranType int
	db := config.GetDB() // here we are getting the database instance and then we are checking if there is any active transaction going on.

	err := PingDatabase(db) // here we are using PingDatabase to check if there is any active transaction

	if err == nil {
		tm.Tran = db.Begin() // Transaction begins here. this line represents something unique. when we call db.Begin() it returns a a new *gorm.db instance and we are suppose to use this instance through out the transaction. this is because we have to maintain the 'isolation'.
		if tm.Tran.Error != nil {
			errMsg := fmt.Sprintf("tpbegin error: %v", tm.Tran.Error)
			log.Println("L31200", errMsg)
			return -1
		}
		TranType = LOCAL_TRNSCTN
		log.Printf("EBA : tranType:%d:", TranType)
		return TranType
	} else {
		TranType = REMOTE_TRNSCTN
		log.Printf("EBA : tranType:%d:", TranType)
		return TranType
	}
}

func PingDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB() // Get the underlying *sql.DB
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (tm *TransactionManager) FnCommitTran(tranType int) int {

	if tranType == LOCAL_TRNSCTN {
		if err := tm.Tran.Commit().Error; err != nil { // it is suppose to commit but here we are encountering the error
			errMsg := fmt.Sprintf("tpcommit error: %v", err)
			log.Println(errMsg)
			tm.Tran.Rollback()
			return -1
		}
		log.Println("Transaction committed")
	} else if tranType == REMOTE_TRNSCTN {
		// No log needed
		// here we are not commiting the transaction because the transaction type is "REMOTE_TRNSCTN"
	} else {
		errMsg := fmt.Sprintf("Invalid Transaction Number |%d|", tranType)
		log.Println(errMsg)

	}
	return 0
}
