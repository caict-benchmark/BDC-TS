package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FilterLoadLog(fileName string) ([]string, int, float64, int, float64, float64, float64) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, 0
	}
	buf := bufio.NewReader(f)
	var result []string
	var items = int(0)
	var timeToken = float64(0)
	var workers = int(0)
	var points = float64(0)
	var values = float64(0)
	var data = float64(0)
	var count = int(0)
	const regx = `^loaded\s([\d]+)\sitems\sin\s(\d+\.\d+)sec\swith\s(\d+)\sworkers.*rate\s(\d+\.\d+)\sitems.*rate\s(\d+\.\d+)/s,\s(\d+\.\d+)MB/sec\sfrom\sstdin\)$`
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				return result, items, timeToken, workers / count, points / float64(count), values / float64(count), data / float64(count)
			}
			return result, items, timeToken, workers / count, points / float64(count), values / float64(count), data / float64(count)
		}
		result = append(result, line)
		flySnowRegexp := regexp.MustCompile(regx)
		params := flySnowRegexp.FindStringSubmatch(line)
		if params != nil {
			count += 1
			value, _ := strconv.Atoi(params[1])
			items = items + value
			token, _ := strconv.ParseFloat(params[2], 32)
			timeToken = timeToken + token
			worker, _ := strconv.Atoi(params[3])
			workers = workers + worker
			point, _ := strconv.ParseFloat(params[4], 32)
			points = points + point
			v, _ := strconv.ParseFloat(params[5], 32)
			values = values + v
			d, _ := strconv.ParseFloat(params[6], 32)
			data = data + d
		}
	}
	return result, items, timeToken, workers / count, points / float64(count), values / float64(count), data / float64(count)
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "filePath", "unknown", "Input result file path")
	flag.Parse()
	_, items, timeToken, workers, itemsRate, valueRate, dataRate := FilterLoadLog(filePath)
	fmt.Printf("Items: %d\n", items)
	fmt.Printf("Time token: %f sec\n", timeToken)
	fmt.Printf("Workers: %d\n", workers)
	fmt.Printf("Items rate: %f items/sec\n", itemsRate)
	fmt.Printf("Values rate: %f values/sec\n", valueRate)
	fmt.Printf("Data rate: %f MB/sec\n", dataRate)
}
