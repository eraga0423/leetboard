package leetboard

import (
	"database/sql"
	"fmt"
)

func TxAfter(tx *sql.Tx, getError error) error {
	p := recover()
	if p != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}

	} else if getError != nil {
		err := tx.Rollback()
		if err != nil {
			return fmt.Errorf("Main error: %w and Error when Rollback %w", getError, err)
		}
	}
	return nil
}
