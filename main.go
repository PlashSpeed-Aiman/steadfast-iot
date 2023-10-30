package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/stats", StatsHandler)
	r.HandleFunc("/ws", WsHandler)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", r)
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
func reader(conn *websocket.Conn) {
	messageType, p, err := conn.ReadMessage()
	for {
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Println(err)
			return
		}
		reply := string(p) + " " + time.Now().Format("Monday-01-2006 15:04:05")
		if err := conn.WriteMessage(messageType, []byte(reply)); err != nil {
			log.Println(err)
			return
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
	go reader(ws)
}
