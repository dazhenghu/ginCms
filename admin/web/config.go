package web

import (
    "io/ioutil"
    "github.com/dazhenghu/util/fileutil"
    "path/filepath"
    "github.com/go-yaml/yaml"
)

type AdminConfig struct {
    Secret string `yaml:"secret"`
}

var AdminConf *AdminConfig = &AdminConfig{};

func init()  {
    rootPath, _ := fileutil.GetCurrentDirectory()
    adminCnfFile := filepath.Join(rootPath, "config/main.yaml")

    exists, err := fileutil.PathExists(adminCnfFile)
    if err == nil && exists {
        cnfBytes, _ := ioutil.ReadFile(adminCnfFile)
        yaml.Unmarshal(cnfBytes, AdminConf)
    }
}

