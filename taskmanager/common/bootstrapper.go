package common

func StartUp() {
	// Initialize Config
	initConfig()
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Start a MongoDB session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}
