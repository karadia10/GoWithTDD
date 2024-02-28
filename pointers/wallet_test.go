package pointers

import "testing"

func TestWallet(t *testing.T) {
  assertBalance:= func (t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got!=want {
      t.Errorf("got %s want %s", got, want)
    }

  }
  assertNoError:= func (t testing.TB, got error)  {
    t.Helper()
    if got != nil {
      t.Fatal("Error unexpected")
    }
  }
  assertError:= func (t testing.TB, err, msg error)  {
    t.Helper()
    if err == nil {
      t.Error("expected error")
    }
    if err != ErrInsufficientBalance {
      t.Errorf("expected error : %q but got : %q", err, msg)
    }
  }
  t.Run("testing deposit", func(t *testing.T) {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(20))
    assertBalance(t, wallet, Bitcoin(20))    
  })
  t.Run("withdraw", func(t *testing.T) {
    wallet := Wallet{Bitcoin(10)}
    err:=wallet.Withdraw(5)
    assertBalance(t, wallet, Bitcoin(5))
    assertNoError(t, err)
  })
  t.Run("withdraw mor than balance", func(t *testing.T) {
    wallet := Wallet{Bitcoin(10)}
    err:=wallet.Withdraw(20)
    assertBalance(t, wallet, Bitcoin(10))
    assertError(t, err, ErrInsufficientBalance)
  })
}
