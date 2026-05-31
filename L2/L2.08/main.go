/* Создать программу, печатающую точное текущее время с использованием NTP-сервера.
Реализовать проект как модуль Go.
Использовать библиотеку ntp для получения времени.
Программа должна выводить текущее время, полученное через NTP (Network Time Protocol).
Необходимо обрабатывать ошибки библиотеки: в случае ошибки вывести её текст в STDERR и вернуть ненулевой код выхода.
Код должен проходить проверки (vet и golint), т.е. быть написан идиоматически корректно. */

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const DefaultNTPHost = "0.beevik-ntp.pool.ntp.org"

func GetCurrentTime() (time.Time, error) {
	return GetCurrentTimeFromHost(DefaultNTPHost)
}

func GetCurrentTimeFromHost(host string) (time.Time, error) {
	if host == "" {
		host = DefaultNTPHost
	}

	currTime, err := ntp.Time(host)
	if err != nil {
		return time.Time{}, err
	}
	return currTime, nil
}

func main() {
	currTime, err := GetCurrentTime()
	if err != nil { // Выводим ошибку в STDERR и выходим с ненулевым кодом
		fmt.Fprintf(os.Stderr, "Ошибка получения времени: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Точное время (NTP): %s\n", currTime.Format(time.RFC3339))
}