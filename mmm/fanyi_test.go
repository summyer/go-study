package mmm

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestGetArchivesCopyMap(t *testing.T) {
	var url = "pig:pLlGJSpdg1RCp^80@tcp(10.20.152.213:3306)/mmm?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB = GetConnection(url)
	var archivesCopyMap map[string]ArchivesCopy
	archivesCopyMap = GetArchivesCopyMap(db, []string{"DM-202203-000001", "DM-202203-000002"})
	for k, v := range archivesCopyMap {
		fmt.Println(k, v)
	}
}
