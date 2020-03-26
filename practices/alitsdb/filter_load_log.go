package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/caict-benchmark/BDC-TS/bulk_data_gen/common"
)

// TimeSlice is the wrapper of the time.Time slice for sorting
type TimeSlice []time.Time

func (is TimeSlice) Swap(i, j int) {
	is[j], is[i] = is[i], is[j]
}

func (is TimeSlice) Less(i, j int) bool {
	return is[i].Before(is[j])
}

func (is TimeSlice) Len() int {
	return len(is)
}

func filterLoadLog(fileName string) ([]string, int64, int64, float64, time.Time, time.Time, int, int, float64, float64) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("failed to open \"%s\", error: %v\n", fileName, err)
		return nil, 0, 0, 0, time.Unix(0, 0), time.Unix(0, 0), 0, 0, 0, 0
	}
	buf := bufio.NewReader(f)
	var result []string
	var items = int64(0)
	var values = int64(0)
	var timeToken = float64(0)
	var workers = int(0)
	var pointsRate = float64(0)
	var valuesRate = float64(0)
	var count = int(0)
	startTimes := TimeSlice(make([]time.Time, 0, 32))
	endTimes := TimeSlice(make([]time.Time, 0, 32))
	var earlyStartTime, lateEndTime time.Time

	const regx = `^loaded\s([\d]+)\sitems\sand\s([\d]+)\svalues\sin\s(\d+\.\d+)sec\s.*start\s(\d+),\send\s(\d+)\)\swith\s(\d+)\sworkers.*rate\s(\d+\.\d+)\sitems.*rate\s(\d+\.\d+)/s\)$`
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				earlyStartTime, lateEndTime = getEarlyStartAndLateEnd(startTimes, endTimes, fileName)
				return result, items, values, timeToken, earlyStartTime, lateEndTime, count, workers / count, pointsRate / float64(count), valuesRate / float64(count)
			}
			fmt.Printf("error occurred: %v\n", err)
			earlyStartTime, lateEndTime = getEarlyStartAndLateEnd(startTimes, endTimes, fileName)
			return result, items, values, timeToken, earlyStartTime, lateEndTime, count, workers / count, pointsRate / float64(count), valuesRate / float64(count)
		}
		result = append(result, line)
		flySnowRegexp := regexp.MustCompile(regx)
		params := flySnowRegexp.FindStringSubmatch(line)
		if params != nil {
			count++
			item, _ := strconv.ParseInt(params[1], 10, 64)
			items += item
			value, _ := strconv.ParseInt(params[2], 10, 64)
			values += value
			token, _ := strconv.ParseFloat(params[3], 32)
			timeToken += token
			start, _ := strconv.ParseInt(params[4], 10, 64)
			startTimes = append(startTimes, time.Unix(start, 0))
			end, _ := strconv.ParseInt(params[5], 10, 64)
			endTimes = append(endTimes, time.Unix(end, 0))
			worker, _ := strconv.Atoi(params[6])
			workers += worker
			point, _ := strconv.ParseFloat(params[7], 32)
			pointsRate += point
			rate, _ := strconv.ParseFloat(params[8], 32)
			valuesRate += rate
		}
	}
}

func getEarlyStartAndLateEnd(startTimes, endTimes TimeSlice, fileName string) (start, end time.Time) {
	if startTimes.Len() == 0 || endTimes.Len() == 0 {
		panic("did not retrieved any start time or end time from " + fileName)
	}

	sort.Sort(startTimes)
	sort.Sort(endTimes)

	return startTimes[0], endTimes[endTimes.Len()-1]
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "filePath", "unknown", "Input result file path")
	flag.Parse()
	_, items, values, _, start, end, processes, workers, itemsRate, valueRate := filterLoadLog(filePath)
	startTimestamp := start.Format(common.DateTimeStdFormat)
	endTimestamp := end.Format(common.DateTimeStdFormat)
	fmt.Printf("Items written: %d\n", items)
	fmt.Printf("Values written: %d\n", values)
	fmt.Printf("Writing started at %s\n", startTimestamp)
	fmt.Printf("Writing ended at %s\n", endTimestamp)
	fmt.Printf("Overall time cost: %d secs\n", end.Sub(start)/time.Second)
	fmt.Printf("Count of processes: %d\n", processes)
	fmt.Printf("Workers (per-process): %d\n", workers)
	fmt.Printf("Overall items rate: %f items/sec\n", float64(items)/float64(end.Sub(start)/time.Second))
	fmt.Printf("Overall values rate: %f values/sec\n", float64(values)/float64(end.Sub(start)/time.Second))
	fmt.Printf("Average items rate (per-process): %f items/sec\n", itemsRate)
	fmt.Printf("Average values rate (per-process): %f values/sec\n", valueRate)
}
