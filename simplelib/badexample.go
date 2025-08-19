package simplelib

import (
	"database/sql"
	"fmt"
)

func main() {
	var tx *sql.Tx // uninitialized transaction

	// Attempting to use Exec before Tx is properly started and wrapped
	res, err := tx.Exec("DELETE FROM users WHERE id = ?", 42)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	fmt.Println("Rows affected:", res)
}
