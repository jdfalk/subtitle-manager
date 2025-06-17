package webserver

// AppVersion holds the application version set by the cmd package.
var AppVersion = "dev"

// SetVersion stores the application version for the webserver package.
func SetVersion(v string) { AppVersion = v }
