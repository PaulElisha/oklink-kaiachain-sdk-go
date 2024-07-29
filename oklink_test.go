package oklink

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// Mock server to simulate API responses
func setupMockServer(responseBody string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}))
}

func TestFetchApi(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"chainFullName": "KLAYTN"}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	apiUrl := server.URL
	response, err := fetchApi[AddressInformation](apiUrl)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	if response.Data.ChainFullName != "KLAYTN" {
		t.Errorf("Expected chainFullName KLAYTN, got %s", response.Data.ChainFullName)
	}
}

func TestAddressInfo(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"chainFullName": "KLAYTN"}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	BASE_URL = server.URL // Override the base URL to point to the mock server

	address := Address("0x1234567890abcdef")
	response, err := AddressInfo(address)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	if response.Data.ChainFullName != "KLAYTN" {
		t.Errorf("Expected chainFullName KLAYTN, got %s", response.Data.ChainFullName)
	}
}

func TestAddressTokenBalance(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"balance": "1000"}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	BASE_URL = server.URL // Override the base URL to point to the mock server

	address := Address("0x1234567890abcdef")
	protocolType := Token20
	response, err := AddressTokenBalance(address, protocolType, nil, nil, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	if response.Data.(map[string]interface{})["balance"] != "1000" {
		t.Errorf("Expected balance 1000, got %s", response.Data.(map[string]interface{})["balance"])
	}
}

func TestEvmAddressInfo(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"chainFullName": "KLAYTN"}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	BASE_URL = server.URL // Override the base URL to point to the mock server

	address := Address("0x1234567890abcdef")
	response, err := EvmAddressInfo(address)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	if response.Data.(map[string]interface{})["chainFullName"] != "KLAYTN" {
		t.Errorf("Expected chainFullName KLAYTN, got %s", response.Data.(map[string]interface{})["chainFullName"])
	}
}

func TestAddressActiveChain(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"chainFullName": "KLAYTN"}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	BASE_URL = server.URL // Override the base URL to point to the mock server

	address := Address("0x1234567890abcdef")
	response, err := AddressActiveChain(address)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	if response.Data.(map[string]interface{})["chainFullName"] != "KLAYTN" {
		t.Errorf("Expected chainFullName KLAYTN, got %s", response.Data.(map[string]interface{})["chainFullName"])
	}
}

func TestAddressBalanceDetails(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"balance": "1000"}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	BASE_URL = server.URL // Override the base URL to point to the mock server

	address := Address("0x1234567890abcdef")
	protocolType := Token20
	response, err := AddressBalanceDetails(address, protocolType, nil, nil, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	if response.Data.(map[string]interface{})["balance"] != "1000" {
		t.Errorf("Expected balance 1000, got %s", response.Data.(map[string]interface{})["balance"])
	}
}

func TestAddressTransactionList(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"transactionCount": "10"}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	BASE_URL = server.URL // Override the base URL to point to the mock server

	address := Address("0x1234567890abcdef")
	response, err := AddressTransactionList(address, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	if response.Data.(map[string]interface{})["transactionCount"] != "10" {
		t.Errorf("Expected transactionCount 10, got %s", response.Data.(map[string]interface{})["transactionCount"])
	}
}

func TestBatchAddressBalances(t *testing.T) {
	mockResponse := `{"code": 0, "data": {"balances": [{"address": "0x1234567890abcdef", "balance": "1000"}]}, "msg": "success"}`
	server := setupMockServer(mockResponse, http.StatusOK)
	defer server.Close()

	BASE_URL = server.URL // Override the base URL to point to the mock server

	addresses := []Address{"0x1234567890abcdef"}
	response, err := BatchAddressBalances(addresses)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Code != 0 {
		t.Errorf("Expected code 0, got %d", response.Code)
	}

	balances := response.Data.(map[string]interface{})["balances"].([]interface{})
	if balances[0].(map[string]interface{})["balance"] != "1000" {
		t.Errorf("Expected balance 1000, got %s", balances[0].(map[string]interface{})["balance"])
	}
}