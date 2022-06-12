package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/anonutopia/gowaves"
	"github.com/mr-tron/base58"
)

func sendAnote(recipient string, amount int) {
	atr := &gowaves.AssetsTransferRequest{
		Recipient: recipient,
		Amount:    amount,
		Fee:       AnoteFee,
		Sender:    NodeAddress,
	}
	_, err := gowaves.WNC.AssetsTransfer(atr)
	if err != nil {
		log.Println(err.Error())
	}
}

func checkFlags() bool {
	init := flag.Bool("init", false, "Initializes your Anote Node.")
	flag.Parse()

	if *init {
		seedStr := ""
		seed, encoded := generateSeed()

		seedStr += fmt.Sprintf("export SEED='%s'\n", seed)
		seedStr += fmt.Sprintf("export ENCODED='%s'", encoded)

		f, _ := os.Create("seed")
		defer f.Close()
		f.Write([]byte(seedStr))

		return false
	}

	return true
}

func urlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return linesFromReader(resp.Body)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func generateSeed() (seed string, encoded string) {
	var words []string
	seed = ""
	encoded = ""

	lines, err := urlToLines(SeedWordsURL)
	if err != nil {
		log.Println(err.Error())
	}

	for _, line := range lines {
		words = append(words, line)
	}

	for i := 1; i <= 15; i++ {
		seed += words[getRandNum()]
		if i < 15 {
			seed += " "
		}
	}

	data := []byte(seed)
	encoded = base58.Encode(data)

	return seed, encoded
}

func getRandNum() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 2048
	rn := rand.Intn(max-min+1) + min
	return rn
}
