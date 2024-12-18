package main

import (
	"hr-system/common/config"
	"hr-system/common/dao/models"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gen"
	"gorm.io/gorm"
)

var dataMap = map[string]func(gorm.ColumnType) (dataType string){
	"varchar": func(columnType gorm.ColumnType) (dataType string) {
		if n, ok := columnType.Nullable(); ok && n {
			return "sql.NullString"
		}
		return "string"
	},

	"date": func(columnType gorm.ColumnType) (dataType string) {
		if n, ok := columnType.Nullable(); ok && n {
			return "sql.NullTime"
		}
		return "time.Time"
	},

	"timestamp": func(columnType gorm.ColumnType) (dataType string) {
		if n, ok := columnType.Nullable(); ok && n {
			return "sql.NullTime"
		}
		return "time.Time"
	},

	"bigint": func(columnType gorm.ColumnType) (dataType string) {
		if n, ok := columnType.Nullable(); ok && n {
			return "sql.NullInt64"
		}

		return "uint"
	},

	"decimal": func(columnType gorm.ColumnType) (dataType string) {
		if n, ok := columnType.Nullable(); ok && n {
			return "sql.NullFloat64"
		}

		return "float64"
	},
}

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id uint) (*gen.T, error)
}

func main() {
	g := newGenerator()

	// 應用基本 CRUD
	g.ApplyBasic(
		models.User{},
		models.Position{},
		models.Employee{},
		models.Department{},
	)

	// 應用自訂接口
	g.ApplyInterface(func(Querier) {},
		models.User{},
		models.Position{},
		models.Employee{},
		models.Department{})

	g.Execute()
}

func newGenerator() *gen.Generator {
	cfg := config.Get()
	_db, err := gorm.Open(mysql.Open(cfg.Mysql.DSN), &gorm.Config{
		AllowGlobalUpdate: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./common/dao/query",
		ModelPkgPath: "./common/dao/models",

		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,

		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(_db)

	g.WithDataTypeMap(dataMap)
	return g
}
