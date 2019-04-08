package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"middleware"
	"mymux"
	"net/http"
	"net/http/httputil"
	car "pb"

	"github.com/golang/protobuf/proto"
)

func init() {
	mux := mymux.GetMux()
	mux.HandleFunc("/requestbin", middleware.WithMiddleware(
		requestbin,
		middleware.Time()))
}

type CarMessage struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int32  `json:"year"`
	Color string `json:"color"`
}

func requestbin(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Not OK")
		return
	}
	fmt.Println(string(requestDump))
	/*buf, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Not OK")
		return
	}
	//carmessage := bytesToProto(buf)
	//carmessage := bytesToJSON(buf)
	*/
	carmessage := CarMessage{}
	carmessage.Make = r.PostFormValue("make")
	carmessage.Model = r.PostFormValue("model")
	temp, _ := strconv.ParseInt(r.PostFormValue("year"), 10, 32)
	carmessage.Year = int32(temp)
	carmessage.Color = r.PostFormValue("color")
	fmt.Println(carmessage.Make, carmessage.Model, carmessage.Year, carmessage.Color)
	fmt.Fprintf(w, "OK")
}

func bytesToProto(buf []byte) *CarMessage {
	message := &car.CarMessage{}
	if err := proto.Unmarshal(buf, message); err != nil {
		log.Fatalln("Failed to parse message:", err)
	}
	carmessage := &CarMessage{}
	carmessage.Make = message.GetMake()
	carmessage.Model = message.GetModel()
	carmessage.Year = message.GetYear()
	carmessage.Color = message.GetColor()
	return carmessage
}

func bytesToJSON(buf []byte) *CarMessage {
	carmessage := &CarMessage{}
	if err := json.Unmarshal(buf, carmessage); err != nil {
		fmt.Println(err)
	}
	return carmessage
}
