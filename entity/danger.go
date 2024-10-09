package entity

import (
	"Bangseungjae/insurance/util"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Danger struct {
	AgeRange   string
	Gender     string
	Occupation string
	RiskRage   float64
	FromAge    int
	ToAge      int
}

var (
	Dangers []Danger
)

func init() {
	log.Println("danger init")
	filename := "./config/danger_ratio.csv"
	readToDangers(filename)
	fmt.Println(Dangers)
}

func readToDangers(filename string) {
	file, err := os.Open(filename)
	util.Check(err)
	read := csv.NewReader(file)
	read.FieldsPerRecord = -1
	var tmp Danger
	data, err := read.ReadAll()
	util.Check(err)
	for _, d := range data[1:] {
		tmp.AgeRange = d[0]
		tmp.Gender = d[1]
		tmp.Occupation = d[2]
		tmp.RiskRage, err = strconv.ParseFloat(d[3], 64)
		ages := strings.Split(tmp.AgeRange, "-")
		tmp.FromAge, err = strconv.Atoi(ages[0])
		util.Check(err)
		tmp.ToAge, err = strconv.Atoi(ages[1])
		util.Check(err)
		Dangers = append(Dangers, tmp)
	}
}
