# easytx
East Text Template Processing

[![GitHub version](https://badge.fury.io/gh/xelwarto%2Feasytx.svg)](http://badge.fury.io/gh/xelwarto%2Feasytx)

<dl>
  <dt>Author</dt><dd>Ted Elwartowski (<a href="mailto:xelwarto.pub@gmail.com">xelwarto.pub@gmail.com</a>)</dd>
  <dt>Copyright</dt><dd>Copyright © 2016 Ted Elwartowski</dd>
  <dt>License</dt><dd>Apache 2.0 - http://www.apache.org/licenses/LICENSE-2.0</dd>
</dl>

## Description

**easytx** is a simple text template processing script written in GO. **easytx** provides an easy command line interface for processing textual templates using a JSON data file.

Building an executable version of **easytx** provides a portable application that is easily distributed to different linux based hosts.

Template processing utilizes the standard GO text template package. Template file actions (Arguements and Pipelines) can be found here: https://golang.org/pkg/text/template

## Building easytx

### Build Requirements:

#### Installation of required packages

* Note: Once **easytx** has been built there is no requirement to install the GO software on other systems.

```bash
# Ubuntu
apt-get update && apt-get -y install golang

# RHEL/CentOS
yum makecache && yum -y install golang
```

### Build easytx

````bash
git clone git@github.com:xelwarto/easytx.git
easytx/build.sh
````

## Usage

````bash
easytx/bin/easytx --help
````

### Examples

#### Simple template example

````
# test.tmpl
First Name: {{.fname}}
Last Name:  {{.lname}}
Email:      {{.email}}
````

````
# test.json
{
  "fname": "John",
  "lname": "Doe",
  "email": "john.doe@gmail.com"
}
````

````bash
easytx/bin/easytx --tmpl=test.tmpl --file=test.json
````

#### Range template example

````
# test.tmpl
{{range .people}}{{.name}}
{{end}}
````

````
# test.json
{
  "people": [
    { "name": "John Doe" },
    { "name": "Jane Doe" }
  ]
}
````

````bash
easytx/bin/easytx --tmpl=test.tmpl --file=test.json
````

#### With template example

````
# test.tmpl
{{with .people}}{{range .names}}{{.name}}
{{end}}{{end}}
````

````
# test.json
{
  "people": {
    "names": [
      { "name": "John Doe" },
      { "name": "Jane Doe" }
    ]
  }
}
````

````bash
easytx/bin/easytx --tmpl=test.tmpl --file=test.json
````

#### Inline JSON example

````
# test.tmpl
{{with .people}}{{range .names}}{{.name}}
{{end}}{{end}}
````

````bash
easytx/bin/easytx --tmpl=test.tmpl --json='{ "people": {"names": [{ "name": "John Doe" },{ "name": "Jane Doe" }]} }'
````


### Environment Examples

* The JSON data file can contain environments that have data specific to an environment. Using the **--env** cli paramater allows **easytx** to select data for that specific environment.

#### Simple environment template example

````
# test.tmpl
{{range .servers}}{{.host}}:{{.port}}
{{end}}
````

````
# test.json
{
  "vagrant": {
    "servers": [
      { "host": "server1", "port": "1234" },
      { "host": "server2", "port": "6789" }
    ]
  },
  "dev": {
    "servers": [
      { "host": "server3", "port": "8080" },
      { "host": "server4", "port": "8080" }
    ]
  }
}
````

````bash
easytx/bin/easytx --tmpl=test.tmpl --file=test.json --env=vagrant
easytx/bin/easytx --tmpl=test.tmpl --file=test.json --env=dev
````
