# LocVis
**In Development!**

Simple command line tool to display lines of code analytics

## Features
- Show LOC (Lines of Code) for the top files
- Displays paths to help accessing them directly
- Median

## Install
```sh
# Download binary

# Move to bin folder
sudo mv ./locvis /usr/bin
```


## Config
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

## Supported Languages:
- Python: python
- C: c
- Java: java
- C#: c#
- JavaScript: js
- TypeScript: ts
- Go: go
- Rust: rs