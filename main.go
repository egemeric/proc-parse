package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileToLines(filepath string) (lines []string, err error) {
	f, err := os.Open(filepath)
	check(err)
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	err = scan.Err()
	return
}

func ReadCpuInfo(result chan []byte) {
	cpu, err := FileToLines("/proc/cpuinfo")
	check(err)
	specs := make(map[string][]string)
	for _, line := range cpu {
		if line == "" {
			break
		}
		tmp := strings.Split(line, ":")
		var splited [2][]string
		splited[1] = strings.Split(tmp[1], " ")
		key := strings.TrimSpace(tmp[0])
		specs[key] = splited[1][1:]

	}
	ConvertToJson(specs, result)
	<-result

}

func ReadMemInfo(result chan []byte) {
	mem, err := FileToLines("/proc/meminfo")
	check(err)
	meminfo := make(map[string][]string)
	for _, line := range mem {
		tmp := strings.Split(line, ":")
		var splited [2][]string
		splited[1] = strings.Split(tmp[1], " ")
		for index, _ := range splited[1] {
			if index >= len(splited[1])-2 {
				break
			}
			splited[1][index] = ""
		}
		splited[1] = splited[1][len(splited[1])-2 : len(splited[1])]
		key := strings.TrimSpace(tmp[0])
		meminfo[key] = splited[1][:]
	}
	ConvertToJson(meminfo, result)
	<-result
}

func ConvertToJson(data map[string][]string, ch chan []byte) {
	jsonString, err := json.Marshal(data)
	check(err)
	ch <- jsonString
}

func ReadLoop(done chan bool) {
	for {
		mem := make(chan []byte)
		go ReadMemInfo(mem)
		fmt.Println(string(<-mem))
		close(mem)
		time.Sleep(time.Second)
	}
}
func init() {
	cpu := make(chan []byte)
	go ReadCpuInfo(cpu)
	fmt.Println(string(<-cpu))
	close(cpu)

}

func main() {
	done := make(chan bool)
	go ReadLoop(done)
	<-done
}
