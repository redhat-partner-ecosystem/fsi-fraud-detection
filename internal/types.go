package internal

/*
IDX,TRANSACTION_ID,TX_DATETIME,CUSTOMER_ID,TERMINAL_ID,TX_AMOUNT,
0,0,2020-04-01 00:00:31,596,3156,57.16,

TX_TIME_SECONDS,TX_TIME_DAYS,TX_FRAUD,TX_FRAUD_SCENARIO,TX_DURING_WEEKEND,TX_DURING_NIGHT,
31,0,0,0,0,1

CUSTOMER_ID_NB_TX_1DAY_WINDOW,CUSTOMER_ID_AVG_AMOUNT_1DAY_WINDOW,CUSTOMER_ID_NB_TX_7DAY_WINDOW,CUSTOMER_ID_AVG_AMOUNT_7DAY_WINDOW,CUSTOMER_ID_NB_TX_30DAY_WINDOW,CUSTOMER_ID_AVG_AMOUNT_30DAY_WINDOW,TERMINAL_ID_NB_TX_1DAY_WINDOW,TERMINAL_ID_RISK_1DAY_WINDOW,TERMINAL_ID_NB_TX_7DAY_WINDOW,TERMINAL_ID_RISK_7DAY_WINDOW,TERMINAL_ID_NB_TX_30DAY_WINDOW,TERMINAL_ID_RISK_30DAY_WINDOW
1.0,57.16,1.0,57.16,1.0,57.16,0.0,0.0,0.0,0.0,0.0,0.0


{24 1585700797000 2323 1257 63.19 1597 0 1 42 0 1 0 0 0 0 0 0 0 0 0 0 0 0}

*/

type (
	Transaction struct {
		TRANSACTION_ID                      int64
		TX_DATETIME                         int64
		CUSTOMER_ID                         int
		TERMINAL_ID                         int
		TX_AMOUNT                           float64
		TX_TIME_SECONDS                     int
		TX_TIME_DAYS                        int
		TX_FRAUD                            int
		TX_FRAUD_SCENARIO                   int
		TX_DURING_WEEKEND                   int
		TX_DURING_NIGHT                     int
		CUSTOMER_ID_NB_TX_1DAY_WINDOW       float64
		CUSTOMER_ID_AVG_AMOUNT_1DAY_WINDOW  float64
		CUSTOMER_ID_NB_TX_7DAY_WINDOW       float64
		CUSTOMER_ID_AVG_AMOUNT_7DAY_WINDOW  float64
		CUSTOMER_ID_NB_TX_30DAY_WINDOW      float64
		CUSTOMER_ID_AVG_AMOUNT_30DAY_WINDOW float64
		TERMINAL_ID_NB_TX_1DAY_WINDOW       float64
		TERMINAL_ID_RISK_1DAY_WINDOW        float64
		TERMINAL_ID_NB_TX_7DAY_WINDOW       float64
		TERMINAL_ID_RISK_7DAY_WINDOW        float64
		TERMINAL_ID_NB_TX_30DAY_WINDOW      float64
		TERMINAL_ID_RISK_30DAY_WINDOW       float64
	}
)
