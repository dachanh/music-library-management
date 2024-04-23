package conf

import (
	"os"
)

type config struct {
	Env      string
	Protocol string
	Version  string

	MainDBName  string
	MainAuthDB  string
	LogDBName   string
	LogAuthDB   string
	JobDBName   string
	JobAuthDB   string
	CacheDBName string

	APIHost string
	APIKey  string

	SearchProductModel string
}

type ERPConfig struct {
	Host   string
	APIKey string
}

// Config main config object
var Config *config

var SearchVersion = ""

func init() {
	env := os.Getenv("env")
	protocol := os.Getenv("protocol")
	version := os.Getenv("version")

	switch env {

	// config for staging
	case "dev":
		Config = &config{
			Env:      "dev",
			Protocol: protocol,
			Version:  version,

			MainDBName:  "seller_dev_purchasing",
			MainAuthDB:  "admin",
			LogDBName:   "seller_dev_purchasing_log",
			LogAuthDB:   "admin",
			JobDBName:   "seller_dev_purchasing_job",
			JobAuthDB:   "admin",
			CacheDBName: "seller_dev_purchasing",

			APIHost: "https://api.v2-dev.thuocsi.vn",
			APIKey:  "Basic UEFSVE5FUi92Mi5jdXN0b21lci5jdXN0b21lcjpWNHRqTDI5UVQ0",

			SearchProductModel: "dev.seller.purchasing.quotation",
		}

	case "stg":
		Config = &config{
			Env:      "stg",
			Protocol: protocol,
			Version:  version,

			MainDBName:  "seller_stg_purchasing",
			MainAuthDB:  "admin",
			LogDBName:   "seller_stg_purchasing_log",
			LogAuthDB:   "admin",
			JobDBName:   "seller_stg_purchasing_job",
			JobAuthDB:   "admin",
			CacheDBName: "seller_stg_purchasing",

			APIHost: "https://api.v2-stg.thuocsi.vn",
			APIKey:  "Basic UEFSVE5FUi92Mi5jdXN0b21lci5jdXN0b21lcjpWNHRqTDI5UVQ0",

			SearchProductModel: "stg.seller.purchasing.quotation",
		}

	case "uat":
		Config = &config{
			Env:      "uat",
			Protocol: protocol,
			Version:  version,

			MainDBName:  "seller_prd_purchasing",
			MainAuthDB:  "seller_prd_purchasing",
			LogDBName:   "seller_uat_purchasing_log",
			LogAuthDB:   "admin",
			JobDBName:   "seller_uat_purchasing_job",
			JobAuthDB:   "admin",
			CacheDBName: "seller_uat_purchasing",

			APIHost: "https://api.v2-uat.thuocsi.vn",
			APIKey:  "Basic UEFSVE5FUi92Mi5wdXJjaGFzaW5nOjJEa3hsd1BzaXJWdXdwOTIwV3JnZw==",

			SearchProductModel: "prd.seller.purchasing.quotation",
		}

	case "prd":
		Config = &config{
			Env:      "prd",
			Protocol: protocol,
			Version:  version,

			MainDBName:  "seller_prd_purchasing",
			MainAuthDB:  "seller_prd_purchasing",
			LogDBName:   "seller_prd_purchasing_log",
			LogAuthDB:   "admin",
			JobDBName:   "seller_prd_purchasing_job",
			JobAuthDB:   "admin",
			CacheDBName: "seller_prd_purchasing",

			APIHost: "http://proxy-service.frontend-prd",
			APIKey:  "Basic UEFSVE5FUi92Mi5wdXJjaGFzaW5nOjJEa3hsd1BzaXJWdXdwOTIwV3JnZw==",

			SearchProductModel: "prd.seller.purchasing.quotation",
		}

	}
}
