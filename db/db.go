// Copyright (c) 2017-2026 Onur Yaşar
// Licensed under AGPL v3 + Commercial Exception
// See LICENSE.txt

// https://github.com/rymory/goutils/blob/main/LICENSE
// https://github.com/rymory/goutils/blob/main/README.md
// rymory.org 
// onuryasar.org
// onxorg@proton.me 

package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		panic("database url is empty")
	}

	conn, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
