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

var userRoleCmd = &cobra.Command{
	Use:   "role [username] [role]",
	Args:  cobra.ExactArgs(2),
	Short: "Set user role",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, role := args[0], args[1]
		db, err := database.Open(viper.GetString("db_path"))
		if err != nil {
			return err
		}
		defer db.Close()
		return auth.SetUserRole(db, username, role)
	},
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

var userTokenCmd = &cobra.Command{
	Use:   "token [email]",
	Args:  cobra.ExactArgs(1),
	Short: "Generate one time login token",
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
		db, err := database.Open(viper.GetString("db_path"))
		if err != nil {
			return err
		}
		defer db.Close()
		var id int64
		row := db.QueryRow(`SELECT id FROM users WHERE email = ?`, email)
		if err := row.Scan(&id); err != nil {
			return err
		}
		token, err := auth.GenerateOneTimeToken(db, id, time.Hour)
		if err != nil {
			return err
		}
		cmd.Println("Token:", token)
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

var loginTokenCmd = &cobra.Command{
	Use:   "login-token [token]",
	Args:  cobra.ExactArgs(1),
	Short: "Authenticate using a one time token",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := args[0]
		db, err := database.Open(viper.GetString("db_path"))
		if err != nil {
			return err
		}
		defer db.Close()
		id, err := auth.ConsumeOneTimeToken(db, t)
		if err != nil {
			return err
		}
		session, err := auth.GenerateSession(db, id, 24*time.Hour)
		if err != nil {
			return err
		}
		cmd.Println("Session Token:", session)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(apiKeyCmd)
	userCmd.AddCommand(userRoleCmd)
	userCmd.AddCommand(userTokenCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(loginTokenCmd)
}
