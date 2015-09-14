package main

import (
	"core"
	"fmt"
	json "github.com/bitly/go-simplejson"
	"html/template"
	"log"
	"net/http"
	"os"
	//"strconv"
	_ "expvar"
	"strings"
)

func jsonCmd(w http.ResponseWriter, q *http.Request) {
	tabs := strings.Split(q.URL.Path, "/")
	cmd := tabs[len(tabs)-1]
	if cmd == "mem" {
		mem, err := core.Memory()
		if err != nil {
			log.Println(err)
			return
		}
		js := json.New()
		js.Set("free", fmt.Sprintf("%d", mem.Free/1024))
		js.Set("used", fmt.Sprintf("%d", mem.Used/1024))
		js.Set("total", fmt.Sprintf("%d", mem.Total/1024))
		jsstr, _ := js.Encode()
		w.Write(jsstr)
	} else if cmd == "cpu" {
		cpu, err := core.CPU()
		if err != nil {
			log.Println(err)
			return
		}
		js := json.New()
		js.Set("msg", cpu.Message)
		js.Set("usr", fmt.Sprintf("%.2f", cpu.Usr))
		js.Set("sys", fmt.Sprintf("%.2f", cpu.Sys))
		js.Set("idle", fmt.Sprintf("%.2f", cpu.Idle))
		jsstr, _ := js.Encode()
		w.Write(jsstr)
	} else {
		w.Write([]byte("unknow cmd"))
	}
}

func htmlParse(w http.ResponseWriter, q *http.Request) {
	tabs := strings.Split(q.URL.Path, "/")
	filepath := fmt.Sprintf("%s/%s", os.Getenv("WWW"), tabs[len(tabs)-1])
	t, err := template.ParseFiles(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, nil)
}

func VarsCmd(w http.ResponseWriter, q *http.Request) {
	obj, err := core.RequestVars("http://localhost:9999/debug/vars")
	if err != nil {
		log.Println(err)
		return
	}
	jsobj := json.New()
	//log.Println(obj.Alloc)
	jsobj.Set("alloc", obj.Alloc)
	jsobj.Set("sys", obj.Sys)
	jsobj.Set("heapAlloc", obj.HeapAlloc)
	jsobj.Set("heapInuse", obj.HeapInuse)
	jsobj.Set("objects", obj.Objects)

	jsstr, _ := jsobj.Encode()
	w.Write(jsstr)
}

func main() {

	http.HandleFunc("/json/", jsonCmd)
	http.HandleFunc("/html/", htmlParse)
	http.HandleFunc("/vars/", VarsCmd)
	http.ListenAndServe(":9999", nil)

	//	var mem core.MemState
	//	var err error
	//	mem, err = core.Memory()
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	log.Println(mem.Total/(1024), mem.Used/1024, mem.Free/1024)
	//
	//	cpu, err := core.CPU()
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	log.Println(cpu.Message)
	//	log.Println(cpu)
}
