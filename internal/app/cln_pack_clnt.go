package app

import (
	"DATA_FWD_TAP/models/structures"
	"log"

	"gorm.io/gorm"
)

var sql_pipe_id [3]byte
var sql_c_xchng_cd [4]byte

func fn_bat_init(args []string, Db *gorm.DB) {

	arg := args[3]
	copy(sql_pipe_id[:], arg)

}

func CLN_PACK_CLNT(args []string, Db *gorm.DB) {
	fnGetNxtCnt(args, Db)
}

func fnGetNxtCnt(args []string, Db *gorm.DB) {
	exchngbook := &structures.Vw_xchngbook{}
	fnSeqToOmd(Db, exchngbook)
}

func fnSeqToOmd(db *gorm.DB, xchngbook *structures.Vw_xchngbook) int {
	log.Println("Inside the 'fnSeqToOmd'")
	log.Println("Before extracting the data from the 'fxb_ordr_rfrnc' and storing it in the 'xchngbook' structure")

	query := `
		SELECT fxb_ordr_rfrnc,
			fxb_lmt_mrkt_sl_flg,
			fxb_dsclsd_qty,
			fxb_ordr_tot_qty,
			fxb_lmt_rt,
			fxb_stp_lss_tgr,
			to_char(fxb_ordr_valid_dt, 'dd-mon-yyyy') as valid_dt,
			DECODE(fxb_ordr_type, 'V', 'T', fxb_ordr_type) as ord_typ,
			fxb_rqst_typ,
			fxb_ordr_sqnc
			FROM FXB_FO_XCHNG_BOOK
			WHERE 
			fxb_xchng_cd = ?
			AND fxb_pipe_id = ?
			AND fxb_mod_trd_dt = to_date(?, 'dd-Mon-yyyy')
			AND fxb_ordr_sqnc = (
				SELECT min(fxb_b.fxb_ordr_sqnc)
				FROM FXB_FO_XCHNG_BOOK fxb_b
				WHERE fxb_b.fxb_xchng_cd = ?
					AND fxb_b.fxb_mod_trd_dt = to_date(?, 'dd-Mon-yyyy')
					AND fxb_b.fxb_pipe_id = ?
					AND fxb_b.fxb_plcd_stts = 'R'
			)
	`

	row := db.Raw(query, xchngbook.C_xchng_cd, xchngbook.C_pipe_id[:], xchngbook.C_mod_trd_dt[:], xchngbook.L_ord_seq, xchngbook.C_pipe_id[:], xchngbook.C_mod_trd_dt[:], xchngbook.L_ord_seq).Row()

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
		log.Println("Error scanning row:", err)
		return -1
	}

	log.Println("Data extracted and stored in the 'xchngbook' structure successfully")
	return 0
}
