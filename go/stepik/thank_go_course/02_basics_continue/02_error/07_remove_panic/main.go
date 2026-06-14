/*
 * Необходимо добавить корректную обработку ошиибок через возврат её из функции.
 * Программа не должна завершаться паникой или exit(1) ни при каких обстоятельствах.
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// не удаляйте, они нужны для проверки
var _ = errors.As
var _ = reflect.Append
var _ = runtime.Gosched

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	acc, trans, err := parseInput()
	fmt.Print("-> ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(acc, trans)
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)
	if err != nil {
		return acc, []int{}, err
	}
	trans, err := parseTransactions(transSrc)
	if err != nil {
		return acc, []int{}, err
	}
	return acc, trans, nil
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return account{}, err
	}
	overdraft, err := strconv.Atoi(parts[1])
	if err != nil {
		return account{}, err
	}
	if overdraft < 0 {
		return account{}, fmt.Errorf("expect overdraft >= 0")
	}
	if balance < -overdraft {
		return account{}, fmt.Errorf("balance cannot exceed overdraft")
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			return []int{}, err
		}
		trans[idx] = t
	}
	return trans, nil
}
