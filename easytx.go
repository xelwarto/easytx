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

var tmplFile string
var jsonFile string
var jsonMap interface{}
var jsonData interface{}
var envParam string

func init() {
  flag.StringVar(&tmplFile, "tmpl", "", "Template File")
  flag.StringVar(&jsonFile, "json", "", "JSON Data File")
	flag.StringVar(&envParam, "env", "", "Environment Parameter")
  flag.Parse()
}

func main() {
	file, err := ioutil.ReadFile(jsonFile)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading JSON file: %v\n", err)
    os.Exit(1)
  }

  err = json.Unmarshal(file, &jsonMap)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error parsing JSON file: %v\n", err)
    os.Exit(1)
  }
	jsonData := jsonMap.(map[string]interface{})

	if envParam != "" {
		jsonData = jsonData[envParam].(map[string]interface{})
	}

  tmpl, err := template.ParseFiles(tmplFile)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
    os.Exit(1)
  }
  err = tmpl.Execute(os.Stdout, jsonData)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
    os.Exit(1)
  }
}
