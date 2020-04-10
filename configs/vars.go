package configs

// Port is number defined by PORT in .env file (default to 5000)
var Port int = getPortFromEnv()

// MongoURI is string defined by MONGODB_URI in .env file
var MongoURI = getMongoURIFromEnv()

// MongoDBName is defined by MONGO_DB_NAME in .env file
var MongoDBName = getMongoDBNameFromEnv()
