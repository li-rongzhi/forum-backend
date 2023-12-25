// package database

// type Database struct {
// }

// func GetDB() (*Database, error) {
// 	return &Database{}, nil
// }
// package database

// import (
//     "gorm.io/gorm"
//     // other imports
// )

// var db *gorm.DB // This should be your GORM database instance

// func GetDB() *gorm.DB {
//     return db
// }

// package database

// import (
//     "gorm.io/gorm"
// )

// type Database struct {
//     GormDB *gorm.DB
// }

// func GetDB() (*Database, error) {
// 	return &Database{}, nil
// }
// package database

// import (
//     // "gorm.io/driver/mysql"
//     "gorm.io/gorm"
// )

// type Database struct {
//     GormDB *gorm.DB
// }

// // GetDB should return an existing *Database instance with an established connection
// var dbInstance *Database

// func GetDB() *Database {
//     return dbInstance
// }
package database

import "gorm.io/gorm"

var DB *gorm.DB  // This holds the global database connection

