/*
Copyright 2016 Ted Elwartowski <xelwarto.pub@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
  "fmt"
	"os"
  "io/ioutil"
  "encoding/json"
  "text/template"
)

var version = string("v0.2.0")

var disVersion bool
var tmplFile string
var jsonFile string
var jsonString string
var jsonMap interface{}
var jsonData map[string]interface{}
var envParam string

func init() {
	flag.BoolVar(&disVersion, "version", false, "Display version")
  flag.StringVar(&tmplFile, "tmpl", "", "Template File")
  flag.StringVar(&jsonFile, "file", "", "JSON Data File")
	flag.StringVar(&jsonString, "json", "", "JSON Input String")
	flag.StringVar(&envParam, "env", "", "Environment Parameter")
  flag.Parse()
}

func main() {
	if disVersion {
		fmt.Fprintf(os.Stdout, "Version: %v\n", version)
		os.Exit(1)
	}

	var err error
	data := []byte(jsonString)

	if jsonFile != "" {
		data, err = ioutil.ReadFile(jsonFile)
	  if err != nil {
	    fmt.Fprintf(os.Stderr, "Error reading JSON file: %v\n", err)
	    os.Exit(1)
	  }
	}

	if data != nil {
		err = json.Unmarshal(data, &jsonMap)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON file: %v\n", err)
			os.Exit(1)
		}
		jsonData = jsonMap.(map[string]interface{})

		if envParam != "" {
			jsonData = jsonData[envParam].(map[string]interface{})
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error reading JSON data: %v\n", err)
		os.Exit(1)
	}

  tmpl, err := template.ParseFiles(tmplFile)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading template file: %v\n", err)
    os.Exit(1)
  }
  err = tmpl.Execute(os.Stdout, jsonData)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error executing template file: %v\n", err)
    os.Exit(1)
  }
}
