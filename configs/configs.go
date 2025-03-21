package configs

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
)

type Config struct {
	PORT    string
	IS_TEST bool

	DB_TYPE        string
	DB_HOST        string
	DB_PORT        string
	DB_USER        string
	DB_PASS        string
	DB_NAME        string
	DB_AUTOMIGRATE string
	DB_LOGMODE     string
	DB_SSL         string
	DB_TZ          string

	JWT_SECRET            string
	JWT_EXP_MINUTES       string
	REFRESH_JWT_EXP_HOURS string
	JWT_ECHO_CONFIG       echojwt.Config
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	isTest := flag.Lookup("test.v") != nil

	port, declared := os.LookupEnv("PORT")
	if !declared && !isTest {
		log.Fatal("\"PORT\" need to be defined")
	}

	// Database
	dbType, declared := os.LookupEnv("DB_TYPE")
	if !declared {
		dbType = "postgres"
	}
	dbHost, declared := os.LookupEnv("DB_HOST")
	if !declared && !isTest {
		log.Fatal("\"DB_HOST\" need to be defined")
	}
	dbPort, declared := os.LookupEnv("DB_PORT")
	if !declared {
		dbPort = "5432"
	}
	dbUser, _ := os.LookupEnv("DB_USER")
	dbPass, _ := os.LookupEnv("DB_PASS")
	dbName, declared := os.LookupEnv("DB_NAME")
	if !declared && !isTest {
		log.Fatal("\"DB_NAME\" need to be defined")
	}
	dbAutoMigrate, declared := os.LookupEnv("DB_AUTOMIGRATE")
	if !declared {
		dbAutoMigrate = "disable"
	}
	dbLogMode, _ := os.LookupEnv("DB_LOGMODE")
	dbSSL, _ := os.LookupEnv("DB_SSL")
	dbTZ, _ := os.LookupEnv("DB_TZ")

	return &Config{
		PORT:    port,
		IS_TEST: isTest,

		DB_TYPE:        dbType,
		DB_HOST:        dbHost,
		DB_PORT:        dbPort,
		DB_USER:        dbUser,
		DB_PASS:        dbPass,
		DB_NAME:        dbName,
		DB_AUTOMIGRATE: dbAutoMigrate,
		DB_LOGMODE:     dbLogMode,
		DB_SSL:         dbSSL,
		DB_TZ:          dbTZ,
	}
}
