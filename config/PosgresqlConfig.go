package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**********************************************************************************/
/*                                                                                 */
/*  Description       : This package defines a configuration structure for         */
/*                      connecting to a PostgreSQL database. The `PostgreSQLConfig`*/
/*                      struct holds essential details required for establishing   */
/*                      a connection, including username, password, database name, */
/*                      host, port, and SSL mode.                                  */
/*                                                                                 */
/*                      The `GetDatabaseConnection` function formats these values  */
/*                      into a Data Source Name (DSN) string, which is used to     */
/*                      open a connection to the PostgreSQL database using GORM.   */
/*                                                                                 */
/*  Functions         :                                                            */
/*                        - GetDatabaseConnection: Constructs a DSN string from    */
/*                          the provided `PostgreSQLConfig` struct, establishes a  */
/*                          connection to the PostgreSQL database using GORM, and  */
/*                          returns a `*gorm.DB` instance.                         */
/*                                                                                 */
/**********************************************************************************/

type PostgreSQLConfig struct {
	Username string
	Password string
	DBName   string
	Host     string
	Port     int
	SSLMode  string
}

var database struct {
	Db *gorm.DB
}

func LoadPostgreSQLConfig(fileName string) (*PostgreSQLConfig, error) {
	cfg, err := ini.Load(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	section, err := cfg.GetSection("database")
	if err != nil {
		return nil, fmt.Errorf("failed to get 'database' section: %v", err)
	}

	config := &PostgreSQLConfig{
		Host:     section.Key("host").String(),
		Port:     section.Key("port").MustInt(),
		Username: section.Key("username").String(),
		Password: section.Key("password").String(),
		DBName:   section.Key("dbname").String(),
		SSLMode:  section.Key("sslmode").String(),
	}

	return config, nil
}

func GetDatabaseConnection(cfg PostgreSQLConfig) int {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return -1
	}
	database.Db = db
	return 0
}

func GetDB() *gorm.DB {
	return database.Db
}
