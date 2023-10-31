package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const file string = "data.db"

var ch chan string = make(chan string, 10)

func main() {
	/*	serialPort, _ := serial.Open("COM7", &serial.Mode{BaudRate: 9600})
		defer serialPort.Close()*/
	fmt.Println("Connected to serial port")
	//simple buffered channel
	//read from serial

	go func(ch *chan string) {
		for {
			fmt.Println(len(*ch))
			time.Sleep(2 * time.Second)
			if len(*ch) != 10 {
				*ch <- string("BISHOP")
				fmt.Println("Data Sent")
			}
			fmt.Println("Data In Transit")

		}

	}(&ch)

	r := mux.NewRouter()
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("C:\\Users\\Aiman\\Desktop\\projectzero\\dist\\assets")))
	fs2 := http.FileServer(http.Dir("./dist"))
	r.Handle("/", fs2)
	r.PathPrefix("/assets/").Handler(fs)
	r.HandleFunc("/stats", StatsHandler)
	r.HandleFunc("/ws", WsHandler)
	http.Handle("/", r)
	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		log.Fatalf(err.Error())

	}
}

func StatsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	_, err := writer.Write(
		[]byte(
			`<p style="">21 Things Have Been Collected</p>
			 <p>Hoping For More</p>
				`))
	if err != nil {

	}
	return
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))

}
func reader(conn *websocket.Conn, c *chan string) {
	//websocket reader that sends data received from go channel to the client
	for {
		select {
		case msg := <-*c:
			fmt.Println(len(ch))
			err := conn.WriteMessage(1, []byte(msg+" "+time.Now().String()))
			if err != nil {
				log.Println(err)
				return
			}
		default:
			fmt.Println(len(ch))
			time.Sleep(1 * time.Second)
		}
	}
}

func WsHandler(writer http.ResponseWriter, request *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)

	}
	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	go reader(ws, &ch)
}
