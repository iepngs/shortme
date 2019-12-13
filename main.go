package main

import (
	shortme "shortme/app"
)

func main() {
	a := shortme.App{}
	a.Initialize(shortme.GetEnv())
	a.Run(":8000")
}
