package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/viper"
	abciclient "github.com/tendermint/tendermint/abci/client"
	cfg "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/types"
)

var homeDir string

func init() {
	flag.StringVar(&homeDir, "tm-home", "", "Path to the tendermint config directory (if empty, uses $HOME/.tendermint)")
}

func main() {
	flag.Parse()
	if homeDir == "" {
		homeDir = os.ExpandEnv("$HOME/Projects/node/.tendermint")
	}
	config := cfg.DefaultValidatorConfig()

	// =================== STEP 1 =======================
	// we load in the tendermint core configuration files
	config.SetRoot(homeDir)

	viper.SetConfigFile(fmt.Sprintf("%s/%s", homeDir, "config/config.toml"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Reading config: %v", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Decoding config: %v", err)
	}
	if err := config.ValidateBasic(); err != nil {
		log.Fatalf("Invalid configuration data: %v", err)
	}
	gf, err := types.GenesisDocFromFile(config.GenesisFile())
	if err != nil {
		log.Fatalf("Loading genesis document: %v", err)
	}

	// fmt.Println("config :", config)
	// fmt.Println("genesis :", gf)

	// =================== STEP 2 =======================
	// create a database handle and use it to construct out ABCI app
	dbPath := filepath.Join(homeDir, "badger")
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Fatalf("Opening database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Closing database: %v", err)
		}
	}()
	app := NewKVStoreApplication(db)
	acc := abciclient.NewLocalCreator(app)

	// =================== STEP 3 =======================
	// construct a logger (tendermint custom logger)

	logger := tmlog.MustNewDefaultLogger(tmlog.LogFormatPlain, tmlog.LogLevelInfo, false)

	// =================== STEP 4 =======================
	// construct new node by using read config(genesis.json, config.toml)
	node, err := nm.New(config, logger, acc, gf)
	if err != nil {
		log.Fatalf("Creating node: %v", err)
	}

	// =================== STEP 5 =======================
	// start the node just created

	node.Start()
	defer func() {
		node.Stop()
		node.Wait()
	}()

	// =================== STEP 6 =======================
	// signal terminer when a user want to shutdown the node
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
