package models

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"os"
	"reflect"
	"golang-starter/config"
	"golang-starter/helpers"
	"strconv"
	"strings"
)

type MigrationTables struct{}

/**
* first loop in all migrations files
* get all migration methods
* drop related table if .env have delete attribute
* Call migration function
 */
func MigrateAllTable(path string) {
	var t MigrationTables
	migrateFiles := helpers.ReadAllFiles(path)
	for _, file := range migrateFiles {
		filepath := strings.Split(file, ".")
		fileName := filepath[0]
		if fileName != "migration" {
			functionName := strcase.ToCamel(filepath[0]) + "Migrate"
			deleteTables, _ := strconv.ParseBool(os.Getenv("DROP_ALL_TABLES"))
			if deleteTables {
				config.DB.DropTableIfExists(inflection.Plural(filepath[0]))
			}
			reflect.ValueOf(&t).MethodByName(functionName).Call([]reflect.Value{})
		}
	}
}
