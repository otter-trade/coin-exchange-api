package engine

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/otter-trade/coin-exchange-api/config"
	"github.com/otter-trade/coin-exchange-api/exchanges/account"
	"github.com/otter-trade/coin-exchange-api/exchanges/orderbook"
	"github.com/otter-trade/coin-exchange-api/exchanges/ticker"
)

// Const vars for websocket
const (
	WebsocketResponseSuccess = "OK"
	restIndexResponse        = "<html>GoCryptoTrader RESTful interface. For the web GUI, please visit the <a href=https://github.com/otter-trade/coin-exchange-api/blob/master/web/README.md>web GUI readme.</a></html>"
	DeprecatedName           = "deprecated_rpc"
	WebsocketName            = "websocket_rpc"
)

var (
	wsHub              *websocketHub
	wsHubStarted       bool
	errNilRemoteConfig = errors.New("received nil remote config")
	errNilPProfConfig  = errors.New("received nil pprof config")
	errNilBot          = errors.New("received nil engine bot")
	errEmptyConfigPath = errors.New("received empty config path")
	errServerDisabled  = errors.New("server disabled")
	errAlreadyRunning  = errors.New("already running")
	// ErrWebsocketServiceNotRunning occurs when a message is sent to be broadcast via websocket
	// and its not running
	ErrWebsocketServiceNotRunning = errors.New("websocket service not started")
)

var (
	errExchangeNotLoaded       = errors.New("exchange is not loaded/doesn't exist")
	errExchangeNotEnabled      = errors.New("exchange is not enabled")
	errExchangeBaseNotFound    = errors.New("cannot get exchange base")
	errInvalidArguments        = errors.New("invalid arguments received")
	errExchangeNameUnset       = errors.New("exchange name unset")
	errCurrencyPairUnset       = errors.New("currency pair unset")
	errInvalidTimes            = errors.New("invalid start and end times")
	errAssetTypeDisabled       = errors.New("asset type is disabled")
	errAssetTypeUnset          = errors.New("asset type unset")
	errDispatchSystem          = errors.New("dispatch system offline")
	errCurrencyNotEnabled      = errors.New("currency not enabled")
	errCurrencyNotSpecified    = errors.New("a currency must be specified")
	errCurrencyPairInvalid     = errors.New("currency provided is not found in the available pairs list")
	errNoTrades                = errors.New("no trades returned from supplied params")
	errUnexpectedResponseSize  = errors.New("unexpected slice size")
	errNilRequestData          = errors.New("nil request data received, cannot continue")
	errNoAccountInformation    = errors.New("account information does not exist")
	errShutdownNotAllowed      = errors.New("shutting down this bot instance is not allowed via gRPC, please enable by command line flag --grpcshutdown or config.json field grpcAllowBotShutdown")
	errGRPCShutdownSignalIsNil = errors.New("cannot shutdown, gRPC shutdown channel is nil")
	errInvalidStrategy         = errors.New("invalid strategy")
	errSpecificPairNotEnabled  = errors.New("specified pair is not enabled")
)

// apiServerManager holds all relevant fields to manage both REST and websocket
// api servers
type apiServerManager struct {
	restStarted            int32
	websocketStarted       int32
	restListenAddress      string
	websocketListenAddress string
	gctConfigPath          string
	restHTTPServer         *http.Server
	websocketHTTPServer    *http.Server
	wgRest                 sync.WaitGroup
	wgWebsocket            sync.WaitGroup

	restRouter      *mux.Router
	websocketRouter *mux.Router
	websocketHub    *websocketHub

	remoteConfig     *config.RemoteControlConfig
	pprofConfig      *config.Profiler
	exchangeManager  iExchangeManager
	bot              iBot
	portfolioManager iPortfolioManager
}

// websocketClient stores information related to the websocket client
type websocketClient struct {
	Hub              *websocketHub
	Conn             *websocket.Conn
	Authenticated    bool
	authFailures     int
	Send             chan []byte
	username         string
	password         string
	maxAuthFailures  int
	exchangeManager  iExchangeManager
	bot              iBot
	portfolioManager iPortfolioManager
	configPath       string
}

// websocketHub stores the data for managing websocket clients
type websocketHub struct {
	Clients    map[*websocketClient]bool
	Broadcast  chan []byte
	Register   chan *websocketClient
	Unregister chan *websocketClient
}

// WebsocketEvent is the struct used for websocket events
type WebsocketEvent struct {
	Exchange  string `json:"exchange,omitempty"`
	AssetType string `json:"assetType,omitempty"`
	Event     string
	Data      interface{}
}

// WebsocketEventResponse is the struct used for websocket event responses
type WebsocketEventResponse struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// WebsocketOrderbookTickerRequest is a struct used for ticker and orderbook
// requests
type WebsocketOrderbookTickerRequest struct {
	Exchange  string `json:"exchangeName"`
	Currency  string `json:"currency"`
	AssetType string `json:"assetType"`
}

// WebsocketAuth is a struct used for
type WebsocketAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Route is a sub type that holds the request routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// AllEnabledExchangeOrderbooks holds the enabled exchange orderbooks
type AllEnabledExchangeOrderbooks struct {
	Data []EnabledExchangeOrderbooks `json:"data"`
}

// EnabledExchangeOrderbooks is a sub type for singular exchanges and respective
// orderbooks
type EnabledExchangeOrderbooks struct {
	ExchangeName   string           `json:"exchangeName"`
	ExchangeValues []orderbook.Base `json:"exchangeValues"`
}

// AllEnabledExchangeCurrencies holds the enabled exchange currencies
type AllEnabledExchangeCurrencies struct {
	Data []EnabledExchangeCurrencies `json:"data"`
}

// EnabledExchangeCurrencies is a sub type for singular exchanges and respective
// currencies
type EnabledExchangeCurrencies struct {
	ExchangeName   string         `json:"exchangeName"`
	ExchangeValues []ticker.Price `json:"exchangeValues"`
}

// AllEnabledExchangeAccounts holds all enabled accounts info
type AllEnabledExchangeAccounts struct {
	Data []account.Holdings `json:"data"`
}

type wsCommandHandler struct {
	authRequired bool
	handler      func(client *websocketClient, data interface{}) error
}
