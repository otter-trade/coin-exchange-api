package engine

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/otter-trade/coin-exchange-api/common"
	"github.com/otter-trade/coin-exchange-api/currency"
	exchange "github.com/otter-trade/coin-exchange-api/exchanges"
	"github.com/otter-trade/coin-exchange-api/exchanges/asset"
	"github.com/otter-trade/coin-exchange-api/exchanges/currencystate"
)

func TestSetupCurrencyStateManager(t *testing.T) {
	t.Parallel()
	_, err := SetupCurrencyStateManager(0, nil)
	if !errors.Is(err, errNilExchangeManager) {
		t.Fatalf("received: '%v' but expected: '%v'", err, errNilExchangeManager)
	}

	cm, err := SetupCurrencyStateManager(0, &ExchangeManager{})
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	if cm.sleep != DefaultStateManagerDelay {
		t.Fatal("unexpected value")
	}
}

var (
	errManager  = errors.New("manager level error")
	errExchange = errors.New("exchange level error")
)

type fakeExchangeManagerino struct {
	ErrorMeOne bool
	ErrorMeTwo bool
}

func (f *fakeExchangeManagerino) GetExchanges() ([]exchange.IBotExchange, error) {
	if f.ErrorMeOne {
		return nil, errManager
	}
	return []exchange.IBotExchange{&fakerino{errorMe: f.ErrorMeTwo}}, nil
}

func (f *fakeExchangeManagerino) GetExchangeByName(_ string) (exchange.IBotExchange, error) {
	if f.ErrorMeOne {
		return nil, errManager
	}
	return &fakerino{errorMe: f.ErrorMeTwo}, nil
}

type fakerino struct {
	exchange.IBotExchange
	errorMe                bool
	GetAvailablePairsError bool
	GetBaseError           bool
}

func (f *fakerino) UpdateCurrencyStates(_ context.Context, _ asset.Item) error {
	if f.errorMe {
		return common.ErrNotYetImplemented
	}
	return nil
}

func (f *fakerino) GetAssetTypes(_ bool) asset.Items {
	return asset.Items{asset.Spot}
}

func (f *fakerino) GetName() string {
	return "testssssssssssssss"
}

func (f *fakerino) GetCurrencyStateSnapshot() ([]currencystate.Snapshot, error) {
	if f.errorMe {
		return nil, errExchange
	}
	return []currencystate.Snapshot{
		{Code: currency.SHORTY, Asset: asset.Spot},
	}, nil
}

func (f *fakerino) CanWithdraw(_ currency.Code, _ asset.Item) error {
	if f.errorMe {
		return errExchange
	}
	return nil
}

func (f *fakerino) CanDeposit(_ currency.Code, _ asset.Item) error {
	if f.errorMe {
		return errExchange
	}
	return nil
}

func (f *fakerino) CanTrade(_ currency.Code, _ asset.Item) error {
	if f.errorMe {
		return errExchange
	}
	return nil
}

func (f *fakerino) CanTradePair(_ currency.Pair, _ asset.Item) error {
	if f.errorMe {
		return errExchange
	}
	return nil
}

func (f *fakerino) GetAvailablePairs(_ asset.Item) (currency.Pairs, error) {
	if f.GetAvailablePairsError {
		return nil, errExchange
	}
	return currency.Pairs{currency.NewPair(currency.BTC, currency.USD)}, nil
}

func (f *fakerino) GetBase() *exchange.Base {
	if f.GetBaseError {
		return nil
	}
	return &exchange.Base{States: currencystate.NewCurrencyStates()}
}

func TestCurrencyStateManagerIsRunning(t *testing.T) {
	t.Parallel()
	err := (*CurrencyStateManager)(nil).Stop()
	if !errors.Is(err, ErrNilSubsystem) {
		t.Fatalf("received: '%v' but expected: '%v'", err, ErrNilSubsystem)
	}

	err = (&CurrencyStateManager{}).Stop()
	if !errors.Is(err, ErrSubSystemNotStarted) {
		t.Fatalf("received: '%v' but expected: '%v'", err, ErrSubSystemNotStarted)
	}

	err = (&CurrencyStateManager{started: 1, shutdown: make(chan struct{})}).Stop()
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	err = (*CurrencyStateManager)(nil).Start()
	if !errors.Is(err, ErrNilSubsystem) {
		t.Fatalf("received: '%v' but expected: '%v'", err, ErrNilSubsystem)
	}

	err = (&CurrencyStateManager{started: 1}).Start()
	if !errors.Is(err, ErrSubSystemAlreadyStarted) {
		t.Fatalf("received: '%v' but expected: '%v'", err, ErrSubSystemAlreadyStarted)
	}

	man := &CurrencyStateManager{
		shutdown:         make(chan struct{}),
		iExchangeManager: &fakeExchangeManagerino{ErrorMeOne: true},
		sleep:            time.Minute}
	err = man.Start()
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	time.Sleep(time.Millisecond)

	err = man.Stop()
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	man.iExchangeManager = &fakeExchangeManagerino{ErrorMeOne: true}
	err = man.Start()
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	time.Sleep(time.Millisecond)

	err = man.Stop()
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	man.iExchangeManager = &fakeExchangeManagerino{ErrorMeOne: true}
	err = man.Start()
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	time.Sleep(time.Millisecond)

	if !man.IsRunning() {
		t.Fatal("this should be running")
	}

	err = man.Stop()
	if !errors.Is(err, nil) {
		t.Fatalf("received: '%v' but expected: '%v'", err, nil)
	}

	if man.IsRunning() {
		t.Fatal("this should be stopped")
	}
}

func TestUpdate(_ *testing.T) {
	man := &CurrencyStateManager{}
	var wg sync.WaitGroup
	wg.Add(3)
	man.update(&fakerino{errorMe: true, GetAvailablePairsError: true}, &wg, asset.Items{asset.Spot})
	man.update(&fakerino{errorMe: true, GetBaseError: true}, &wg, asset.Items{asset.Spot})
	man.update(&fakerino{errorMe: true}, &wg, asset.Items{asset.Spot})
}
