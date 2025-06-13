package main

import (
	"fmt"
	"os"

	"github.com/MasonDye/CC-Attack-Rewrite/pkg/attack"
	"github.com/MasonDye/CC-Attack-Rewrite/pkg/config"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	attacker, err := attack.NewAttacker(cfg)
	if err != nil {
		fmt.Println("Failed to initialize attacker:", err)
		os.Exit(1)
	}

	attacker.StartAttack()
}
