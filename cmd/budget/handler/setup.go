package handler

import (
	"log"

	internal_alert "github.com/ibilalkayy/flow/internal/app/alert"
	"github.com/ibilalkayy/flow/internal/common/functions"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/spf13/cobra"
)

// SetupCmd represents the setup command
var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup for alert notification",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		frequency, _ := cmd.Flags().GetString("frequency")
		method, _ := cmd.Flags().GetString("method")
		day, _ := cmd.Flags().GetString("day")
		weekday, _ := cmd.Flags().GetString("weekday")
		hour, _ := cmd.Flags().GetString("hour")
		minute, _ := cmd.Flags().GetString("minute")
		second, _ := cmd.Flags().GetString("second")

		dayInt := functions.StringToInt(day)
		hourInt := functions.StringToInt(hour)
		minuteInt := functions.StringToInt(minute)
		secondInt := functions.StringToInt(second)

		av := structs.AlertVariables{
			Category:  category,
			Frequency: frequency,
			Method:    method,
			Days:      dayInt,
			Weekdays:  weekday,
			Hours:     hourInt,
			Minutes:   minuteInt,
			Seconds:   secondInt,
		}

		err := internal_alert.AlertSetup(&av)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	SetupCmd.Flags().StringP("category", "c", "", "Write the category name to take its budget amount")
	SetupCmd.Flags().StringP("frequency", "f", "", "Write the frequency of notifications (e.g., hourly, daily, weekly, monthly)")
	SetupCmd.Flags().StringP("method", "t", "", "Write the preferred method of notification [email or CLI] message")
	SetupCmd.Flags().StringP("day", "d", "", "Write the day to set the notification")
	SetupCmd.Flags().StringP("weekday", "w", "", "Write the minute to set the notification")
	SetupCmd.Flags().StringP("hour", "o", "", "Write the hour to set the notification")
	SetupCmd.Flags().StringP("minute", "m", "", "Write the minute to set the notification")
	SetupCmd.Flags().StringP("second", "s", "", "Write the second to set the notification")
}
