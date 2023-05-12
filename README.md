# CLI Command using Golang

Usage for those Laravel and NGIX user.

---

Only for Linux environment

1. Install [GO](https://go.dev/doc/install)
2. Clone this repo
3. Run
    ```
    go build
    ```
4. Run
    ```
    sudo mv nstool /usr/local/bin/nstool
    ```

---

## Features

### NGINX Configuration:

1. Able to display existing NGINX configuration file
2. Able to create NGINX configuration file using standard Laravel suggestion (require sudo)
3. Able to remove existing NGINX configuration file (require sudo)

### Laravel .env:

1. Able to copy .env from existing .env.example file or standard Laravel suggestion
2. Able to change value in .env file

---

NOTE:

1. This is meant for local and personal usage. Any usage on production we will not be responsible.
2. This repo purpose for educational only and exercise on Go language.
