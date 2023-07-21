package db

import (
	"database/sql"
	"fmt"
	"project_api/configs"

	_ "github.com/lib/pq"
)

//Função responsável por abrir a conexão com o banco de dados
func OpenConnection() (*sql.DB, error) {

	//pega as configuraçoes do banco para poder conectar no banco de dados
	conf := configs.GetDB()

	//string de conexao com o banco; sc = string connection
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
