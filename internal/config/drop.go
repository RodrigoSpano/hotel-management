package database

import "gorm.io/gorm"

func DropDbTable(db *gorm.DB, table string) {
	db.Exec("DROP TABLE " + table)
}
