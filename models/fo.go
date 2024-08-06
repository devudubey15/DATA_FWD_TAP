package models

/*********************************************************************/
/*                                                                   */
/*  Description : Sets a null character ('\0') at a specified        */
/*                position in a byte slice, marking the end of the   */
/*                string in the slice.                               */
/*                                                                   */
/*********************************************************************/

func SETNULL(a []byte, length int) {
	if length < len(a) {
		a[length] = 0
	}
}
