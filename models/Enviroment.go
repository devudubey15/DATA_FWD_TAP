package models

import (
	"log"
	"strconv"

	"gopkg.in/ini.v1"
)

/**********************************************************************************/
/*                                                                                 */
/*  Description       : This program initializes a process space configuration     */
/*                      by reading values from a specified section of an INI file. */
/*                      The configuration values are stored in a global map,       */
/*                      allowing easy retrieval of values based on keys.           */
/*                                                                                 */
/*  Functions		  :                                 					       */
/*                        - InitProcessSpace: Loads the INI file, retrieves the    */
/*                          specified section, and stores key-value pairs in the   */
/*                          configMap. Validates the number of tokens against      */
/*                          MaxToken limit.                                        */
/*											                                       */
/*                        - GetProcessSpaceValue: Fetches the value associated     */
/*                          with a given key (token) from the configMap.           */
/*                                                                                 */
/*  Constants         :                                                            */
/*                        - MaxToken: The maximum number of tokens allowed.        */
/*                        												           */
/*                                                                                 */
/**********************************************************************************/

const (
	MaxToken = 50
)

var configMap map[string]string
var serviceName string

func InitProcessSpace(serviceName1 string, FileName string, ProcessName string) int {
	serviceName = serviceName1
	log.Printf("[%s] Initializing process space", serviceName)

	cfg, err := ini.Load(FileName)
	if err != nil {
		log.Printf("[%s] Error loading INI file: %s, Error: %v", serviceName, FileName, err)
		return -1
	}

	log.Printf("[%s] Successfully loaded INI file: %s", serviceName, FileName)

	section, err := cfg.GetSection(ProcessName)
	if err != nil {
		log.Printf("[%s] Section '%s' not specified in INI file: %s, Error: %v", serviceName, ProcessName, FileName, err)
		return -1
	}

	log.Printf("[%s] Successfully retrieved section: %s from INI file: %s", serviceName, ProcessName, FileName)

	configMap = make(map[string]string)

	for _, key := range section.Keys() {
		configMap[key.Name()] = key.String()
		log.Printf("[%s] Loaded key: %s, value: %s", serviceName, key.Name(), key.String())
	}

	if len(configMap) > MaxToken {
		log.Printf("[%s] Exceeding max token limit: MaxToken: %d, Count of tokens: %d", serviceName, MaxToken, len(configMap))
		return -1
	}

	log.Printf("[%s] Process space initialized successfully with %d tokens", serviceName, len(configMap))
	return 0
}

func GetProcessSpaceValue(token string) string {
	value, found := configMap[token]
	if !found {
		log.Printf("[%s] Token '%s' not found in configuration map", serviceName, token)
		return ""
	}
	log.Printf("[%s] Retrieved value for token '%s': %s", serviceName, token, value)
	return value
}

func GetProcessSpaceValueAsInt(token string) int {
	valueStr, found := configMap[token]
	if !found {
		log.Printf("[%s] Token '%s' not found in configuration map", serviceName, token)
		return -1
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("[%s] Failed to convert token '%s' value '%s' to integer: %v", serviceName, token, valueStr, err)
		return -1
	}
	log.Printf("[%s] Retrieved integer value for token '%s': %d", serviceName, token, value)
	return value
}
