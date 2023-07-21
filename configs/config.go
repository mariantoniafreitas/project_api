package configs

import "github.com/spf13/viper"

var cfg *config //ponteiro de config

//Struct Principal - une as duas structs para que leia apenas um arquivo
type config struct {
	API APIConfig
	DB  DBConfig
}

//Struct que possui porta que o serviço vai levantar
type APIConfig struct {
	Port string
}

//Struct de banco de dados
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

//funcao que é chamada no start das aplicaçoes
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

//funcao que carrega as configuracoes e retorna um erro caso nao consiga
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { //validando o erro
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { //validando o tipo de erro: fazendo um casting do erro; se o erro for diferent de nao enocontrei o arquivo returna erro
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
