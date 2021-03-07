# ocrserver

Simple small ocr demo developed by Golang, based on gin framework and [gosseract](https://github.com/otiai10/gosseract)

Support PostgreSql storage and Docker deployment.

## Install

```sh
docker build -t ocrserver .
docker-compose -f docker-compose.yml up -d world
```
