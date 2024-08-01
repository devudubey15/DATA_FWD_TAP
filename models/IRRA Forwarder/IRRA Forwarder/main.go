package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	host          = "localhost"
	port          = 5432
	user          = "postgres"
	password      = "Pass@123"
	dbname        = "IRRA"
	serverAddress = "localhost:8080"
	retryDelay    = 5 * time.Second
)

type st_spd_oe_reqres struct {
	C_ordr_rfrnc    string
	C_slm_flg       string
	L_dsclsd_qty    int64
	L_ord_tot_qty   int64
	L_ord_lmt_rt    int64
	L_stp_lss_tgr   int64
	L_mdfctn_cntr   int64
	C_valid_dt      string
	C_ord_typ       string
	C_sprd_ord_ind  string
	C_req_typ       string
	L_quote         int64
	C_qt_tm         string
	C_rqst_tm       string
	C_frwrd_tm      string
	C_plcd_stts     string
	C_rms_prcsd_flg string
	L_ors_msg_typ   int64
	C_ack_tm        string
	C_rmrks         string
	C_ex_ordr_typ   string
	L_xchng_can_qty int64
	C_spl_flg       string
	L_ord_seq       int64
	C_ip_addrs      string
	C_prcimpv_flg   string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach the database:", err)
	}

	messageChannel := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go sender(db, messageChannel, &wg)

	wg.Add(1)
	go receiver(messageChannel, &wg)

	wg.Wait()
}

// Sender will send the TCP messages to the TAP
func sender(db *sql.DB, messageChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn := manageConnection()
	defer conn.Close()

	for {
		records := fetchNewRecords(db)
		for _, record := range records {
			message := formatRecord(record)

			_, err := conn.Write([]byte(message))
			if err != nil {
				log.Println("Error writing:", err)
				conn = manageConnection()
			}

			messageChannel <- message
			time.Sleep(1 * time.Second)
		}

		time.Sleep(10 * time.Second) // Check for new records every 10 seconds
	}
}

// Receiver will receive the messages from the TAP
func receiver(messageChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for message := range messageChannel {
		fmt.Println("Received from sender:", message)
	}
}

func manageConnection() net.Conn {
	for {
		conn, err := net.Dial("tcp", serverAddress)
		if err == nil {
			return conn
		}

		log.Println("Error connecting to server, retrying in", retryDelay)
		time.Sleep(retryDelay)
	}
}

func fetchNewRecords(db *sql.DB) []st_spd_oe_reqres {
	query := `
		SELECT fxb_ordr_rfrnc, fxb_lmt_mrkt_sl_flg, fxb_dsclsd_qty, fxb_ordr_tot_qty,
		       fxb_lmt_rt, fxb_stp_lss_tgr, fxb_mdfctn_cntr, to_char(fxb_ordr_valid_dt, 'dd-mon-yyyy'),
		       CASE WHEN fxb_ordr_type = 'V' THEN 'T' ELSE fxb_ordr_type END, 
		       fxb_sprd_ord_ind, fxb_rqst_typ, fxb_quote,
		       to_char(fxb_qt_tm, 'dd-mon-yyyy hh24:mi:ss'), 
		       to_char(fxb_rqst_tm, 'dd-mon-yyyy hh24:mi:ss'),
		       to_char(fxb_frwd_tm, 'dd-mon-yyyy hh24:mi:ss'), 
		       fxb_plcd_stts, fxb_rms_prcsd_flg,
		       fxb_ors_msg_typ, to_char(fxb_ack_tm, 'dd-mon-yyyy hh24:mi:ss'), 
		       fxb_xchng_rmrks, fxb_ex_ordr_typ,
		       fxb_xchng_cncld_qty, fxb_spl_flag, fxb_ordr_sqnc,
		       COALESCE(fxb_ip, 'NA'), COALESCE(fxb_prcimpv_flg, 'N')
		FROM public.fxb_fo_xchng_book
		WHERE fxb_processed = false
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	var records []st_spd_oe_reqres
	for rows.Next() {
		var record st_spd_oe_reqres
		err := rows.Scan(
			&record.C_ordr_rfrnc, &record.C_slm_flg, &record.L_dsclsd_qty, &record.L_ord_tot_qty,
			&record.L_ord_lmt_rt, &record.L_stp_lss_tgr, &record.L_mdfctn_cntr, &record.C_valid_dt,
			&record.C_ord_typ, &record.C_sprd_ord_ind, &record.C_req_typ, &record.L_quote,
			&record.C_qt_tm, &record.C_rqst_tm, &record.C_frwrd_tm, &record.C_plcd_stts, &record.C_rms_prcsd_flg,
			&record.L_ors_msg_typ, &record.C_ack_tm, &record.C_rmrks, &record.C_ex_ordr_typ,
			&record.L_xchng_can_qty, &record.C_spl_flg, &record.L_ord_seq, &record.C_ip_addrs,
			&record.C_prcimpv_flg)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error with rows:", err)
	}

	// Mark records as processed
	for _, record := range records {
		_, err = db.Exec("UPDATE public.fxb_fo_xchng_book SET fxb_processed = true WHERE fxb_ordr_rfrnc = $1", record.C_ordr_rfrnc)
		if err != nil {
			log.Fatal("Error updating record as processed:", err)
		}
	}

	return records
}

func formatRecord(record st_spd_oe_reqres) string {
	var sb strings.Builder
	sb.WriteString(record.C_ordr_rfrnc + ",")
	sb.WriteString(record.C_slm_flg + ",")
	sb.WriteString(fmt.Sprintf("%d,", record.L_dsclsd_qty))
	sb.WriteString(fmt.Sprintf("%d,", record.L_ord_tot_qty))
	sb.WriteString(fmt.Sprintf("%d,", record.L_ord_lmt_rt))
	sb.WriteString(fmt.Sprintf("%d,", record.L_stp_lss_tgr))
	sb.WriteString(fmt.Sprintf("%d,", record.L_mdfctn_cntr))
	sb.WriteString(record.C_valid_dt + ",")
	sb.WriteString(record.C_ord_typ + ",")
	sb.WriteString(record.C_sprd_ord_ind + ",")
	sb.WriteString(record.C_req_typ + ",")
	sb.WriteString(fmt.Sprintf("%d,", record.L_quote))
	sb.WriteString(record.C_qt_tm + ",")
	sb.WriteString(record.C_rqst_tm + ",")
	sb.WriteString(record.C_frwrd_tm + ",")
	sb.WriteString(record.C_plcd_stts + ",")
	sb.WriteString(record.C_rms_prcsd_flg + ",")
	sb.WriteString(fmt.Sprintf("%d,", record.L_ors_msg_typ))
	sb.WriteString(record.C_ack_tm + ",")
	sb.WriteString(record.C_rmrks + ",")
	sb.WriteString(record.C_ex_ordr_typ + ",")
	sb.WriteString(fmt.Sprintf("%d,", record.L_xchng_can_qty))
	sb.WriteString(record.C_spl_flg + ",")
	sb.WriteString(fmt.Sprintf("%d,", record.L_ord_seq))
	sb.WriteString(record.C_ip_addrs + ",")
	sb.WriteString(record.C_prcimpv_flg)

	return sb.String()
}
