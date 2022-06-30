package server

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
)

//go:embed index.json
var indexJson []byte

const IndexName = "market"
const login = "elastic"
const password = "elastic"

func (g *GServer) ES(method, path string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, e := http.NewRequest(method, fmt.Sprintf("http://%s/%s", g.ESaddr, path), bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
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
		g.ESCreateIndex()
		return nil
	} else {
		fmt.Println("\nerror connected to es")
		return errors.New("error connected to es")
	}
}

type SuggPart struct {
	LinkUrl string
	Title   string
	Query   []string
}

func (g *GServer) ESDeleteIndex() {
	g.ES(http.MethodDelete, IndexName, nil)
}
func (g *GServer) ESCreateIndex() {
	g.ES(http.MethodPut, IndexName, indexJson)
}
func (g *GServer) ESAdd(suggs []*pb.Suggest) error {
	body := []byte{}
	for _, x := range suggs {
		body = append(body, []byte(fmt.Sprintf("{\"index\": { \"_id\": %d }}\n", x.ID))...)
		s := SuggPart{LinkUrl: x.LinkUrl, Title: x.Title, Query: x.Queries}
		b, _ := json.Marshal(s)
		body = append(body, b...)
		body = append(body, '\n')
	}
	res, e := g.ES(http.MethodPost, fmt.Sprintf("%s/_bulk", IndexName), body)
	if e != nil {
		return e
	}
	var m map[string]interface{}
	e = json.Unmarshal(res, &m)
	if e != nil {
		return e
	}
	status, ok := m["errors"]
	if !ok {
		return errors.New("field ERRORS not found")
	}
	if status.(bool) {
		return errors.New("error when add suggest")
	}
	return nil
}

func (g *GServer) ESDeleteOne(ind int32) error {
	response, e := g.ES(http.MethodDelete, fmt.Sprintf("%s/_doc/%d", IndexName, ind), nil)
	if e != nil {
		return e
	}

	var fields map[string]interface{}
	e = json.Unmarshal(response, &fields)
	if e != nil {
		return e
	}
	result, ok := fields["result"]
	if !ok {
		return errors.New("expect field \"result\"")
	}
	if result == "deleted" {
		return nil
	} else {
		return errors.New("result not delete")
	}
}

func (g *GServer) ESGetOne(ind int32) (*pb.Suggest, error) {
	response, e := g.ES(http.MethodGet, fmt.Sprintf("%s/_doc/%d", IndexName, ind), nil)
	if e != nil {
		return nil, e
	}

	var fields map[string]interface{}
	e = json.Unmarshal(response, &fields)
	if e != nil {
		return nil, e
	}
	found, ok := fields["found"]
	if !ok {
		return nil, errors.New("expect field \"found\"")
	}
	if !found.(bool) {
		return nil, errors.New(fmt.Sprintf("suggest with ID = %d not exist", ind))
	}
	obj, ok := fields["_source"].(map[string]interface{})
	if !ok {
		return nil, errors.New("expect field \"_source\"")
	}
	sugg := &pb.Suggest{}
	sugg.ID = ind
	sugg.LinkUrl = obj["LinkUrl"].(string)
	sugg.Title = obj["Title"].(string)
	query := obj["Query"].([]interface{})
	for _, q := range query {
		sugg.Queries = append(sugg.Queries, q.(string))
	}
	return sugg, nil
}

func (g *GServer) ESGet() ([]*pb.Suggest, error) {
	res, e := g.ES(http.MethodGet, fmt.Sprintf("%s/_search?pretty=true&q=*:*&size=100", IndexName), nil)
	if e != nil {
		return nil, e
	}

	var hits_ map[string]interface{}
	e = json.Unmarshal(res, &hits_)
	if e != nil {
		return nil, e
	}
	hits, ok := hits_["hits"].(map[string]interface{})
	if !ok {
		return nil, errors.New("expect field \"hits\"(1)")
	}
	objs, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, errors.New("expect field \"hits\"(2)")
	}
	suggs := make([]*pb.Suggest, 0)
	for _, x := range objs {
		o := &pb.Suggest{}
		hit, ok := x.(map[string]interface{})
		if !ok {
			return nil, errors.New("cant convert \"hit\" to map[string]interface{}")
		}
		source, ok := hit["_source"].(map[string]interface{})
		if !ok {
			return nil, errors.New("expect field \"_source\"")
		}
		id, _ := strconv.Atoi(hit["_id"].(string))
		o.ID = int32(id)
		o.LinkUrl = source["LinkUrl"].(string)
		o.Title = source["Title"].(string)
		query := source["Query"].([]interface{})
		for _, q := range query {
			o.Queries = append(o.Queries, q.(string))
		}
		suggs = append(suggs, o)
	}

	return suggs, nil
}

func (g *GServer) ESSearch(q string) []*pb.Suggest {
	body := []byte(fmt.Sprintf(`
	{ 
		"query": {
		"bool": {
		  "should": [
			{"match": {"query": {"query": "%[1]s","analyzer": "keyboard"}}},
			{"match": {"query": {"query": "%[1]s","analyzer": "translit"}}}
		  ],
		  "minimum_should_match": 1
		}
	  }
	}`, q))
	res, e := g.ES(http.MethodGet, fmt.Sprintf("%s/_search", IndexName), body)
	fmt.Println(string(res))

	var hits_ map[string]interface{}
	e = json.Unmarshal(res, &hits_)
	if e != nil {
		return nil
	}
	hits, ok := hits_["hits"].(map[string]interface{})
	if !ok {
		return nil
	}
	objs, ok := hits["hits"].([]interface{})
	if !ok {
		return nil
	}
	suggs := make([]*pb.Suggest, 0)
	for _, x := range objs {
		o := &pb.Suggest{}
		hit, ok := x.(map[string]interface{})
		if !ok {
			return nil
		}
		source, ok := hit["_source"].(map[string]interface{})
		if !ok {
			return nil
		}
		id, _ := strconv.Atoi(hit["_id"].(string))
		o.ID = int32(id)
		o.LinkUrl = source["LinkUrl"].(string)
		o.Title = source["Title"].(string)
		query := source["Query"].([]interface{})
		for _, q := range query {
			o.Queries = append(o.Queries, q.(string))
		}
		suggs = append(suggs, o)
	}

	return suggs
}
