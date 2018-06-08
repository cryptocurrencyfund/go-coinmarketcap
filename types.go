package coinmarketcap

import (
	"fmt"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Coin struct
type Coin struct {
	ID               string    `json:"id"`
	Date             time.Time `json:"date"`
	DateStr          string    `json:"dateStr,string"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Rank             int       `json:"rank,string"`
	PriceUsd         float64   `json:"price_usd,string"`
	PriceBtc         float64   `json:"price_btc,string"`
	Usd24hVolume     float64   `json:"24h_volume_usd,string"`
	MarketCapUsd     float64   `json:"market_cap_usd,string"`
	AvailableSupply  float64   `json:"available_supply,string"`
	TotalSupply      float64   `json:"total_supply,string"`
	PercentChange1h  float64   `json:"percent_change_1h,string"`
	PercentChange24h float64   `json:"percent_change_24h,string"`
	PercentChange7d  float64   `json:"percent_change_7d,string"`
	LastUpdated      string    `json:"last_updated"`
}

func (c Coin) String() string {
	return fmt.Sprintf("[%s(%s)]\n", c.Name, c.Symbol) +
		fmt.Sprintf("Rank: %d\n", c.Rank) +
		fmt.Sprintf("Price: $%.2f (%.4f btc)\n", c.PriceUsd, c.PriceBtc) +
		fmt.Sprintf("Market Cap: $%.0f\n", c.MarketCapUsd) +
		fmt.Sprintf("Volume Cap: $%.0f\n", c.Usd24hVolume) +
		fmt.Sprintf("Changes: {1h %.2f%%} {24h %.2f%%} {7d %.2f%%}\n", c.PercentChange1h, c.PercentChange24h, c.PercentChange7d)
}

// MarkdownPrice MarkdownPrice
func (c Coin) MarkdownPrice() string {
	return fmt.Sprintf("* **%s** (%s) - $%.4f (%.4f btc) `%.2f%%` weekly: `%.2f%%`\n",
		c.Symbol, c.Name, c.PriceUsd, c.PriceBtc, c.PercentChange24h, c.PercentChange7d)
}

// MarkdownVolume MarkdownVolume
func (c Coin) MarkdownVolume() string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("* **%s** (%s) `$%.0f`\n", c.Symbol, c.Name, c.Usd24hVolume)
}

// GlobalMarketData struct
type GlobalMarketData struct {
	TotalMarketCapUsd            float64 `json:"total_market_cap_usd"`
	Total24hVolumeUsd            float64 `json:"total_24h_volume_usd"`
	BitcoinPercentageOfMarketCap float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies             int     `json:"active_currencies"`
	ActiveAssets                 int     `json:"active_assets"`
	ActiveMarkets                int     `json:"active_markets"`
}

func (m GlobalMarketData) String() string {
	return fmt.Sprintf("Total Market Cap: $%.2f\n", m.TotalMarketCapUsd) +
		fmt.Sprintf("Total Market Volume: $%.2f\n", m.Total24hVolumeUsd)
}

// CoinGraph struct
type CoinGraph struct {
	MarketCapByAvailableAupply [][]float64 `json:"market_cap_by_available_supply"`
	PriceBtc                   [][]float64 `json:"price_btc"`
	PriceUsd                   [][]float64 `json:"price_usd"`
	VolumeUsd                  [][]float64 `json:"volume_usd"`
}

// Market struct
type Market struct {
	Rank          int
	Exchange      string
	Pair          string
	Volume        int
	Price         float64
	PercentVolume float64
	Updated       bool
}
