package wallet

import (
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if want != got {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func assertErr(t testing.TB, got, expected error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error, but none thrown")
	}
	if got.Error() != expected.Error() {
		t.Errorf("expected %s but got %s", expected.Error(), got.Error())
	}
}

func assertNoErr(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("expected error, but none thrown")
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(1))
		assertNoErr(t, err)
		assertBalance(t, wallet, Bitcoin(9))
	})
	t.Run("withdraw negative", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{}
		wallet.Deposit(startingBalance)
		err := wallet.Withdraw(Bitcoin(11))
		assertBalance(t, wallet, startingBalance)
		assertErr(t, err, ErrInsufficientFunds)
	})
}
