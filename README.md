[![Build](https://github.com/jidicula/random-standup/actions/workflows/build.yml/badge.svg)](https://github.com/jidicula/random-standup/actions/workflows/build.yml) [![Latest Release](https://github.com/jidicula/random-standup/actions/workflows/release-draft.yml/badge.svg)](https://github.com/jidicula/random-standup/actions/workflows/release-draft.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/jidicula/random-standup.svg)](https://pkg.go.dev/github.com/jidicula/random-standup)

# 🎲random-standup🎲
Do you have awkward pauses in your standups because no one wants to give their
update next? Why not have a defined order? To make it fair, why not also
🎲randomize🎲 that order!

### Do you find this useful?

Star this repo!

### Do you find this *really* useful?

You can sponsor me [here](https://github.com/sponsors/jidicula)!

## Usage

1. Get the tool with `go install github.com/jidicula/random-standup@latest`

2. Create a team roster in a TOML file, following the format in
`example-roster.toml`:
```toml
[Subteam-1]
members = [
        "Alice",                # TOML spec allows whitespace to break arrays
        "Bob",
        "Carol",
        "David"
        ]

["Subteam 2"]                   # Keys can have whitespace in quoted strings
members = ["Erin", "Frank", "Grace", "Heidi"]

["Empty Subteam"]               # Subteam with 0 members won't be printed

["Subteam 3"]
members = [
        "Ivan",
        "Judy",
        "Mallory",
        "Niaj"
]
```

3. `random-standup example-roster.toml`

## Output
```
$ random-standup example-roster.toml
# 2021-03-27
## Subteam-1
Alice
David
Bob
Carol

## Subteam 2
Grace
Heidi
Frank
Erin

## Subteam 3
Judy
Niaj
Ivan
Mallory
```

## Building from `main`

1. Clone and `cd` into the repo.
2. `go build -v`
3. `./random-standup example-roster.toml`

You can run tests with `go test -v -cover`.
