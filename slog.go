package slog

import(
	"fmt"
)

func init() {
	fmt.Println("init")
}

func LogInfo(entry string) {
	fmt.Println("Info", entry)
}