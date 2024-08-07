/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	ac "github.com/CarlsonYuan/agora-chat-cli/agora-chat"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Manage push notifications",
	Long:  `Commands to manage push notifications.`,
}

var testPushCmd = &cobra.Command{
	Use:   "test",
	Short: "Test push notification",
	Example: heredoc.Doc(`
		# Send a test push notification for a specific user
		$ agchat push test --user <user-id>
	`),

	RunE: func(cmd *cobra.Command, args []string) error {
		userID, _ := cmd.Flags().GetString("user")
		if userID == "" {
			cmd.Println("Usage: agchat push test --user <user-id>")
			return nil
		}

		message, _ := cmd.Flags().GetString("message")
		var msg ac.PushMessage
		if err := json.Unmarshal([]byte(message), &msg); err != nil {
			return fmt.Errorf("failed to parse push message JSON: %v", err)
		}

		client := ac.NewClient()

		res, err := client.Push().SyncPush(userID, ac.OnlyPushPrivider, msg)
		if err != nil {
			return fmt.Errorf("failed to send push notification: %w", err)
		}

		for i, pushResult := range res {
			i++
			switch pushResult.PushStatus {
			case "SUCCESS":
				if pushResult.Data != nil {
					if pushResult.Data != nil {
						cmd.Printf("[Device %d] Success - Result from push provider(s)(Firebase/APN): %+v\n", i, pushResult.Data)
					}
				} else {
					cmd.Printf("[Device %d] Success - Data is nil", i)
				}
			case "FAIL":
				if pushResult.Desc != nil {
					cmd.Printf("[Device %d] Failure - Desc: %s\n", i, *pushResult.Desc)
				} else if pushResult.Data != nil {
					cmd.Printf("[Device %d] Failure - Result from push provider(s)(Firebase/APN): %+v\n", i, pushResult.Data)
				} else {
					cmd.Printf("[Device %d] Failure - No description provided", i)
				}
			default:
				cmd.Printf("[Device %d] Unknown pushStatus:%s", i, pushResult.PushStatus)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.AddCommand(testPushCmd)

	testPushCmd.Flags().StringP("user", "u", "", "the user ID of the target user")
	testPushCmd.MarkFlagsRequiredTogether()
	testPushCmd.Flags().StringP("message", "m", `{"title": "Admin sent you a message", "content": "For push notification testing", "sub_title": "Test message is sent"}`, "JSON string for the push message")

	testPushCmd.MarkFlagRequired("user")
}
