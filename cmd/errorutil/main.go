package main

import (
	"github.com/fenixvlabs/meshkit/cmd/errorutil/cli"
	"github.com/fenixvlabs/meshkit/utils"
)

func main() {
	utils.SetupLogrusFormatter()
	utils.SetupMeshkitLogger(false)
	cli.Execute()
}
