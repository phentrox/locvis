# LocVis
A simple command line tool for linux to display the top 3 source files with the most lines of code

## Supported Languages
- Java (Maven)

## Install
```sh
# Download binary

# Create
sudo mv ./locvis /usr/local/bin
```

## Build from Source and Install
    ./sh/build-and-upgrade.sh

## Local Config for Repositories
### Create file
    touch .locvism.yaml
### Yaml file
```yaml
---
# Programming Language
lang: java

# Top number of files with the highest count (preconfig = 10)
top: 20

# Directories to ignore
ignored-dirs:
  - autoconfigure
  - .ide
```
