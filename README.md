# 🎲random-standup🎲
Do you have awkward pauses in your standups because no one wants to give their
update? Why not have a defined order? To make it fair, why not also
🎲randomize🎲 that order!

## Usage

1. Create a team roster in a TOML file, following the format in `example.toml`:
```toml
[Subteam-1]
members = ["Alice", "Bob", "Carol", "David"]

[Subteam-2]
members = ["Erin", "Frank", "Grace", "Heidi"]
```

2. `go run Main.go example.toml`
