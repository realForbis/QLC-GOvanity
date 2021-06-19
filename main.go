package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"

	"github.com/qlcchain/qlc-go-sdk/pkg/types"
	"github.com/urfave/cli"
)

var (
	iterations float64
)

const (
	addressEncodingAlphabet = "13456789abcdefghijkmnopqrstuwxyz"
)

func main() {
	app := cli.NewApp()
	app.Name = "vanity address generator for qlcchain"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "prefix, p",
			Usage: "Prefix to search for at the start of address",
		},
		cli.IntFlag{
			Name:  "count, n",
			Value: 1,
			Usage: "Number of valid addresses to generate before exiting, or 0 for infinite (default=1).",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "Do not output progress message.",
		},
	}
	app.Action = func(c *cli.Context) {

		iterations = estimatedIterations(c.String("prefix"))
		quiet := c.Bool("quiet")

		if !quiet {
			fmt.Println("Estimated number of iterations needed:", int(iterations))
		}
		for i := 0; i < c.Int("count") || c.Int("count") == 0; i++ {
			seed, addr, err := generateVanityAddress(c.String("prefix"), quiet)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println("Found matching address!")
			fmt.Println("seed = ", seed)
			fmt.Println("address = ", addr)
		}
	}
	app.Run(os.Args)
}

func estimatedIterations(prefix string) float64 {
	return math.Pow(32, float64(len(prefix))) / 2
}

// GenerateSeed generate seed
func GenerateSeed() (string, error) {
	seed, err := types.NewSeed()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(seed[:]), nil
}

//GenerateAccount generate account
func GenerateAccount() (string, types.Address, error) {
	seed, err := GenerateSeed()
	if err != nil {
		fmt.Println(err)
	}

	_, priv, err := types.KeypairFromSeed(seed, 0)
	if err != nil {
		fmt.Println(err)
	}

	acc := types.NewAccount(priv)
	addr := acc.Address()

	return seed, addr, nil
}

func isValidPrefix(prefix string) bool {
	for _, char := range prefix {
		if !strings.Contains(addressEncodingAlphabet, string(char)) {
			return false
		}
	}
	return true
}

func generateVanityAddress(prefix string, quiet bool) (string, string, error) {
	if !isValidPrefix(prefix) {
		return "", "", fmt.Errorf("Invalid character in prefix")
	}

	c := make(chan string, 100)
	progress := make(chan int, 100)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(c chan string, progress chan int) {
			defer func() {
				recover()
			}()
			count := 0
			for {
				count++
				if count%(500+i) == 0 {
					progress <- count
					count = 0
				}

				seed, addr, _ := GenerateAccount()
				s := fmt.Sprintf("%s", addr)
				if strings.HasPrefix(s[5:], prefix) {
					c <- seed
					break
				}
			}
		}(c, progress)
	}

	go func(progress chan int) {
		total := 0
		fmt.Println()
		for {
			count, ok := <-progress
			if !ok {
				break
			}
			total += count
			if !quiet {
				fmt.Printf("\033[1A\033[KTried %d (~%.2f%%)\n", total, float64(total)/iterations*100)
			}
		}
	}(progress)

	seed := <-c
	pub, _, err := types.KeypairFromSeed(seed, 0)
	if err != nil {
		fmt.Println(err)
	}

	account := types.PubToAddress(pub)
	address := fmt.Sprintf("%s", account)

	close(c)
	close(progress)

	return seed, address, nil
}
