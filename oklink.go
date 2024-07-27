package oklink

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	BASE_URL string = "https://www.oklink.com/";
	CHAIN_ID string = "8217";
	CHAIN_FULLNAME string = "KLAYTN";
	CHAIN_SHORTNAME string = "KLAYTN";
)

type Address string

var address Address = fmt.Sprintf("%xstring", 0)

type ProtocolType string

const (
	Token20   ProtocolType = "token_20"
	Token721  ProtocolType = "token_721"
	Token1155 ProtocolType = "token_1155"
)

type AddressInformation struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data AddressData `json:"data"`
}

type AddressData struct {
    ChainFullName               string `json:"chainFullName"`
    ChainShortName              string `json:"chainShortName"`
    Address                     string `json:"address"`
    ContractAddress             string `json:"contractAddress"`
    Balance                     string `json:"balance"`
    BalanceSymbol               string `json:"balanceSymbol"`
    TransactionCount            string `json:"transactionCount"`
    Verifying                   string `json:"verifying"`
    SendAmount                  string `json:"sendAmount"`
    ReceiveAmount               string `json:"receiveAmount"`
    TokenAmount                 string `json:"tokenAmount"`
    TotalTokenValue             string `json:"totalTokenValue"`
    CreateContractAddress       string `json:"createContractAddress"`
    CreateContractTransactionHash string `json:"createContractTransactionHash"`
    FirstTransactionTime        string `json:"firstTransactionTime"`
    LastTransactionTime         string `json:"lastTransactionTime"`
    Token                       string `json:"token"`
    Bandwidth                   string `json:"bandwidth"`
    Energy                      string `json:"energy"`
    VotingRights                string `json:"votingRights"`
    UnclaimedVotingRewards      string `json:"unclaimedVotingRewards"`
    IsAaAddress                 bool   `json:"isAaAddress"`
}

type ApiResponse[T any] struct {
	Code 	int 	`json:"code"`
	Data 	T 		`json:"code"`
	Msg 	string 	`json:"code"`
}

func main()  {
	info, err := AddressInfo("klaytn", "0x1234567890abcdef")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Address Information: %+v\n", info)
}

func fetchApi[T any](url string) (*ApiResponse[T], error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return nil,fmt.Errorf("error making request: %w", err);
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error! status: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	var apiResponse ApiResponse[T]
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	if apiResponse.Code != 0 {
		return nil, fmt.Errorf("API error! code: %d, message: %s", apiResponse.Code, apiResponse.Msg)
	}
	return &apiResponse, nil

}

	// response, err := http.Get(url);
	// if err != nil {
	// 	return nil, err
	// };
	// defer response.Body.Close();
	// if response.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("failed to fetch address info: status code %d", response.StatusCode)
	// }
	// var result ApiResponse[AddressInfo]
	// err = json.NewDecoder(response.Body).Decode(&result)
	// if err != nil {
	// 	return nil, err
	// }
	// return &result, nil


func AddressInfo(address Address) (*ApiResponse[AddressInformation], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)

	url := fmt.Sprintf("%sapi/v5/explorer/address/address-summary?%s", BASE_URL, params.Encode())
	return fetchApi[AddressInformation](url)
}

func EvmAddressInfo(address Address) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)

	url := fmt.Sprintf("%sapi/v5/explorer/address/information-evm?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressActiveChain(address Address) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)

	url := fmt.Sprintf("%sapi/v5/explorer/address/address-active-chain?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressTokenBalance(address Address, protocolType ProtocolType, tokenContractAddress *Address, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)
	params.Add("protocolType", protocolType)

	if tokenContractAddress != nil {
		params.Add("tokenContractAddress", string(*tokenContractAddress))
	}

	if page != nil {
		params.Add("page", *page)
	}
	
	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/token-balance?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressBalanceDetails(address Address, protocolType ProtocolType, tokenContractAddress *Address, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)
	params.Add("protocolType", protocolType)

	if tokenContractAddress != nil {
		params.Add("tokenContractAddress", string(*tokenContractAddress))
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/address-balance-fills?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressTokenTransactionList(address Address, protocolType ProtocolType, tokenContractAddress *Address, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)
	params.Add("protocolType", protocolType)

	if tokenContractAddress != nil {
		params.Add("tokenContractAddress", string(*tokenContractAddress))
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/token-transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressBalanceHistory(address Address, height string, tokenContractAddress *Address) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)
	params.Add("height", height)

	if tokenContractAddress != nil {
		params.Add("tokenContractAddress", string(*tokenContractAddress))
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/address-balance-fills?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressTransactionList(address Address, protocolType *ProtocolType, symbol *string, startBlockHeight *string, endBlockHeight *string, isFromOrTo *string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)

	if protocolType != nil {
		params.Add("protocolType", *string(protocolType))
	}

	if symbol != nil {
		params.Add("symbol", *symbol)
	}

	if startBlockHeight != nil {
		params.Add("startBlockHeight", *startBlockHeight)
	}

	if endBlockHeight != nil {
		params.Add("endBlockHeight", *endBlockHeight)
	}

	if isFromOrTo != nil {
		params.Add("isFromOrTo", *isFromOrTo)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressNormalTransactionList(address Address, startBlockHeight *string, endBlockHeight *string, isFromOrTo *string, page *string, limit *string) (*ApiResponse[any], error)  {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)

	if startBlockHeight != nil {
		params.Add("startBlockHeight", *startBlockHeight)
	}

	if endBlockHeight != nil {
		params.Add("endBlockHeight", *endBlockHeight)
	}

	if isFromOrTo != nil {
		params.Add("isFromOrTo", *isFromOrTo)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/normal-transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressInternalTransactionList(address Address, startBlockHeight *string, endBlockHeight *string, isFromOrTo *string, page *string, limit *string) (*ApiResponse[any], error)  {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)

	if startBlockHeight != nil {
		params.Add("startBlockHeight", *startBlockHeight)
	}

	if endBlockHeight != nil {
		params.Add("endBlockHeight", *endBlockHeight)
	}

	if isFromOrTo != nil {
		params.Add("isFromOrTo", *isFromOrTo)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/internal-transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressTokenTransactionList(address Address, protocolType ProtocolType, tokenContractAddress *Address, page *string, limit *string) (*ApiResponse[any], error)  {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)
	params.Add("protocolType", protocolType)

	if tokenContractAddress != nil {
		params.Add("tokenContractAddress", *tokenContractAddress)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/token-transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func AddressEntityLabels(address Address) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("address", address)

	url := fmt.Sprintf("%sapi/v5/explorer/address/entity-labels?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}


func RichList(address *Address) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)

	if address != nil {
		params.Add("address", *address)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/rich-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func NativeTokenRanking(page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/native-token-position-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TransactionList(blockhash *string, height *string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)

	if blockhash != nil {
		params.Add("blockhash", *blockhash)
	}

	if height != nil {
		params.Add("height", *height)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func LargeTransactionList(type *string, height *string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)

	if type != nil {
		params.Add("type", *type)
	}

	if height != nil {
		params.Add("height", *height)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func UnconfirmedTransactionList(page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/unconfirmed-transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func InternalTransactionDetails(txId string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("txId", txId)

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/internal-transaction-detail?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TokenTransactionDetails(txId string, protocolType ProtocolType, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("txId", txId)

	if protocolType != nil {
		params.Add("protocolType", protocolType)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/token-transaction-detail?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TransactionDetails(txId string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("txId", txId)

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/transaction-fills?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TokenSupplyHistory(tokenContractAddress Address, height string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("tokenContractAddress", tokenContractAddress)
	params.Add("height", height)

	url := fmt.Sprintf("%sapi/v5/explorer/token/supply-history?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TokenList(protocolType *ProtocolType, tokenContractAddress Address, startTime *string, endTime *string, orderBy *string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)

	if protocolType != nil {
		params.Add("protocolType", protocolType)
	}

	if tokenContractAddress != nil {
		params.Add("tokenContractAddress", tokenContractAddress)
	}

	if startTime != nil {
		params.Add("startTime", startTime)
	}

	if endTime != nil {
		params.Add("endTime", *endTime)
	}

	if orderBy != nil {
		params.Add("orderBy", *orderBy)
	}

	if page != nil {
		params.Add("page", page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/token/token-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TokenPositionList(tokenContractAddress *Address, holderAddress *Address, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("tokenContractAddress", tokenContractAddress)
	
	if holderAddress != nil {
		params.Add("holderAddress", *holderAddress)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/token/position-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TokenPositionStatistics(tokenContractAddress *Address, holderAddress *Address, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("tokenContractAddress", tokenContractAddress)
	
	if holderAddress != nil {
		params.Add("holderAddress", *holderAddress)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/token/position-statistics?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TokenTransferDetails(tokenContractAddress Address, maxAmount *string, minAmount *string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("tokenContractAddress", tokenContractAddress)
	
	if maxAmount != nil {
		params.Add("maxAmount", *maxAmount)
	}

	if minAmount != nil {
		params.Add("minAmount", *minAmount)
	}

	if page != nil {
		params.Add("page", *page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/token/transaction-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func TokenTransactionStatistics(tokenContractAddress Address, orderBy *string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("tokenContractAddress", tokenContractAddress)

	if orderBy != nil {
		params.Add("orderBy", *orderBy)
	}

	if page != nil {
		params.Add("page", page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/token/token-list?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func BatchAddressBalances(addresses []Address) (*ApiResponse[any], error) {
	if len(addresses) > 100 {
		return "", errors.New("The maximum number of addresses is 100")
	}

	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)	
	params.Add("addresses", strings.Join(addresses, ","))

	url := fmt.Sprintf("%sapi/v5/explorer/address/balance-multi?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func BatchAddressTokenBalances(addresses []Address, protocolType *ProtocolType, page *string, limit *string) (*ApiResponse[any], error) {
	if len(addresses) > 50 {
		return "", errors.New("The maximum number of addresses is 50")
	}

	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)	
	params.Add("addresses", strings.Join(addresses, ","))

	if protocolType != nil {
		params.Add("protocolType", protocolType)
	}

	if page != nil {
		params.Add("page", page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/token-balance-multi?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func BatchAddressNormalTransactionList(addresses []Address, startBlockHeight *string, endBlockHeight *string, isFromOrTo *string, page *string, limit *string) (*ApiResponse[any], error) {
	if len(addresses) > 50 {
		return "", errors.New("The maximum number of addresses is 50")
	}

	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)	
	params.Add("addresses", strings.Join(addresses, ","))

	if startBlockHeight != nil {
		params.Add("startBlockHeight", *startBlockHeight)
	}

	if endBlockHeight != nil {
		params.Add("endBlockHeight", *endBlockHeight)
	}

	if isFromOrTo != nil {
		params.Add("isFromOrTo", *isFromOrTo)
	}

	if page != nil {
		params.Add("page", page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/normal-transaction-list-multi?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func BatchAddressInternalTransactionList(addresses []Address, startBlockHeight *string, endBlockHeight *string, isFromOrTo *string, page *string, limit *string) (*ApiResponse[any], error) {
	if len(addresses) > 20 {
		return "", errors.New("The maximum number of addresses is 20")
	}

	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)	
	params.Add("addresses", strings.Join(addresses, ","))

	if startBlockHeight != nil {
		params.Add("startBlockHeight", *startBlockHeight)
	}

	if endBlockHeight != nil {
		params.Add("endBlockHeight", *endBlockHeight)
	}

	if isFromOrTo != nil {
		params.Add("isFromOrTo", *isFromOrTo)
	}

	if page != nil {
		params.Add("page", page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/internal-transaction-list-multi?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func BatchAddressTokenTransactionList(addresses []Address, startBlockHeight string, endBlockHeight string, protocolType *ProtocolType, tokenContractAddress *Address, isFromOrTo *string, page *string, limit *string) (*ApiResponse[any], error) {
	if len(addresses) > 20 {
		return "", errors.New("The maximum number of addresses is 20")
	}

	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)	
	params.Add("addresses", strings.Join(addresses, ","))
	params.Add("startBlockHeight", startBlockHeight)
	params.Add("endBlockHeight", endBlockHeight)

	if tokenContractAddress != nil {
		params.Add("tokenContractAddress", *tokenContractAddress)
	}

	if protocolType != nil {
		params.Add("protocolType", *protocolType)
	}

	if isFromOrTo != nil {
		params.Add("isFromOrTo", *isFromOrTo)
	}

	if page != nil {
		params.Add("page", page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/address/token-transaction-list-multi?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func BatchTokenTransaction(tokenContractAddress Address, startBlockHeight string, endBlockHeight string, page *string, limit *string) (*ApiResponse[any], error) {
	params := url.Values{}
	params.Add("CHAIN_SHORTNAME", CHAIN_SHORTNAME)
	params.Add("tokenContractAddress", tokenContractAddress)
	params.Add("startBlockHeight", startBlockHeight)
	params.Add("endBlockHeight", endBlockHeight)

	if page != nil {
		params.Add("page", page)
	}

	if limit != nil {
		params.Add("limit", limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/token/token-transaction-list-multi?%s", BASE_URL, params.Encode())
	return fetchApi[any](url)
}

func BatchTransactionDetails(txIds []string) (*ApiResponse[any], error) {
	if len(txIds) > 20 {
		return "", errors.New("the maximum number of transactions is 20")
	}

	params := url.Values{}
	params.Add("chainShortName", chainShortName)
	params.Add("txIds", strings.Join(txIds, ","))

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/transaction-multi?%s", baseUrl, params.Encode())

	return fetchApi(url)
}

func BatchInternalTransactionDetails(txIds []string) (*ApiResponse[any], error) {
	if len(txIds) > 20 {
		return "", errors.New("the maximum number of transactions is 20")
	}

	params := url.Values{}
	params.Add("chainShortName", chainShortName)
	params.Add("txIds", strings.Join(txIds, ","))

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/internal-transaction-multi?%s", baseUrl, params.Encode())

	return fetchApi(url)
}

func BatchTokenTransactionDetails(txIds []string, protocolType *ProtocolType, page *string, limit *string) (*ApiResponse[any], error) {
	if len(txIds) > 20 {
		return "", errors.New("the maximum number of transactions is 20")
	}

	params := url.Values{}
	params.Add("chainShortName", chainShortName)
	params.Add("txIds", strings.Join(txIds, ","))

	if protocolType != nil {
		params.Add("protocolType", *protocolType)
	}

	if page != nil {
		params.Add("protocolType", *page)
	}

	if limit != nil {
		params.Add("protocolType", *limit)
	}

	url := fmt.Sprintf("%sapi/v5/explorer/transaction/token-transfer-multi?%s", baseUrl, params.Encode())

	return fetchApi(url)
}
//