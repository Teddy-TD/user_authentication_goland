package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql" 
    "log"
)

func Connect(){

	_, err := gorm.Open(mysql.Open("root:@/golang"), &gorm.Config{})

    if err != nil {
        panic("Couldn't open database connection")
    } else {  
        log.Println("Database connection established")
    }
}