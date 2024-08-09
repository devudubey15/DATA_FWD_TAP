package common

import (
	"gorm.io/gorm"
)

// DBConfigLoader defines an interface for loading database configurations.
type DBConfigLoader interface {
	LoadPostgreSQLConfig(serviceName string, fileName string) (*PostgreSQLConfig, error)
}

// DBConnector defines an interface for establishing database connections.
type DBConnector interface {
	GetDatabaseConnection(serviceName string, cfg PostgreSQLConfig) int
}

// DBAccessor defines an interface for accessing database instances.
type DBAccessor interface {
	GetDB(serviceName string) *gorm.DB
}

// PostgreSQLConfig holds the necessary configuration details for connecting to a PostgreSQL database.
type PostgreSQLConfig struct {
	Username string // Database username
	Password string // Database password
	DBName   string // Database name
	Host     string // Database host
	Port     int    // Database port
	SSLMode  string // SSL mode for the database connection
}
