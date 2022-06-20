package main

import (
	"fmt"
	"log"
	"os"

	"github.com/femibiwoye/image-upload-service/internal/grpc"
	"github.com/spf13/viper"
)

// init reads in config file and ENV variables if set.
func init() {
	// Search config in current directory with name ".config" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	setAWSEnv()
}

func main() {
	s := grpc.NewServer()
	if err := s.Start(viper.GetString("PORT")); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}

func setAWSEnv() {
	if _, ok := os.LookupEnv("AWS_ACCESS_KEY_ID"); !ok {
		os.Setenv("AWS_ACCESS_KEY_ID", viper.GetString("AWSAccessKeyID"))
	}

	if _, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); !ok {
		os.Setenv("AWS_SECRET_ACCESS_KEY", viper.GetString("AWSSecretKey"))
	}

	if _, ok := os.LookupEnv("AWS_REGION"); !ok {
		os.Setenv("AWS_REGION", viper.GetString("AWSRegion"))
	}

	if _, ok := os.LookupEnv("AWS_ACCESS_KEY_ID"); !ok {
		os.Setenv("AWS_ACCESS_KEY_ID", viper.GetString("AWSAccessKeyID"))
	}
}
