package common

// init bootstrapps the application
func init() {
	// Initialize AppConfig variable
	initConfig()
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Initialize Logger objects with Log Level
	setLogLevel(Level(AppConfig.LogLevel))
	// Start a MongoDB session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}
