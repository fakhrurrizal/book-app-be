package config

import (
	"fmt"
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

	if err := godotenv.Load(RootPath() + `\.env`); err != nil {
		fmt.Println("error", err)
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
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath)
}
