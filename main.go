package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func expand(field string, min, max int) []int {
	expanded := make(map[int]bool)
	parts := strings.Split(field, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			for i := start; i <= end; i++ {
				expanded[i] = true
			}
		} else if strings.Contains(part, "/") {
			stepParts := strings.Split(part, "/")
			step, _ := strconv.Atoi(stepParts[1])
			for i := 0; i <= max; i += step {
				expanded[i] = true
			}
		} else if part == "*" {
			for i := min; i <= max; i++ {
				expanded[i] = true
			}
		} else {
			num, _ := strconv.Atoi(part)
			expanded[num] = true
		}
	}

	result := make([]int, 0, len(expanded))
	for key := range expanded {
		result = append(result, key)
	}

	return result
}

func parse(cron string) map[string][]int {
	fields := strings.Fields(cron)
	if len(fields) != 6 {
		fmt.Println("Invalid cron string: it should contain 6 fields")
		os.Exit(1)
	}

	minute := expand(fields[0], 0, 59)
	hour := expand(fields[1], 0, 23)
	dayOfMonth := expand(fields[2], 1, 31)
	month := expand(fields[3], 1, 12)
	dayOfWeek := expand(fields[4], 0, 6)

	return map[string][]int{
		"minute":       minute,
		"hour":         hour,
		"day of month": dayOfMonth,
		"month":        month,
		"day of week":  dayOfWeek,
	}
}

func main() {
	l := len(os.Args)
	log.Println(l)
	if len(os.Args) < 7 {
		fmt.Println("Usage: go run main.go minutes hours days-of-month months days-of-week command")
		os.Exit(1)
	}

	cronString := strings.Join(os.Args[1:], " ")
	schedule := parse(cronString)
	for key, values := range schedule {
		//sort values
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		fmt.Printf("%-14s %s\n", key, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), " "), "[]"))
	}
	fmt.Printf("%-14s %s\n", "command", strings.Join(os.Args[6:], " "))
}
