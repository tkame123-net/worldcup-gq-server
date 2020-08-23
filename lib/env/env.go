package env

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type LoadOption func(c *loadConf) error

type loadConf struct {
	searchFilename                     string
	searchAncestorDirectoryRecursively bool
}

func defaultLoadConf() *loadConf {
	return &loadConf{
		searchFilename:                     ".env",
		searchAncestorDirectoryRecursively: false,
	}
}

func LoadWithSearchAncestorDirectory() LoadOption {
	return func(c *loadConf) error {
		c.searchAncestorDirectoryRecursively = true
		return nil
	}
}

func Load(opts ...LoadOption) error {
	conf := defaultLoadConf()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return err
		}
	}

	// try to load from DOTENV_BODY variable
	dotenvBody := os.Getenv("DOTENV_BODY")
	if len(dotenvBody) > 0 {
		log.Println("DOTENV_BODY detected, try to load.")
		envMap, err := godotenv.Parse(bytes.NewBufferString(dotenvBody))
		if err != nil {
			return err
		}
		for k, v := range envMap {
			if err := os.Setenv(k, v); err != nil {
				return err
			}
		}
		return nil
	}

	// try to load from env file
	dir, err := os.Getwd()
	if err != nil {
		log.Print(err)
		// fail sort
		return nil
	}
	for {
		envpath := filepath.Join(dir, conf.searchFilename)
		envFile := os.Getenv("ENV_FILE")
		if len(envFile) > 0 {
			envpath = envFile
		}

		if _, err := os.Stat(envpath); err == nil {
			log.Println(".env configuration detected. try to load")
			if err := godotenv.Overload(envpath); err == os.ErrPermission {
				// fail soft
				log.Print(err)
			} else if err != nil {
				return err
			}
		} else if err == os.ErrNotExist {
			if !conf.searchAncestorDirectoryRecursively {
				break
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				// root directory
				break
			}
			dir = parent
			continue
		} else {
			// fail soft
			log.Print(err)
		}
		break
	}

	return nil
}
