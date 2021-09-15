package Controller

import (
	"net/http"
	"strings"
)

func ReadCpuInfo(w http.ResponseWriter, r *http.Request) {
	cpu, err := FileToLines("/proc/cpuinfo")
	Check(err)
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
	w.Write(ConvertToJson(specs))

}
