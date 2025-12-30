package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func HackerQuotes() string {
	quotes := []string{
		"Accessing mainframe...",
		"Bypassing firewall...",
		"Decrypting data packets...",
		"Establishing secure connection...",
		"Initializing quantum processor...",
		"Hacking the Gibson...",
	}
	return quotes[rand.Intn(len(quotes))]
}

func GenerateHackingLog(step int, total int) string {
	progress := float64(step) / float64(total)
	bars := int(progress * 40)

	bar := "[" + strings.Repeat("█", bars) + strings.Repeat("░", 40-bars) + "]"
	percentage := fmt.Sprintf("%.0f%%", progress * 100)

	return fmt.Sprintf("%s %s %s", bar, percentage, HackerQuotes())
}

func TypewriterSpeed() time.Duration {
	return time.Duration(20+rand.Intn(30)) * time.Millisecond
}

func RandomTerminalCommand() string {
	commands := []string {
		"ssh root@mainframe.corp",
		"sudo rm -rf /suspicion",
		"curl https://secrets.gov/data.json",
		"cat /etc/shadow | grep admin",
		"nmap -sV 192.168.1.1",
	}

	return commands[rand.Intn(len(commands))]
}