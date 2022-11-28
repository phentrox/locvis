# LocVis
**In Development!**

A simple command line tool to display repository stats

## Supported Languages
- java
- go

## Features
- Shows LOC (Lines of Code) for the top files
- Displays paths to help accessing them directly

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
    touch .locvis.yaml
### Yaml file
```yaml
---
# Programming Language
lang: java

# Top number of files with the highest loc count (default = 10)
top: 20

# Show paths for files (default = true)
paths: true

# Additional Directories to ignore
ignore-dirs:
  - autoconfigure
  - .ide
```
