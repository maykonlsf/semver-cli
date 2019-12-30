# Semantic Versioning Tool

An easy to use CLI tool to manage your projects current version and its upgrades according to the Semantic Versioning specification.


## Getting Started

### Install
If you already have golang installed you can install by running the command:
```sh
go get -u github.com/maykonlf/semver-cli/cmd/semver
```

### check install
Check if the semver was instaled running the command:
```sh
semver 
```


### Init semver
To start managing your project versions enter into the project folder, then run:
```sh
semver init
```

This command will start the versioning based on release version v1.0.0. If you want to start with another version number you car run instead:
```sh
semver init  \
    --release [base release version] \
    [ --alpha [curent alpha number] ] \
    [ --beta [current beta number] ] \
    [ --rc [current release candiate number] ] \
    [--force] # to override an already initialized semver in the current directory.
```
