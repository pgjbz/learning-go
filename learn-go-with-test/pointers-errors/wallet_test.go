package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw unsufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(0)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(10000))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "no enough balance")
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}


func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}


func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}