package leetboard

import "database/sql"

func TxAfter(tx *sql.Tx, err error) {
	p := recover()
	if p != nil {
		tx.Rollback()
	} else if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
