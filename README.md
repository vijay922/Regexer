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
git clone https://github.com/vijay922/Regexer.git
cd Regexer

# Build the executable
go build -o regexer
mv regexer /usr/local/bin/
```

## Usage
### Scan a Single URL
```sh
regexer -u https://example.com -w "apikey=,access_token=,secret=,auth=,password=,session=,jwt=,bearer=,Authorization=,Bearer,eyJ,AWS_ACCESS_KEY_ID,AWS_SECRET_ACCESS_KEY"
```

### Scan a List of URLs
```sh
regexer -l urls.txt -w "wp-content,wp-login,wp-admin,wp-includes,wp-json,xmlrpc.php,wordpress,wp-config,wp-cron.php" -c 10 -o results.txt
```
- `-u` : Specify a single URL.
- `-l` : Path to a file containing URLs (one per line).
- `-w` : Comma-separated list of keywords to search for.
- `-c` : Number of concurrent workers (default: 10).
- `-o` : Output file to save results.

## Example Output
```
https://bigbang-sf.example.com.ar contains: wp-content
https://example.com contains: wp-content, wp-admin, wp-includes, wp-json, xmlrpc.php, wordpress
https://ltw.example.com contains: wp-content, wp-admin, xmlrpc.php, wordpress
https://pin-library.example-hub.com contains: wordpress
https://yourcareer.example.com contains: wp-content, wp-admin, wp-includes, wp-json, xmlrpc.php
```

## License
This project is licensed under the MIT License.

## Contributions
Contributions are welcome! Feel free to submit issues or pull requests.

<h2 id="donate" align="center">⚡️ Support</h2>

<details>
<summary>☕ Buy Me A Coffee</summary>

<p align="center">
  <a href="https://buymeacoffee.com/vijay922">
    <img src="https://img.shields.io/badge/Buy%20Me%20a%20Coffee-ffdd00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black"/>
  </a>
</p>

</details>

<p align="center">
  <b><i>"Keep pushing forward. Never surrender."</i></b>
</p>

<p align="center">🌱</p>



```
Inspired by buggedout-1/regexer and modified with a few changes.
```
