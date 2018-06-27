# Chronos

[![Go Report Card](https://goreportcard.com/badge/github.com/bagaking/chronos)](https://goreportcard.com/badge/github.com/bagaking/chronos)

## Usage
#### 1.requirement
go : ver >= 1.9

#### 2.install
> go get
```go get github.com/bagaking/chronos```
> or build
```git clone https://github.com/bagaking/chronos.git```
```cd chronos```
```make all```

#### 3.configuration
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

#### 4.run
```chronos```