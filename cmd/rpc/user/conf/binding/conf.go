package binding

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

const (
	devConfFile = "dev_conf.yaml"
	proConfFile = "pro_conf.yaml"
	relPath     = "conf/file" // config yaml files' dir name
)

var (
	localConf  *LocalConfig
	remoteConf *RemoteConfig
	once       sync.Once
)

// LocalConfig below
type LocalConfig struct {
	Nacos NacosConfig `yaml:"nacos"`
}

type NacosConfig struct {
	IP        string `yaml:"ip"`
	Port      uint64 `yaml:"port"`
	Username  string `yaml:"username"`
	Namespace string `yaml:"namespace"`
	DataId    string `yaml:"data_id"`
	Group     string `yaml:"group"`
	LogDir    string `yaml:"log_dir"`
	CacheDir  string `yaml:"cache_dir"`
}

// RemoteConfig below
type RemoteConfig struct {
	Microservice MicroserviceConfig `yaml:"microservice"`
	Log          LogConfig          `yaml:"log"`
	Mysql        DBConfig           `yaml:"mysql"`
	Redis        DBConfig           `yaml:"redis"` // todo redis db type maybe int
}

type MicroserviceConfig struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
}

type LogConfig struct {
	LogLevel      string `yaml:"log_level"`
	LogFileName   string `yaml:"log_file_name"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxAge     int    `yaml:"log_max_age"`
	LogMaxBackups int    `yaml:"log_max_backups"`
}
type DBConfig struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

// GetLocalConf gets configuration instance
func GetLocalConf() *LocalConfig {
	once.Do(initLocalConf)
	return localConf
}

// GetRemoteConf gets configuration instance
func GetRemoteConf() *RemoteConfig {
	once.Do(loadConfigFromNacos)
	return remoteConf
}

func initLocalConf() {
	var confFileRelPath string
	switch LoadEnv().GetEnv() {
	case "dev":
		confFileRelPath = filepath.Join(relPath, devConfFile)
	case "pro":
		confFileRelPath = filepath.Join(relPath, proConfFile)
	}

	content, err := os.ReadFile(confFileRelPath)

	if err != nil {
		klog.Fatalf("failed to read local file: %s", err)
	}

	err = yaml.Unmarshal(content, &localConf)
	if err != nil {
		klog.Fatalf("parse yaml error - %v", err)
		panic(err)
	}
	if err = validator.Validate(localConf); err != nil {
		klog.Fatalf("validate config error - %v", err)
		panic(err)
	}
}

// loadConfigFromNacos loads configs including mysql, redis, consul configs and so on
// from nacos config center
func loadConfigFromNacos() {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(GetLocalConf().Nacos.Namespace),
		constant.WithLogDir(GetLocalConf().Nacos.LogDir),
		constant.WithCacheDir(GetLocalConf().Nacos.CacheDir),
		constant.WithUpdateCacheWhenEmpty(true),
		constant.WithUsername(GetLocalConf().Nacos.Username),
		constant.WithPassword(LoadEnv().GetEnvVar("nacos_password")),
	)

	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig(
			GetLocalConf().Nacos.IP,
			GetLocalConf().Nacos.Port,
		),
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	)
	if err != nil {
		klog.Fatalf("failed to create a nacos config client: %s", err)
	}

	content, err := configClient.GetConfig(
		vo.ConfigParam{
			DataId: GetLocalConf().Nacos.DataId,
			Group:  GetLocalConf().Nacos.Group,
		},
	)
	if err != nil {
		klog.Fatalf("failed to get config from remote: %s", err)
	}

	err = yaml.Unmarshal([]byte(content), &remoteConf)
	if err != nil {
		klog.Fatalf("failed to unmarshal remote config: %s", err)
	}
	// todo validate yaml
}

//	type LocalConfig struct {
//		envVar      string
//		Kitex    Kitex    `yaml:"kitex"`
//		MySQL    MySQL    `yaml:"mysql"`
//		Redis    Redis    `yaml:"redis"`
//		Registry Registry `yaml:"registry"`
//	}

//type MySQL struct {
//	DSN string `yaml:"dsn"`
//}
//
//type Redis struct {
//	Address  string `yaml:"address"`
//	Username string `yaml:"username"`
//	Password string `yaml:"password"`
//	DB       int    `yaml:"db"`
//}
//
//type Kitex struct {
//	Service         string `yaml:"service"`
//	Address         string `yaml:"address"`
//	EnablePprof     bool   `yaml:"enable_pprof"`
//	EnableGzip      bool   `yaml:"enable_gzip"`
//	EnableAccessLog bool   `yaml:"enable_access_log"`
//	LogLevel        string `yaml:"log_level"`
//	LogFileName     string `yaml:"log_file_name"`
//	LogMaxSize      int    `yaml:"log_max_size"`
//	LogMaxBackups   int    `yaml:"log_max_backups"`
//	LogMaxAge       int    `yaml:"log_max_age"`
//}
//
//type Registry struct {
//	RegistryAddress []string `yaml:"registry_address"`
//	Username        string   `yaml:"username"`
//	Password        string   `yaml:"password"`
//}
