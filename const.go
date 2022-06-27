package main

const (
	// SatInBTC represents number of satoshis in 1 bitcoin
	SatInBTC = uint64(100000000)

	// MonitorTick interval in seconds
	MonitorTick = 10

	// PingTick interval in seconds
	PingTick = 60

	// Main network node address
	NetworkNode = "3AVTze8bR1SqqMKv3uLedrnqCuWpdU7GZwX"

	// AnoteFee is Anote regular fee amount
	AnoteFee = 100000

	// RewardFee is regular fee amount for sending mining rewards
	RewardFee = 5 * AnoteFee

	// SeedWordsURL contains all words for generating seed
	SeedWordsURL = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt"

	// AnoteNodeURL is an URL for Anote Node
	AnoteNodeURL = "https://nodes.aint.digital"

	// AnoteAddress is Anote smart contract address
	AnoteAddress = "3AVkEwYsZeooN1GEc81a66N2zmnKFw1ZxyB"

	// MasterNodeUrl is URL for master node
	MasterNodeUrl = "http://146.190.23.217:5000"
)
