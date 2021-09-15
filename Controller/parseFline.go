package Controller

import (
	"bufio"
	"os"
)

func FileToLines(filepath string) (lines []string, err error) {
	f, err := os.Open(filepath)
	Check(err)
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	err = scan.Err()
	return
}
