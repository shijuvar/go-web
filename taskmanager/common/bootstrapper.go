package common

func StartUp() {
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Start a MongoDB session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}
