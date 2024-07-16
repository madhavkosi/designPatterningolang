package nse

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

// Helper function to set up a mock server
func setupMockServer(responseBody string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseBody))
	})
	return httptest.NewServer(handler)
}

func TestFetchCookies(t *testing.T) {
	mockResponse := ``
	server := setupMockServer(mockResponse)
	defer server.Close()

	client := &http.Client{Jar: nil}
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	client.Jar = jar

	// Override BASE_URL with the mock server URL

	cookies := fetchCookies(client)
	if len(cookies) == 0 {
		t.Errorf("Expected cookies, got none")
	}
}

func TestScrapePreMarketData(t *testing.T) {
	mockResponse := `{"data": [{"metadata": {"symbol": "ABC", "companyName": "ABC Corp", "lastPrice": 123.45}}]}`
	server := setupMockServer(mockResponse)
	defer server.Close()

	client := &http.Client{Jar: nil}
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	client.Jar = jar

	cookies := fetchCookies(client)

	// Override PRE_MARKET_URL with the mock server URL

	metadata, err := scrapePreMarketData(client, cookies)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	printMetadataTable(metadata)

}
