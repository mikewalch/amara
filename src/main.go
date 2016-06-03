package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	AccumuloVersion string
	HdfsRoot        string
	MarathonRoot    string
	User            string
}

type App struct {
	Id          string     `json:"id"`
	Instances   int        `json:"instances"`
	Cmd         string     `json:"cmd"`
	Constraints [][]string `json:"constraints"`
	Cpus        float32    `json:"cpus"`
	Uris        []string   `json:"uris"`
	User        string     `json:"user"`
}

func createApp(config Config, name string, instances int, cpus float32) App {
	cmd := "chmod +x start.sh; ./start.sh " + name
	hdfsAmara := config.HdfsRoot + "/amara"
	var uris [4]string
	uris[0] = hdfsAmara + "/accumulo-" + config.AccumuloVersion + "-bin.tar.gz"
	uris[1] = hdfsAmara + "/accumulo-site.xml"
	uris[2] = hdfsAmara + "/accumulo-env.sh"
	uris[3] = hdfsAmara + "/start.sh"
	fmt.Println(uris)
	constraints := [][]string{}
	if name == "tserver" {
		c := []string{"hostname", "UNIQUE"}
		constraints = append(constraints, c)
		fmt.Println(constraints)
	}
	app := App{"accumulo-" + name, instances, cmd, constraints, cpus, uris[0:4], config.User}
	return app
}

func launchApp(config Config, name string, instances int, cpus float32) {
	fmt.Println("Launching ", instances, " Accumulo ", name)
	app := createApp(config, name, instances, cpus)
	b, err := json.Marshal(app)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	response, err := http.Post(config.MarathonRoot+"/v2/apps", "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}

func destroyApp(config Config, name string) {
	req, err := http.NewRequest("DELETE", config.MarathonRoot+"/v2/apps/accumulo-"+name, nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}

func getConfig(configPath string) Config {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal("open config", err)
	}
	var config Config
	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	return config
}

func printConfigValue(configPath string, key string) {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal("open config", err)
	}

	var data map[string]interface{}
	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(&data)
	if err != nil {
		log.Fatal("decode data", err)
	}
	f.Close()
	fmt.Println(data[key])
}

func main() {

	configPath := os.Args[1]
	command := os.Args[2]

	if command == "config" {
		if len(os.Args) == 4 {
			key := os.Args[3]
			printConfigValue(configPath, key)
			os.Exit(0)
		}
		log.Fatal("Missing arguments")
	} else if command == "start" {
		config := getConfig(configPath)
		launchApp(config, "master", 1, 0.2)
		launchApp(config, "monitor", 1, 0.2)
		launchApp(config, "tserver", 3, 0.5)
		launchApp(config, "gc", 1, 0.2)
	} else if command == "destroy" {
		config := getConfig(configPath)
		for _, app := range []string{"gc", "tserver", "monitor", "master"} {
			destroyApp(config, app)
		}
	}
}
