package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

type CoinData struct {
    Name string
    Rank string
    Symbol string
    Price_usd string
    Percent_change_24h string
    Percent_change_1h string
    Market_cap_usd string
}

type ApiResponse struct {
    Collection []CoinData
}

func main() {
    // Program Name is always the first (implicit) argument
    cmd := os.Args[0]

    fmt.Printf("Program Name: %s\n", cmd)

    url := "https://api.coinmarketcap.com/v1/ticker"

    spaceClient := http.Client{
        Timeout: time.Second * 2, // Maximum of 2 secs
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("User-Agent", "spacecount-tutorial")

    res, getErr := spaceClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    keys := make([]CoinData,0)
    json.Unmarshal(body, &keys)

    // Price (USD)   │ Change (24H)  │ Change (1H)   │ Market Cap (USD)  

    fmt.Println("--------------------------------------------------------------------------------------------")
    fmt.Printf("|%-6s|%-10s|%-15s|%-20s|%-20s|%-25s|\n", "RANK", "COIN", "Price (USD)", "Change (24H)", "Change (1H)", "Market Cap (USD)")

    for i := 0; i < len(keys); i += 1 {
    	v := keys[i]
    	fmt.Printf("|%-6s|%-10s|%-15s|%-20s|%-20s|%-25s|\n", v.Rank, v.Symbol, v.Price_usd, v.Percent_change_24h, v.Percent_change_1h, v.Market_cap_usd)
	}

	fmt.Println("--------------------------------------------------------------------------------------------")
}