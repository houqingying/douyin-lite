package main

import (
	"github.com/houqingying/douyin-lite/dal/db"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/db/gen",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(db.GetDsn()))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//
	//g.ApplyBasic(
	//	// Generate struct `User` based on table `users`
	//	g.GenerateModel("users"),
	//
	//	// Generate struct `Employee` based on table `users`
	//	g.GenerateModelAs("users", "Employee"),
	//
	//
	//	// Generate struct `User` based on table `users` and generating options
	//	g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	//
	//)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
