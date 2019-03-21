package main

//go:generate sqlboiler --wipe --output "../../pkg/models" -c ../../sqlboiler.toml psql

import (
	"github.com/chuckleheads/builder-api/internal/cmd"
)

func main() {
	cmd.Execute()
}
