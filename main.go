package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"sync"
	"context"
)

func KeySearch(data map[string]interface{},searchKey string ) string{
	if value, ok := data[searchKey]; ok {
	     return value.(string)
	} 
	     return ""
	}

func sendToWebhook(data OutgoingData) {	
	log.Println("sendToWebhook Finished....")
	jsonData, err := json.Marshal(data)
	
	if err != nil {
	fmt.Println("Error encoding JSON:", err)
	return
	}
	URL := "https://webhook.site/fc198e4d-5eb5-4650-8798-926bad5845e9"
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("ERROR",err)
		return
	}
	defer resp.Body.Close()
	log.Println("sendToWebhook Finished",resp.Status)
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		defer r.Body.Close()
		ctx:=context.Background()
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err!=nil{
			log.Println("Failed to decode request body:", err)
		}
		ctx=context.WithValue(ctx,"body",bodyBytes)
        var wg sync.WaitGroup

		wg.Add(1)
		go func (ctx context.Context,w http.ResponseWriter){
			  var bodyData map[string]interface{}
			  if body := ctx.Value("body"); body!= nil {
			  err := json.Unmarshal( body.([]byte), &bodyData)
			  if err != nil {
				  	log.Println("Failed to decode request body:", err)
			     }
				 transformedData := TransformData(bodyData)
				 sendToWebhook(transformedData)
				 w.WriteHeader(http.StatusOK)
				 w.Write([]byte("Request Recieved"))	
				 wg.Done()
				return
		      }
			 log.Println("BODY EMPTY")
		}(ctx,w)
        //GONE
	    wg.Wait()
	    })
       
		
	http.ListenAndServe(":8001", nil)
}