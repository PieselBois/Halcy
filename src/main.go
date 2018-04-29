package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alexflint/go-arg"
	"gopkg.in/yaml.v2"
)

var args struct {
	Debug   bool     `arg:"-d" help:"debug mode"`
	Verbose bool     `arg:"-v" help:"verbose mode"`
	Config  string   `arg:"-c" help:"set path for config file"`
	APIkey  string   `arg:"env"`
	Modules []string `arg:"-m" help:"set modules to load"`
	URL     string   `arg:"-u" help:"set server url"`
}

var cfg struct {
	CppCheck struct {
		CompileCommands string `yaml:"compile-commands"`
	}
}

func postResult(json []byte) {

	req, err := http.NewRequest("POST", args.URL, bytes.NewBuffer(json))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "applcation/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response Body:", string(body))
}

func main() {
	modules := map[string]module{
		"cppcheck": cppcheck{},
	}

	arg.MustParse(&args)

	data, err := ioutil.ReadFile(args.Config)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		log.Fatal(err)
	}

	warns := make([]warningInfo, 0)

	for _, mn := range args.Modules {
		m := modules[mn]
		wparse := m.warnings()
		warns = append(warns, wparse...)
	}
	j, _ := json.Marshal(warns)

	if args.Debug {
		fmt.Println(string(j))
	} else {
		postResult(j)
	}

}
