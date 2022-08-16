package main

import "github.com/HariKrizz/mypackage"

func main()  {
	mypackage.LogInfo("This is an Info Message")
	mypackage.LogWarning("This is a Warning Message!")
	mypackage.LogError("This code Broke!")
}