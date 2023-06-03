package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
	"errors"
    "io/ioutil"
)

type Client struct{
    ID  uint32
    Name string
    Phone string
    Adress string
    Ordre string 
    State uint32
} 

func AddAnOrdre() error {
	url := "http://localhost:8080/Commande"
    data := Client{Name: "hamidhamid", Phone: "0632343698", Adress: "Benguerir", Ordre: "Kabba chamiya", State:  0}
    payload, err := json.Marshal(data)
    if err != nil {
        return errors.New("Json Encoding")
    }
    fmt.Println(payload)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        return errors.New("Setuping The Request")
	}

    req.Header.Set("Content-Type", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
		return errors.New("Sending The request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return errors.New("Reading The response body")
    }
    fmt.Println(resp.Status)
    fmt.Println(string(body))
	return nil
}

func Get_the_DB() error {
	url := "http://localhost:8080/Admin"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return errors.New("Setuping The Request")
	}
    req.Header.Set("Authorization", "bearer mLGV?uxpQ0rfsy-))CGW.1!=r#dKOiZe5#G]|B4Xi@tPuVF2SdubREaTAK#}DO3I")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
		return errors.New("Sending The request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return errors.New("Reading The response body")
    }
    fmt.Println(resp.Status)
    fmt.Println(string(body))
	return nil
}

func ChangeState() error {
    url := "http://localhost:8080/Admin"
    data := 4
    payload, err := json.Marshal(data)
    if err != nil {
        return errors.New("Json Encoding")
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        return errors.New("Setuping The Request")
	}
    req.Header.Set("Authorization", "bearer mLGV?uxpQ0rfsy-))CGW.1!=r#dKOiZe5#G]|B4Xi@tPuVF2SdubREaTAK#}DO3I")
    req.Header.Set("Content-Type", "application/json")
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
		return errors.New("Sending The request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return errors.New("Reading The response body")
    }
    fmt.Println(resp.Status)
    fmt.Println(string(body))
	return nil
}

var err error

func main() {
    // err = Get_the_DB()
    // err = ChangeState()
    // err = Get_the_DB()
    err = AddAnOrdre()
	if (err != nil){
		fmt.Println("Error: ", err)
	}
}
// {"Name": "Hamid", "Phone": "0623313463", "Adress": "Rabat", "Commande": "Tacos"}