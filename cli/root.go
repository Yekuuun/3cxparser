package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "3cxparser",
	SilenceUsage: true,
	Short:        "A simple 3cx phone number parser for DID pool config",
}

var MainContext context.Context

func Execute() {
	// Code ANSI pour afficher du texte en bleu
	blue := "\033[34m"
	reset := "\033[0m"

	// Texte ASCII à afficher
	asciiArt := `
	  _____                                          
	 |___ /  _____  ___ __   __ _ _ __ ___  ___ _ __ 
	   |_ \ / __\ \/ / '_ \ / _` + "`" + ` | '__/ __|/ _ \ '__|
	  ___) | (__ >  <| |_) | (_| | |  \__ \  __/ |   
	 |____/ \___/_/\_\ .__/ \__,_|_|  |___/\___|_|   
					 |_|                             
		`

	// Affichage en bleu
	fmt.Println(blue + asciiArt + reset)

	var cancel context.CancelFunc
	MainContext, cancel = context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan:
			fmt.Println("\n[ERROR] Keyboard interrupt detected, terminating...")
			cancel()
		case <-MainContext.Done():
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		// Leaving this in results in the same error appearing twice
		// Once before and once after the help output. Not sure if
		// this is going to be needed to output other errors that
		// aren't automatically outputted.
		// fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\n\n")
	fmt.Print("Ciao....")
}