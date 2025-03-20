package database

import (
	"fmt"

	"gorm.io/gorm"
)

func DropDbTable(db *gorm.DB, table string) {
	db.Exec("DROP TABLE " + table)
	fmt.Printf("table %s was dropped successfully", table)
}
