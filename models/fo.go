package models

import "log"

/*********************************************************************/
/*                                                                   */
/*  Description : Sets a null character ('\0') at a specified        */
/*                position in a byte slice, marking the end of the   */
/*                string in the slice.                               */
/*                                                                   */
/*********************************************************************/

func SETNULL(serviceName string, a []byte, length int) {
	if length < len(a) {
		a[length] = 0
		log.Printf("[%s] Set null character at position %d in byte slice", serviceName, length)
	} else {
		log.Printf("[%s] Length %d exceeds byte slice size %d. No action taken.", serviceName, length, len(a))
	}
}
