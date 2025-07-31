# golang_json_html_convert — JSON2HTML web server 

A small HTTP server that converts JSON into HTML using an HTML template.

## Table of contents
- [Features](#features)
- [Structure](#structure)
    - [Used patterns](#used-patterns)
    - [RFCs](#rfcs-and-standards-followed)
    - [Robustness](#robustness)
- [Run](#run)
    - [Run](#run)
    - [Docker](#docker)
- [Notes](#notes)
- [License](#license)

## Features
- `GET /` — simple form (`<textarea>`) to paste JSON.
- `POST /render` — accepts **raw `application/json`** or form field `json_input`.
- Missing keys use Go zero-values; unknown keys are ignored (via `encoding/json`).
- Response: `text/html; charset=UTF-8`.



## Structure
```
golang_json_html_convert/
├── cmd/
│   └── server/
│       └── main.go          # CLI + HTTP entry point
├── internal/
│   ├── handler/             # HTTP handlers and logic
│   │   └── handler.go
│   ├── model/               # Data structs
│   │   └── threat.go
│   └── render/              # Template loader
│       └── render.go
├── templates/
│   └── threat.html.tmpl     # HTML template file
```

## Used patterns
- Strategy Pattern (Gang of Four)
- Factory Method
- Single Responsibility Principle
- Facade pattern

## RFCs and standards followed

| Standard  | Description |
| ------------- | ------------- |
| RFC 7231	| HTTP/1.1 semantics (GET, POST, status codes) |
| RFC 8259	| JSON (application/json) format |
| W3C HTML5 |	HTML output and form structure |
| Go net/http |	Go's HTTP server behavior |

## Robustness
- Parse template **once** at startup; safe concurrent `Execute`.
- Limit request body size via `http.MaxBytesReader` (default 1 MiB).
- Strict method checks: `/` GET, `/render` POST.
- JSON/form decoding strategies; extra keys ignored; missing keys -> zero values.
- Server timeouts (read/write/idle).

---

## Run
```bash
go build -o golang_json_html_convert.exe ./cmd/server

golang_json_html_convert.exe -t templates/threat.html.tmpl -p 8080
# Open http://localhost:8080/
```

## Docker
```bash
docker build -t golang_json_html_convert:latest .

docker run --rm -p 8080:8080 golang_json_html_convert:latest -p 8080 -t templates/threat.html.tmpl
```

## Notes
- Only standard library is used (as per assignment).
- Form field name is `json_input`.

## License
[MIT](https://choosealicense.com/licenses/mit/)
