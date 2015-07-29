package common

func init() {
	//Initialize private/public keys got JWT authentication
	initKeys()
	//Start a MongoDB session
	createDbSession()
}
