package config

import (
	"DATA_FWD_TAP/models"
	"fmt"
	"log"

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
var resultTmp int

// LoadPostgreSQLConfig reads the configuration from an INI file and returns a PostgreSQLConfig instance.
// It takes the file name of the INI file as an argument.
func LoadPostgreSQLConfig(serviceName string, fileName string) (*PostgreSQLConfig, error) {
	// Initialize the process space with the 'database' section of the INI file
	resultTmp := models.InitProcessSpace(serviceName, fileName, "database")
	if resultTmp != 0 {
		log.Printf("[%s] Failed to read config file: %v", serviceName, resultTmp)
		return nil, fmt.Errorf("[%s] Failed to read config file: %v", serviceName, resultTmp)
	}

	// Create a PostgreSQLConfig instance with values from the configMap
	config := &PostgreSQLConfig{
		Host:     models.GetProcessSpaceValue("host"),
		Port:     models.GetProcessSpaceValueAsInt("port"),
		Username: models.GetProcessSpaceValue("username"),
		Password: models.GetProcessSpaceValue("password"),
		DBName:   models.GetProcessSpaceValue("dbname"),
		SSLMode:  models.GetProcessSpaceValue("sslmode"),
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
	log.Printf("[%s] GetDb is called and Returning the Database instance ", serviceName)
	return database.Db // Return the GORM database connection instance
}
