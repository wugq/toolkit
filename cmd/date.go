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
}

var dateCmdData DateCmdData

const TimeFormat = time.RFC3339

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Show date.",
	Long:  `Show date.`,
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

	fmt.Printf("Date:      %v\n", tm.Format(TimeFormat))
	fmt.Printf("Unix time: %v\n", tm.Unix())
	return nil
}
