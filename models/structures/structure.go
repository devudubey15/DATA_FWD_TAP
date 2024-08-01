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

type vw_xchngbook struct {
	c_pipe_id      string // null="*"
	c_mod_trd_dt   string // null="*"
	c_slm_flg      byte   // null='*'
	l_dsclsd_qty   int64  // null=-1
	l_ord_tot_qty  int64  // null=-1
	l_ord_lmt_rt   int64  // null=-1
	l_stp_lss_tgr  int64  // null=-1
	c_ord_typ      byte   // null='*'
	c_sprd_ord_ind byte   // null='*'
	c_req_typ      byte   // null='*'
}

type vw_orderbook struct {
	c_cln_mtch_accnt string // null="*"
	c_ordr_flw       byte   // null='*'
	l_ord_tot_qty    int64  // null=-1
	l_exctd_qty      int64  // null=-1
	l_exctd_qty_day  int64  // null=-1
	c_settlor        string // null="*"
	c_xchng_ack      string // null="*"
	c_spl_flg        byte   // null='*'
	c_ack_tm         string // null="*"
	c_prev_ack_tm    string // null="*"
	c_pro_cli_ind    byte   // null='*'
	c_ctcl_id        string // null="*"
}

type vw_nse_cntrct struct {
	c_xchng_cd     string // null="*"
	c_ctgry_indstk byte   // null='*'
	c_symbol       string // null="*"
	l_ca_lvl       int64  // null=-1
	l_token_id     int64  // null=-1
}

type vw_spdordbk struct {
	c_sprd_ord_rfrnc [3]string // null="*"
}

type st_opm_pipe_mstr struct {
	li_opm_brnch_id int64  // null=-1
	c_xchng_brkr_id string // null="*"
	c_opm_trdr_id   string // null="*"
	si_user_typ_glb int    // null=0   (i think , it can be 0 (for trader) , 4 (for CORPORATE_MANAGER) , 5 (for BRANCH_MANAGER) )
}

type st_req_q_data struct {
	li_msg_type      int64
	st_exch_msg_data st_exch_msg // it is union in original file
}

type st_exch_msg struct {
	st_net_header  st_net_hdr
	st_exch_sndmsg st_exch_snd_msg
}

type st_net_hdr struct {
	si_message_length int16
	i_seq_num         int32
	c_checksum        [16]byte
}

type st_exch_snd_msg struct {
	st_spdoe_reqres st_spd_oe_reqres
}

type st_spd_oe_reqres struct {
	St_hdr                     st_int_header
	C_participant_type         byte
	C_filler1                  byte
	Si_competitor_period       int16
	Si_solicitor_period        int16
	C_mod_cxl_by               byte
	C_filler2                  byte
	Si_reason_code             int16
	C_start_alpha              [2]byte
	C_end_alpha                [2]byte
	L_token                    int32 // V1.4
	St_cntrct_desc             st_contract_desc
	C_op_broker_id             [LEN_BROKER_ID]byte
	C_filler3                  byte
	C_filler_options           [LEN_FILLER_OPTIONS]byte
	C_filler4                  byte
	Si_order_type              int16
	D_order_number             float64
	C_account_number           [LEN_ACCOUNT_NUMBER]byte
	Si_book_type               int16
	Si_buy_sell                int16
	Li_disclosed_vol           int32
	Li_disclosed_vol_remaining int32
	Li_total_vol_remaining     int32
	Li_volume                  int32
	Li_volume_filled_today     int32
	Li_price                   int32
	Li_trigger_price           int32
	Li_good_till_date          int32
	Li_entry_date_time         int32
	Li_min_fill_aon            int32
	Li_last_modified           int32
	St_order_flgs              st_order_flags
	Si_branch_id               int16
	Li_trader_id               int32 // Ver 1.8 Data Type Changed From Short Int To Long
	C_broker_id                [LEN_BROKER_ID]byte
	C_oe_remarks               [LEN_REMARKS]byte
	C_open_close               byte
	C_settlor                  [LEN_SETTLOR]byte
	Si_pro_client              int16
	Si_settlement_period       int16
	C_cover_uncover            byte
	C_give_up_flag             byte
	// Ver 1.3 fields omitted
	I_sprd_seq_no      int
	D_nnf_field        float64
	D_filler19         float64
	C_pan              [10]byte // Ver 2.7
	L_algo_id          int64    // Ver 2.7
	Si_algo_category   int16    // Ver 2.7
	Ll_lastactivityref int64    // Ver 2.9
	C_reserved         [52]byte // Ver 2.9 (Not used)
	Li_spd_price_diff  int32
	St_spd_lg_inf      [2]st_spd_leg_info
}

type st_int_header struct {
	Si_transaction_code int16
	Li_log_time         int32

	C_alpha_char      [LEN_ALPHA_CHAR]byte
	Li_trader_id      int32
	Si_error_code     int16
	C_filler_2        [8]byte
	C_time_stamp_1    [LEN_TIME_STAMP]byte
	C_time_stamp_2    [LEN_TIME_STAMP]byte
	Si_message_length int16
}

type st_spd_leg_info struct {
	L_token                    int32
	St_cntrct_desc             st_contract_desc
	C_op_broker_id             [LEN_BROKER_ID]byte
	C_filler1                  byte
	Si_order_type              int16
	Si_buy_sell                int16
	Li_disclosed_vol           int32
	Li_disclosed_vol_remaining int32
	Li_total_vol_remaining     int32
	Li_volume                  int32
	Li_volume_filled_today     int32
	Li_price                   int32
	Li_trigger_price           int32
	Li_min_fill_aon            int32
	St_order_flgs              st_order_flags
	C_open_close               byte
	C_cover_uncover            byte
	C_giveup_flag              byte
	C_filler2                  byte
}

type st_contract_desc struct {
	C_instrument_name [LEN_INSTRUMENT_NAME]byte
	C_symbol          [LEN_SYMBOL_NSE]byte
	Li_expiry_date    int32
	Li_strike_price   int32
	C_option_type     [LEN_OPTION_TYPE]byte
	Si_ca_level       int16
}

type st_order_flags struct {
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

type st_pk_sequence struct {
	c_pipe_id  [3]byte
	c_trd_dt   [23]byte
	c_rqst_typ [3]byte
	i_seq_num  int32
}

type st_addtnal_order_flags struct {
	c_cover_uncover byte
}

type st_oe_reqres struct {
	st_hdr                        st_int_header
	c_participant_type            byte
	c_filler_1                    byte
	si_competitor_period          int16
	si_solicitor_period           int16
	c_modified_cancelled_by       byte
	c_filler_2                    byte
	si_reason_code                int16
	c_filler_3                    [4]byte
	l_token_no                    int32
	st_con_desc                   st_contract_desc
	c_counter_party_broker_id     [LEN_BROKER_ID]byte
	c_filler_4                    byte
	c_filler_5                    [2]byte
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
	st_ord_flg                    st_order_flags
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
	c_pan                         [10]byte // Added in Ver 2.7
	l_algo_id                     int32    // Added in Ver 2.7
	si_algo_category              int16    // Added in Ver 2.7
	ll_lastactivityref            int64    // Added in Ver 2.9
	c_reserved                    [52]byte // Updated in Ver 2.9
}
