package cmd

import (
	"fmt"

	"github.com/schoolboybru/git-notifier/internal/service"
	"github.com/spf13/cobra"
)

// getNotificationsCmd represents the getNotifications command
var getNotificationsCmd = &cobra.Command{
	Use:   "notify",
	Short: "Get all of your notifications from Github",
	Run: func(cmd *cobra.Command, args []string) {
		notificationService := service.New()

		n, err := notificationService.GetNotifications()

		if err != nil {
			fmt.Print(err)
		}

		fmt.Print(n)

	},
}

func init() {
	rootCmd.AddCommand(getNotificationsCmd)
}
