package main

import "time"
import "./vandp"

func main() {
	vandp.Dummy = time.Now().Nanosecond() < 100
}
