package database

import (
	"controle-de-gastos/pkg/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createTables(db *sqlx.DB) error {
	const query = `
		CREATE TABLE IF NOT EXISTS usuarios(
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			hash_pass TEXT NOT NULL,
        	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	return err
}

func ConnectDB(cfg *config.DB_Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting database: %v", err)
	}

	//testar a conexao
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error testing database connection: %v", err)
	}

	log.Println("Successfully connected to database: " + cfg.DBHost)

	//Criar as tabelas se não existir
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("error creating database tables: %v", err)
	}

	return db, nil
}
