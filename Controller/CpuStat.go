package Controller

import (
	"strconv"
	"strings"
)

var CpuStatRead []byte
var tmp_uint uint64

func CpuStat() {
	cpu, err := FileToLines("/proc/stat")
	Check(err)
	Stat := make(map[string][]uint)
	for _, line := range cpu {
		splited := strings.Split(line, " ")
		for _, data := range splited {
			tmp_uint, err = strconv.ParseUint(data, 10, 32)
			if err != nil {
				continue
			}
			Stat[splited[0]] = append(Stat[splited[0]], uint(tmp_uint))
		}

	}
}
