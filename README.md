# Chronos

[![Go Report Card](https://goreportcard.com/badge/github.com/bagaking/chronos)](https://goreportcard.com/badge/github.com/bagaking/chronos)

## Usage
--------
#### requirement
go : ver >= 1.9

#### install
```go get github.com/bagaking/chronos```

#### config
there should be a config file which named ```.kh.chronos.json```

> example:
```
  {
    "version": "1.0",
    "workers":[
      {
        "workername" : "echo g",
        "srcpath" : "./cmds/echog.sh",
        "timespan" : "10s" 
      },
      {
        "workername" : "echo d",
        "srcpath" : "./cmds/echod.sh",
        "timespan" : "4s" 
      }
    ]
  }
```

#### run
```chronos```