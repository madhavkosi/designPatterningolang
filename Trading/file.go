package nse

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tealeg/xlsx"
)

const (
	HISTORICAL_DATA_URL  = "https://www.nseindia.com/api/historical/cm/equity?series=[%22EQ%22]&"
	BASE_URL             = "https://www.nseindia.com/"
	PRE_MARKET_URL       = "https://www.nseindia.com/api/market-data-pre-open?key=ALL&date=12-07-2024"
	CORPORATE_EVENTS_URL = "https://www.nseindia.com/api/corporate-announcements?"
	USER_AGENT           = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"
)

type HistoricalData struct {
	Data []struct {
		Timestamp        string `json:"TIMESTAMP"`
		Symbol           string `json:"CH_SYMBOL"`
		Series           string `json:"CH_SERIES"`
		HighPrice        string `json:"CH_TRADE_HIGH_PRICE"`
		LowPrice         string `json:"CH_TRADE_LOW_PRICE"`
		OpenPrice        string `json:"CH_OPENING_PRICE"`
		ClosePrice       string `json:"CH_CLOSING_PRICE"`
		LastTradedPrice  string `json:"CH_LAST_TRADED_PRICE"`
		PrevClosePrice   string `json:"CH_PREVIOUS_CLS_PRICE"`
		TotalTradedQty   string `json:"CH_TOT_TRADED_QTY"`
		TotalTradedValue string `json:"CH_TOT_TRADED_VAL"`
		WeekHighPrice    string `json:"CH_52WEEK_HIGH_PRICE"`
		WeekLowPrice     string `json:"CH_52WEEK_LOW_PRICE"`
	} `json:"data"`
}

type Metadata struct {
	Symbol         string  `json:"symbol"`
	Purpose        string  `json:"purpose"`
	PChange        float64 `json:"pChange"`
	TotalTurnover  float64 `json:"totalTurnover"`
	QuantityToBuy  float64 `json:"quantityToBuy"`
	PrevClosePrice float64 `json:"previousClose"`
}

type PreMarketData struct {
	Declines  int `json:"declines"`
	Unchanged int `json:"unchanged"`
	Data      []struct {
		Metadata Metadata `json:"metadata"`
	} `json:"data"`
}

func getAdjustedHeaders() http.Header {
	headers := http.Header{}
	headers.Set("User-Agent", USER_AGENT)
	headers.Set("Accept", "*/*")
	headers.Set("Accept-Language", "en-US,en;q=0.5")
	headers.Set("Accept-Encoding", "gzip, deflate, br")
	headers.Set("X-Requested-With", "XMLHttpRequest")
	headers.Set("DNT", "1")
	headers.Set("Connection", "keep-alive")
	return headers
}

func fetchCookies(client *http.Client) map[string]string {
	req, err := http.NewRequest("GET", BASE_URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = getAdjustedHeaders()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch cookies: %v", err)
	}
	defer resp.Body.Close()

	cookies := make(map[string]string)
	for _, cookie := range resp.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}
	return cookies
}

func fetchURL(client *http.Client, url string, cookies map[string]string) ([]Metadata, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = getAdjustedHeaders()
	for k, v := range cookies {
		req.AddCookie(&http.Cookie{Name: k, Value: v})
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var data PreMarketData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	var metadataList []Metadata
	for _, item := range data.Data {
		if item.Metadata.TotalTurnover > 1000000 && item.Metadata.PChange > 1.5 { // 10 crores in local currency units
			item.Metadata.QuantityToBuy = 500000 / item.Metadata.PrevClosePrice
			metadataList = append(metadataList, item.Metadata)
		}
	}
	return metadataList, nil
}
func printMetadataTable(metadataList []Metadata) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol", "Purpose", "PChange", "Total Turnover", "Quantity to Buy", "Prev Close Price"})

	for _, data := range metadataList {
		row := []string{
			data.Symbol,
			data.Purpose,
			formatFloat(data.PChange),
			formatFloat(data.TotalTurnover),
			formatFloat(data.QuantityToBuy),
			formatFloat(data.PrevClosePrice),
		}
		table.Append(row)
	}

	table.SetBorder(true)
	table.SetRowLine(true)
	table.Render()
}

func formatFloat(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func formatDataFrameResult(allData []HistoricalData) (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Historical Data")
	if err != nil {
		return nil, err
	}

	columnsRequired := []string{"Timestamp", "Symbol", "Series", "HighPrice", "LowPrice", "OpenPrice", "ClosePrice", "LastTradedPrice", "PrevClosePrice", "TotalTradedQty", "TotalTradedValue", "WeekHighPrice", "WeekLowPrice"}
	header := sheet.AddRow()
	for _, col := range columnsRequired {
		cell := header.AddCell()
		cell.Value = col
	}

	for _, data := range allData {
		for _, record := range data.Data {
			row := sheet.AddRow()
			row.AddCell().Value = record.Timestamp
			row.AddCell().Value = record.Symbol
			row.AddCell().Value = record.Series
			row.AddCell().Value = record.HighPrice
			row.AddCell().Value = record.LowPrice
			row.AddCell().Value = record.OpenPrice
			row.AddCell().Value = record.ClosePrice
			row.AddCell().Value = record.LastTradedPrice
			row.AddCell().Value = record.PrevClosePrice
			row.AddCell().Value = record.TotalTradedQty
			row.AddCell().Value = record.TotalTradedValue
			row.AddCell().Value = record.WeekHighPrice
			row.AddCell().Value = record.WeekLowPrice
		}
	}
	return file, nil
}
func scrapePreMarketData(client *http.Client, cookies map[string]string) ([]Metadata, error) {
	return fetchURL(client, PRE_MARKET_URL, cookies)
}
func main() {
	// Example metadata list
	metadataList := []Metadata{
		{
			Symbol:         "ABC",
			Purpose:        "Annual General Meeting",
			PChange:        12.34,
			TotalTurnover:  123456789.12,
			QuantityToBuy:  1000,
			PrevClosePrice: 45.67,
		},
		{
			Symbol:         "XYZ",
			Purpose:        "Board Meeting",
			PChange:        5.67,
			TotalTurnover:  987654321.45,
			QuantityToBuy:  2000,
			PrevClosePrice: 89.01,
		},
	}

	// Print the metadata table
	printMetadataTable(metadataList)
}
