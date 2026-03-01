package cmd

import (
	"fmt"
	"os"
	"time"
	"toolkit/runner/dateRunner"

	"github.com/spf13/cobra"
)

type DateCmdData struct {
	TimeStamp int64
	TimeText  string
	AddTime   dateRunner.AddDate
	Timezone  string
	DiffText  string
}

var dateCmdData DateCmdData

const TimeFormat = time.RFC3339

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Show date.",
	Long: `Show the current date and Unix timestamp, or convert a given time.

Input (pick one):
  -u / --timestamp  Convert a Unix epoch timestamp
  -t / --time       Parse a date string in yyyyMMdd or yyyyMMddHHmmss format

Date arithmetic (can be combined with any input):
  -y / --add-year    Add N years
  -M / --add-month   Add N months
  -d / --add-day     Add N days
  -H / --add-hour    Add N hours
  -m / --add-minute  Add N minutes
  -s / --add-second  Add N seconds

Examples:
  toolkit date
  toolkit date -u 1700000000
  toolkit date -t 20240101 -d 7`,
	Run: func(cmd *cobra.Command, args []string) {
		if dateCmdData.TimeStamp != 0 && dateCmdData.TimeText != "" {
			fmt.Println("-timestamp and -time can not be used at the same time")
			os.Exit(2)
		}
		err := runDate()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	dateCmd.Flags().Int64VarP(&dateCmdData.TimeStamp, "timestamp", "u", 0, "Unix time")
	dateCmd.Flags().StringVarP(&dateCmdData.TimeText, "time", "t", "", "Time text in yyyyMMdd or yyyyMMddHHmmss")

	dateCmd.Flags().StringVarP(&dateCmdData.Timezone, "timezone", "z", "", "Output timezone (e.g. UTC, Asia/Tokyo)")
	dateCmd.Flags().StringVarP(&dateCmdData.DiffText, "diff", "D", "", "Date to diff against, in yyyyMMdd or yyyyMMddHHmmss format")

	dateCmd.Flags().IntVarP(&dateCmdData.AddTime.Year, "add-year", "y", 0, "Add Year")
	dateCmd.Flags().IntVarP(&dateCmdData.AddTime.Month, "add-month", "M", 0, "Add Month")
	dateCmd.Flags().IntVarP(&dateCmdData.AddTime.Day, "add-day", "d", 0, "Add Day")
	dateCmd.Flags().IntVarP(&dateCmdData.AddTime.Hour, "add-hour", "H", 0, "Add Hour")
	dateCmd.Flags().IntVarP(&dateCmdData.AddTime.Minute, "add-minute", "m", 0, "Add Minute")
	dateCmd.Flags().IntVarP(&dateCmdData.AddTime.Second, "add-second", "s", 0, "Add Second")
}

func runDate() error {
	var tm time.Time
	var err error
	if dateCmdData.TimeStamp != 0 {
		tm = time.Unix(dateCmdData.TimeStamp, 0)
	} else if len(dateCmdData.TimeText) > 0 {
		tm, err = dateRunner.ParseDateText(dateCmdData.TimeText)
		if err != nil {
			return err
		}
	} else {
		tm = time.Now()
	}

	tm, err = dateRunner.Add(tm, dateCmdData.AddTime)
	if err != nil {
		return err
	}

	if dateCmdData.Timezone != "" {
		loc, err := time.LoadLocation(dateCmdData.Timezone)
		if err != nil {
			return fmt.Errorf("unknown timezone %q: %v", dateCmdData.Timezone, err)
		}
		tm = tm.In(loc)
	}

	if dateCmdData.DiffText != "" {
		tm2, err := dateRunner.ParseDateText(dateCmdData.DiffText)
		if err != nil {
			return fmt.Errorf("invalid --diff date: %v", err)
		}
		diff := tm2.Sub(tm)
		if diff < 0 {
			diff = -diff
		}
		days := int(diff.Hours()) / 24
		hours := int(diff.Hours()) % 24
		minutes := int(diff.Minutes()) % 60
		seconds := int(diff.Seconds()) % 60
		fmt.Printf("Difference: %d days, %d hours, %d minutes, %d seconds\n", days, hours, minutes, seconds)
		return nil
	}

	fmt.Printf("Date:      %v\n", tm.Format(TimeFormat))
	fmt.Printf("Unix time: %v\n", tm.Unix())
	return nil
}
