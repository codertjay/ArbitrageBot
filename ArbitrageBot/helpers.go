package ArbitrageBot

import (
	"github.com/ethereum/go-ethereum/common"
)

type TokenInfo struct {
	Address  common.Address
	Decimals int
	Symbol   string
	Name     string
}

func main() {

	//b, err := ioutil.ReadFile("tokens.json")
	//if err != nil {
	//	fmt.Println("Failed to read deployer account:", err)
	//	return
	//}
	//
	//key, err := keystore.DecryptKey(b, "password")
	//if err != nil {
	//	fmt.Println("Failed to decrypt key:", err)
	//	return
	//}
	//
	//client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	//if err != nil {
	//	fmt.Println("Failed to connect to the Ethereum client:", err)
	//	return
	//}
	//defer client.Close()
	//
	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	fmt.Println("Failed to get network ID:", err)
	//	return
	//}
	//
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	fmt.Println("Failed to get gas price:", err)
	//	return
	//}

}
