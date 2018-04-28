package main

import (
	"github.com/gorilla/mux"
	"log"
	jsonparse "encoding/json"
	"os"
	"io/ioutil"
	"github.com/gorilla/rpc/json"
	"net/http"
)


type Args struct{
	Id string
}

type Book struct{
	Id string '"json:string,omitempty"'
	Name string '"json:name,omitempty"'
	Author string '"json:author,omitempty"'
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error{
	var books []Book
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil{
		log.Println("error:",readerr)
		os.Exit(1)
	}
	marshalerr := jsonparse.Unmashal(raw, &books)
	if marshalerr != nil{
		log.Println("error:",marshalerr)
	}
	for _, book := range books{
		if book.Id == args.Id{
			*reply = book
			break
		}
	}
	return nil

}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(),"application/json")
	s.RegisterService(new(JSONServer),"")
	r := mux.NewRouter()
	r.Handle("/rpc",s)
	http.ListenAndServe(":123",r)
}
