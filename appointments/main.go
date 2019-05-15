package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type AppointmentReq struct {
	Patient  string `csv:"Patient"`
	Provider string `csv:"Provider"`
	Appdate  string `csv:"Adate"`
	Status   string `csv:"Status"`
}

func main() {
	file, err := os.OpenFile("../DailyAppointments.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	appts := []*AppointmentReq{}

	if err := gocsv.UnmarshalFile(file, &appts); err != nil {
		panic(err)
	}
	for _, a := range appts {
		fmt.Println("Appointment", a)
	}

	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}

}
