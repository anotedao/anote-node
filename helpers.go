package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/anonutopia/gowaves"
	"github.com/mr-tron/base58"
	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func sendAnote(amount uint64) error {
	var networkByte = byte(55)
	var nodeURL = AnoteNodeURL
	var assetBytes []byte

	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Current time in milliseconds
	ts := uint64(time.Now().Unix() * 1000)

	assetBytes = []byte{}

	asset, err := proto.NewOptionalAssetFromBytes(assetBytes)
	if err != nil {
		log.Println(err)
		return err
	}

	rec, err := proto.NewRecipientFromString(AnoteAddress)
	if err != nil {
		log.Println(err)
		return err
	}

	args := proto.Arguments{}
	arg := proto.StringArgument{
		Value: conf.OwnerAddress,
	}
	args = append(args, arg)

	fc := proto.FunctionCall{
		Default:   false,
		Name:      "distributeMinerReward",
		Arguments: args,
	}

	sp := proto.ScriptPayment{
		Amount: amount,
		Asset:  *asset,
	}

	sps := proto.ScriptPayments{}
	sps = append(sps, sp)

	// tr := proto.NewUnsignedTransferWithSig(sender, *asset, *assetW, uint64(ts), amount, RewardFee, proto.Recipient{Address: &rec}, nil)
	tr := proto.NewUnsignedInvokeScriptWithProofs(
		2,
		55,
		sender,
		rec,
		fc,
		sps,
		*asset,
		RewardFee,
		ts)

	err = tr.Sign(networkByte, sk)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: nodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Mined: %d\n", amount)

	return nil
}

func checkFlags() bool {
	init := flag.Bool("init", false, "Initializes your Anote Node.")
	flag.Parse()

	if *init {
		seedStr := ""
		seed, encoded := generateSeed()
		key, encKey := generateApiKey()
		ip := getIP()

		seedStr += fmt.Sprintf("export SEED='%s'\n", seed)
		seedStr += fmt.Sprintf("export ENCODED='%s'\n", encoded)
		seedStr += fmt.Sprintf("export KEY='%s'\n", key)
		seedStr += fmt.Sprintf("export KENCODED='%s'\n", encKey)
		seedStr += fmt.Sprintf("export PUBLICIP='%s'", ip)

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

func generateApiKey() (key string, encodedKey string) {
	key = generatePassword(15, 3, 2, 3)
	uhsr, err := gowaves.WNC.UtilsHashSecure(key)
	if err != nil {
		log.Println(err.Error())
	}
	if err != nil {
		log.Println(err.Error())
	}
	return key, uhsr.Hash
}

func getRandNum() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 2048
	rn := rand.Intn(max-min+1) + min
	return rn
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	// Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func getIP() string {
	url := "https://api.ipify.org?format=text"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return string(ip)
}

func balanceC() uint64 {
	abr, err := gowaves.WNC.AddressesBalanceConfirmations(NodeAddress, 2)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return uint64(abr.Balance)
}

func balance() uint64 {
	abr, err := gowaves.WNC.AddressesBalance(NodeAddress)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return uint64(abr.Balance)
}
