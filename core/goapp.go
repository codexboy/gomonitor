package core

import (
	//"bytes"
	"fmt"
	//"runtime/pprof"
	"errors"
	json "github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	//"log"
	
)

type GoAPP struct {
	Alloc     int64
	Sys       int64
	HeapAlloc int64
	HeapInuse int64
	Objects   int64
}

func RequestVars(url string) (objGoApp GoAPP, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return objGoApp, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return objGoApp, errors.New("vars not found")
	} else if resp.StatusCode != http.StatusOK {
		return objGoApp, errors.New(fmt.Sprintf("request error code[%d]", resp.StatusCode))
	}
	jsonstr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return objGoApp, err
	}
	objson, err := json.NewJson(jsonstr)
	if err != nil {
		return objGoApp, err
	}
	objGoApp.Alloc, _ = objson.Get("memstats").Get("Alloc").Int64()
	objGoApp.Sys, _ = objson.Get("memstats").Get("Sys").Int64()
	objGoApp.HeapAlloc, _ = objson.Get("memstats").Get("HeapAlloc").Int64()
	objGoApp.HeapInuse, _ = objson.Get("memstats").Get("HeapInuse").Int64()
	objGoApp.Objects, _ = objson.Get("memstats").Get("HeapObjects").Int64()
	return objGoApp, err
}
