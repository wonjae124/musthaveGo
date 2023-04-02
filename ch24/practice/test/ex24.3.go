package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func DepositAndWithdraw(account *Account) {
	// 뮤텍스는 고루틴 하나만 확보할 수 있다. 하나의 고루틴만 공유 자원에 접근하도록 제한한다.
	// lock으로 뮤텍스를 확보할 때까지 대기
	mutex.Lock()

	// defer을 통해, 오류가 발생하더라도 Unlock 메서드 호출 보
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value : %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
