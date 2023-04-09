# LocVis
**In Development!**

A simple command line tool to display repository stats

## Features
- Shows LOC (Lines of Code) for the top files
- Displays paths to help accessing them directly
- Median

## Install
```sh
# Download binary

# Move to bin folder
sudo mv ./locvis /usr/bin
```


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

## Run
    locvis


## Dev
### Build from Source and Install (Linux)
    make build-and-replace
