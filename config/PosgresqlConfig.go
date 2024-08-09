package config

import (
	"DATA_FWD_TAP/common"
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

// ConfigManager implements the DBConfigLoader, DBConnector, and DBAccessor interfaces.
type ConfigManager struct {
	database struct {
		Db *gorm.DB // GORM database connection instance
	}
}

var resultTmp int

// LoadPostgreSQLConfig reads the configuration from an INI file and returns a PostgreSQLConfig instance.
// It takes the file name of the INI file as an argument.
func (cm *ConfigManager) LoadPostgreSQLConfig(serviceName string, fileName string) (*common.PostgreSQLConfig, error) {
	// Initialize the process space with the 'database' section of the INI file

	environmentManager := &models.EnvironmentManager{}

	resultTmp = environmentManager.InitProcessSpace(serviceName, fileName, "database")
	if resultTmp != 0 {
		log.Printf("[%s] Failed to read config file: %v", serviceName, resultTmp)
		return nil, fmt.Errorf("[%s] Failed to read config file: %v", serviceName, resultTmp)
	}

	// Create a PostgreSQLConfig instance with values from the configMap
	config := &common.PostgreSQLConfig{
		Host:     environmentManager.GetProcessSpaceValue("host"),
		Port:     environmentManager.GetProcessSpaceValueAsInt("port"),
		Username: environmentManager.GetProcessSpaceValue("username"),
		Password: environmentManager.GetProcessSpaceValue("password"),
		DBName:   environmentManager.GetProcessSpaceValue("dbname"),
		SSLMode:  environmentManager.GetProcessSpaceValue("sslmode"),
	}

	return config, nil // Return the PostgreSQLConfig instance
}

// GetDatabaseConnection constructs a DSN string from the provided PostgreSQLConfig struct,
// establishes a connection to the PostgreSQL database using GORM, and returns an int status code.
// It returns 0 on success and -1 on failure.
func (cm *ConfigManager) GetDatabaseConnection(serviceName string, cfg common.PostgreSQLConfig) int {
	// Construct the DSN string for PostgreSQL connection
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)
	// Open the database connection using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[%s] failed to connect database: %v", serviceName, err) // Log the error if the connection fails
		return -1                                                           // Return -1 on failure
	}
	cm.database.Db = db // Assign the database connection to the package-level variable
	return 0            // Return 0 on success
}

// GetDB returns the currently established database connection as a *gorm.DB instance.

func (cm *ConfigManager) GetDB(serviceName string) *gorm.DB {
	log.Printf("[%s] GetDb is called and Returning the Database instance ", serviceName)
	return cm.database.Db // Return the GORM database connection instance
}
