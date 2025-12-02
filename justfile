_run year day:
    @go run main.go --year {{ year }} --day {{ day }}

_run_part year day part:
    @go run main.go --year {{ year }} --day {{ day }} --part {{ part }}

run day year="2025": (_run year day)

list:
    @go run main.go --list

bench day iterations="1000" $YEAR="2025":
    @go run main.go --year {{ YEAR }} --day {{ day }} --bench --iterations {{ iterations }}

test day year="2025":
    @go test  -v ./solutions/{{ year }}/day$(printf "%02d" {{ day }})/...

new day year="2025":
    @go run main.go --year {{ year }} --day {{ day }} --new

clean:
    go clean
    rm -rf advent-of-go

fmt:
    go fmt ./...

lint:
    go vet ./...
