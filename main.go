/*
**	write by chamber1ain 09.03.21
 */

package main

import (
	"fmt"
	nt "github.com/beevik/ntp"
	"time"
)

func main() {
	response, _ := nt.Query("0.beevik-ntp.pool.ntp.org")
	time := time.Now().Add(response.ClockOffset)
	fmt.Println(time)
}
