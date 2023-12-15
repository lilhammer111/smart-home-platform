package binding

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type EnvVar map[string]string

type ReadOnlyEnv struct {
	sync.RWMutex
	envVar EnvVar
	env    string
}

func newReadOnly(envVar EnvVar, env string) *ReadOnlyEnv {
	return &ReadOnlyEnv{
		envVar: envVar,
		env:    env,
	}
}

func (r *ReadOnlyEnv) GetEnvVar(key string) string {
	r.RLock()
	defer r.RUnlock()
	val, _ := r.envVar[key]
	return val
}

// GetEnv exposes the read-only environment variable which indicates
// whether the runtime is in development or production,
// the return value is either "dev" or "pro".
func (r *ReadOnlyEnv) GetEnv() string {
	return r.env
}

func LoadEnv() *ReadOnlyEnv {
	envVar, err := godotenv.Read(".env")
	if err != nil {
		log.Panicf("failed to read envVar from .envVar file: %s", err)
		// omit here: return nil, err
	}
	env := os.Getenv("PET_SERVICE_ENV")
	return newReadOnly(envVar, env)
}
