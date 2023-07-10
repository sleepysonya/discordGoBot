package birthday

import (
	"encoding/csv"
	"log"
	"os"

	storage "github.com/sleepysonya/discordGoBot/util"
)

func AddCol(name string, day string, month string) string {
	var response string = ""
	f, err := os.Open("birthday.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	f.Close()
	wr, err := os.OpenFile("birthday.csv", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer wr.Close()
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(wr)

	for _, record := range records {
		if record[0] == name {
			response = "Birthday has been updated"
			record[1] = day
			record[2] = month
			w.Write([]string{record[0], record[1], record[2]})
		} else {
			w.Write([]string{record[0], record[1], record[2]})
		}
	}
	if response == "" {
		response = "Birthday has been added"
		w.Write([]string{name, day, month})
	}
	w.Flush()

	return response
}

func GetCol(name string) string {
	f, err := os.Open("birthday.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		if record[0] == name {
			return record[1]
		}
	}
	return ""
}

func GetAllCols() []storage.BirthType {
	var response []storage.BirthType = []storage.BirthType{}
	f, err := os.Open("birthday.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		response = append(response, storage.BirthType{
			Id: record[0], Day: record[1], Month: record[2]})
	}
	return response
}
