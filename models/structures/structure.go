package structures

const (
	LEN_BROKER_ID       = 5
	LEN_FILLER_OPTIONS  = 3
	LEN_ACCOUNT_NUMBER  = 10
	LEN_REMARKS         = 24
	LEN_SETTLOR         = 12
	LEN_ALPHA_CHAR      = 2
	LEN_TIME_STAMP      = 8
	LEN_INSTRUMENT_NAME = 6
	LEN_SYMBOL_NSE      = 10
	LEN_OPTION_TYPE     = 2
)

type Vw_xchngbook struct {
	C_ordr_rfrnc  string // this valiable we are getting from the table
	C_pipe_id     string // this variable we are getting from the args[3]
	C_mod_trd_dt  string // checkout this one, we are getting from table "exg_xchng_mstr"
	L_ord_seq     int32  // this variable we are gettign from the table
	C_slm_flg     string //byte   // this variable we are gettign from the table
	L_dsclsd_qty  int32  // this variable we are gettign from the table
	L_ord_tot_qty int32  // this variable we are gettign from the table
	L_ord_lmt_rt  int32  // this variable we are gettign from the table
	L_stp_lss_tgr int32  // this variable we are gettign from the table
	C_ord_typ     string //byte   // this variable we are gettign from the table
	C_req_typ     string //byte   // this variable we are gettign from the table
	C_valid_dt    string // this variable we are gettign from the table
	C_xchng_cd    string // this we are getting from "opm_ord_pipe_mstr"
}

/*
	value of this structure "Vw_xchngbook" we are extracting from "FXB_FO_XCHNG_BOOK". and all this is happening in "fn_seq_to_omd" in "cln_pack_clnt".
*/

type Vw_orderbook struct {
	c_cln_mtch_accnt string
	c_ordr_flw       byte
	l_ord_tot_qty    int32
	l_exctd_qty      int32
	l_exctd_qty_day  int32
	c_settlor        string

	c_spl_flg     byte
	c_ack_tm      string
	c_prev_ack_tm string
	c_pro_cli_ind byte
	c_ctcl_id     string
}

/*
	values of this structure  "Vw_orderbook". we are extracting from "fod_fo_ordr_dtls" . all this is happening in "fn_ref_to_ord" in "cln_pack_clnt"
*/

type Vw_nse_cntrct struct {
	c_prd_typ      byte
	c_ctgry_indstk byte
	c_symbol       string
	l_ca_lvl       int32
	l_token_id     int32
}

/*
	this value of this structure we are getting from "Vw_contract". which is in "fn_get_ext_cnt" in "cln_pack_clnt"

	 **** problem here is i am not able find the source like from where we are getting the value of the structure "Vw_contract".
*/

type St_opm_pipe_mstr struct {
	/*
		long li_opm_brnch_id;	 // --------------------- (this one is used) from (OPM_ORD_PIPE_MSTR)
		char c_xchng_brkr_id[6]; //--------------------- (this one is used)	 from (exg_xchng_mstr)
		char c_opm_trdr_id[6];	 //--------------------- (this one is used) from (OPM_ORD_PIPE_MSTR)
		int si_user_typ_glb;	 //----------------------(this one is used) "this one we are getting from "cofiguration file" "GetProcessSpaceValue()""

	*/
	li_opm_brnch_id int64  // null=-1
	c_xchng_brkr_id string // null="*"
	c_opm_trdr_id   string // null="*"
	si_user_typ_glb int    // null=0   (i think , it can be 0 (for trader) , 4 (for CORPORATE_MANAGER) , 5 (for BRANCH_MANAGER) )

}

/*
	fields of this structure we are getting from different sources.
*/

// ================================ till now all the strctures we are getting as parameters ========================

type St_req_q_data struct {
	li_msg_type int64
	/* we are changeing this field multiple times.
		1. it is changing, where we are packing the structure which is in "fn_pack_ordnry_ord_to_nse" here we are setting the the variable based on the request type like (if request is "new" then "li_msg_type = BOARD_LOT_IN" and if request is "modify" the "li_msg_type = ORDER_MOD_IN" and if request is "cancel" then "li_msg_type = ORDER_CANCEL_IN")
	 	2.
	*/
	st_exch_msg_data St_exch_msg // in the original structure here a union is used . but i am directly using the structure.
}

type St_exch_msg struct {
	st_net_header St_net_hdr
	st_oe_res     St_oe_reqres
}

type St_net_hdr struct {
	si_message_length int16
	i_seq_num         int32
	c_checksum        [16]byte
}

type St_oe_reqres struct {
	st_hdr                        St_int_header
	c_participant_type            byte
	c_filler_1                    byte
	si_competitor_period          int16
	si_solicitor_period           int16
	c_modified_cancelled_by       byte
	c_filler_2                    byte
	si_reason_code                int16
	c_filler_3                    string
	l_token_no                    int32
	st_con_desc                   St_contract_desc
	c_counter_party_broker_id     [LEN_BROKER_ID]byte
	c_filler_4                    byte
	c_filler_5                    string
	c_closeout_flg                byte
	c_filler_6                    byte
	si_order_type                 int16
	d_order_number                float64
	c_account_number              [LEN_ACCOUNT_NUMBER]byte
	si_book_type                  int16
	si_buy_sell_indicator         int16
	li_disclosed_volume           int32
	li_disclosed_volume_remaining int32
	li_total_volume_remaining     int32
	li_volume                     int32
	li_volume_filled_today        int32
	li_price                      int32
	li_trigger_price              int32
	li_good_till_date             int32
	li_entry_date_time            int32
	li_minimum_fill_aon_volume    int32
	li_last_modified              int32
	st_ord_flg                    St_order_flags
	si_branch_id                  int16
	li_trader_id                  int32 // Changed from int16 to int32 in Ver 1.8
	c_broker_id                   [LEN_BROKER_ID]byte
	c_remarks                     [LEN_REMARKS]byte
	c_open_close                  byte
	c_settlor                     [LEN_SETTLOR]byte
	si_pro_client_indicator       int16
	si_settlement_period          int16
	c_cover_uncover               byte
	c_giveup_flag                 byte
	i_order_seq                   int32 // Changed from i_ordr_rfrnc to i_ordr_sqnc in Ver 1.7
	d_nnf_field                   float64
	d_filler19                    float64
	c_pan                         string // Added in Ver 2.7
	l_algo_id                     int32  // Added in Ver 2.7
	si_algo_category              int16  // Added in Ver 2.7
	ll_lastactivityref            int64  // Added in Ver 2.9
	c_reserved                    string // Updated in Ver 2.9
}

type St_int_header struct {
	Si_transaction_code int16
	Li_log_time         int32

	C_alpha_char      [LEN_ALPHA_CHAR]byte
	Li_trader_id      int32
	Si_error_code     int16
	C_filler_2        string
	C_time_stamp_1    [LEN_TIME_STAMP]byte
	C_time_stamp_2    [LEN_TIME_STAMP]byte
	Si_message_length int16
}

type St_contract_desc struct {
	C_instrument_name [LEN_INSTRUMENT_NAME]byte
	C_symbol          [LEN_SYMBOL_NSE]byte
	Li_expiry_date    int32
	Li_strike_price   int32
	C_option_type     [LEN_OPTION_TYPE]byte
	Si_ca_level       int16
}

type St_order_flags struct {
	flg_ato         uint32 // 1 bit
	flg_market      uint32 // 1 bit
	flg_sl          uint32 // 1 bit
	flg_mit         uint32 // 1 bit
	flg_day         uint32 // 1 bit
	flg_gtc         uint32 // 1 bit
	flg_ioc         uint32 // 1 bit
	flg_aon         uint32 // 1 bit
	flg_mf          uint32 // 1 bit
	flg_matched_ind uint32 // 1 bit
	flg_traded      uint32 // 1 bit
	flg_modified    uint32 // 1 bit
	flg_frozen      uint32 // 1 bit
	flg_filler1     uint32 // 3 bits
}

//-------------------------------------------------------------------------------------------------------------------

type St_pk_sequence struct {
	c_pipe_id  string
	c_trd_dt   string
	c_rqst_typ string
	i_seq_num  int32
}

type St_addtnal_order_flags struct {
	c_cover_uncover byte
}
