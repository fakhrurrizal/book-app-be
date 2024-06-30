package config

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName                     string
	AppKey                      string
	BaseUrl                     string
	Environtment                string
	DatabaseUsername            string
	DatabasePassword            string
	DatabaseHost                string
	DatabasePort                string
	DatabaseName                string
	DatabasePlannerName         string
	PathDB                      string
	CacheURL                    string
	CachePassword               string
	LoggerLevel                 string
	ContextTimeout              int
	Port                        string
	EnableDatabaseAutomigration bool
	APIKey                      string
}

func LoadConfig() (config *Config) {

	if err := godotenv.Load(RootPath() + `/.env`); err != nil {
		fmt.Println(err)
	}

	appName := os.Getenv("APP_NAME")
	appKey := os.Getenv("APP_KEY")
	baseurl := os.Getenv("BASE_URL")
	environment := strings.ToUpper(os.Getenv("ENVIRONMENT"))
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePlannerName := os.Getenv("DATABASE_PLANNER_NAME")
	PathDB := os.Getenv("PATH_DB")
	cacheURL := os.Getenv("CACHE_URL")
	cachePassword := os.Getenv("CACHE_PASSWORD")
	loggerLevel := os.Getenv("LOGGER_LEVEL")
	contextTimeout, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	enableDatabaseAutomigration, _ := strconv.ParseBool(os.Getenv("ENABLE_DATABASE_AUTOMIGRATION"))
	port := os.Getenv("PORT")
	apiKey := os.Getenv("API_KEY")

	return &Config{
		AppName:                     appName,
		AppKey:                      appKey,
		BaseUrl:                     baseurl,
		Environtment:                environment,
		DatabaseUsername:            databaseUsername,
		DatabasePassword:            databasePassword,
		DatabaseHost:                databaseHost,
		DatabasePort:                databasePort,
		DatabaseName:                databaseName,
		DatabasePlannerName:         databasePlannerName,
		PathDB:                      PathDB,
		CacheURL:                    cacheURL,
		CachePassword:               cachePassword,
		LoggerLevel:                 loggerLevel,
		ContextTimeout:              contextTimeout,
		EnableDatabaseAutomigration: enableDatabaseAutomigration,
		Port:                        port,
		APIKey:                      apiKey,
	}
}

func RootPath() string {
	projectDirName := os.Getenv("DIR_NAME")
	if projectDirName == "" {
		log.Fatalf("Environment variable DIR_NAME is not set")
	}

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	rootPath := projectName.Find([]byte(currentWorkDirectory))
	if rootPath == nil {
		log.Fatalf("Error finding project root path based on DIR_NAME: %v", projectDirName)
	}

	return string(rootPath)
}
