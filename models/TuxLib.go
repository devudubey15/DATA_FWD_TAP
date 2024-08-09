package models

import (
	"DATA_FWD_TAP/common"
	"fmt"
	"log"

	"gorm.io/gorm"
)

/**********************************************************************************/
/*                                                                                */
/*  Constants:                                                                    */
/*    - TRAN_TIMEOUT: Placeholder for transaction timeout.                        */
/*    - REMOTE_TRNSCTN: Constant for remote transaction type.                     */
/*    - LOCAL_TRNSCTN: Constant for local transaction type.                       */
/*                                                                                */
/*  Structs:                                                                      */
/*    - TransactionManager: Holds a GORM DB transaction instance.                 */
/*                                                                                */
/*  Functions:                                                                    */
/*    - FnBeginTran: Begins a new transaction and determines if it is local or    */
/*      remote based on the database connection status.                           */
/*                                                                                */
/*    - PingDatabase: Checks if the database connection is alive by pinging it.   */
/*                                                                                */
/*    - FnCommitTran: Commits the transaction if it is local, or skips commit     */
/*      for remote transactions.                                                  */
/*                                                                                */
/**********************************************************************************/

const (
	TRAN_TIMEOUT   = 300 // placeholder for transaction timeout . i have to use this .
	REMOTE_TRNSCTN = 2
	LOCAL_TRNSCTN  = 1
)

type TransactionManager struct {
	Tran       *gorm.DB
	DbAccessor common.DBAccessor
}

func (tm *TransactionManager) FnBeginTran(serviceName string) int {

	// Check if a transaction is already active (GORM does not provide tpgetlev equivalent, manage this externally if needed)

	var TranType int
	db := tm.DbAccessor.GetDB(serviceName) // Get the database instance on this instance we will check if there is any active transaction.

	err := PingDatabase(serviceName, db) // Use PingDatabase to check if there is any active transaction.

	if err == nil {
		tm.Tran = db.Begin() // Begin the transaction here. This line represents something unique. When we call db.Begin(), it returns a new *gorm.DB instance, and we are supposed to use this instance throughout the transaction. This is because we have to maintain the 'isolation'.
		if tm.Tran.Error != nil {
			errMsg := fmt.Sprintf("[%s] tpbegin error: %v", serviceName, tm.Tran.Error)
			log.Println("L31200", errMsg)
			return -1
		}
		TranType = LOCAL_TRNSCTN
		log.Printf("[%s] EBA : tranType:%d", serviceName, TranType)
		return TranType
	} else {
		TranType = REMOTE_TRNSCTN
		log.Printf("[%s] EBA : tranType:%d", serviceName, TranType)
		return TranType
	}
}

func PingDatabase(serviceName string, db *gorm.DB) error {
	sqlDB, err := db.DB() // Get the underlying *sql.DB
	if err != nil {
		log.Printf("[%s] Failed to get underlying *sql.DB: %v", serviceName, err)
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		log.Printf("[%s] Failed to ping database: %v", serviceName, err)
		return err
	}
	return nil
}

func (tm *TransactionManager) FnCommitTran(serviceName string, tranType int) int {

	if tranType == LOCAL_TRNSCTN {
		if err := tm.Tran.Commit().Error; err != nil {
			errMsg := fmt.Sprintf("[%s] tpcommit error: %v", serviceName, err)
			log.Println(errMsg)
			tm.Tran.Rollback()
			return -1
		}
		log.Println("[%s] Transaction committed", serviceName)
	} else if tranType == REMOTE_TRNSCTN {
		// No log needed
		// here we are not commiting the transaction because the transaction type is "REMOTE_TRNSCTN"
		log.Printf("[%s] Skipping commit for remote transaction", serviceName)
	} else {
		errMsg := fmt.Sprintf("[%s] Invalid Transaction Number |%d|", serviceName, tranType)
		log.Println(errMsg)

	}
	return 0
}
