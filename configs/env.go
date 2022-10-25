package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvDbHost() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("DB_HOST")
}
func EnvDbPort() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("DB_PORT")
}
func EnvDbName() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("DB_DATABASE")
}
func EnvDbUsername() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("DB_USERNAME")
}
func EnvDbPassword() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("DB_PASSWORD")
}