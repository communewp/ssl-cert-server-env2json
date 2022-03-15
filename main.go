package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"encoding/json"

	"github.com/caarlos0/env/v6"
)

type sslServerConfig struct {
	Listen      string                     `env:"SSL_LISTEN" envDefault:"0.0.0.0:8999" json:"listen"`
	PidFile     string                     `env:"SSL_PID_FILE" envDefault:"/ssl-cert-server.pid" json:"pid_file"`
	Storage     sslServerStorageConfig     `json:"storage"`
	LetsEncrypt sslServerLetsEncryptConfig `json:"lets_encrypt"`
}

type sslServerStorageConfig struct {
	Type     string                      `env:"SSL_STORAGE_TYPE" envDefault:"dir_cache" json:"type"`
	DirCache string                      `env:"SSL_STORAGE_DIR_CACHE" envDefault:"/data" json:"dir_cache,omitempty"`
	Redis    sslServerStorageRedisConfig `json:"redis,omitempty"`
}

type sslServerStorageRedisConfig struct {
	Addr string `env:"SSL_STORAGE_REDIS_ADDR" envDefault:"redis:6379" json:"addr"`
}

// TBD: sslServerManagedConfig

type sslServerLetsEncryptConfig struct {
	Staging       bool     `env:"SSL_LE_STAGING" envDefault:"false" json:"staging"`
	ForceRSA      bool     `env:"SSL_LE_FORCE_RSA" envDefault:"false" json:"force_rsa"`
	RenewBefore   int      `env:"SSL_LE_RENEW_BEFORE" envDefault:"30" json:"renew_before"`
	Email         string   `env:"SSL_LE_EMAIL,notEmpty" json:"email"`
	Domains       []string `env:"SSL_LE_DOMAINS" envSeparator:"," json:"domains,omitempty"`
	RegexPatterns []string `env:"SSL_LE_REGEX_PATTERNS" envSeparator:"," json:"re_patterns,omitempty"`
}

// TBD: sslServerSelfSignedConfig

func main() {

	outPtr := flag.String("o", "", "File to write the generated configuration to")
	execPtr := flag.String("e", "", "A command to run after config generation is complete")

	flag.Parse()

	cfg := sslServerConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Storage.Type == "dir_cache" {
		cfg.Storage.Redis.Addr = ""
	} else if cfg.Storage.Type == "redis" {
		cfg.Storage.DirCache = ""
	}

	out, err := json.Marshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	if *outPtr == "" {
		fmt.Print(string(out))
	} else {
		err = os.WriteFile(*outPtr, out, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *execPtr == "" {
		return
	}

	execArgs := strings.Split(*execPtr, " ")
	execCmd, execArgs := execArgs[0], execArgs[1:]
	cmd := exec.Command(execCmd, execArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
