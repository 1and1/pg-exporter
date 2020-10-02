package main

import (
	"os"

	// "k8s.io/gengo/generator"
	"k8s.io/gengo/args"
	"k8s.io/klog"

	"github.com/1and1/pg-exporter/gen/generator"
)

func main() {
	klog.InitFlags(nil)
	arguments := args.Default()
	if err := arguments.Execute(
		generator.NameSystems(),
		generator.DefaultNameSystem(),
		generator.Packages,
	); err != nil {
		klog.Errorf("Error: %v", err)
		os.Exit(1)
	}
	klog.Info("Completed successfully.")
}



