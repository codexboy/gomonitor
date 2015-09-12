package core

import (
	"errors"
	//"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	//"github.com/shirou/gopsutil/mem"
)

func ServeCpuNum() int {
	return runtime.NumCPU()
}

type MemState struct {
	Msg   string
	Total uint64
	Free  uint64
	Used  uint64
}

func Memory() (memstate MemState, err error) {
	out, err := exec.Command("free").Output()
	if err != nil {
		return MemState{}, err
	}
	var memstr string = string(out)
	mems := strings.Split(memstr, "\n")
	if len(mems) < 2 {
		return MemState{}, errors.New("memset error")
	}
	//tabs := strings.Split(, " ")
	tmp := delSpace(mems[1])
	tabs := strings.Split(tmp, " ")
	memstate.Total ,_ = strconv.ParseUint(tabs[1], 0, 0)
	memstate.Used ,_ = strconv.ParseUint(tabs[2], 0, 0)
	memstate.Free ,_ = strconv.ParseUint(tabs[3], 0, 0)
	return
}

type CPUInfo struct {
	Message string
	Usr     float64
	Nice    float64
	Sys     float64
	Iowait  float64
	Irq     float64
	Soft    float64
	Steal   float64
	Guest   float64
	Idle    float64
}

func delSpace(str string) string {
	b := []byte(str)
	by := make([]byte, len(str))
	var flag bool = true
	for _, c := range b {
		if c == ' ' && flag {
			flag = false
			by = append(by, ' ')
		}
		if c != ' ' {
			by = append(by, c)
			flag = true
		}
	}
	return string(by)
}

func CPU() (cpuInfo CPUInfo, err error) {
	out, err := exec.Command("mpstat").Output()
	if err != nil {
		return CPUInfo{}, err
	}
	var cpustr string = string(out)
	cpus := strings.Split(cpustr, "\n")
	if len(cpus) < 4 {
		return CPUInfo{}, errors.New("mpstat error")
	}
	cpuInfo.Message = cpus[0]
	tmp := delSpace(cpus[3])
	tabs := strings.Split(tmp, " ")
	cpuInfo.Usr, _ = strconv.ParseFloat(tabs[2], 0)
	cpuInfo.Nice, _ = strconv.ParseFloat(tabs[3], 0)
	cpuInfo.Sys, _ = strconv.ParseFloat(tabs[4], 0)
	cpuInfo.Iowait, _ = strconv.ParseFloat(tabs[5], 0)
	cpuInfo.Irq, _ = strconv.ParseFloat(tabs[6], 0)
	cpuInfo.Soft, _ = strconv.ParseFloat(tabs[7], 0)
	cpuInfo.Steal, _ = strconv.ParseFloat(tabs[8], 0)
	cpuInfo.Guest, _ = strconv.ParseFloat(tabs[9], 0)
	cpuInfo.Idle, _ = strconv.ParseFloat(tabs[10], 0)
	return
}
