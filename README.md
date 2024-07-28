# jsonenv
jsonenv is a go package to load environment variables from a json file. Supported are hashmapish json files.
Arrays will not be parsed and recognized. jsonenv sets environment variables based on JSON path of your configuraion file.

## Installation
```bash
go get -u github.com/sinemah/jsonenv
```
## Usage
### Config File
Create your configuration file in JSON format e.g. `config.json`.
The keys in the JSON file will be used to set the environment variables.
The values of the keys will be set as the values of the environment variables as strings.
```json
{
  "port": 1111,
  "port_as_string":"1110",
  "foo": {
    "bar": "baz"
  },
  "is_true": true
}
```

load any JSON file which can be parsed by json.Unmarshal.
The default configuration file name is `env.json`.
Also you can load multiple configuration files by passing the file names as arguments to the `Load` function.

Import the package and load your configuration file.
### Example
```go
// main.go

err := jsonenv.Load("config.json")
if err != nil {
	log.Fatal(err)
}

os.Getenv("port") // "1111"
os.Getenv("port_as_string") // "1110"
os.Getenv("foo.bar") // "baz"
os.Getenv("is_true") // "true"
```
