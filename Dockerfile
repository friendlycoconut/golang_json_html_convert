FROM golang:1.22-alpine AS build

WORKDIR /src
COPY . .
RUN cd cmd/server && CGO_ENABLED=0 go build -o /out/golang_json_html_convert .

FROM scratch
COPY --from=build /out/golang_json_html_convert /golang_json_html_convert
COPY templates/threat.html.tmpl /templates/threat.html.tmpl
EXPOSE 8080
ENTRYPOINT ["/golang_json_html_convert", "-p", "8080", "-t", "/templates/threat.html.tmpl"]
