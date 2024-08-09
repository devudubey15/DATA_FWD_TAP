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

type EnvironmentManager struct {
	ConfigMap   map[string]string
	ServiceName string
}

func (Em *EnvironmentManager) InitProcessSpace(serviceName1 string, FileName string, ProcessName string) int {
	Em.ServiceName = serviceName1
	log.Printf("[%s] Initializing process space", Em.ServiceName)

	cfg, err := ini.Load(FileName)
	if err != nil {
		log.Printf("[%s] Error loading INI file: %s, Error: %v", Em.ServiceName, FileName, err)
		return -1
	}

	log.Printf("[%s] Successfully loaded INI file: %s", Em.ServiceName, FileName)

	section, err := cfg.GetSection(ProcessName)
	if err != nil {
		log.Printf("[%s] Section '%s' not specified in INI file: %s, Error: %v", Em.ServiceName, ProcessName, FileName, err)
		return -1
	}

	log.Printf("[%s] Successfully retrieved section: %s from INI file: %s", Em.ServiceName, ProcessName, FileName)

	Em.ConfigMap = make(map[string]string)

	for _, key := range section.Keys() {
		Em.ConfigMap[key.Name()] = key.String()
		log.Printf("[%s] Loaded key: %s, value: %s", Em.ServiceName, key.Name(), key.String())
	}

	if len(Em.ConfigMap) > MaxToken {
		log.Printf("[%s] Exceeding max token limit: MaxToken: %d, Count of tokens: %d", Em.ServiceName, MaxToken, len(Em.ConfigMap))
		return -1
	}

	log.Printf("[%s] Process space initialized successfully with %d tokens", Em.ServiceName, len(Em.ConfigMap))
	return 0
}

func (Em *EnvironmentManager) GetProcessSpaceValue(token string) string {
	value, found := Em.ConfigMap[token]
	if !found {
		log.Printf("[%s] Token '%s' not found in configuration map", Em.ServiceName, token)
		return ""
	}
	log.Printf("[%s] Retrieved value for token '%s': %s", Em.ConfigMap, token, value)
	return value
}

func (Em *EnvironmentManager) GetProcessSpaceValueAsInt(token string) int {
	valueStr, found := Em.ConfigMap[token]
	if !found {
		log.Printf("[%s] Token '%s' not found in configuration map", Em.ServiceName, token)
		return -1
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("[%s] Failed to convert token '%s' value '%s' to integer: %v", Em.ServiceName, token, valueStr, err)
		return -1
	}
	log.Printf("[%s] Retrieved integer value for token '%s': %d", Em.ServiceName, token, value)
	return value
}
