/*
 * Напишите логику работы с лицевым счетом в банке. У счета есть баланс (сумма на счете) и овердрафт (на какую сумму можно уйти в минус). К счету последовательно применяются транзакции списания и пополнения. Транзакции изменяют баланс счета, овердрафт не меняется.
 * Со счета нельзя списать больше, чем (баланс + овердрафт). Зачислить или списать нулевую сумму тоже нельзя.
 * Гарантируется, что сумма на счете, овердрафт и размер транзакции не выйдут за пределы типа int, как по отдельности, так и все вместе. Гарантируется, что овердрафт больше либо равен 0.
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// errInsufficientFunds сигнализирует,
// что на счете недостаточно денег,
// чтобы выполнить списание
var errInsufficientFunds = errors.New("insufficient funds")

// errInvalidAmount сигнализирует,
// что указана некорректная сумма транзакции
var errInvalidAmount = errors.New("invalid transaction amount")

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

// deposit зачисляет деньги на счет.
func (acc *account) deposit(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	acc.balance += amount
	return nil
}

// withdraw списывает деньги со счета.
func (acc *account) withdraw(amount int) error {
	if amount <= 0 {
		return errInvalidAmount
	}
	if acc.balance+acc.overdraft >= amount {
		acc.balance -= amount
		return nil
	}
	return errInsufficientFunds
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

type test struct {
	acc   account
	trans []int
}

var tests = map[string]test{
	"{100 10} [10 -50 20]":  {account{100, 10}, []int{10, -50, 20}},
	"{30 0} [-20 -10]":      {account{30, 0}, []int{-20, -10}},
	"{30 0} [-20 -10 -10]":  {account{30, 0}, []int{-20, -10, -10}},
	"{30 0} [-100]":         {account{30, 0}, []int{-100}},
	"{0 0} [10 20 30]":      {account{0, 0}, []int{10, 20, 30}},
	"{0 0} [10 -10 20 -20]": {account{0, 0}, []int{10, -10, 20, -20}},
	"{20 10} [-20 -10]":     {account{20, 10}, []int{-20, -10}},
	"{20 10} [-20 -10 -10]": {account{20, 10}, []int{-20, -10, -10}},
	"{0 100} [-20 -10]":     {account{0, 100}, []int{-20, -10}},
	"{0 30} [-20 -10]":      {account{0, 30}, []int{-20, -10}},
	"{0 30} [-20 -10 -10]":  {account{0, 30}, []int{-20, -10, -10}},
	"{70 30} [-100 100]":    {account{70, 30}, []int{-100, 100}},
	"{100 10} [10 0 20]":    {account{100, 10}, []int{10, 0, 20}},
}

func main() {
	var err error
	name, err := readString()
	if err != nil {
		log.Fatal(err)
	}
	testCase, ok := tests[name]
	if !ok {
		log.Fatalf("Test case '%v' not found", name)
	}
	for _, t := range testCase.trans {
		if t >= 0 {
			err = testCase.acc.deposit(t)
		} else {
			err = testCase.acc.withdraw(-t)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if err == nil {
		fmt.Println(testCase.acc)
	}
}

// readString считывает и возвращает строку из os.Stdin
func readString() (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
