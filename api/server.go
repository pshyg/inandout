package main

import (
	"fmt"

	"github.com/labstack/echo"
)

const (
	banner = `
   ____                    __  ____       __ 
  /  _/___    ___ ____ ___/ / / __ \__ __/ /_
 _/ / / _ \  / _ \/ _ | _  / / /_/ / // / __/
/___//_//_/  \_,_/_//_|_,_/  \____/\_,_/\__/

`
)

func initServer(e *echo.Echo) {
	fmt.Println(banner)

	if err := e.Start("0.0.0.0:2327"); err != nil {
		fmt.Println(err)
	}
}
