package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	"github.com/stephenzsy/doc-locker/server/common/configurations"
)

func main() {
	newFs := flag.NewFlagSet("new", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'new' commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		newFs.Parse(os.Args[2:])
		// create new deployment Id
		deploymentId := time.Now().UTC().Format("20060102150405")
		serviceContext := app_context.NewInitializeAppServiceContext(context.Background(), app_context.WellKnownCallerdBootstrap, deploymentId).Elevate()
		rootDir, err := configurations.GetConfigurationsRootDir(serviceContext)
		if err != nil {
			log.Fatal(err)
		}
		deploymentDir := path.Join(rootDir, deploymentId)
		os.Mkdir(deploymentDir, 0700)
		os.Mkdir(path.Join(deploymentDir, "certs"), 0700)
		os.Mkdir(path.Join(deploymentDir, "envoy"), 0700)
		os.Mkdir(path.Join(deploymentDir, "certsk"), 0700)
		os.Mkdir(path.Join(deploymentDir, "scripts"), 0700)
		os.Mkdir(path.Join(deploymentDir, "server"), 0700)
		os.Mkdir(path.Join(deploymentDir, "site"), 0700)
		os.Mkdir(path.Join(deploymentDir, "tmp"), 0700)
		os.Remove(path.Join(rootDir, "latest"))
		os.Symlink(deploymentDir, path.Join(rootDir, "latest"))
		ioutil.WriteFile(path.Join(deploymentDir, "deployment-id"), []byte(deploymentId), 0600)
	default:
		log.Fatal("expected 'new' commands")
		os.Exit(1)
	}
}
