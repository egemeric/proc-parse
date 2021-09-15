package Controller

import (
	"net/http"
	"strings"
)

func ReadMemInfo(w http.ResponseWriter, r *http.Request) {
	mem, err := FileToLines("/proc/meminfo")
	Check(err)
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
	w.Write(ConvertToJson(meminfo))
}
