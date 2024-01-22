package main

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var infoLog, errorLog, debugLog *log.Logger

var db *gorm.DB

var err error

func CreateLoggers() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	infoLog = log.New(file, "INFO ", log.LstdFlags|log.Lshortfile)
	errorLog = log.New(file, "ERROR ", log.LstdFlags|log.Lshortfile)
	debugLog = log.New(file, "DEBUG ", log.LstdFlags|log.Lshortfile)
}

func OpenDBConnection() {
	sql := "root:admin123@tcp(localhost:3306)/assignment"
	db, err = gorm.Open(mysql.Open(sql), &gorm.Config{})
	if err != nil {
		errorLog.Println("Error Connecting to the Database!")
		panic(err)
	}
	debugLog.Println("DB Connection Established Successfully")
	result := db.Debug().AutoMigrate(Player{})
	if result != nil {
		debugLog.Println(result.Error())
		panic("Table not created successfully")
	}
	debugLog.Println("Player table created successfully")
}

func GetAddPlayerList() []*Player {
	p1 := &Player{18, "Virat Kohli", 34, 78}
	p2 := &Player{45, "Rohit Sharma", 36, 45}
	p3 := &Player{7, "M. S. Dhoni", 42, 16}
	p4 := &Player{77, "Shubman Gill", 24, 6}
	p5 := &Player{63, "Surya Kumar Yadav", 33, 3}
	p6 := &Player{31, "Ruturaj Gaikwad", 26, 0}

	return []*Player{p1, p2, p3, p4, p5, p6}
}

func GetUpdatedPlayerList() []*Player {
	p1 := &Player{18, "Virat Kohli", 35, 79}         //changed age and No. of centuries
	p2 := &Player{45, "Rohit Sharma", 36, 46}        // changed No. of Centuries
	p3 := &Player{7, "Mahendra Singh Dhoni", 42, 17} // Changed No. of centuries and name

	/*p[0].NoOfCenturies++
	p[0].Age++
	p[1].NoOfCenturies++
	p[2].Name = "Mahendra Singh Dhoni"*/

	return []*Player{p1, p2, p3}
}
