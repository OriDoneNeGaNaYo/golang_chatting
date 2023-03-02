package main

import (
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "frontend/home.html")
}

func main() {
	router := httprouter.New()
	router.GET("/", serveHome)

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run(":8080")
	
	//world := newWorld()
	//go world.run()

	//http.HandleFunc("/", serveHome)
	//
	//http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
	//	conn, err := upgrader.Upgrade(w, r, nil)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	client := NewClient(world, conn)
	//	client.world.ChanEnter <- client
	//})
	//
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
