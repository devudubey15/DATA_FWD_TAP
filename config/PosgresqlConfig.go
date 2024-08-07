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
/*                        - LoadPostgreSQLConfig: Reads the configuration from     */
/*                          an INI file and returns a `PostgreSQLConfig` instance. */
/*                                                                                 */
/*                        - GetDatabaseConnection: Constructs a DSN string from    */
/*                          the provided `PostgreSQLConfig` struct, establishes a  */
/*                          connection to the PostgreSQL database using GORM, and  */
/*                          returns an int status code.                            */
/*                                                                                 */
/*                        - GetDB: Returns the currently established database      */
/*                          connection as a `*gorm.DB` instance.                   */
/*                                                                                 */
/**********************************************************************************/

// PostgreSQLConfig holds the necessary configuration details for connecting to a PostgreSQL database.
type PostgreSQLConfig struct {
	Username string // Database username
	Password string // Database password
	DBName   string // Database name
	Host     string // Database host
	Port     int    // Database port
	SSLMode  string // SSL mode for the database connection
}

// database is a package-level variable to hold the GORM database connection.
var database struct {
	Db *gorm.DB // GORM database connection instance
}

// LoadPostgreSQLConfig reads the configuration from an INI file and returns a PostgreSQLConfig instance.
// It takes the file name of the INI file as an argument.
func LoadPostgreSQLConfig(serviceName string, fileName string) (*PostgreSQLConfig, error) {
	cfg, err := ini.Load(fileName) // Load the INI file
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to read config file: %v", serviceName, err) // Return error if the file can't be read
	}

	section, err := cfg.GetSection("database") // Get the 'database' section from the INI file
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to get 'database' section: %v", serviceName, err) // Return error if the section is missing
	}

	// Create a PostgreSQLConfig instance with values from the INI file
	config := &PostgreSQLConfig{
		Host:     section.Key("host").String(),     // Read host
		Port:     section.Key("port").MustInt(),    // Read port (must be an integer)
		Username: section.Key("username").String(), // Read username
		Password: section.Key("password").String(), // Read password
		DBName:   section.Key("dbname").String(),   // Read database name
		SSLMode:  section.Key("sslmode").String(),  // Read SSL mode
	}

	return config, nil // Return the PostgreSQLConfig instance
}

// GetDatabaseConnection constructs a DSN string from the provided PostgreSQLConfig struct,
// establishes a connection to the PostgreSQL database using GORM, and returns an int status code.
// It returns 0 on success and -1 on failure.
func GetDatabaseConnection(serviceName string, cfg PostgreSQLConfig) int {
	// Construct the DSN string for PostgreSQL connection
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)
	// Open the database connection using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[%s] failed to connect database: %v", serviceName, err) // Log the error if the connection fails
		return -1                                                           // Return -1 on failure
	}
	database.Db = db // Assign the database connection to the package-level variable
	return 0         // Return 0 on success
}

// GetDB returns the currently established database connection as a *gorm.DB instance.

func GetDB(serviceName string) *gorm.DB {
	log.Printf("[%s] Returning the Database instance ", serviceName)
	return database.Db // Return the GORM database connection instance
}
