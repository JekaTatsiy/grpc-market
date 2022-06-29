package server

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

//go:embed index.json
var indexJson []byte

const IndexName = "market"
const login = "elastic"
const password = "elastic"

func (g *GServer) ES(method, path string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, e := http.NewRequest(method, fmt.Sprintf("http://%s/%s", g.ESaddr, path), nil)
	if e != nil {
		return nil, e
	}
	req.SetBasicAuth(login, password)
	resp, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	b, _ := io.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}
	return b, e
}

func (g *GServer) IndexCreateIfNotExist() error {
	ok := false
	var resp []byte
	var e error
	for i := 0; i < 60; i++ {
		resp, e = g.ES(http.MethodGet, IndexName, []byte{})
		if e != nil {
			fmt.Println(i, resp, e, time.Now())
			time.Sleep(time.Second)
		} else {
			ok = true
			break
		}
	}
	if ok {
		fmt.Println("\nconnected to es")
		var obj map[string]interface{}

		e := json.Unmarshal(resp, &obj)
		if e != nil {
			return e
		}
		_, ok := obj["error"]
		if !ok {
			return nil
		}

		g.ES(http.MethodPut, IndexName, indexJson)
		return nil
	} else {
		fmt.Println("\nerror connected to es")
		return errors.New("error connected to es")
	}
}

func (g *GServer) ESAddOne() {

}

func (g *GServer) ESAdd() {

}

func (g *GServer) ESDeleteOne() {

}

func (g *GServer) ESDelete() {

}
