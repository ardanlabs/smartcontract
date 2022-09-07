![Github CI/CD](https://img.shields.io/github/workflow/status/DeOne4eg/eth-unit-converter/Tests)
![Go Report](https://goreportcard.com/badge/github.com/DeOne4eg/eth-unit-converter)
![Repository Top Language](https://img.shields.io/github/languages/top/DeOne4eg/eth-unit-converter)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/DeOne4eg/eth-unit-converter)
![Github Repository Size](https://img.shields.io/github/repo-size/DeOne4eg/http-multiplexer)
![Github Open Issues](https://img.shields.io/github/issues/DeOne4eg/eth-unit-converter)
![Lines of code](https://img.shields.io/tokei/lines/github/DeOne4eg/eth-unit-converter)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/DeOne4eg/eth-unit-converter)
![GitHub contributors](https://img.shields.io/github/contributors/DeOne4eg/eth-unit-converter)

<img align="right" width="300px" src="./images/go.png">

# Ethereum Unit Converter

It is a simple library for conversion units between themselves. The library is written in Golang.

## Install
`go get github.com/DeOne4eg/eth-unit-converter`

## Usage

```
// first
amount := ethUnit.NewWei(big.NewInt(1000000000000000000))

fmt.Printf("Wei: %v\n", amount.Wei())
fmt.Printf("GWei: %v\n", amount.GWei())
fmt.Printf("Ether: %v\n", amount.Ether())

// second (the same of first)
amount, err := ethUnit.ParseUnit(big.NewFloat(1000000000000000000), "wei")
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Wei: %v\n", amount.Wei())
fmt.Printf("GWei: %v\n", amount.GWei())
fmt.Printf("Ether: %v\n", amount.Ether())
```

## Tests
+ run with `make test`