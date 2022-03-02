package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var TimeStamp int64
var TimeString string

var AddYear int
var AddMonth int
var AddDay int
var AddHour int
var AddMinute int
var AddSecond int

const TimeFormat = time.RFC3339

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Show date.",
	Long:  `Show date.`,
	Run: func(cmd *cobra.Command, args []string) {
		if TimeStamp != 0 && TimeString != "" {
			fmt.Println("-timestamp and -time can not be used at the same time")
			os.Exit(2)
		}
		runDate()
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	dateCmd.Flags().Int64VarP(&TimeStamp, "timestamp", "u", 0, "Unix time")
	dateCmd.Flags().StringVarP(&TimeString, "time", "t", "", "Time text in yyyyMMdd or yyyyMMddHHmmss")

	dateCmd.Flags().IntVarP(&AddYear, "add-year", "y", 0, "Add Year")
	dateCmd.Flags().IntVarP(&AddMonth, "add-month", "M", 0, "Add Month")
	dateCmd.Flags().IntVarP(&AddDay, "add-day", "d", 0, "Add Day")
	dateCmd.Flags().IntVarP(&AddHour, "add-hour", "H", 0, "Add Hour")
	dateCmd.Flags().IntVarP(&AddMinute, "add-minute", "m", 0, "Add Minute")
	dateCmd.Flags().IntVarP(&AddSecond, "add-second", "s", 0, "Add Second")
}

func parseDate(timeString string) time.Time {
	var layout string
	if len(timeString) == 8 {
		layout = "20060102"
	} else if len(timeString) == 14 {
		layout = "20060102150405"
	} else {
		fmt.Println("Time text is not correct")
		os.Exit(2)
	}

	t, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func getTime() time.Time {
	var tm time.Time
	if TimeStamp != 0 {
		tm = time.Unix(TimeStamp, 0)
	} else if len(TimeString) > 0 {
		tm = parseDate(TimeString)
	} else {
		tm = time.Now()
	}

	return tm
}

func addTime(tm time.Time) time.Time {
	var newTime time.Time
	newTime = tm

	if AddYear != 0 || AddMonth != 0 || AddDay != 0 {
		newTime = newTime.AddDate(AddYear, AddMonth, AddDay)
	}

	if AddHour != 0 {
		var durationString = strconv.Itoa(AddHour) + "h"
		var duration, _ = time.ParseDuration(durationString)
		newTime = newTime.Add(duration)
	}

	if AddMinute != 0 {
		var durationString = strconv.Itoa(AddMinute) + "m"
		var duration, _ = time.ParseDuration(durationString)
		newTime = newTime.Add(duration)
	}

	if AddSecond != 0 {
		var durationString = strconv.Itoa(AddSecond) + "s"
		var duration, _ = time.ParseDuration(durationString)
		newTime = newTime.Add(duration)
	}

	return newTime
}

func runDate() {
	var tm = getTime()
	tm = addTime(tm)

	fmt.Printf("Date:      %v\n", tm.Format(TimeFormat))
	fmt.Printf("Unix time: %v\n", tm.Unix())
}
