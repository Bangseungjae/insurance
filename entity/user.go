package entity

import (
	"Bangseungjae/insurance/util"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	Id           int
	Name         string
	Age          int
	Gender       string
	Occupation   string
	HealthStatus string
}

var (
	Users []User
)

func init() {
	log.Println("user init")
	filename := "./config/user.csv"
	readToUsers(filename)
	fmt.Println(Users)
}

func readToUsers(filename string) {
	file, err := os.Open(filename)
	util.Check(err)
	read := csv.NewReader(file)
	read.FieldsPerRecord = -1
	var tmp User
	data, err := read.ReadAll()
	util.Check(err)
	for _, d := range data[1:] {
		tmp.Id, _ = strconv.Atoi(d[0])
		tmp.Name = d[1]
		tmp.Age, _ = strconv.Atoi(d[2])
		tmp.Gender = d[3]
		tmp.Occupation = d[4]
		tmp.HealthStatus = d[5]
		Users = append(Users, tmp)
	}
}
