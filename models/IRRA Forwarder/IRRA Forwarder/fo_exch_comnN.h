/**** Structure according to NNF 4.9 ver.  ******/

/****Ver 1.1 Infotech(SR) for NNF 4.9.3 ****/
/****Ver 1.2 Infotech(MM) for NNF (Broker Name is added in Login Structure ****/
/****Ver 1.3 Infotech(MM) for NNF (CTCL ID is added in Order Structure ****/
/****Ver 1.4 Infotech(SS) for NNF (Token no datatype chaned in ver 704 ****/
/****Ver 1.5 Infotech(Shailesh Hinge) for NNF (Extended Facility of closeout flg in case TURNOVER EXCEEDED****/
/****Ver 1.6 Infotech(Shailesh Hinge) for TAP ****/
/****Ver 1.7 23-Jun-2009 ORS2 ****/
/****Ver 1.8 Infotech(Sandeep Patil) for User Id Data Type Change			***/
/****VEr 1.9 INfotech(Sachin Birje)  For Global Indices ****/ 
/****Ver 2.0 Infotech(Sandip Tambe) 2l_Order_OERemarks_Changes**/
/****Ver 2.1 Infotech(Sachin Birje) NSE Broadcast Changes**/
/****Ver 2.2 Infotech(Sachin Birje) NSE Broadcast Changes**/
/****Ver 2.3 Infotech(Mahesh Shinde) CR_ISEC14_48665_Introduction_of_VIX_India Broadcast Changes 26-Feb-2014 **/
/****Ver 2.5 Infotech(Navina Dhumal) for NNF (Introduction of Trade execution range broadcast data) 28-Oct-2014 **/
/****Ver 2.6 Infotech(Parag Kanojia) Changes for Exchange Exception Handling 13-Apr-2018 ****/
/****Ver 2.7 Infotech(Parag Kanojia) for NNF 7.21 Changes 16-Mar-2018****/
/****Ver 2.8 Infotech(Parag Kanojia) Changes for Stream No. Exchange Exception Handling in Download 01-Aug-2018 ****/
/****Ver 2.9	Infotech(Parag Kanojia) Direct Connectivity ****/
/****Ver 3.0 Infotech(Parag Kanojia) Rollover with Spread 30-Aug-2019 ****/
/****Ver 3.1 (Parag) ***/
/****Ver 3.2 Profit order in FPSL (Naveena) ***/
/****Ver 3.3  Flash Trade Changes  (Ravi Malla) ***/
#pragma once  /* VER TOL : TUX on LINUX */
# pragma pack ( 2 )

/********************************************************************/
/* Length                                                           */
/********************************************************************/
# define LEN_ALPHA_CHAR            2
# define LEN_TIME_STAMP            8
# define LEN_ERROR_KEY             14		/**V1.4**/
# define LEN_ERROR_MESSAGE         128
# define LEN_PASSWORD              8
# define LEN_TRADER_NAME           26
# define LEN_BROKER_ID             5
# define LEN_BROKER_NAME           25   /* ver 1.2  */
# define LEN_COLOUR                50
# define LEN_WS_CLASS_NAME         14
# define LEN_ACCOUNT_NUMBER        10
# define LEN_REMARKS               24
# define LEN_REMARKS_1             30
# define LEN_REMARKS_2             25
# define LEN_INSTRUMENT_NAME       6
# define LEN_INSTRUMENT_DESC       25
# define LEN_SYMBOL_NSE            10
# define LEN_OPTION_TYPE           2
# define LEN_ACTIVITY_TYPE         2
# define LEN_BROADCAST_MESSAGE     239
# define LEN_FIRM_NAME             25
# define LEN_INDICATOR             4
# define LEN_BRANCH_NAME           25
# define LEN_SETTLOR               12
# define LEN_ACTION_CODE           3
# define LEN_CREDIT_RATING         12
# define LEN_SECURITY_NAME         25
# define LEN_ASSET_NAME            10
#define LEN_PAN                    10 /* Added in Ver TOL */
/***Reffered Wrong as 12 bytes till date, Changed for Ver 1.6 TAP
# define LEN_PARTICIPANT_NAME      25
**/

# define LEN_PARTICIPANT_NAME			 12		/**Ver 1.6 **/
# define LEN_SERIES_NSE            2
# define LEN_INDEX_NAME            15
# define LEN_INDUSTRY_NAME				 15
# define LEN_BCAST_NAME						 26
# define LEN_CHANGED_NAME 				 10
# define LEN_FILLER_OPTIONS 				3


# define MAX_NO_OPEN_ORDERS        2
# define MAX_NO_TRADES_TODAY       4
# define MAX_RCRDS_MKT_STATS_DATAT 6
# define MAX_NO_SEC                4
# define MAX_TICKER_INDEX          29
# define MAX_MBO                   10
# define MAX_MBP                   10
# define MAX_MKT_WTCH              5
# define MAX_MKT_INFO              3
# define MAX_INDICES               15
# define MAX_INDEX_MAP             10
# define MAX_INDUSTRY_INDICES			 20
# define MAX_MBP_BUYS							 5
# define MAX_MBP_SELLS						 5
# define MAX_BP_DOWNLOAD					 40
# define MAX_TANDE					       43
# define MAX_EQ_OI	               77
# define MAX_SEQ_NUM	         		 2147483647 				/**Added for TAP ****/

/********************************************************************/
/* Message codes                                                    */
/********************************************************************/
/********************************************************************/
/*  2. The logon process                                            */
/********************************************************************/
# define SIGN_ON_REQUEST_IN             2300
# define SIGN_ON_REQUEST_OUT            2301

# define SIGN_OFF_REQUEST_IN            2320
# define SIGN_OFF_REQUEST_OUT           2321

/********************************************************************/
/* 3. The download process                                          */
/********************************************************************/
# define SYSTEM_INFORMATION_IN          1600
# define SYSTEM_INFORMATION_OUT         1601

# define UPDATE_LOCALDB_IN              7300
# define PARTIAL_SYSTEM_INFORMATION     7321
# define UPDATE_LOCALDB_HEADER          7307
# define UPDATE_LOCALDB_DATA            7304
# define UPDATE_LOCALDB_TRAILER         7308

# define EXCH_PORTF_OUT         				1776

# define DOWNLOAD_REQUEST               7000
# define HEADER_RECORD                  7011
# define MESSAGE_RECORD                 7021
# define TRAILER_RECORD                 7031

# define BCAST_INDEX_MSTR_CHG           7325
# define BCAST_BASE_PRICE               7327
# define BCAST_INDEX_MAP_TABLE          7326

/********************************************************************/
/* 4. Order & Trade management                                      */
/********************************************************************/
# define BOARD_LOT_IN                   2000
# define BOARD_LOT_IN_TR                20000		/*** Ver 2.9 ***/
# define BOARD_LOT_OUT                  2001
# define ORDER_CONFIRMATION_OUT         2073
# define ORDER_CONFIRMATION_OUT_TR      20073		/*** Ver 2.9 ***/
# define ORDER_ERROR_OUT                2231

# define ORDER_MOD_IN                   2040
# define ORDER_MOD_IN_TR               	20040		/*** Ver 2.9 ***/
# define ORDER_MOD_OUT                  2041
# define ORDER_MOD_CONFIRM_OUT          2074
# define ORDER_MOD_CONFIRM_OUT_TR       20074		/*** Ver 2.9 ***/
# define ORDER_MOD_REJ_OUT              2042

# define ORDER_CANCEL_IN                2070
# define ORDER_CANCEL_IN_TR             20070		/*** Ver 2.9 ***/
# define ORDER_CANCEL_OUT               2071
# define ORDER_CANCEL_CONFIRM_OUT       2075
# define ORDER_CXL_CONFIRMATION_TR	    20075		/*** Ver 2.9 ***/
# define ORDER_CXL_REJ_OUT              2072

# define PRICE_CONFIRMATION             2012
# define FREEZE_TO_CONTROL              2170

# define TRADE_MOD_IN                   5445
# define TRADE_MOD_OUT                  5446
# define TRADE_CANCEL_IN                5440
# define TRADE_CANCEL_OUT               5441
# define TRADE_ERROR                    2223

/********************************************************************/
/* 5. Spread order & trade management                               */
/********************************************************************/
# define SP_BOARD_LOT_IN                2100
# define SP_BOARD_LOT_OUT               2101
# define SP_ORDER_CONFIRMATION          2124
# define SP_ORDER_ERROR                 2154

# define SP_ORDER_MOD_IN                2118
# define SP_ORDER_MOD_OUT               2119
# define SP_ORDER_MOD_CON_OUT           2136
# define SP_ORDER_MOD_REJ_OUT           2133

# define SP_ORDER_CXL_CONFIRM_OUT       2130
# define SP_ORDER_CANCEL_IN             2106   /** ver 3.0 **/
# define SP_ORDER_CANCEL_REJ_OUT        2127   /** ver 3.0 **/

/********************************************************************/
/* 6. 2L and 3L order entry                                         */
/********************************************************************/
# define TWOL_BOARD_LOT_IN              2102 
# define TWOL_BOARD_LOT_OUT             2103
# define TWOL_ORDER_CONFIRMATION        2125
# define TWOL_ORDER_ERROR               2155
# define TWOL_ORDER_CXL_CONFIRMATION    2131

# define THRL_BOARD_LOT_IN              2104
# define THRL_BOARD_LOT_OUT             2105
# define THRL_ORDER_CONFIRMATION        2126
# define THRL_ORDER_ERROR               2156
# define THRL_ORDER_CXL_CONFIRMATION    2132

/********************************************************************/
/* 7. Unsolicited messages                                          */
/********************************************************************/
# define ON_STOP_NOTIFICATION           2212

# define TRADE_CONFIRMATION             2222
# define TRADE_CONFIRMATION_TR          20222		/*** Ver 2.9 ***/

# define BATCH_ORDER_CXL_OUT		        9002
# define BATCH_SPREAD_CXL_OUT           9004 /** ver 3.0 suchita ***/

# define CTRL_MSG_TO_TRADER             5295

# define TRADE_MODIFY_CONFIRM           2287
# define TRADE_MODIFY_REJECT            2288

# define TRADE_CANCEL_CONFIRM           2282
# define TRADE_CANCEL_REJECT            2286

# define INVITATION_PACKET							15000					/***Added for TAP ***/
/********************************************************************/
/* 8. Bhav copy                                                     */
/********************************************************************/
# define RPRT_MARKET_STATS_OUT_RPT      1833

/********************************************************************/
/* 9. Reports                                                       */
/********************************************************************/
# define REPORT_REQUEST_IN              1900
# define REPORT_REQUEST_OUT             1901
# define OPEN_OREDER_REPORT             1821
# define ORDER_LOG_REPORT               1824
# define TRADE_REPORT                   1827
# define SPREAD_ORDER_LOG_REPORT        1992
# define SPREAD_TRADE_REPORT            1993

/********************************************************************/
/* 10. Broadcast                                                    */
/********************************************************************/
# define BCAST_JRNL_VCT_MSG             6501

# define BCAST_SYSTEM_INFORMATION_OUT		7206

# define BCAST_SECURITY_MSTR_CHG        7305
# define BCAST_INSTR_MSTR_CHG           7324
# define BCAST_PART_MSTR_CHG            7306

# define BCAST_STOCK_STATUS_CHG         7320
# define BCAST_STOCK_STATUS_CHG_PREOPEN 7322

# define BCAST_TURNOVER_EXCEEDED        9010
# define BROADCAST_BROKER_REACTIVATED   9011

# define BC_OPEN_MSG                    6511
# define BC_CLOSE_MSG                   6521
# define BC_PRE_OR_POST_DAY_MSG         6531
# define BC_PRE_OPEN_ENDED              6571
# define EQUAL_BC_POSTCLOSE_MSG         6522

# define BCAST_TICKER_AND_MKT_INDEX     7202
# define BCAST_MBO_MBP_UPDATE           7200
# define BCAST_MW_ROUND_ROBIN           7201
# define BCAST_ONLY_MBP									7208
# define BCAST_TRADE_EXECUTION_RANGE    7220	/*** Ver 2.5 ***/

# define SECURITY_OPEN_PRICE            6013

# define BCAST_CIRCUIT_MSG              6541

# define BCAST_INDICES                  7207
# define BCAST_INDUSTRY_INDEX_UPDATE		7203

# define BCAST_SPD_MBP_DELTA						7211
# define MKT_MVMT_CM_OI_IN						  7130

# define BCAST_GI_INDICES_ASSETS        7732   /*** Ver 1.9 ***/   
# define BCAST_INDICES_VIX 							7216   /*** Ver 2.3 ***/
# define BCAST_CONT_MSG									5294	 /*** Ver 2.6 ***/
# define BCAST_SPD_MSTR_CHG             7309   /*** Ver 3.0 ***/
/********************************************************************/
/* 11. Exercise and position liquidation                            */
/********************************************************************/
# define EX_PL_ENTRY_IN                 4000
# define EX_PL_ENTRY_OUT                4001
# define EX_PL_CONFIRMATION             4002

# define EX_PL_MOD_IN                   4005
# define EX_PL_MOD_OUT                  4006
# define EX_PL_MOD_CONFIRMATION         4007

# define EX_PL_CXL_IN                   4008
# define EX_PL_CXL_OUT                  4009
# define EX_PL_CXL_CONFIRMATION         4010

/********************************************************************/
/* 12. Give up trades                                               */
/********************************************************************/
# define GIVEUP_APPROVED_IN             4500
# define GIVEUP_APPROVED_OUT            4501
# define GIVEUP_APP_CONFIRM             4502
# define GIVEUP_REJECTED_IN             4503
# define GIVEUP_REJECTED_OUT            4504
# define GIVEUP_APPROVE_ALL             4513

/********************************************************************/
/*	Market types                                                    */
/********************************************************************/
# define NORMAL_MARKET	1

/********************************************************************/
/*	Market status																										*/
/********************************************************************/
# define PRE_OPEN				0
# define OPEN						1
# define CLOSED					2
# define PRE_OPEN_ENDED 3
# define POST_CLOSE			4


/********************************************************************/
/*	Security Status																									*/
/********************************************************************/
# define SEC_PRE_OPEN					1
# define SEC_OPEN							2
# define SEC_SUSPENDED				3
# define SEC_PREOPEN_EXTENDED	4
# define SEC_OPEN_WITH_MARKET	5

/********************************************************************/
/*	Indicator whether security is deleted or not                    */
/********************************************************************/
# define SEC_ACTIVE	 'N'
# define SEC_DELETED 'Y'

/********************************************************************/
/*	Book type                                                       */
/********************************************************************/
# define REGULAR_LOT_ORDER    1
# define STOP_LOSS_MIT_ORDER  3

/********************************************************************/
/* Structure definitions                                            */
/********************************************************************/
/********************************************************************/
/*  Common structures on NNF                                        */
/********************************************************************/
/** 2 **/
/*** Ver 2.7 Starts ***/
/********************************************************************/
/*	Algo Id definitions																							*/
/********************************************************************/
#	define FO_AUTO_MTM_ALG_ID	24264
# define FO_AUTO_MTM_ALG_CAT_ID	99
# define FO_PRICE_IMP_ALG_ID 37383
# define FO_PRICE_IMP_ALG_CAT_ID 1
# define FO_PRFT_ORD_ALG_CAT_ID 1  /** ver 3.2 **/
# define FO_PRFT_ORD_ALG_ID 82686 /** ver 3.2 **/
# define FO_FLASH_TRD_ALG_CAT_ID 1  /** ver 3.3 **/
# define FO_FLASH_TRD_ALG_ID 89385 /** ver 3.3 **/
# define FO_NON_ALG_ID 0
# define FO_NON_ALG_CAT_ID 0
/*** Ver 2.7 Ends ***/

/* COMMENTED IN VER TOL : TUX on LINUX * 
 * struct st_broker_eligibility_per_mkt
{
	unsigned int flg_normal_mkt  :1;
	unsigned int flg_oddlot_mkt  :1;
	unsigned int flg_spot_mkt    :1;
	unsigned int flg_auction_mkt :1;
	unsigned int flg_filler1     :4;
	unsigned int flg_filler2     :8;
}; * COMMENTED IN VER TOL */

/* ADDED IN VER TOL : TUX on LINUX */
struct st_broker_eligibility_per_mkt
{
    unsigned short flg_filler1     :4;
    unsigned short flg_auction_mkt :1;
    unsigned short flg_spot_mkt    :1;
    unsigned short flg_oddlot_mkt  :1;
    unsigned short flg_normal_mkt  :1;
    unsigned short flg_filler2     :8;
};

/** 3 **/
struct st_market_status
{
	short int si_normal;
	short int si_oddlot;
	short int si_spot;
	short int si_auction;
};

struct st_ex_market_status
{
	short int si_normal;
	short int si_oddlot;
	short int si_spot;
	short int si_auction;
};

struct st_pl_market_status
{
	short int si_normal;
	short int si_oddlot;
	short int si_spot;
	short int si_auction;
};

struct st_stock_eligible_indicators
{
	unsigned int flg_aon          : 1;
	unsigned int flg_minimum_fill : 1;
	unsigned int flg_books_merged : 1;
	unsigned int flg_filler1      : 5;
	unsigned int flg_filler2      : 8;
};

struct st_index_dtls
{
	char c_index_name[LEN_INDEX_NAME];
	long l_token;														/**V1.4**/
	long int li_last_upd_dt_tm;
};

struct st_download_bp_tok
{
	long l_token;									/**V1.4**to be checked**/
	long int li_baseprice;
	long int li_filler;
};

struct st_bcast_index_map_dtls
{
	char c_bcast_name[LEN_BCAST_NAME];
	char c_changed_name[LEN_CHANGED_NAME];
	char c_delete_flag;
	long int li_last_upd_dt_tm;
};

/** 4 **/
struct st_contract_desc
{
	char      c_instrument_name [ LEN_INSTRUMENT_NAME ];
	char      c_symbol [ LEN_SYMBOL_NSE ];
	long int  li_expiry_date;
	long int  li_strike_price;
	char      c_option_type [ LEN_OPTION_TYPE ];
	short int si_ca_level;
};

struct st_order_flags
{
	unsigned int flg_ato         :1;
	unsigned int flg_market      :1;
	unsigned int flg_sl          :1;
	unsigned int flg_mit         :1;
	unsigned int flg_day         :1;
	unsigned int flg_gtc         :1;
	unsigned int flg_ioc         :1;
	unsigned int flg_aon         :1;
	unsigned int flg_mf          :1;
	unsigned int flg_matched_ind :1;
	unsigned int flg_traded      :1;
	unsigned int flg_modified    :1;
	unsigned int flg_frozen      :1;
	unsigned int flg_filler1     :3;
};

/** 5 **/
struct st_spd_leg_info
{
	long     l_token;				/**V1.4**/
	struct    st_contract_desc st_cntrct_desc;
	char      c_op_broker_id[LEN_BROKER_ID];
	char      c_filler1;
  short int si_order_type;
	short int si_buy_sell;
	long int	li_disclosed_vol;
	long int  li_disclosed_vol_remaining;
	long int  li_total_vol_remaining;
	long int  li_volume;
	long int  li_volume_filled_today;   
	long int  li_price;
	long int  li_trigger_price;
  long int  li_min_fill_aon;
  struct st_order_flags st_order_flgs;
	char   		c_open_close;
	char 			c_cover_uncover;
	char 			c_giveup_flag;
	char 			c_filler2;
};

/** 8 **/
struct st_stats_data 
{
	struct    st_contract_desc st_con_desc;
	short int si_market_type;
	long int	li_open_price;
	long int 	li_high_price;
	long int 	li_low_price;
	long int	li_closing_price;
	long int	li_total_quantity_traded;
	double		d_total_value_traded;
	long int	li_previous_close_price;
	long int	li_open_interest;
	long int	li_chgopeninterest;
	char			c_indicator [ LEN_INDICATOR ];
};

/** 10 **/
struct st_bcast_destination
{
	unsigned int flg_trdr_ws  :1;
	unsigned int flg_cntrl_ws :1;
	unsigned int flg_tandem   :1;
	unsigned int flg_jrnl_req :1;
	unsigned int flg_filler1  :4;
	unsigned int flg_filler2  :1;   /** Ver 2.2 changed from 8 to 1 **/
};

struct st_sec_eligibility_per_market
{ 
	unsigned int	flg_eligibility :1;
	unsigned int flg_filler      :7;
	short int si_stts;
};

struct st_eligibility_indicators
{
	unsigned int flg_participate_in_mkt_indx :1;
	unsigned int flg_aon                     :1;
	unsigned int flg_minimum_fill            :1;
	unsigned int flg_filler1                 :5;
	unsigned int flg_filler2                 :8;
};

struct st_purpose
{
	unsigned int flg_dividend              :1;
	unsigned int flg_rights                :1;
	unsigned int flg_bonus                 :1;
	unsigned int flg_interest              :1;
	unsigned int flg_agm                   :1;
	unsigned int flg_egm                   :1;
	unsigned int flg_filler1               :1;
	unsigned int flg_exer_style            :1;
	unsigned int flg_ex_alwd               :1;
	unsigned int flg_ex_rjctn_alwd         :1;
	unsigned int flg_pl_alwd               :1;
	unsigned int flg_is_this_asset         :1;
	unsigned int flg_is_corporate_adjusted :1;
	unsigned int flg_filler2               :3;
};

struct st_sec_info
{
	char      c_instrument_name [ LEN_INSTRUMENT_NAME ];
	char      c_symbol [ LEN_SYMBOL_NSE ];
	char			c_series [ LEN_SERIES_NSE ];
	long int  li_expiry_date;
	long int  li_strike_price;
	char      c_option_type [ LEN_OPTION_TYPE ];
	short int si_ca_level;
};

struct st_sec_status_per_market
{
	short int si_stts;
};

struct st_token_and_elgibility 
{
	long 			l_token;				/**V1.4**/
	struct    st_sec_status_per_market st_sec_stts[ MAX_NO_SEC ];
};

struct st_ticker_index_info
{
	long 			l_token;				/**V1.4**/
	short int si_market_type;
	long int  li_fill_price;
	long int  li_fill_volume;
	long int  li_open_interest;
	long int  li_high_oi;
	long int  li_low_oi;
};

struct st_mbo_mbp_terms
{
	unsigned int flg_mf       :1;
	unsigned int flg_aon      :1;
	unsigned int flg_filler1  :6;
	unsigned int flg_filler2  :8;
};

struct st_mbo_info
{
/*short int si_trader_id;			Commented In Ver 1.8	***/
	long	int	li_trader_id;			/***	Ver 1.8	Data Type Changed From Short Int To Long	***/
	long int  li_qty;
	long int  li_price;
	struct    st_mbo_mbp_terms st_terms;
	long int  li_min_fill_qty;
};

struct st_mbp_info
{
  long int li_qty;
  long int li_price;
  short int si_no_of_ordrs;
};

struct st_interactive_mbo_data
{
	long 			l_token;			/**V1.4**/
	short int si_book_type;
	short int si_trading_status;
	long int  li_volume_traded_today;
	long int  li_last_traded_price;
	char      c_net_change_indicator;
	long int  li_net_change_from_closing_price;
	long int  li_last_traded_quantity;
	long int  li_last_traded_time;
	long int  li_average_trade_price;
	short int si_auction_number;
	short int si_auction_status;
	short int si_initiator_type;
	long int  li_initiator_price;
	long int  li_initiator_quantity;
	long int  li_auction_price;
	long int  li_auction_quantity;
	struct    st_mbo_info st_mboinfo [ MAX_MBO ];
};

struct st_indicator
{
	unsigned int flg_lst_trd_more  :1;
	unsigned int flg_lst_trd_less  :1;
	unsigned int flg_buy           :1;
	unsigned int flg_sell          :1;
	unsigned int flg_filler1       :4;
	unsigned int flg_filler2       :8;
};

/****Ver 1.1 for NNF 4.9.3****/
struct st_mbp_information
{
	long int	li_quantity;
	long int	li_price;
	short int	si_number_of_orders;
	short int	si_bb_buy_sell_flag;
};

struct st_interactive_only_mbp_data
{
	long 			l_token;			/**V1.4**/
	short int si_book_type;
	short int si_trading_status;
	long int  li_volume_traded_today;
	long int  li_last_traded_price;
	char      c_net_change_indicator;
	long int  li_net_change_from_closing_price;
	long int  li_last_traded_quantity;
	long int  li_last_traded_time;
	long int  li_average_trade_price;
	short int si_auction_number;
	short int si_auction_status;
	short int si_initiator_type;
	long int  li_initiator_price;
	long int  li_initiator_quantity;
	long int  li_auction_price;
	long int  li_auction_quantity;
	struct    st_mbp_information st_mbpinfo [ MAX_MBO ];
	short int	si_bb_total_buy_flag;
	short int	si_bb_total_sell_flag;
	double		d_total_buy_quantity;
	double		d_total_sell_quantity;
	struct		st_indicator st_ind;
	long int	li_closing_price;
	long int	li_open_price;
	long int	li_high_price;
	long int	li_low_price;
};

struct st_oi_data
{
	long l_token_no;				/**V1.4**/
	long int li_current_oi;
};

/****End of Ver 1.1 for NNF 4.9.3****/

/*** Ver 3.0 starts here  ***/

struct st_message_header
{
  char       c_i_api_tcode;
  char       c_i_api_funcld;
  long int   li_log_time;
  char       c_alpha_char[2];
  short int  si_trasaction_code;
  short int  si_error_code;
  char       c_timestamp[8];
  char       c_timestamp1[8];
  char       c_timestamp2[8];
  short int  si_message_length;
};

struct st_spd_eligibility
{
  unsigned int flg_eligibility  :1;
  unsigned int flg_filler       :7;
};

/*** Ver 3.0 ends here ***/

struct st_mkt_wise_info
{
	struct    st_indicator st_ind;
	long int  li_buy_volume;
	long int  li_buy_price;
	long int  li_sell_volume;
	long int  li_sell_price;
	long int  li_last_trade_price;
	long int  li_last_trade_time;
};

struct st_market_watch_bcast
{
	long 			l_token;					/**V1.4**/
	struct    st_mkt_wise_info st_mkt_info [ MAX_MKT_INFO ];
	long int li_open_interest;
};

struct st_filler_1
{
	unsigned int is_filler_1 : 1;
	unsigned int is_filler_2 : 1;
	unsigned int is_filler_3 : 1;
	unsigned int is_filler_4 : 1;
};

struct st_indices
{
	char      c_index_name [ LEN_INDEX_NAME ];
	long int  li_index_value;
	long int  li_high_index_value;
	long int  li_low_index_value;
	long int  li_opening_index;
	long int  li_closing_index;
	long int  li_percent_change;
	long int  li_yearly_high;
	long int  li_yearly_low;
	long int  li_no_of_upmoves;
	long int  li_no_of_downmoves;
	double    d_market_capitalisation;
	char      c_net_change_indicator;
  char      c_filter  ;     														/*****changed by sangeet******/
};

struct st_industry_indices
{
	char			c_industry_name [ LEN_INDUSTRY_NAME ];
	long int	li_iindex_value;
}; 

struct st_mbp_buys
{
	short int		si_no_orders;
  long int		li_volume;
	long int		li_price;
};

struct st_mbp_sells
{
	short int   si_no_orders;
  long int    li_volume;
  long int    li_price;
};

struct st_total_order_volume
{
	double	d_buy;
	double	d_sell;
};

struct st_bcast_data 
{
	long				l_itoken_1;					/**V1.4**/
	long				l_itoken_2;					/**V1.4**/
	short int		si_mbp_buy;
	short int		si_mbp_sell;
	long int		li_last_active_time;
	long int		li_traded_volume;
	double			d_total_traded_value;
	struct			st_mbp_buys st_mbpbuys [ MAX_MBP_BUYS ];
	struct			st_mbp_sells st_mbpsells [ MAX_MBP_SELLS ];
	struct			st_total_order_volume st_totalordrvol;
	long int    li_open_price_difference;                 /** Ver 3.0 **/
  long int    li_day_high_price_difference;             /** Ver 3.0 **/
  long int    li_day_low_price_difference;              /** Ver 3.0 **/
  long int    li_last_traded_price_difference;          /** Ver 3.0 **/
	long int		li_last_updated_time;
};	

/** 11 **/

struct st_ex_pl_info
{
  long 	l_token;						/**V1.4**/
	struct st_contract_desc st_cntrct_desc;
  short	int si_expl_flag;
  double    d_expl_number;
  short int si_market_type;
  char      c_account_number [ LEN_ACCOUNT_NUMBER ];
  long	int  li_quantity;
  short int si_pro_cli;
  short int si_exercise_type;
  long 	int  li_entry_date_tm;
  short int si_branch_id;
/*short int si_trader_id;		Commented In Ver 1.8	***/
	long	int	li_trader_id;			/***	Ver 1.8 Data Type Chaged From Short Int To Long	***/
  char      c_broker_id [ LEN_BROKER_ID ];
  char      c_remarks [ LEN_REMARKS_1 ];
  char      c_participant [ LEN_SETTLOR ];
};

/********************************************************************/
/* 1. Header                                                        */
/********************************************************************/
/** 1.1 **/
/*** Commented in Ver 2.9 ***
struct st_int_header 
{
*char      c_filler_1 [ 2 ];		**	Commented In Ver	1.8	***
	char      c_filler_1 [ 4 ];			***	Ver	1.8 Size Increased From 2 To 4	***
	long int  li_log_time;
	char      c_alpha_char [ LEN_ALPHA_CHAR ];
	short int si_transaction_code;
	short int si_error_code;
	char      c_filler_2 [ 8 ];
	char      c_time_stamp_1 [ LEN_TIME_STAMP ];
	char      c_time_stamp_2 [ LEN_TIME_STAMP ];
	short int si_message_length;
};
	*** Ver 2.9 Comment Ends ***/

/*** Ver 2.9 Starts ***/
struct st_int_header
{
	short int si_transaction_code;
	long int  li_log_time;
	char      c_alpha_char [ LEN_ALPHA_CHAR ];
	long int  li_trader_id;
	short int si_error_code;
	char      c_filler_2 [ 8 ];
	char      c_time_stamp_1 [ LEN_TIME_STAMP ];
	char      c_time_stamp_2 [ LEN_TIME_STAMP ];
	short int si_message_length;
};

struct st_inner_header
{
  char      c_filler_1 [ 4 ];
  long int  li_log_time;
  char      c_alpha_char [ LEN_ALPHA_CHAR ];
  short int si_transaction_code;
  short int si_error_code;
  char      c_filler_2 [ 8 ];
  char      c_time_stamp_1 [ LEN_TIME_STAMP ];
  char      c_time_stamp_2 [ LEN_TIME_STAMP ];
  short int si_message_length;
};

/*** Ver 3.1 Starts ***/
struct st_cmn_header_brd
{
  char      c_filler_1 [ 4 ];
  long int  li_log_time;
  char      c_alpha_char [ LEN_ALPHA_CHAR ];
  short int si_transaction_code;
  short int si_error_code;
  char      c_filler_2 [ 8 ];
  char      c_time_stamp_1 [ LEN_TIME_STAMP ];
  char      c_time_stamp_2 [ LEN_TIME_STAMP ];
  short int si_message_length;
};
/*** Ver 3.1 Ends ***/

struct st_cmn_header_tr
{
	short int si_transaction_code;
};
/*** Ver 2.9 Ends ***/
/** 1.2 **/
struct st_bcast_header
{
	char      c_filler_1 [ 2 ];
	char      c_filler_4 [ 2 ];			/***	Ver	1.8	***/
	long int  li_log_time;
	char      c_alpha_char [ LEN_ALPHA_CHAR ];
	short int si_trans_code;
	short int si_error_code;
	long int  li_bc_seq_no;
	char      c_crc;
	char      c_filler_2 [ 3 ];
	char      c_time_stamp_2 [ LEN_TIME_STAMP ];
	char      c_filler_3 [ 8 ];
	short int si_message_length;
};

/** 1.3 **/
struct st_error_response
{
	struct st_int_header st_hdr;
	char   c_key [ LEN_ERROR_KEY ];									/**V1.4**/
	char   c_error_message [ LEN_ERROR_MESSAGE ];
};

/*** Added for local use not in NNF ***/
/***	struct st_cmn_header	*** Commented in Ver 2.9 ***
{
*char      c_filler_1 [ 2 ];			Coomented In Ver	1.8	***
	char      c_filler_1 [ 4 ];
	long int  li_log_time;
	char      c_alpha_char [ LEN_ALPHA_CHAR ];
	short int si_transaction_code;
	short int si_error_code;
};	*** Ver 2.9 ***/

/*** Ver 2.9 Starts ***/
struct st_cmn_header
{
	short int si_transaction_code;
  long int  li_log_time;
  char      c_alpha_char [ LEN_ALPHA_CHAR ];
  long int  li_trader_id;
  short int si_error_code;
};
/*** Ver 2.9 Ends ***/

struct st_cmp_header
{
	short int si_reserved;
	short int si_comp_len;
};
/********************************************************************/
/*  2. The logon process                                            */
/********************************************************************/
/** 2.1 **/
struct st_sign_on_req
{
	struct    st_int_header st_hdr;
/*short int si_user_id;				Commented In Ver	1.8	***/
	long	int	li_user_id;										/***	Ver	1.8 Data Type Changed From Short Int To Long	***/
	char      c_reserved_1 [ 8 ];     /*** Ver 2.7 ***/
	char      c_password [ LEN_PASSWORD ];
	char      c_reserved_2 [ 8 ];     /*** Ver 2.7 ***/
	char      c_new_password [ LEN_PASSWORD ];
	char      c_trader_name [ LEN_TRADER_NAME ];
	long      li_last_password_change_date;
	char      c_broker_id [ LEN_BROKER_ID ];
	char      c_filler_1;
	short int si_branch_id;
	long int  li_version_number;
	long int  li_batch_2_start_time;
	char      c_host_switch_context;
	char      c_colour [ LEN_COLOUR ];
	char      c_filler_2;
	short int si_user_type;
	double    d_sequence_number;
	char      c_ws_class_name [ LEN_WS_CLASS_NAME ];
	char      c_broker_status;
	char      c_show_index;
	struct    st_broker_eligibility_per_mkt st_mkt_allwd_lst;
	short int si_member_type;
	char      c_clearing_status;
	char			c_broker_name [ LEN_BROKER_NAME ];   /* ver 1.2 */
/******Commented in Ver 1.2
	char      c_filler_3; */
	char      c_reserved_3 [ 16 ];     /*** Ver 2.7 ***/
	char      c_reserved_4 [ 16 ];     /*** Ver 2.7 ***/
	char      c_reserved_5 [ 16 ];     /*** Ver 2.7 ***/
};

/** 2.3 **/
struct st_sign_on_res
{
	struct st_int_header st_hdr;
/*short int si_user_id;			Commented In Ver 1.8	***/
	long	int	li_user_id;			/***  Ver 1.8 Data Type Changed From Short Int To Long  ***/
	char      c_reserved_1 [ 8 ];     /*** Ver 2.7 ***/
	char      c_password [ LEN_PASSWORD ];
	char      c_reserved_2 [ 8 ];     /*** Ver 2.7 ***/
	char      c_new_password [ LEN_PASSWORD ];
	char      c_trader_name [ LEN_TRADER_NAME ];
	long      li_last_password_change_date;
	char      c_broker_id [ LEN_BROKER_ID ];
	char      c_filler_1;
	short int si_branch_id;
	long int  li_version_number;
	long int  li_end_time;
	char      c_filler_2 [ 52 ];
	short int si_user_type;
	double    d_sequence_number;
	char      c_filler_3 [ 14 ];
	char      c_broker_status;
	char      c_show_index;
	struct    st_broker_eligibility_per_mkt st_mkt_allwd_lst;
	short int si_member_type;
	char      c_clearing_status;
	char			c_broker_name [ LEN_BROKER_NAME ]; /* ver 1.2 */
	char      c_reserved_3 [ 16 ];     /*** Ver 2.7 ***/
	char      c_reserved_4 [ 16 ];     /*** Ver 2.7 ***/
	char      c_reserved_5 [ 16 ];     /*** Ver 2.7 ***/
};

/** 2.5 **/
struct st_sign_off_req
{
	struct st_int_header st_hdr;
};

/** 2.6 **/
struct st_sign_off_res
{
/*short int si_user_id;					Comented In Ver 1.8	***/
	long	int	li_user_id;					/***	Ver 1.8 Data Type Changed From Short Int To Long  ***/
	char      c_filler_1 [ 145 ];
};

/********************************************************************/
/* 3. The download process                                          */
/********************************************************************/
/** 3.1.1 **/
struct st_system_info_req
{
	struct   st_int_header st_hdr;
	long int li_last_update_protfolio_time;
};

/** 3.1.2 **/
struct st_system_info_data
{
	struct    st_int_header st_hdr; 
	struct    st_market_status st_mkt_stts;
	struct    st_ex_market_status st_ex_mkt_stts;
	struct    st_pl_market_status st_pl_mkt_stts;
	char      c_update_portfolio;
	long int  li_market_index;
	short int si_default_sttlmnt_period_nm;
	short int si_default_sttlmnt_period_sp;
	short int si_default_sttlmnt_period_au;
	short int si_competitor_period;
	short int si_solicitor_period;
	short int si_warning_percent;
	short int si_volume_freeze_percent;
	char      c_filler_1 [ 4 ];
	long int  li_board_lot_quantity;
	long int  li_tick_size;
	short int si_maximum_gtc_days;
	struct    st_stock_eligible_indicators st_stk_elg_ind;
	short int si_disclosed_quantity_percent_allowed;
	long int  li_risk_free_interest_rate;
};

/** 3.2.1 **/
struct st_update_local_database
{
	struct   st_int_header st_hdr;
	long int li_last_update_security_time;
	long int li_last_update_participant_time;
	long int li_last_update_instrument_time;
	long int li_last_update_index_time;
	char     c_request_for_open_orders;
	char     c_filler_1;
	struct   st_market_status st_mkt_stts;
	struct   st_ex_market_status st_ex_mkt_stts;
	struct   st_pl_market_status st_pl_mkt_stts;
};

/** 3.2.3 **/
/* Same as 3.1.2 **/

/** 3.2.5 **/
struct st_update_ldb_header
{
	struct   st_int_header st_hdr;
	char 		 c_filler_1 [ 2 ];
};

/** 3.2.6 **/
/* Declared in the end as it involves union */

/** 3.2.7 **/
struct st_download_index
{
  /**struct st_int_header st_hdr; *** Ver 2.9 **/
  struct  st_inner_header st_hdr; /*** Ver 2.9 **/	
	short si_no_recs;
	struct st_index_dtls index_dtls;
};

/** 3.2.8 **/
struct st_ms_download_bp_tok
{
  /**struct st_int_header st_hdr; *** Ver 2.9 **/
  struct  st_inner_header st_hdr; /*** Ver 2.9 **/	
	short si_no_recs;
	struct st_download_bp_tok sbasep[MAX_BP_DOWNLOAD];
};

/** 3.2.9 **/
struct st_ms_download_index_map
{
  /**struct st_int_header st_hdr; *** Ver 2.9 **/
  struct  st_inner_header st_hdr; /*** Ver 2.9 **/	
	short si_no_recs;
	struct st_bcast_index_map_dtls indices_map[MAX_INDEX_MAP];
};

/** 3.2.10 **/
/* Same as 3.2.5 */

/** 3.2.11.1 **/
/* To_be_typed */

/** 3.2.11.4 **/
/* To_be_typed */

/** 3.3.1 **/
struct st_message_download
{
	struct st_int_header st_hdr;
	double d_sequence_number;
};

/** 3.3.3 **/
/* Same as 1.1 */

/** 3.3.4 **/
/* Declared in the end as it involves union */

/** 3.3.5 **/
/* Same as 1.1 */

/********************************************************************/
/* 4. Order & Trade management                                      */
/********************************************************************/
/** 4.3.1 **/
struct st_oe_reqres
{
	struct    st_int_header st_hdr;
	char      c_participant_type;
	char      c_filler_1;
	short int si_competitor_period;
	short int si_solicitor_period;
	char      c_modified_cancelled_by;
	char      c_filler_2;
	short int si_reason_code;
	char      c_filler_3 [ 4 ];
	long 			l_token_no;											/**V1.4**/
	struct    st_contract_desc st_con_desc;
	char      c_counter_party_broker_id [ LEN_BROKER_ID ];
	char      c_filler_4 ;										/**Ver 1.5**/
	char      c_filler_5 [ 2 ];								/**Ver 1.5**/
	char      c_closeout_flg;									/**Ver 1.5**/
	char      c_filler_6;											/**Ver 1.5**/
	short int si_order_type;
	double    d_order_number;
	char      c_account_number [ LEN_ACCOUNT_NUMBER ];
	short int si_book_type;
	short int si_buy_sell_indicator;
	long int  li_disclosed_volume;
	long int  li_disclosed_volume_remaining;
	long int  li_total_volume_remaining;
	long int  li_volume;
	long int  li_volume_filled_today;
	long int  li_price;
	long int  li_trigger_price;
	long int  li_good_till_date;
	long int  li_entry_date_time;
	long int  li_minimum_fill_aon_volume;
	long int  li_last_modified;
	struct    st_order_flags st_ord_flg;
	short int si_branch_id;
/*short int si_trader_id;			Commented In Ver 1.8	***/
	long int	li_trader_id;			/***	Ver	1.8	Data Type Changed From Short Int To Long	***/
	char      c_broker_id [ LEN_BROKER_ID ];
	char      c_remarks [ LEN_REMARKS ];
	char      c_open_close;
	char      c_settlor [ LEN_SETTLOR ];
	short int si_pro_client_indicator;
	short int si_settlement_period;
	char      c_cover_uncover;
	char      c_giveup_flag;
	/* Ver 1.3 */
	/*Commented for Ver 1.6
	unsigned int us_filler1			:1;
	unsigned int us_filler2			:1;
	unsigned int us_filler3			:1;
	unsigned int us_filler4			:1;
	unsigned int us_filler5			:1;
	unsigned int us_filler6			:1;
	unsigned int us_filler7			:1;
	unsigned int us_filler8			:1;
	unsigned int us_filler9			:1;
	unsigned int us_filler10		:1;
	unsigned int us_filler11		:1;
	unsigned int us_filler12		:1;
	unsigned int us_filler13		:1;
	unsigned int us_filler14		:1;
	unsigned int us_filler15		:1;
	unsigned int us_filler16		:1;
	char c_filler17;
	char c_filler18; 
	*****Comment ends***/
	/**int i_ordr_rfrnc;									**Ver 1.6 added to retain Async Mech**/
	int i_ordr_sqnc;											/**Ver 1.7 name changed**/
	double	d_nnf_field;
	double	d_filler19;
	/* till here ...ver 1.3 */
	char c_pan [ 10 ];  /*** Ver 2.7 ***/
	long l_algo_id;     /*** Ver 2.7 ***/
	short si_algo_category;   /*** Ver 2.7 ***/
	/*** char c_reserved [ 60 ];   *** Ver 2.7 *** *** commented in Ver 2.9 ***/
  long long ll_lastactivityref;  	/*** Ver 2.9 ***/
  char c_reserved[52];           	/*** Ver 2.9 ***/
};

/** 4.3.4 **/
/* Same as 4.3.1 */

/** 4.3.5 **/
/* Same as 4.3.1 */

/** 4.3.6 **/
/* Same as 4.3.1 */

/** 4.3.7 **/
/* Same as 4.3.1 */

/** 4.3.8 **/
/* Same as 4.3.1 */

/** 4.3.4 **/
/* Same as 4.3.1 */

/** 4.4.1 **/
/* Same as 4.3.1 */

/** 4.4.3 **/
/* Same as 4.3.1 */

/** 4.4.4 **/
/* Same as 4.3.1 */

/** 4.4.5 **/
/* Same as 4.3.1 */

/** 4.5.1 **/
/* Same as 4.3.1 */

/** 4.5.2 **/
/* Same as 4.3.1 */

/** 4.5.3 **/
/* Same as 4.3.1 */

/** 4.5.4 **/
/* Same as 4.3.1 */

/** 4.5.7 **/
struct st_trade_inq_reqres
{
	struct    st_int_header st_hdr;
	long 			l_token_no;												/**V1.4**/
	struct    st_contract_desc st_con_desc;
	long int  li_fill_number;
	long int  li_fill_quantity;
	long int  li_fill_price;
	char      c_mkt_type;
	char      c_buy_open_close;
	long int  li_new_volume;
	char      c_buy_broker_id [ LEN_BROKER_ID ];
	char      c_sell_broker_id [ LEN_BROKER_ID ];
/*short int si_trader_id;			Commented In Ver 1.8	***/
	long	int	li_trader_id;			/***	Ver 1.8 Data Type Changed From Short Int To Long	***/
	char      c_requested_by;
	char      c_sell_open_close;
	char      c_buy_account_number [ LEN_ACCOUNT_NUMBER ];
	char      c_sell_account_number [ LEN_ACCOUNT_NUMBER ];
	char      c_buy_participant [ LEN_PARTICIPANT_NAME ];
	char      c_sell_participant [ LEN_PARTICIPANT_NAME ];
	char      c_buy_cover_uncover;
	char      c_sell_cover_uncover;
	char      c_buy_giveup_flag;
	char      c_sell_giveup_flag;
	char      c_buypan [ 10 ];    /*** Ver 2.7 ***/
	char      c_sellpan [ 10 ]; /*** Ver 2.7 ***/
	char      c_reserved [ 60 ];  /*** Ver 2.7 ***/
};

/** 4.5.8 **/
/* Same as 4.5.7 */

/** 4.5.9 **/
/* Same as 4.5.7 */

/** 4.6.1 **/
/* Same as 4.5.7 */

/** 4.6.2 **/
/* Same as 4.5.7 */

/** 4.6.4 **/
/* Same as 4.5.7 */

/**** Ver 2.9 Starts ****/
# pragma pack (  )
# pragma pack ( 1 )
struct st_addtnal_order_flags
{
  char c_cover_uncover;
};
# pragma pack (  )
# pragma pack ( 2 )
struct st_contract_desc_tr
{
  char      c_instrument_name [ LEN_INSTRUMENT_NAME ];
  char      c_symbol [ LEN_SYMBOL_NSE ];
  long int  li_expiry_date;
  long int  li_strike_price;
  char      c_option_type [ LEN_OPTION_TYPE ];
};

struct st_oe_reqres_tr
{
short int si_transaction_code;
long int li_userid;
short int si_reason_code;
long l_token_no;
struct  st_contract_desc_tr st_con_desc_tr;
char  c_account_number [ LEN_ACCOUNT_NUMBER ];
short int si_book_type;
short int si_buy_sell_indicator;
long int  li_disclosed_volume;
long int  li_volume;
long int  li_price;
long int  li_good_till_date;
struct  st_order_flags st_ord_flg;
short int si_branch_id;
long int  li_trader_id;
char      c_broker_id [ LEN_BROKER_ID ];
char      c_open_close;
char      c_settlor [ LEN_SETTLOR ];
short int si_pro_client_indicator;
struct  st_addtnal_order_flags st_add_ord_flg;
long int  li_ordr_sqnc;   /*** li_filler ***/
double  d_nnf_field;
char c_pan [ 10 ];
long l_algo_id;
short si_algo_category;
char c_reserved [ 32 ];
};

struct st_om_rqst_tr
{
short int si_transaction_code;
long int li_userid;
char      c_modified_cancelled_by;
long l_token_no;
struct  st_contract_desc_tr st_con_desc_tr;
double    d_order_number;
char  c_account_number [ LEN_ACCOUNT_NUMBER ];
short int si_book_type;
short int si_buy_sell_indicator;
long int  li_disclosed_volume;
long int  li_disclosed_volume_remaining;
long int  li_total_volume_remaining;
long int  li_volume;
long int  li_volume_filled_today;
long int  li_price;
long int  li_good_till_date;
long int  li_entry_date_time;
long int  li_last_modified;
struct    st_order_flags st_ord_flg;
short int si_branch_id;
long int  li_trader_id;
char      c_broker_id [ LEN_BROKER_ID ];
char      c_open_close;
char      c_settlor [ LEN_SETTLOR ];
short int si_pro_client_indicator;
struct  st_addtnal_order_flags st_add_ord_flg;
long int  li_ordr_sqnc;
double  d_nnf_field;
char c_pan [ 10 ];
long l_algo_id;
short si_algo_category;
long long ll_lastactivityref;
char c_reserved[24];
};


struct st_oe_rspn_tr
{
short int si_transaction_code;
long int  li_log_time;
long int li_userid;
short int si_error_code;
char      c_time_stamp_1 [ LEN_TIME_STAMP ] ;
char      c_time_stamp_2 ;
char      c_modified_cancelled_by;
short int si_reason_code;
long l_token_no;
struct  st_contract_desc_tr st_con_desc_tr;
char      c_closeout_flg;
double    d_order_number;
char  c_account_number [ LEN_ACCOUNT_NUMBER ];
short int si_book_type;
short int si_buy_sell_indicator;
long int  li_disclosed_volume;
long int  li_disclosed_volume_remaining;
long int  li_total_volume_remaining;
long int  li_volume;
long int  li_volume_filled_today;
long int  li_price;
long int  li_good_till_date;
long int  li_entry_date_time;
long int  li_last_modified;
struct    st_order_flags st_ord_flg;
short int si_branch_id;
long int  li_trader_id;
char      c_broker_id [ LEN_BROKER_ID ];
char      c_open_close;
char      c_settlor [ LEN_SETTLOR ];
short int si_pro_client_indicator;
struct  st_addtnal_order_flags st_add_ord_flg;
long int  li_ordr_sqnc;
double  d_nnf_field;
long long l_timestamp;
char c_pan [ 10 ];
long l_algo_id;
short si_algo_category;
long long ll_lastactivityref;
char c_reserved[52];
};

struct st_trade_confirm_tr
{
	short int si_transaction_code;
  long int  li_log_time;
  long  int li_trader_number;
  long long ll_timestamp;
  char c_timestamp1 [ LEN_TIME_STAMP ];
  double d_timestamp2;
  double    d_response_order_number;
  char      c_broker_id [ LEN_BROKER_ID ];
  char      c_reserved1;
  char      c_account_number [ LEN_ACCOUNT_NUMBER ];
  short int si_buy_sell_indicator;
  long int  li_original_volume;
  long int  li_disclosed_volume;
  long int  li_remaining_volume;
  long int  li_disclosed_volume_remaining;
  long int  li_price;
  struct    st_order_flags st_ord_flg;
  long int  li_good_till_date;
  long int  li_fill_number;
  long int  li_fill_quantity;
  long int  li_fill_price;
  long int  li_volume_filled_today;
  char      c_activity_type [ LEN_ACTIVITY_TYPE ];
  long int  li_activity_time;
  long      l_token;
  struct    st_contract_desc_tr st_con_desc_tr;
  char      c_open_close;
  char      c_book_type;
  char      c_participant [ LEN_PARTICIPANT_NAME ];
  struct  st_addtnal_order_flags st_add_ord_flg;
  char c_pan [ 10 ];
  long l_algo_id;
  short si_algo_category;
  long long ll_lastactivityref;
  char c_reserved2[52];
};
/**** Ver 2.9 Ends ****/

/********************************************************************/
/* 5. Spread order & trade management                               */
/********************************************************************/
/** 5.2 **/
struct st_spd_oe_reqres
{
	struct    st_int_header st_hdr;
	char      c_participant_type;
	char      c_filler1;
  short int si_competitor_period;
  short int si_solicitor_period;  
  char      c_mod_cxl_by;
	char      c_filler2;
  short int si_reason_code;
  char 			c_start_alpha[2];
	char			c_end_alpha[2];
	long     l_token;				/**V1.4**/
	struct st_contract_desc st_cntrct_desc;
	char      c_op_broker_id[LEN_BROKER_ID];
	char      c_filler3;
	char 			c_filler_options[LEN_FILLER_OPTIONS];
	char 			c_filler4;
	short int si_order_type;
	double    d_order_number;
	char 			c_account_number[LEN_ACCOUNT_NUMBER];
	short int si_book_type;
	short int si_buy_sell;
	long int  li_disclosed_vol;
	long int  li_disclosed_vol_remaining;
	long int  li_total_vol_remaining;
	long int  li_volume;
	long int  li_volume_filled_today;	
	long int  li_price;
	long int  li_trigger_price;	
	long int  li_good_till_date;
	long int  li_entry_date_time;
	long int  li_min_fill_aon;
	long int  li_last_modified;
	struct st_order_flags st_order_flgs;
	short int si_branch_id;	
/*short int si_trader_id;			Commented In Ver 1.8	***/
	long	int	li_trader_id;			/***	Ver 1.8 Data Type Changed From Short Int To Long	***/
	char      c_broker_id[LEN_BROKER_ID];
	char      c_oe_remarks[LEN_REMARKS];
	char      c_open_close;
	char      c_settlor[LEN_SETTLOR];
	short int si_pro_client;
	short int si_settlement_period;
	char      c_cover_uncover;
	char 			c_give_up_flag;
  /* Ver 1.3 */
	/*** Ver 2.0 Coment Starts ****
  unsigned int us_filler1     :1;
  unsigned int us_filler2     :1;
  unsigned int us_filler3     :1;
  unsigned int us_filler4     :1;
  unsigned int us_filler5     :1;
  unsigned int us_filler6     :1;
  unsigned int us_filler7     :1;
  unsigned int us_filler8     :1;
  unsigned int us_filler9     :1;
  unsigned int us_filler10    :1;
  unsigned int us_filler11    :1;
  unsigned int us_filler12    :1;
  unsigned int us_filler13    :1;
  unsigned int us_filler14    :1;
  unsigned int us_filler15    :1;
  unsigned int us_filler16    :1;
  char c_filler17;
  char c_filler18;							*** Ver 2.0 Comment Ends **/
	int	i_sprd_seq_no;			/*** Ver 2.0 ***/
  double  d_nnf_field;					
  double  d_filler19;
  /* till here ...ver 1.3 */
	char c_pan [ 10 ];  /*** Ver 2.7 ***/
	long l_algo_id;     /*** Ver 2.7 ***/
	short si_algo_category;   /*** Ver 2.7 ***/
/***	char c_reserved [ 60 ];   *** Ver 2.7 ***	*** Commented in Ver 2.9 ***/
	long long ll_lastactivityref;   /*** Ver 2.9 ***/
  char c_reserved[52];            /*** Ver 2.9 ***/
	long int 	li_spd_price_diff;
	struct st_spd_leg_info st_spd_lg_inf [ 2 ];
};

/** 5.4.1 **/
/* Same as 5.2 */

/** 5.4.2 **/
/* Same as 5.2 */

/** 5.4.3 **/
/* Same as 5.2 */

/** 5.4.4 **/
/* Same as 5.2 */

/** 5.5.1 **/
/* Same as 5.2 */

/** 5.5.3 **/
/* Same as 5.2 */

/** 5.5.4 **/
/* Same as 5.2 */

/** 5.5.5 **/
/* Same as 5.2 */

/********************************************************************/
/* 6. 2L and 3L order entry                                         */
/********************************************************************/
/** 6.1.2 **/
/* Same as 5.2 */

/** 6.3.1 **/
/* Same as 5.2 */

/** 6.3.2 **/
/* Same as 5.2 */

/** 6.3.3 **/
/* Same as 5.2 */

/** 6.3.4 **/
/* Same as 5.2 */

/** 6.3.5 **/
/* Same as 5.2 */

/********************************************************************/
/* 7. Unsolicited messages                                          */
/********************************************************************/
/** 7.1 **/
/* Same as 4.3.1 */

/** 7.2 **/
/* Same as 4.3.1 */

/** 7.3 **/
/* Same as 4.3.1 */

/** 7.4 **/
/* Same as 4.3.1 */

/** 7.5 **/
struct st_trade_confirm
{
	struct    st_int_header st_hdr;
	double    d_response_order_number;
	char      c_broker_id [ LEN_BROKER_ID ];
	char      c_filler_1;
/*short int si_trader_number;			Commented In Ver 1.8	***/
	long	int	li_trader_number;			/***	Ver	1.8 Data Type Changed From Short Int To Long	***/
	char      c_account_number [ LEN_ACCOUNT_NUMBER ];
	short int si_buy_sell_indicator;
	long int  li_original_volume;
	long int  li_disclosed_volume;
	long int  li_remaining_volume;
	long int  li_disclosed_volume_remaining;
	long int  li_price;
	struct    st_order_flags st_ord_flg;
	long int  li_good_till_date;
	long int  li_fill_number;
	long int  li_fill_quantity;
	long int  li_fill_price;
	long int  li_volume_filled_today;
  char      c_activity_type [ LEN_ACTIVITY_TYPE ];
	long int  li_activity_time;
	double    d_counter_trade_order_number;
	char      c_counter_broker_id [ LEN_BROKER_ID ];
	long 			l_token;															/**V1.4**/
	struct    st_contract_desc st_con_desc;
	char      c_open_close;
	char      c_old_open_close;
	char      c_book_type;
	long int  li_new_volume;
	char      c_old_account_number [ LEN_ACCOUNT_NUMBER ];
	char      c_participant [ LEN_PARTICIPANT_NAME ];
	char      c_old_participant [ LEN_PARTICIPANT_NAME ];
	char      c_cover_uncover;
	char      c_old_cover_uncover;
	char      c_giveup_trade;
	char			c_reserved_filler2;	/*** Ver 2.7 ***/
	char c_pan [ 10 ];  /*** Ver 2.7 ***/
  char c_old_pan [ 10 ];  /*** Ver 2.7 ***/
  long l_algo_id;     /*** Ver 2.7 ***/
  short si_algo_category;   /*** Ver 2.7 ***/
/***  char c_reserved [ 60 ];   *** Ver 2.7 ***	Commented in Ver 2.9 ***/
	long long ll_lastactivityref;   /*** Ver 2.9 ***/
  char c_reserved[52];            /*** Ver 2.9 ***/
};

/** 7.6 **/
struct st_trader_int_msg
{
	struct    st_int_header st_hdr;
/*short int si_trader_id;			Commented In Ver 1.8	***/
	long	int	li_trader_id;			/***	Ver	1.8	Data Type Changed From Short Int To Long	***/
	char      c_filler_1 [ 3 ];
	char      c_filler_2;
	short int si_broadcast_message_length;
	char      c_broadcast_message [ LEN_BROADCAST_MESSAGE ];
};

/** 7.7 **/
/* Same as 7.5 */

/** 7.8 **/
/* Same as 7.5 */

/** 7.9 **/
/* Same as 7.6 */

/** 7.10 **/
/* Same as 7.5 */

/** 7.11 **/
/* Same as 7.5 */

/** 7.12 **/
/* Same as 7.6 */

/********************************************************************/
/* 8. Bhav copy                                                     */
/********************************************************************/
/** 8.1 **/
/* Same as 10.1 */

/** 8.2 **/
struct st_rp_hdr
{
	struct 		st_int_header st_hdr;
	char	 		c_message_type;
	long int 	li_report_date;
	short int si_user_type;
	char			c_broker_id [ LEN_BROKER_ID ];
	char 			c_firm_name [ LEN_FIRM_NAME ];
/*short int si_trader_number;			Commented In Ver	1.8	***/
	long	int	li_trader_number;			/***	Ver	1.8	Data Type Changed From Short Int To Long	***/
	char  		c_trader_name [ LEN_TRADER_NAME ];
};

/** 8.3 **/
struct st_rp_market_stats
{
	struct 		st_int_header st_hdr;
	char      c_message_type;
	char 			c_filler_1;
	short int si_number_of_records;
	struct 		st_stats_data	st_stats_dat [ MAX_RCRDS_MKT_STATS_DATAT ];
};

/** 8.4 **/
/* Same as 8.3 */

/** 8.5 **/
/* Same as 8.3 */

/** 8.6 **/
struct st_rp_trailer
{
	struct 		st_int_header st_hdr;
	char			c_message_type;
	long int	li_number_of_packets;
	char			c_filler_1;
};

/********************************************************************/
/* 9. Reports                                                       */
/********************************************************************/

/********************************************************************/
/* 10. Broadcast                                                    */
/********************************************************************/
/** 10.1 **/
struct st_bcast_vct_msgs
{
	struct st_bcast_header st_hdr;
	long 			l_token;				/**V1.4**/
	struct    st_sec_info st_sc_info;
	short int si_market_type;
	struct    st_bcast_destination st_bcst_dest;
	short int si_broadcast_message_length;
	char      c_broadcast_message [ LEN_BROADCAST_MESSAGE ];
};

/** 10.2 **/
/* Same as 3.1.2 */

/** 10.3 **/
struct st_security_update_info
{
	/**struct st_int_header st_hdr; *** Ver 2.9 **/
  struct  st_inner_header st_hdr; /*** Ver 2.9 **/  
	long 		l_token;						/**V1.4**/
	struct    st_sec_info st_sc_info;
	short int si_permited_trade;
	double    d_issued_capital;
	long int  li_warning_quantity;
	long int  li_freeze_quantity;
	char      c_credit_rating [ LEN_CREDIT_RATING ];
	struct    st_sec_eligibility_per_market st_sec_elg [ MAX_NO_SEC ];
	short int si_issue_rate;
	long int  li_issue_start_date;
	long int  li_interest_payment_date;
	long int  li_issue_maturity_date;
	long int  li_margin_percentage;
	long int  li_minimum_lot_quantity;
	long int  li_board_lot_quantity;
	long int  li_tick_size;
	char  		c_name [ LEN_SECURITY_NAME ];
	char      c_filler_1;
	long int  li_listing_date;
	long int  li_expulsion_date;
	long int  li_readmission_date;
	long int  li_record_date;
	long int  li_low_price_range;
	long int  li_high_price_range;
	long int  li_expiry_date;
	long int  li_nd_start_date;
	long int  li_nd_end_date;
	struct    st_eligibility_indicators st_elg_ind;
	long int  li_book_closure_start_date;
	long int  li_book_closure_end_date;
	long int  li_exercise_start_date;
	long int  li_exercise_end_date;
	long 			l_old_token;		/**V1.4**/
	char      c_asset_instrument [ LEN_INSTRUMENT_NAME ];
	char      c_asset_name [ LEN_ASSET_NAME ]; 
	long 			l_asset_token;				/**V1.4**/
	long int  li_intrinsic_value;
	long int  li_extrinsic_value;
	struct    st_purpose st_pps;
	long int  li_local_update_date_time;
	char      c_delete_flag;
	char      c_remarks [ LEN_REMARKS_2 ];
};

/** 10.4 **/
struct st_instrument_update_info
{
  /**struct st_int_header st_hdr; *** Ver 2.9 **/
  struct  st_inner_header st_hdr; /*** Ver 2.9 **/	
	short si_instrument_id;
	char c_instrument_name[LEN_INSTRUMENT_NAME];
	char c_instrument_desc[LEN_INSTRUMENT_DESC];
	long int li_instrument_upd_dt_tm;
	char c_delete_flag;
};

/** 10.5 **/
struct st_participant_update_info
{
  /**struct st_int_header st_hdr; *** Ver 2.9 **/
  struct  st_inner_header st_hdr; /*** Ver 2.9 **/	
	char      c_participant_id [ LEN_SETTLOR ];
	char      c_participant_name [ LEN_PARTICIPANT_NAME ];
	char      c_participant_status;
	long int  li_participant_update_date_time;
	char      c_delete_flag;
};

/** 10.6 **/
struct st_security_status_update_info
{
  /**struct st_int_header st_hdr; *** Ver 2.9 **/
  struct  st_inner_header st_hdr; /*** Ver 2.9 **/	
	short int si_number_of_records;
	struct    st_token_and_elgibility st_t_and_e [ MAX_TANDE ];
};

/** 10.7 **/
struct st_broadcast_tlimit_exceeded
{
	struct st_bcast_header st_hdr;
	char      c_broker_code [ LEN_BROKER_ID ];
	char      c_counter_broker_code [ LEN_BROKER_ID ];
	short int si_warning_type;
	char      c_symbol [ LEN_SYMBOL_NSE ];
	char      c_series [ LEN_SERIES_NSE ];
	long int  li_trade_number;
	long int  li_trade_price;
	long int  li_trade_volume;
	char      c_final;
};

/** 10.8 **/
struct st_bcast_message
{
	struct st_int_header st_hdr;
	short int si_branch_number;
	char      c_broker_number [ LEN_BROKER_ID ];
	char      c_action_code [ LEN_ACTION_CODE ];
/**	char      c_filler_1 [ 24 ]; **/
	struct    st_bcast_destination st_bcast_dest;
  char      c_filler_1 [ 26 ];  /** Ver 2.2 **/
	short int si_broadcast_message_length;
	char      c_broadcast_message [ LEN_BROADCAST_MESSAGE ];
};

/** 10.9 **/
struct st_ticker_trade_data
{
	struct st_bcast_header st_hdr;
	short int si_number_of_records;
	struct    st_ticker_index_info st_tck_idx_info [ MAX_TICKER_INDEX ];
};

/** 10.10 **/
struct st_bcast_mbo_mbp
{
/*  struct st_int_header st_hdr;  *** Commented in ver 3.1 ***/
  struct    st_cmn_header_brd st_hdr;   /*** Ver 3.1 ***/	
	struct    st_interactive_mbo_data st_mbo_data;
	struct    st_mbp_info st_mbp [ MAX_MBP ];
	double    d_total_buy_quantity;	
	double    d_total_sell_quantity;
	struct    st_indicator st_ind;
	long int  li_closing_price;
	long int  li_open_price;
	long int  li_high_price;
	long int  li_low_price;
};

/****Ver 1.1 for NNF 4.9.3****/
struct st_bcast_only_mbp
{
/*  struct st_int_header st_hdr;  *** Commented in ver 3.1 ***/
  struct    st_cmn_header_brd st_hdr;   /*** Ver 3.1 ***/
	short int	si_no_of_records;
	struct    st_interactive_only_mbp_data st_only_mbp[2];
};

/****End of Ver 1.1 for NNF 4.9.3****/

/*** Ver 3.0 starts here ***/

struct st_spd_update_info
{
  struct     st_message_header st_hdr;
  long int   li_token_1;
  long int   li_token_2;
  struct     st_sec_info  st_secinfo[2];
  long int   li_reference_price;
  long int   li_day_low_price_diff_range;
  long int   li_day_high_price_diff_range;
  long int   li_op_low_price_diff_range;
  long int   li_op_high_price_diff_range;
  struct     st_spd_eligibility  st_eligibile;
  char       c_flg_filler1;
  char       c_delete_flag;
  char       c_flg_filler2;
};

/*** Ver 3.0 ends here ***/

/** 10.11 **/
struct st_bcast_inq_resp_2
{
	struct st_bcast_header st_hdr;
	short int si_number_of_records;
	struct    st_market_watch_bcast st_mkt_wtch [ MAX_MKT_WTCH ];
};

/** 10.12 **/
struct st_sec_open_msg
{
	struct st_int_header st_hdr;
	char      c_symbol [ LEN_SYMBOL_NSE ];
	char      c_series [ LEN_SERIES_NSE ];
	long 			l_token;				/**V1.4**/
	long int  li_open_price;
	struct		st_filler_1 st_filler; 
};

/** 10.13 **/
/* Same as 1.1 **/

/** 10.14 **/
struct st_bcast_indices
{
	struct 		st_int_header st_hdr;
	short int si_number_of_records;
	struct    st_indices st_idx [ MAX_INDICES ];
};

/** 10.15 **/
struct st_bcast_industry_indices
{
/*  struct    st_int_header st_hdr;   *** Commented in Ver 3.1 ***/
  struct    st_cmn_header_brd st_hdr;   /*** Ver 3.1 ***/
	short int	si_inofrecs;
	struct		st_industry_indices st_industry [ MAX_INDUSTRY_INDICES ];
};

/** 10.16 **/
struct st_spd_mkt_info
{
/*  struct st_int_header st_hdr;  *** Commented in Ver 3.1 ***/
  struct    st_cmn_header_brd st_hdr;   /*** Ver 3.1 ***/
	struct st_bcast_data st_bcast_dat;
}; 

/** 10.17 **/
struct st_oi_info
{
/***   struct st_int_header st_hdr; *** Ver 3.1 **/
  struct    st_cmn_header_brd st_hdr;   /*** Ver 3.1 ***/
  struct st_oi_data st_oi_dat[MAX_EQ_OI];
};

/*** Ver 2.5 starts ***/
struct st_trd_exec_rng_dtls
{ 
	long int li_token_num;
	long int li_high_exec_band;
	long int li_low_exec_band; 
};

struct st_trd_exec_rng_data
{
	long	int li_msg_count;
	struct st_trd_exec_rng_dtls st_trd_exec_rng_dtl [25];
};

struct st_bcast_trd_exec_rng
{
	struct		st_bcast_header st_hdr;
	struct    st_trd_exec_rng_data st_trd_exec_rng_dat;
};
/*** Ver 2.5 ends ***/

/********************************************************************/
/* 11. Exercise and position liquidation                            */
/********************************************************************/
/** 11.4 **/
struct st_ex_pl_reqres
{
  struct    st_int_header st_hdr;
  short int si_reason_code;
  struct    st_ex_pl_info st_ex_pl_data;
};

/** 11.6 **/
/* Same as 11.4 */

/** 11.7 **/
/* Same as 11.4 */

/** 11.8 **/
/* Same as 11.4 */

/** 11.9 **/
/* Same as 11.4 */

/** 11.11 **/
/* Same as 11.4 */

/** 11.12 **/
/* Same as 11.4 */

/** 11.13.1 **/
/* Same as 11.4 */

/** 11.13.2 **/
/* Same as 11.4 */

/** 11.13.3 **/
/* Same as 11.4 */

/** 11.13.4 **/
/* Same as 11.4 */

/********************************************************************/
/* 12. Give up trades                                               */
/********************************************************************/


/********************************************************************/
/* Structures involving unions                                      */
/********************************************************************/
/** 3.2.6 **/
union un_ldb_data
{
	struct st_security_update_info st_sec_upd_info;
	struct st_security_status_update_info st_sec_stts_upd_info;
	struct st_participant_update_info st_part_upd_info;
	struct st_instrument_update_info st_instr_upd_info;
	struct st_download_index st_dwnld_indx;
	struct st_ms_download_bp_tok st_dwnld_bp_tok;
	struct st_ms_download_index_map st_dwnld_indx_map;
};

struct st_local_database_data
{
	struct st_int_header st_hdr;
	union un_ldb_data u_ldbmsg;
};

/** 3.3.4 **/
union u_dwld_msgs
{
	struct st_sign_on_res st_lgn_res;
	struct st_sign_off_res st_lgf_res;
	struct st_oe_reqres st_oe_res;
	struct st_oe_rspn_tr st_oe_res_tr;						/*** Ver 2.9 ***/
	struct st_trade_inq_reqres st_trd_inq_res;
	struct st_trade_confirm st_trd_conf;
	struct st_trade_confirm_tr st_trd_conf_tr;		/*** Ver 2.9 ***/
	struct st_trader_int_msg st_trd_msg;
  struct st_bcast_vct_msgs st_brd_vct_msg;
	struct st_broadcast_tlimit_exceeded st_brd_tlmt_exceed;
	struct st_bcast_message st_bcast_msg;
};

struct st_dwnld_msg
{
  struct st_int_header st_hdr;
  union u_dwld_msgs un_dw_msg;
};


/******  Ver 1.9 *********************/

struct st_indexdetails
{
 struct st_bcast_header st_hdr;
 long int li_token;
 char c_name [50];
 long int li_Open;
 long int li_High;
 long int li_Low;
 long int li_Last;
 long int li_Close;
 long int li_PrevClose;
 long int li_LifeHigh;
 long int li_LifeLow;
 long int li_filler1;
 long int li_filler2;
 long int li_filler3;
};


/**************************/

/******* Ver 2.3 **********/

struct st_india_vix
{
	char 			c_index_name [21];         
	long 	int	li_index_value;             
	long 	int li_high_index_value;         
	long 	int li_low_index_value;          
	long 	int li_opening_index;           
	long 	int	li_closing_index;           
	long 	int li_percent_change;          
	long 	int li_yearly_high;             
	long 	int li_yearly_low;              
	long 	int	li_no_of_upmoves;            
	long 	int	li_no_of_downmoves;          
	double 		d_market_capitalisation;           
	char 			c_net_change_indicator;     
	char    	c_filler;               
};

struct st_ind_vix_dtls
{
  struct st_bcast_header st_hdr;
  short int i_num_of_rec;
  struct st_india_vix st_indvix[6];
};

/******* Ver 2.3 **********/

/** Ver 2.1 Starts Here **/
struct st_eqoi_msg
{
/*** struct st_int_header st_hdr; *** Ver 3.1**/
struct st_cmn_header_brd st_hdr;  /*** Ver 3.1 ***/
struct st_oi_info st_eq_oi;
};
/** Ver 2.1 Ends Here **/

/*** Ver 2.6 Starts ***/
struct st_brd_exch_exception
{
  struct   st_bcast_header st_hdr;
  short    streamnumber;
  short    status;
  char     reserved[200];
};
/*** Ver 2.6 Ends***/

/** Union of all messages in the broadcast channel **/
union st_exch_brd_msg
{
	struct st_cmp_header st_cmp;
/*  struct st_cmn_header st_hdr;  *** Commented in Ver 3.1 ***/
  struct st_cmn_header_brd st_hdr;  /*** Ver 3.1 ***/	
	struct st_rp_hdr st_report_hdr;
  struct st_rp_market_stats st_rp_mkt_stts;
	struct st_rp_trailer st_report_trlr;
  struct st_bcast_vct_msgs st_brd_vct_msg;
	struct st_system_info_data st_sys_info_dat;
	struct st_security_update_info st_sec_upd_info;
	struct st_instrument_update_info st_instr_upd_info;
	struct st_participant_update_info st_part_upd_info;
	struct st_security_status_update_info st_sec_stts_upd_info;
	struct st_broadcast_tlimit_exceeded st_brd_tlmt_exceed;
	struct st_bcast_message st_bcast_msg;
	struct st_ticker_trade_data st_ticker;
  struct st_bcast_mbo_mbp st_brd_mbo_mbp;
  struct st_bcast_inq_resp_2 st_brd_inq_res;
	struct st_sec_open_msg st_security_opn;
  struct st_int_header st_interactive_hdr;
  struct st_bcast_indices st_brd_indices;
  struct st_bcast_industry_indices st_brd_indust_indices;
	struct st_spd_mkt_info st_spread_mkt_info;
	struct st_download_index st_dwnld_indx;
	struct st_ms_download_bp_tok st_dwnld_bp_tok;
	struct st_ms_download_index_map st_dwnld_indx_map;
	struct st_bcast_only_mbp st_brd_only_mbp;   /****Ver 1.1 for NNF 4.9.3****/
	struct st_oi_info st_eq_oi;
  struct st_indexdetails st_index_details;        /*** Ver 1.9 ***/
  struct st_ind_vix_dtls st_indvix_dtls;        	/*** Ver 2.3 ***/
  struct st_eqoi_msg st_eqoi_data;   /** Ver 2.1 **/
	struct st_bcast_trd_exec_rng st_brd_trd_exec_rng;  /*** Ver 2.5 ***/
	char c_flr [ 512 ];
	struct st_brd_exch_exception st_brd_exch_exception_msg;		/*** Ver 2.6 ***/
	struct st_spd_update_info st_brd_spd_update_info;  				/*** Ver 3.0 ***/
};

/**Added for TAP Ver 1.6 **/
struct st_invt_pckt
{
 struct st_int_header st_hdr;
 short si_inv_cnt;
};

struct st_net_hdr
{
 short si_message_length;
 int i_seq_num;
 unsigned char c_checksum[16];
};

union st_exch_rcv_msg
{
	struct st_cmn_header st_hdr;
	struct st_cmn_header_tr st_hdr_tr;
  struct st_inner_header st_innr_hdr;  /** Ver 2.9 **/
	struct st_sign_on_res st_lgn_res;
	struct st_sign_off_res st_lgf_res;
	struct st_system_info_data st_sys_inf_dat;
  struct st_update_ldb_header st_upd_ldb_hdr;
  struct st_local_database_data st_ldb_data;
	union un_ldb_data u_ldbmsg;
  struct st_int_header st_interactive_hdr;
	struct st_dwnld_msg st_dwnldmsg;
  union u_dwld_msgs un_dw_msg;
	struct st_oe_reqres st_oe_res;
	struct st_oe_rspn_tr st_oe_res_tr;						/*** Added in Ver 2.9 ***/
	struct st_spd_oe_reqres st_spdoe_reqres;
	struct st_trade_inq_reqres st_trd_inq_res;
	struct st_trade_confirm st_trd_conf;
	struct st_trade_confirm_tr st_trd_conf_tr;		/*** Added in Ver 2.9 ***/
	struct st_trader_int_msg st_trd_msg;
	struct st_ex_pl_reqres st_expl_reqres;
	struct st_error_response st_err_res;
	struct st_invt_pckt	st_inv_pck;												/**Ver 1.6 Added for TAP **/
	struct st_brd_exch_exception st_brd_exch_exception_msg;   /*** Ver 2.8 ***/
	char c_flr [ 512 ];
};

union st_exch_snd_msg
{
	struct st_sign_on_req st_signon_req;
	struct st_system_info_req st_sysinf_req;
	struct st_update_local_database st_updlcdb_req;
	struct st_message_download st_msg_dwld;
	struct st_oe_reqres st_oe_res;
	struct st_oe_reqres_tr st_oe_req_tr;					/*** Added in Ver 2.9 ***/
	struct st_om_rqst_tr st_om_req_tr;						/*** Added in Ver 2.9 ***/
	struct st_spd_oe_reqres st_spdoe_reqres;
	struct st_ex_pl_reqres st_expl_reqres;
};

struct st_exch_msg
{
	struct st_net_hdr st_net_header;
	union st_exch_snd_msg st_exch_sndmsg;
};

struct st_bcast_pack_data
{
	char c_netid[2];
	short int si_no_of_packets;
	char c_pack_data[512];
};

struct st_bcast_cmp_packet
{
	short int si_comp_len;
	char c_comp_data[512];
}; 

/*ver 1.7 */
struct st_req_q_data
{
	long int li_msg_type;
	struct st_exch_msg st_exch_msg_data;
};

/**** Ver 2.1 Starts Here ***/

union st_brd_msgs
{
struct st_bcast_only_mbp st_brd_only_mbp;
};

struct st_brd_q_data
{
long l_msg_typ;
 union st_brd_msgs st_brd_data;
};


struct st_brdcst_msgs
{
 long l_msg_typ;
 union st_exch_brd_msg st_xchng_brdmsg;
};


/**** Ver 2.1 Ends Here ****/

#pragma pack ( )
