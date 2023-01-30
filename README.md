<img src="https://www.ga4gh.org/wp-content/themes/ga4gh-theme/gfx/GA-logo-horizontal-tag-RGB.svg" alt="GA4GH Logo" width="50%"/>

# GA4GH rnaget API demo client

This is a demo client implementation for the GA4GH [RNAget API](https://github.com/ga4gh-rnaseq/schema). The implementation consists in a command line application written in Go and covers the basic functionalities of the API.

## Installation instructions

Use one of the following methods to install the client.

### Binary install

Download the appropriate version for your platform from the GitHub releases page. Once downloaded, the binary can be run from anywhere. Just move it to a folder inside your `PATH` for convenient use.

### Build and install from source

If you are familiar with Go and have the Go toolkit installed, the following command will install the master branch of the client into your `$GOPATH`:

```
go get github.com/guigolab/rnaget-client
```

## Quickstart

To get the client usage run it without any argument:

```
$ rnaget-client
A demo client for the GA4GH RNAget API

Usage:
  rnaget-client [command]

Available Commands:
  continuous   Query continuous data
  expressions  Query expression data
  help         Help about any command
  projects     Query projects
  service-info Get service info
  studies      Query studies

Flags:
  -h, --help              help for rnaget-client
  -l, --location string   Server location (default "crg")
  -V, --verbose count     Verbosity
      --version           version for rnaget-client

Use "rnaget-client [command] --help" for more information about a command.
```

The program takes server information from a configuration file. The file must be called `.rnaget-client.yml` and be stored in the current working directory. If no local configuration file is found the client loads the default configuration corresponding to its version from the [GitHub repository](https://github.com/guigolab/rnaget-client) (which can also be used as a template). The default target location can be specified in the config. A different server can be specified using the `-l,--location` command line option. e.g.:

```
rnaget-client -l gtex ...
```

### Service-info query

You can check the status of the remote server via the `service-info` API endpoint using the corresponding command, e.g.:

```
 $ rnaget-client -l gtex service-info
   Host : gtexportal.org
 Status : 200 OK
Payload :
{
  "contactUrl": "https://gtexportal.org/home/contact",
  "description": "This service provides access to GTEx public RNA-Seq data.",
  "documentationUrl": "https://gtexportal.org/rnaget/docs",
  "environment": "prod",
  "id": "org.gtexportal.api.rnaget",
  "name": "GTEx Portal RNAget",
  "organization": {
    "name": "GTEx Project",
    "url": "https://gtexportal.org"
  },
  "supported": {
    "continuous": false,
    "expressions": true,
    "projects": true,
    "studies": true
  },
  "type": {
    "artifact": "rnaget",
    "version": "1.2.0"
  },
  "version": "1.0.0"
}
```

### Project queries

The client allows to get a list of projects, e.g.:

```
$ rnaget-client projects
   Host : genome.crg.cat
 Status : 200 OK
Payload :
[
  {
    "description": "Test project object used by RNAget compliance testing suite.",
    "id": "9c0eba51095d3939437e220db196e27b",
    "name": "RNAgetTestProject0",
    "version": "1.0"
  }
]
```

The `projects` command allows to filter the query by version, using a command line flag. Please see the command help for more details.

Projects can also be accessed by their identifier by passing it as the command argument, e.g.:

```
$ rnaget-client projects 9c0eba51095d3939437e220db196e27b
   Host : genome.crg.cat
 Status : 200 OK
Payload :
{
  "description": "Test project object used by RNAget compliance testing suite.",
  "id": "9c0eba51095d3939437e220db196e27b",
  "name": "RNAgetTestProject0",
  "version": "1.0"
}
```

Available additional query filters can be fetched with the following command:

```
$ rnaget-client projects filters
   Host : genome.crg.cat
 Status : 200 OK
Payload :
[
  {
    "description": "The name of the project",
    "fieldType": "string",
    "filter": "name"
  }
]
```

### Study queries

The client also allows getting a list of studies, e.g.:

```
$ rnaget-client studies
   Host : genome.crg.cat
 Status : 200 OK
Payload :
[
  {
    "description": "Test study object used by RNAget compliance testing suite.",
    "id": "f3ba0b59bed0fa2f1030e7cb508324d1",
    "name": "RNAgetTestStudy0",
    "parentProjectID": "9c0eba51095d3939437e220db196e27b",
    "version": "1.0"
  }
]
```

The `studies` command also allows to filter the query by version, using a command line flag. Please see the command help for more details.

Studies can be accessed by their identifier too, passing the id as the command argument, e.g.:

```
$ rnaget-client studies f3ba0b59bed0fa2f1030e7cb508324d1
   Host : genome.crg.cat
 Status : 200 OK
Payload :
{
  "description": "Test study object used by RNAget compliance testing suite.",
  "id": "f3ba0b59bed0fa2f1030e7cb508324d1",
  "name": "RNAgetTestStudy0",
  "parentProjectID": "9c0eba51095d3939437e220db196e27b",
  "version": "1.0"
}
```

Available additional query filters can be fetched with the following command:

```
$ rnaget-client studies filters 
   Host : genome.crg.cat
 Status : 200 OK
Payload :
[
  {
    "description": "The name of the study",
    "fieldType": "string",
    "filter": "name"
  }
]
```

### Expression queries

Expression data queries can be performed using the `expressions` command of the client.

The available expression formats can be looked up by using the following command:

```
$ rnaget-client expressions formats
   Host : genome.crg.cat
 Status : 200 OK
Payload :
[
  "tsv"
]
```

The format is a required parameter for `bytes` and `ticket` queries and can be specified as a command line flag for each of the corresponding client commands. If the format is not specified on the command line, the first available format from the server is used. 

`Bytes` queries return the actual content of the expression matrix. Tab separated content can be shown on the screen, while for `loom` content the user is required to specify an output file. An example query with `tsv` data from the `crg` server could be like the following:

```
$ rnaget-client expressions bytes
         Host : genome.crg.cat
       Status : 200 OK
 Content-Type : text/tab-separated-values
      Payload : 
...
```

And one with `loom` data from the `caltech` server, saving the matrix to the file `matrix.loom`:

```
$ rnaget-client -l caltech expressions bytes -o matrix.loom
         Host : felcat.caltech.edu
       Status : 200 OK
 Content-Type : application/vnd.loom
```

The `expressions` command allows to include several query filters and slicing options, by using command line flags. Please see the command help for more details.

Expressions can also be accessed by their identifier too, passing the id as the command argument, e.g.:

```
$ rnaget-client expressions bytes 8beada7b93d5e55aa557138b39c6f930
         Host : genome.crg.cat
       Status : 200 OK
 Content-Type : text/tab-separated-values
      Payload :
...
```

Available additional query filters can be fetched with the following command:

```
$ rnaget-client expressions filters 
   Host : genome.crg.cat
 Status : 200 OK
Payload :
[
  {
    "description": "The name of the expression",
    "fieldType": "string",
    "filter": "name"
  }
]
```

`Ticket` queries have the same command line interface as `bytes` queries while returning a ticket for the expression data instead, e.g.:

```
$ rnaget-client expressions ticket
   Host : genome.crg.cat
 Status : 200 OK
Payload :
{
  "fileType": "tsv",
  "headers": {
    "Authorization": "Bearer abcdefuvwxyz"
  },
  "studyID": "f3ba0b59bed0fa2f1030e7cb508324d1",
  "units": "TPM",
  "url": "https://genome.crg.cat/rnaget/expressions/ac3e9279efd02f1c98de4ed3d335b98e/bytes",
  "version": "1.0"
}
```

### Continuous queries

Continuous data can also be queried via the `continuous` client command. The command looks very similar to the `expressions` one, just having few different command line flags used for slicing and filtering the response output. Please have a look a the command help for more details.