# Regexer

## Overview
Regexer is a simple Go-based tool designed to search for specified keywords in the response bodies of web pages. It supports both single URL and bulk URL processing with concurrent execution.

## Features
- Scan a single URL or a list of URLs.
- Search for specific keywords in response bodies.
- Concurrent execution for improved performance.
- Output results to the console or a file.
- Customizable timeout settings.

## Installation
### Prerequisites
Ensure you have Go installed on your system. If not, install it from [golang.org](https://golang.org/).

### Build the Binary
```sh
# Clone the repository
git clone https://github.com/yourusername/regexer.git
cd regexer

# Build the executable
go build -o regexer
```

## Usage
### Scan a Single URL
```sh
./regexer -u https://example.com -w "keyword1,keyword2"
```

### Scan a List of URLs
```sh
./regexer -l urls.txt -w "keyword1,keyword2" -c 10 -o results.txt
```
- `-u` : Specify a single URL.
- `-l` : Path to a file containing URLs (one per line).
- `-w` : Comma-separated list of keywords to search for.
- `-c` : Number of concurrent workers (default: 10).
- `-o` : Output file to save results.

## Example Output
```
https://example.com contains: keyword1, keyword2
https://another.com contains: keyword3
```

## License
This project is licensed under the MIT License.

## Contributions
Contributions are welcome! Feel free to submit issues or pull requests.

