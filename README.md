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
yum makecache && yum -y install
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

#### Simple template Examples

````
# test.tmpl
First Name: {{.fname}}
Last Name:  {{.lname}}
Email:      {{.email}}
````

````
# test.json
{
  "fname": "Ted",
  "lname": "Elwartowski",
  "email": "xelwarto.pub@gmail.com"
}
````

````bash
easytx/bin/easytx --tmpl=test.tmpl --json=test.json
````
