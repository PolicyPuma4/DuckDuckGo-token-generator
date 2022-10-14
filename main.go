package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.design/x/clipboard"
)

type param struct {
	key   string
	value string
}

type response struct {
	Token string
}

func makeUrl(rawUrl string, params ...param) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	p := url.Values{}
	for _, param := range params {
		p.Add(param.key, param.value)
	}

	u.RawQuery = p.Encode()

	return u.String(), nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Print("Input address without @duck.com suffix, then press enter: ")
	var user string
	fmt.Scanln(&user)

	linkUrl, err := makeUrl("https://quack.duckduckgo.com/api/auth/loginlink", param{"user", user})
	check(err)

	linkResp, err := http.Get(linkUrl)
	check(err)

	if linkResp.StatusCode != 200 {
		log.Fatal(linkResp.Status)
	}

	fmt.Print("Input the passphrase that you were sent via email, then press enter: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	passphrase := scanner.Text()

	tokenUrl, err := makeUrl("https://quack.duckduckgo.com/api/auth/login", param{"otp", passphrase}, param{"user", user})
	check(err)

	tokenResp, err := http.Get(tokenUrl)
	check(err)

	if tokenResp.StatusCode != 200 {
		log.Fatal(tokenResp.Status)
	}

	body, err := io.ReadAll(tokenResp.Body)
	check(err)

	var data response
	json.Unmarshal(body, &data)

	clipboard.Write(clipboard.FmtText, []byte(data.Token))

	fmt.Println("Your token has been copied to your clipboard")

	fmt.Print("Press any key to continue . . . ")
	fmt.Scanln()
}
