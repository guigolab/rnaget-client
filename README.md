# GA4GH rnaget API demo client

This is a demo client implementation for the GA4GH rnaget API. Please have a look at [this](https://github.com/ga4gh-rnaseq/schema/blob/master/TESTING.md) document for more information. The implementation consists in a command line application written in Go and covers the basic functionalities of the API.

## Installation instructions

Use one of the following methods to install the client.

### Binary install

Download the appropriate version for your platform from GitHub release. Once downloaded, the binary can be run from anywhere. Ideally, you should install it somewhere in your PATH for easy use.

### Build and install from source

If you are familiar with Go and have the Go toolkit installed, the following command will install the master branch of the client into your `$GOPATH`:

```
go get github.com/guigolab/rnaget-client/rnaget-client
```

## Quickstart

To get the client usage run it without any argument:

```
$ rnaget-client
A demo client for the GA4gh rnaget API

Usage:
  rnaget-client [command]

Available Commands:
  expression  Expression queries
  help        Help about any command
  project     Project queries
  study       Study queries

Flags:
  -h, --help              help for rnaget-client
  -l, --location string   Server location

Use "rnaget-client [command] --help" for more information about a command.
```

The client allows to run queries against the servers described in [this](https://github.com/ga4gh-rnaseq/schema/blob/master/TESTING.md) document. The default server is the `crg.cat` server. The target server can be changed using the `-l,--location` command line option. Allowed values for the option are `crg` and `caltech`, e.g.:

```
rnaget-client -l caltech ...
```

### Project queries

The client allows searching for projects, e.g.:

```
$ rnaget-client project search
[
  {
    "description": "The Pancancer Analysis of Whole Genomes (PCAWG) study is an international collaboration to identify common patterns of mutation in more than 2,800 cancer whole genomes from the International Cancer Genome Consortium.",
    "id": "E-MTAB-5200",
    "name": "Pancancer Analysis of Whole Genomes",
    "tags": [
      "bulk",
      "RNA-seq",
      "human",
      "cancer"
    ],
    "version": "1.0"
  },
  {
    "description": "Single-Cell Analysis of Human Pancreas Reveals Transcriptional Signatures  of Aging and Somatic Mutation Patterns",
    "id": "E-GEOD-81547",
    "name": "Single cell transcriptome analysis of human pancreas",
    "tags": [
      "single-cell",
      "RNA-seq",
      "human",
      "cancer"
    ],
    "version": "1.0"
  }
]
INFO[0000] Got 2 project(s)  
```

The `search` command allows to include query filters, with the following command line options:

```
  -t, --tags strings     Search for specific tags
  -v, --version string   Search for a specific version
``` 

Projects can also be accessed by their identifier. The `get` command accepts the project id as its only argument, e.g.:

```
$ rnaget-client project get E-MTAB-5200
{
  "description": "The Pancancer Analysis of Whole Genomes (PCAWG) study is an international collaboration to identify common patterns of mutation in more than 2,800 cancer whole genomes from the International Cancer Genome Consortium.",
  "id": "E-MTAB-5200",
  "name": "Pancancer Analysis of Whole Genomes",
  "tags": [
    "bulk",
    "RNA-seq",
    "human",
    "cancer"
  ],
  "version": "1.0"
}
INFO[0000] Got 1 project(s)  
```

### Study queries

The client also allows searching for studies, e.g.:

```
$ rnaget-client study search
[
  {
    "description": "PCAWG study",
    "id": "E-MTAB-5200-ST0",
    "name": "Pancancer Analysis of Whole Genomes",
    "parentProjectID": "E-MTAB-5200",
    "patientList": null,
    "sampleList": null,
    "tags": [
      "bulk",
      "RNA-seq",
      "human",
      "cancer"
    ],
    "version": "1.0"
  }
]
INFO[0000] Got 1 study(s)    
```

The `search` command allows to include query filters, with the following command line options:

```
  -t, --tags strings     Search for specific tags
  -v, --version string   Search for a specific version
``` 

Studies can be accessed by their identifier too. The `get` command accepts the study id as its only argument, e.g.:

Studies can be accessed by their identifier too, e.g.:

```
$ rnaget-client study get E-MTAB-5200-ST0
{
  "description": "PCAWG study",
  "id": "E-MTAB-5200-ST0",
  "name": "Pancancer Analysis of Whole Genomes",
  "parentProjectID": "E-MTAB-5200",
  "patientList": null,
  "sampleList": null,
  "tags": [
    "bulk",
    "RNA-seq",
    "human",
    "cancer"
  ],
  "version": "1.0"
}
INFO[0000] Got 1 study(s)
```

### Expression queries

The expression data can also be searched, e.g.:

```
$ rnaget-client expression search
[
  {
    "URL": "https://genome.crg.cat/rnaget/data/E-MTAB-5200.tpms.tsv",
    "fileType": "tsv",
    "id": "8beada7b93d5e55aa557138b39c6f930",
    "studyID": "E-MTAB-5200-ST0",
    "tags": [
      "bulk",
      "cancer",
      "human",
      "RNA-seq"
    ]
  }
]
INFO[0000] Got 1 file(s)  
```

The expression `search` command allows to include several query filters and slicing options, with the following command line flags:

```
  -a, --feature-accession strings   Slice by feature accession
  -f, --feature-id strings          Slice by feature id
  -n, --feature-name strings        Slice by feature name
  -p, --project-id string           Search for a specific project id
  -i, --sample-id strings           Slice by sample id
  -s, --study-id string             Search for a specific study id
  -t, --tags strings                Search for specific tags
  -v, --version string              Search for a specific version
``` 

And expressions can be accessed by their identifier too. The `get` command accepts the study id as its only argument, e.g.:

```
$ rnaget-client expression get 8beada7b93d5e55aa557138b39c6f930
{
  "URL": "https://genome.crg.cat/rnaget/data/E-MTAB-5200.tpms.tsv",
  "fileType": "tsv",
  "id": "8beada7b93d5e55aa557138b39c6f930",
  "studyID": "E-MTAB-5200-ST0",
  "tags": [
    "bulk",
    "cancer",
    "human",
    "RNA-seq"
  ]
}
INFO[0000] Got 1 file(s) 
```

#### Expression matrix slicing

> Work in progress