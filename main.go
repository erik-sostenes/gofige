package main

import (
	"fmt"
	"os"

	"github.com/erik-sostenes/gofige/internal/cli"
)

const (
	logo = `
      _ __ _ __           _ _ __ _ __ __ _ __ _ __ _  
     /  __ __//_ ___ _   /   _/// //  __ __//     //
    /  / _ _  // _ _ \\ /  /__ / // // _ _    ___//   
   /  / /_ \\/  / _ \  \   _/// // // /_ \\ ___//_ v0.1.0
  /   \_ _///\  \_ _/  / //  / //   \_ / //     //
 /_ __ __ _/  \\ ___ //_//  /_//_ __ __ _/__ _ // 
 Go File Generator
	`
)

func main() {
	fmt.Println(logo)
	
	if err := cli.Execute(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	
}
