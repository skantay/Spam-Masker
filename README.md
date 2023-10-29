### Spam Masker

This repository contains a Go function that masks links in a text message by replacing them with asterisks. The function manipulates bytes directly, focusing on efficiency and adherence to specific rules:

- Do not use standard library functions or external packages.
- Mask only links starting with http://.
- Treat ONLY lowercase letters in link detection.

#### How to Use:

1. Clone the repository.
2. Run `go test` to validate the solution.

#### Author:
[Sanzhar Kantay]