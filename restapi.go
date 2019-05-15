package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "9000", "Port to listen on")
)

var results []string

// GetHandler handles the index route
func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("/ReceiveFile --> File name - %s\n", name[0])
	io.Copy(&Buf, file)
	contents := Buf.String()
	fmt.Println("------------")
	fmt.Println(contents)
	fmt.Println("------------")

	fmt.Println("------------")
	fmt.Println("Struct")
	fmt.Println("------------")
	clients := []*Client{}
	if err := gocsv.UnmarshalBytes([]byte(contents), &clients); err != nil {
		panic(err)
	}
	for _, client := range clients {
		fmt.Println("Hello", client.Name)
	}
	fmt.Println("------------")
	Buf.Reset()
	return
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
}

func UploadAppointments(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("/ReceiveFile --> File name - %s\n", name[0])
	io.Copy(&Buf, file)
	contents := Buf.String()
	fmt.Println("------------")
	fmt.Println(contents)
	fmt.Println("------------")

	fmt.Println("------------")
	fmt.Println("Struct")
	fmt.Println("------------")
	appts := []*AppointmentReq{}
	if err := gocsv.UnmarshalBytes([]byte(contents), &appts); err != nil {
		panic(err)
	}
	for _, app := range appts {
		fmt.Println("Appointment - ", &app)
	}
	fmt.Println("------------")
	Buf.Reset()
	return
}

func main() {
	results = append(results, time.Now().Format(time.RFC3339))

	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHandler)
	mux.HandleFunc("/upload", UploadAppointments)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

type AppointmentReq struct { // Our example struct, you can use "-" to ignore a field
	Patient  string `csv:"Patient"`
	Provider string `csv:"Provider"`
	Appdate  string `csv:"Adate"`
	Status   string `csv:"Status"`
}
