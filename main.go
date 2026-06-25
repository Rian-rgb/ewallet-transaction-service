package main

import (
	"ewallet-transaction/cmd"
	"ewallet-transaction/helper"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	// load config
	helper.SetupConfig()

	// load log
	helper.SetupLogger()

	// load db
	helper.SetupPostgreSQL()

	//// run grpc
	//go cmd.ServeGRPC()

	// run http
	cmd.ServeHTTP()
}
