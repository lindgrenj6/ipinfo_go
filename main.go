package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

//go:embed token.txt
var token string

type IpInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func main() {
	var requestedIp string
	if len(os.Args) > 1 {
		requestedIp = "/" + os.Args[1]
	}

	url := fmt.Sprintf("https://ipinfo.io%v?token=%v", requestedIp, strings.TrimSpace(token))
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		panic("error hitting ipinfo: " + err.Error())
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var ipinfo IpInfo
	err = json.Unmarshal(bytes, &ipinfo)
	if err != nil {
		panic(err)
	}

	t := table.NewWriter()
	fields := reflect.TypeOf(ipinfo)
	values := reflect.ValueOf(ipinfo)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i).Name
		value := values.Field(i)
		t.AppendRow(table.Row{field, value})
	}

	fmt.Println(t.Render())
}
