package app

import (
	"DATA_FWD_TAP/models"
	"DATA_FWD_TAP/models/structures"
	"log"

	"gorm.io/gorm"
)

type ClnPackClntManager struct {
	serviceName string                   // Name of the service being managed
	xchngbook   *structures.Vw_xchngbook // Pointer to the Vw_xchngbook structure
	orderbook   *structures.Vw_orderbook // Pointer to the Vw_orderbook structure
	cPanNo      string                   // Pan number, initialized in the 'fnRefToOrd' method
	cLstActRef  string                   // Last activity reference, initialized in the 'fnRefToOrd' method
	cEspID      string                   // ESP ID, initialized in the 'fnRefToOrd' method
	cAlgoID     string                   // Algorithm ID, initialized in the 'fnRefToOrd' method
	cSourceFlg  string                   // Source flag, initialized in the 'fnRefToOrd' method
}

func (cpcm *ClnPackClntManager) Fn_bat_init(args []string, Db *gorm.DB) int {
	var temp_str string
	var resultTmp int
	cpcm.serviceName = "cln_pack_clnt"
	cpcm.xchngbook = &structures.Vw_xchngbook{}
	log.Printf("[%s] Entering Fn_bat_init", cpcm.serviceName)

	// we are getting the 7 args
	if len(args) < 7 {
		log.Printf("[%s] Error: insufficient arguments provided", cpcm.serviceName)
		return -1
	}

	cpcm.xchngbook.C_pipe_id = args[3]

	log.Printf("[%s] Copied pipe ID from args[3]: %s", cpcm.serviceName, cpcm.xchngbook.C_pipe_id)

	// Fetch exchange code
	queryForOpm_Xchng_Cd := `SELECT opm_xchng_cd
				FROM opm_ord_pipe_mstr
				WHERE opm_pipe_id = ?`

	log.Printf("[%s] Executing query to fetch exchange code with pipe ID: %s", cpcm.serviceName, cpcm.xchngbook.C_pipe_id)

	row := Db.Raw(queryForOpm_Xchng_Cd, cpcm.xchngbook.C_pipe_id).Row()
	temp_str = ""
	err := row.Scan(&temp_str)
	if err != nil {
		log.Printf("[%s] Error scanning row for exchange code: %v", cpcm.xchngbook.C_xchng_cd, err)
		return -1
	}

	cpcm.xchngbook.C_xchng_cd = temp_str
	temp_str = ""
	log.Printf("[%s] Exchange code fetched: %s", cpcm.serviceName, cpcm.xchngbook.C_xchng_cd)

	// Fetch modification trade date
	queryFor_exg_nxt_trd_dt := `SELECT exg_nxt_trd_dt
				FROM exg_xchng_mstr
				WHERE exg_xchng_cd = ?`

	log.Printf("[%s] Executing query to fetch modification trade date with exchange code: %s", cpcm.serviceName, cpcm.xchngbook.C_xchng_cd)
	row2 := Db.Raw(queryFor_exg_nxt_trd_dt, cpcm.xchngbook.C_xchng_cd).Row()

	err2 := row2.Scan(&temp_str)
	if err2 != nil {
		log.Printf("[%s] Error scanning row for modification trade date: %v", cpcm.serviceName, err2)
		return -1
	}

	TmpByteArr := []byte(temp_str) // these i have to look again
	TmpArrLen := len(TmpByteArr)   // this i have to look again

	cpcm.xchngbook.C_mod_trd_dt = temp_str

	models.SETNULL(cpcm.serviceName, []byte(cpcm.xchngbook.C_mod_trd_dt), TmpArrLen)

	log.Printf("[%s] Modification trade date fetched and set in C_mod_trd_dt: %s", cpcm.serviceName, cpcm.xchngbook.C_mod_trd_dt[:TmpArrLen])

	cpcm.xchngbook.L_ord_seq = 0 // I am initially setting it to '0' because it was set that way in 'fn_bat_init' and I have not seen it getting changed anywhere. If I find it being changed somewhere, I will update it accordingly.

	log.Printf("[%s] L_ord_seq initialized to %d", cpcm.serviceName, cpcm.xchngbook.L_ord_seq)

	resultTmp = cpcm.CLN_PACK_CLNT(args, Db)
	if resultTmp != 0 {
		log.Printf("[%s] CLN_PACK_CLNT failed with result code: %d", cpcm.serviceName, resultTmp)
		log.Printf("[%s] Returning to main from fn_bat_init", cpcm.serviceName)
		return -1
	}
	log.Printf("[%s] Exiting Fn_bat_init", cpcm.serviceName)
	return 0
}

func (cpcm *ClnPackClntManager) CLN_PACK_CLNT(args []string, Db *gorm.DB) int {
	var resultTmp int

	log.Printf("[%s] Entering CLN_PACK_CLNT", cpcm.serviceName)
	resultTmp = cpcm.fnGetNxtRec(args, Db)
	if resultTmp != 0 {
		log.Printf("[%s] failed in getting the next record returning with result code: %d", cpcm.serviceName, resultTmp)
		log.Printf("[%s] Exiting CLN_PACK_CLNT with error", cpcm.serviceName)
		return -1
	}
	log.Printf("[%s] Exiting CLN_PACK_CLNT", cpcm.serviceName)
	return 0
}

func (cpcm *ClnPackClntManager) fnGetNxtRec(args []string, Db *gorm.DB) int {
	log.Printf("[%s] Entering fnGetNxtRec", cpcm.serviceName)

	resultTmp := cpcm.fnSeqToOmd(Db)
	if resultTmp != 0 {
		log.Printf("[%s] Failed to fetch data into exchngbook structure", cpcm.serviceName)
		return -1
	}

	if cpcm.xchngbook == nil {
		log.Printf("[%s] Error: xchngbook is nil", cpcm.serviceName)
		return -1
	}

	// Log the assignment of C_ordr_rfrnc
	log.Printf("[%s] Assigning C_ordr_rfrnc from exchngbook to orderbook", cpcm.serviceName)
	cpcm.orderbook.C_ordr_rfrnc = cpcm.xchngbook.C_ordr_rfrnc

	resultTmp = cpcm.fnRefToOrd(Db)
	if resultTmp != 0 {
		log.Printf("[%s] Failed to fetch data into orderbook structure", cpcm.serviceName)
		return -1
	}

	log.Printf("[%s] Exiting fnGetNxtRec", cpcm.serviceName)
	return 0
}

func (cpcm *ClnPackClntManager) fnSeqToOmd(db *gorm.DB) int {
	log.Printf("[%s] Entering fnSeqToOmd", cpcm.serviceName)
	log.Printf("[%s] Before extracting the data from the 'fxb_ordr_rfrnc' and storing it in the 'xchngbook' structure", cpcm.serviceName)

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
WHERE fxb_xchng_cd =?
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

	log.Printf("[%s] Executing query to fetch order details", cpcm.serviceName)

	log.Printf("[%s] C_xchng_cd: %s", cpcm.serviceName, cpcm.xchngbook.C_xchng_cd)
	log.Printf("[%s] C_pipe_id: %s", cpcm.serviceName, cpcm.xchngbook.C_pipe_id)
	log.Printf("[%s] C_mod_trd_dt: %s", cpcm.serviceName, cpcm.xchngbook.C_mod_trd_dt)

	row := db.Raw(query,
		cpcm.xchngbook.C_xchng_cd,
		cpcm.xchngbook.C_pipe_id,
		cpcm.xchngbook.C_mod_trd_dt,
		cpcm.xchngbook.C_xchng_cd,
		cpcm.xchngbook.C_mod_trd_dt,
		cpcm.xchngbook.C_pipe_id).Row()
	/*

	 */

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
		&cpcm.xchngbook.C_ordr_rfrnc,
		&cpcm.xchngbook.C_slm_flg,
		&cpcm.xchngbook.L_dsclsd_qty,
		&cpcm.xchngbook.L_ord_tot_qty,
		&cpcm.xchngbook.L_ord_lmt_rt,
		&cpcm.xchngbook.L_stp_lss_tgr,
		&cpcm.xchngbook.C_valid_dt,
		&cpcm.xchngbook.C_ord_typ,
		&cpcm.xchngbook.C_req_typ,
		&cpcm.xchngbook.L_ord_seq,
	)

	if err != nil {
		log.Printf("[%s] Error scanning row: %v", cpcm.serviceName, err)
		log.Printf("[%s] Exiting fnSeqToOmd with error", cpcm.serviceName)
		return -1
	}

	log.Printf("[%s] Data extracted and stored in the 'xchngbook' structure:", cpcm.serviceName)
	log.Printf("[%s]   C_ordr_rfrnc: 	%s", cpcm.serviceName, cpcm.xchngbook.C_ordr_rfrnc)
	log.Printf("[%s]   C_slm_flg: 		%s", cpcm.serviceName, cpcm.xchngbook.C_slm_flg)
	log.Printf("[%s]   L_dsclsd_qty: 	%d", cpcm.serviceName, cpcm.xchngbook.L_dsclsd_qty)
	log.Printf("[%s]   L_ord_tot_qty: 	%d", cpcm.serviceName, cpcm.xchngbook.L_ord_tot_qty)
	log.Printf("[%s]   L_ord_lmt_rt: 	%d", cpcm.serviceName, cpcm.xchngbook.L_ord_lmt_rt)
	log.Printf("[%s]   L_stp_lss_tgr: 	%d", cpcm.serviceName, cpcm.xchngbook.L_stp_lss_tgr)
	log.Printf("[%s]   C_valid_dt: 		%s", cpcm.serviceName, cpcm.xchngbook.C_valid_dt)
	log.Printf("[%s]   C_ord_typ: 		%s", cpcm.serviceName, cpcm.xchngbook.C_ord_typ)
	log.Printf("[%s]   C_req_typ: 		%s", cpcm.serviceName, cpcm.xchngbook.C_req_typ)
	log.Printf("[%s]   L_ord_seq: 		%d", cpcm.serviceName, cpcm.xchngbook.L_ord_seq)

	log.Printf("[%s] Data extracted and stored in the 'xchngbook' structure successfully", cpcm.serviceName)
	log.Printf("[%s] Exiting fnSeqToOmd", cpcm.serviceName)
	return 0

}

func (cpcm *ClnPackClntManager) fnRefToOrd(db *gorm.DB) int {
	/*
		c_pan_no --> we are simply setting it to 0 using memset in the
		MEMSET(c_ordr_rfrnc);
		MEMSET(c_pan_no);
		MEMSET(c_lst_act_ref);
		MEMSET(c_esp_id);
		MEMSET(c_algo_id);
		char c_source_flg = '\0';
	*/
	log.Printf("[%s] Entering fnFetchOrderDetails", cpcm.serviceName)
	log.Printf("[%s] Before extracting the data from the 'fod_fo_ordr_dtls' and storing it in the 'orderbook' structure", cpcm.serviceName)

	query := `
    SELECT 
        fod_clm_mtch_accnt,
        fod_ordr_flw,
        fod_ordr_tot_qty,
        fod_exec_qty,
        COALESCE(fod_exec_qty_day, 0),
        fod_settlor,
        fod_spl_flag,
        TO_CHAR(fod_ord_ack_tm, 'YYYY-MM-DD HH24:MI:SS') AS fod_ord_ack_tm,
        TO_CHAR(fod_lst_rqst_ack_tm, 'YYYY-MM-DD HH24:MI:SS') AS fod_lst_rqst_ack_tm,
        fod_pro_cli_ind,
        COALESCE(fod_ctcl_id, ' '),
        COALESCE(fod_pan_no, '*'),
        COALESCE(fod_lst_act_ref, '0'),
        COALESCE(FOD_ESP_ID, '*'),
        COALESCE(FOD_ALGO_ID, '*'),
        COALESCE(FOD_SOURCE_FLG, '*')
    FROM 
        fod_fo_ordr_dtls
    WHERE 
        fod_ordr_rfrnc = ?
    FOR UPDATE NOWAIT
    `

	log.Printf("[%s] Executing query to fetch order details", cpcm.serviceName)

	log.Printf("[%s] Order Reference: %s", cpcm.serviceName, cpcm.orderbook.C_ordr_rfrnc)

	row := db.Raw(query, cpcm.orderbook.C_ordr_rfrnc).Row()

	err := row.Scan(
		&cpcm.orderbook.C_cln_mtch_accnt,
		&cpcm.orderbook.C_ordr_flw,
		&cpcm.orderbook.L_ord_tot_qty,
		&cpcm.orderbook.L_exctd_qty,
		&cpcm.orderbook.L_exctd_qty_day,
		&cpcm.orderbook.C_settlor,
		&cpcm.orderbook.C_spl_flg,
		&cpcm.orderbook.C_ack_tm,
		&cpcm.orderbook.C_prev_ack_tm,
		&cpcm.orderbook.C_pro_cli_ind,
		&cpcm.orderbook.C_ctcl_id,
		&cpcm.cPanNo,
		&cpcm.cLstActRef,
		&cpcm.cEspID,
		&cpcm.cAlgoID,
		&cpcm.cSourceFlg,
	)

	if err != nil {
		log.Printf("[%s] Error scanning row: %v", cpcm.serviceName, err)
		log.Printf("[%s] Exiting fnFetchOrderDetails with error", cpcm.serviceName)
		return -1
	}

	log.Printf("[%s] Data extracted and stored in the 'orderbook' structure:", cpcm.serviceName)
	log.Printf("[%s]   C_cln_mtch_accnt:   %s", cpcm.serviceName, cpcm.orderbook.C_cln_mtch_accnt)
	log.Printf("[%s]   C_ordr_flw:        %s", cpcm.serviceName, cpcm.orderbook.C_ordr_flw)
	log.Printf("[%s]   L_ord_tot_qty:     %d", cpcm.serviceName, cpcm.orderbook.L_ord_tot_qty)
	log.Printf("[%s]   L_exctd_qty:       %d", cpcm.serviceName, cpcm.orderbook.L_exctd_qty)
	log.Printf("[%s]   L_exctd_qty_day:   %d", cpcm.serviceName, cpcm.orderbook.L_exctd_qty_day)
	log.Printf("[%s]   C_settlor:         %s", cpcm.serviceName, cpcm.orderbook.C_settlor)
	log.Printf("[%s]   C_spl_flg:         %s", cpcm.serviceName, cpcm.orderbook.C_spl_flg)
	log.Printf("[%s]   C_ack_tm:          %s", cpcm.serviceName, cpcm.orderbook.C_ack_tm)
	log.Printf("[%s]   C_prev_ack_tm:     %s", cpcm.serviceName, cpcm.orderbook.C_prev_ack_tm)
	log.Printf("[%s]   C_pro_cli_ind:     %s", cpcm.serviceName, cpcm.orderbook.C_pro_cli_ind)
	log.Printf("[%s]   C_ctcl_id:         %s", cpcm.serviceName, cpcm.orderbook.C_ctcl_id)
	log.Printf("[%s]   C_pan_no:          %s", cpcm.serviceName, cpcm.cPanNo)
	log.Printf("[%s]   C_lst_act_ref_tmp: %s", cpcm.serviceName, cpcm.cLstActRef)
	log.Printf("[%s]   C_esp_id:          %s", cpcm.serviceName, cpcm.cEspID)
	log.Printf("[%s]   C_algo_id:         %s", cpcm.serviceName, cpcm.cAlgoID)
	log.Printf("[%s]   C_source_flg_tmp:  %s", cpcm.serviceName, cpcm.cSourceFlg)

	log.Printf("[%s] Data extracted and stored in the 'orderbook' structure successfully", cpcm.serviceName)
	log.Printf("[%s] Exiting fnFetchOrderDetails", cpcm.serviceName)
	return 0
}
