package config


import(

	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct{
	Server struct{
		Port int
	}

	Database struct{
		Host string
		User string
		Password string
		Name string
		Port int
	}
}

var AppConfig Config

func InitConfig(){

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err:=viper.ReadInConfig(); err!=nil{
		log.Fatalf("Error when reading config file: %v", err)
	}

	err:=viper.Unmarshal(&AppConfig)

	if err!=nil{
		log.Fatalf("Unable to decode config: %v", err)
	}

	fmt.Println("Configuration loaded successfully")
}