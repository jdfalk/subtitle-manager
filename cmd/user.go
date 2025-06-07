package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"subtitle-manager/pkg/auth"
	"subtitle-manager/pkg/database"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
}

var userAddCmd = &cobra.Command{
	Use:   "add [username] [email] [password]",
	Args:  cobra.ExactArgs(3),
	Short: "Create user",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, email, password := args[0], args[1], args[2]
		db, err := database.Open(viper.GetString("db_path"))
		if err != nil {
			return err
		}
		defer db.Close()
		return auth.CreateUser(db, username, password, email, "user")
	},
}

var apiKeyCmd = &cobra.Command{
	Use:   "apikey [username]",
	Args:  cobra.ExactArgs(1),
	Short: "Generate API key for user",
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		db, err := database.Open(viper.GetString("db_path"))
		if err != nil {
			return err
		}
		defer db.Close()
		var id int64
		row := db.QueryRow(`SELECT id FROM users WHERE username = ?`, username)
		if err := row.Scan(&id); err != nil {
			return err
		}
		key, err := auth.GenerateAPIKey(db, id)
		if err != nil {
			return err
		}
		cmd.Println("API Key:", key)
		return nil
	},
}

var loginCmd = &cobra.Command{
	Use:   "login [username] [password]",
	Args:  cobra.ExactArgs(2),
	Short: "Authenticate and get session token",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, password := args[0], args[1]
		db, err := database.Open(viper.GetString("db_path"))
		if err != nil {
			return err
		}
		defer db.Close()
		id, err := auth.AuthenticateUser(db, username, password)
		if err != nil {
			return err
		}
		token, err := auth.GenerateSession(db, id, 24*time.Hour)
		if err != nil {
			return err
		}
		cmd.Println("Session Token:", token)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(apiKeyCmd)
	rootCmd.AddCommand(loginCmd)
}
