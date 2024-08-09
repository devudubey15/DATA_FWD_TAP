package app

import (
	"DATA_FWD_TAP/models"
	"DATA_FWD_TAP/models/structures"
	"log"

	"gorm.io/gorm"
)

// var sql_pipe_id [3]byte
// var sql_c_xchng_cd [4]byte
var serviceName string = "cln_pack_clnt"

func Fn_bat_init(args []string, Db *gorm.DB) int {
	var temp_str string
	var resultTmp int
	log.Printf("[%s] Entering Fn_bat_init", serviceName)

	// we are getting the 7 args
	if len(args) < 7 {
		log.Printf("[%s] Error: insufficient arguments provided", serviceName)
		return -1
	}

	// here we are fetching the Pipe id
	exchngbook := &structures.Vw_xchngbook{}

	exchngbook.C_pipe_id = args[3]

	log.Printf("[%s] Copied pipe ID from args[3]: %s", serviceName, exchngbook.C_pipe_id[:])

	// Fetch exchange code
	queryForOpm_Xchng_Cd := `SELECT opm_xchng_cd
				FROM opm_ord_pipe_mstr
				WHERE opm_pipe_id = ?`

	log.Printf("[%s] Executing query to fetch exchange code with pipe ID: %s", serviceName, exchngbook.C_pipe_id[:])

	row := Db.Raw(queryForOpm_Xchng_Cd, exchngbook.C_pipe_id).Row()
	temp_str = ""
	err := row.Scan(&temp_str)
	if err != nil {
		log.Printf("[%s] Error scanning row for exchange code: %v", serviceName, err)
		return -1
	}

	exchngbook.C_xchng_cd = temp_str
	temp_str = ""
	log.Printf("[%s] Exchange code fetched: %s", serviceName, exchngbook.C_xchng_cd)

	// Fetch modification trade date
	queryFor_exg_nxt_trd_dt := `SELECT exg_nxt_trd_dt
				FROM exg_xchng_mstr
				WHERE exg_xchng_cd = ?`

	log.Printf("[%s] Executing query to fetch modification trade date with exchange code: %s", serviceName, exchngbook.C_xchng_cd)
	row2 := Db.Raw(queryFor_exg_nxt_trd_dt, exchngbook.C_xchng_cd).Row()

	err2 := row2.Scan(&temp_str)
	if err2 != nil {
		log.Printf("[%s] Error scanning row for modification trade date: %v", serviceName, err2)
		return -1
	}

	TmpByteArr := []byte(temp_str) // these i have to look again
	TmpArrLen := len(TmpByteArr)   // this i have to look again

	exchngbook.C_mod_trd_dt = temp_str

	models.SETNULL(serviceName, []byte(exchngbook.C_mod_trd_dt), TmpArrLen)

	log.Printf("[%s] Modification trade date fetched and set in C_mod_trd_dt: %s", serviceName, exchngbook.C_mod_trd_dt[:TmpArrLen])

	exchngbook.L_ord_seq = 0 // I am initially setting it to '0' because it was set that way in 'fn_bat_init' and I have not seen it getting changed anywhere. If I find it being changed somewhere, I will update it accordingly.

	log.Printf("[%s] L_ord_seq initialized to %d", serviceName, exchngbook.L_ord_seq)

	log.Printf("[%s] Exiting Fn_bat_init", serviceName)

	resultTmp = CLN_PACK_CLNT(args, Db)
	if resultTmp != 0 {
		log.Printf("[%s] CLN_PACK_CLNT failed with result code: %d", serviceName, resultTmp)
		log.Printf("[%s] Returning to main from fn_bat_init", serviceName)
		return -1
	}
	return 0
}

func CLN_PACK_CLNT(args []string, Db *gorm.DB) int {
	var resultTmp int

	log.Printf("[%s] Entering CLN_PACK_CLNT", serviceName)
	resultTmp = fnGetNxtRec(args, Db)
	if resultTmp != 0 {
		log.Printf("[%s] failed in getting the next record returning with result code: %d", serviceName, resultTmp)
		log.Printf("[%s] Exiting CLN_PACK_CLNT with error", serviceName)
		return -1
	}
	log.Printf("[%s] Exiting CLN_PACK_CLNT", serviceName)
	return 0
}

func fnGetNxtRec(args []string, Db *gorm.DB) int {
	log.Printf("[%s] Entering fnGetNxtRec", serviceName)

	exchngbook := &structures.Vw_xchngbook{}

	resultTmp := fnSeqToOmd(Db, exchngbook)
	if resultTmp != 0 {
		log.Printf("[%s] Failed to fetch data into exchngbook structure", serviceName)
		return -1
	}

	log.Printf("[%s] Exiting fnGetNxtRec", serviceName)
	return 0
}

func fnSeqToOmd(db *gorm.DB, xchngbook *structures.Vw_xchngbook) int {
	log.Printf("[%s] Entering fnSeqToOmd", serviceName)
	log.Printf("[%s] Before extracting the data from the 'fxb_ordr_rfrnc' and storing it in the 'xchngbook' structure", serviceName)

	query := `
	SELECT fxb_ordr_rfrnc,
       fxb_lmt_mrkt_sl_flg,
       fxb_dsclsd_qty,
       fxb_ordr_tot_qty,
       fxb_lmt_rt,
       fxb_stp_lss_tgr,
       TO_CHAR(TO_DATE(fxb_ordr_valid_dt, 'YYYY-MM-DD'), 'DD-Mon-YYYY') AS valid_dt,
       CASE
           WHEN fxb_ordr_type = 'V' THEN 'T'
           ELSE fxb_ordr_type
       END AS ord_typ,
       fxb_rqst_typ,
       fxb_ordr_sqnc
FROM FXB_FO_XCHNG_BOOK
WHERE fxb_xchng_cd = ?
  AND fxb_pipe_id = ?
  AND TO_DATE(fxb_mod_trd_dt, 'YYYY-MM-DD') = TO_DATE(?, 'YYYY-MM-DD')
  AND fxb_ordr_sqnc = (
      SELECT MIN(fxb_b.fxb_ordr_sqnc)
      FROM FXB_FO_XCHNG_BOOK fxb_b
      WHERE fxb_b.fxb_xchng_cd = ?
        AND TO_DATE(fxb_b.fxb_mod_trd_dt, 'YYYY-MM-DD') = TO_DATE(?, 'YYYY-MM-DD')
        AND fxb_b.fxb_pipe_id = ?
        AND fxb_b.fxb_plcd_stts = 'R'
  )
`

	log.Printf("[%s] Executing query to fetch order details", serviceName)

	log.Printf("[%s] C_xchng_cd: %s", serviceName, xchngbook.C_xchng_cd)
	log.Printf("[%s] C_pipe_id: %s", serviceName, xchngbook.C_pipe_id)
	log.Printf("[%s] C_mod_trd_dt: %s", serviceName, xchngbook.C_mod_trd_dt)

	row := db.Raw(query,
		xchngbook.C_xchng_cd,
		xchngbook.C_pipe_id,
		xchngbook.C_mod_trd_dt,
		xchngbook.C_xchng_cd,
		xchngbook.C_mod_trd_dt,
		xchngbook.C_pipe_id).Row()

	/* xchngbook.C_xchng_cd
	   The value of 'xchngbook.C_xchng_cd' is obtained from a query in 'fn_bat_init':

	   'EXEC SQL
	   SELECT opm_xchng_cd,
	          opm_max_pnd_ord,
	          opm_stream_no
	   INTO   :sql_c_xchng_cd,
	          :sql_li_max_pnd_ord,
	          :i_stream_count_opm
	   FROM   opm_ord_pipe_mstr
	   WHERE  opm_pipe_id = :sql_c_pipe_id;'
	*/

	/* C_pipe_id

	   The value of this is obtained from args[3] and is initialized in 'fn_bat_init'.
	*/

	/* "C_mod_trd_dt"

	   The value of "C_mod_trd_dt" is obtained from a query in "fn_bat_init" and then set to 'null'. The logic for this is unclear.
	   [
	       SELECT exg_nxt_trd_dt,
	              exg_brkr_id,
	              exg_ctcl_id
	       INTO   :sql_c_nxt_trd_date,
	              :st_opm_mstr.c_xchng_brkr_id,
	              :sql_c_xchng_ctcl_id
	       FROM   exg_xchng_mstr
	       WHERE  exg_xchng_cd = :sql_c_xchng_cd;
	   ]
	*/

	/* "L_ord_seq"

	   It is set to 0 in fn_bat_init.

	   st_rqst1.l_ord_seq = li_seq_nm;  // Set to '0' in fn_bat_init.
	*/

	err := row.Scan(
		&xchngbook.C_slm_flg,
		&xchngbook.L_dsclsd_qty,
		&xchngbook.L_ord_tot_qty,
		&xchngbook.L_ord_lmt_rt,
		&xchngbook.L_stp_lss_tgr,
		&xchngbook.C_valid_dt,
		&xchngbook.C_ord_typ,
		&xchngbook.C_req_typ,
		&xchngbook.L_ord_seq,
	)

	if err != nil {
		log.Printf("[%s] Error scanning row: %v", serviceName, err)
		log.Printf("[%s] Exiting fnSeqToOmd with error", serviceName)
		return -1
	}

	log.Printf("[%s] Data extracted and stored in the 'xchngbook' structure:", serviceName)
	log.Printf("[%s]   C_slm_flg: %c", serviceName, xchngbook.C_slm_flg)
	log.Printf("[%s]   L_dsclsd_qty: %d", serviceName, xchngbook.L_dsclsd_qty)
	log.Printf("[%s]   L_ord_tot_qty: %d", serviceName, xchngbook.L_ord_tot_qty)
	log.Printf("[%s]   L_ord_lmt_rt: %d", serviceName, xchngbook.L_ord_lmt_rt)
	log.Printf("[%s]   L_stp_lss_tgr: %d", serviceName, xchngbook.L_stp_lss_tgr)
	log.Printf("[%s]   C_valid_dt: %s", serviceName, xchngbook.C_valid_dt)
	log.Printf("[%s]   C_ord_typ: %c", serviceName, xchngbook.C_ord_typ)
	log.Printf("[%s]   C_req_typ: %c", serviceName, xchngbook.C_req_typ)
	log.Printf("[%s]   L_ord_seq: %d", serviceName, xchngbook.L_ord_seq)

	log.Printf("[%s] Data extracted and stored in the 'xchngbook' structure successfully", serviceName)
	log.Printf("[%s] Exiting fnSeqToOmd", serviceName)
	return 0

	// comment for first commit of the day (09/08/2024)
}
