// cmd/artificialmasterpiecegeneratorsystemsx/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratorsystemsx/internal/artificialmasterpiecegeneratorsystemsx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratorsystemsx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
