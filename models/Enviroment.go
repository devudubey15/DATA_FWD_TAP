package models

import (
	"log"

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
/*                          MaxToken limit.
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

func InitProcessSpace(FileName string, ProcessName string) int {
	cfg, err := ini.Load(FileName)
	if err != nil {
		log.Printf("INI File: %s: %v", FileName, err)
		return -1
	}

	section, err := cfg.GetSection(ProcessName)
	if err != nil {
		log.Printf("%s Not Specified in the INI File %s: %v", ProcessName, FileName, err)
		return -1
	}

	configMap = make(map[string]string)

	for _, key := range section.Keys() {
		configMap[key.Name()] = key.String()
	}

	if len(configMap) > MaxToken {
		log.Printf("Exceeding max token limit: MAXTOKEN: %d, Count of Token: %d", MaxToken, len(configMap))
		return -1
	}

	return 0
}

func GetProcessSpaceValue(token string) string {
	value, found := configMap[token]
	if !found {
		log.Printf("Token '%s' not found\n", token)
		return ""
	}
	return value // here we are returning the value of given Macro
}
