# Parser

## Tasks

### snapshot

Snapshot all of the tests.

```sh
ls -d tests/* | xargs -I '{}' go run snapshot/main.go -pkg="github.com/charlires/go-openapi-gen/getcomments/parser/{}" -op="./{}/snapshot.json"
```

