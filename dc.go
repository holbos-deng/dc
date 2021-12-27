// Copyright © 2021 holbos Deng <2292861292@qq.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package dc

import (
	"errors"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type configuration struct {
	fileName string
	fileType string
	paths    []string
	err      string
}

var config configuration
var log logrus.Logger

// conf type result string
type conf struct {
	key   string
	value interface{}
}

// Get key: mode1.mode11.mode111
func (c conf) Get(key string) conf {
	i := 0
	var value interface{}
	keys := strings.Split(key, ".")
	keyLen := len(keys)

	if c.value == nil {
		value = viper.Get(keys[0])
		key = keys[0]
		i++
	} else {
		value = c.value
		key = c.key + "." + key
	}

	for ; i < keyLen; i++ {
		if i == keyLen {
			break
		}
		switch value.(type) {
		case map[string]interface{}:
			value, _ = value.(map[string]interface{})[keys[i]]
		default:
			log.Errorf("在配置文件中没有找到键: %s", key)
			return conf{}
		}
	}
	return conf{key: key, value: value}
}

// GetMust key: mode1.mode11.mode111
func (c conf) GetMust(key string, defaultValue string) conf {
	if r := c.Get(key); r.value == nil {
		r.value = defaultValue
		return r
	} else {
		return r
	}
}

func (c conf) Int() (r int) {
	var err error
	switch c.value.(type) {
	case int:
		r = c.value.(int)
	case string:
		r, err = strconv.Atoi(c.value.(string))
		if err != nil {
			log.Errorf("不能将 %v 转成int型, 获取 %s 配置信息时", c.value, c.key)
		}
	default:
		log.Errorf("不能将 %v 转成int型, 获取 %s 配置信息时", c.value, c.key)
	}
	return
}

func (c conf) Value() (r interface{}) {
	return c.value
}

func (c conf) Key() (r string) {
	return c.key
}

func New(filePath string) conf {
	filePath = filepath.ToSlash(filePath)
	fileDir, fileName := filepath.Split(filePath)
	fileType := filepath.Ext(fileName)[1:]
	fileName, fileType = fileName[:len(fileName)-len(fileType)-1], fileName[len(fileName)-len(fileType):]

	if fileType == "" {
		log.Panic("配置文件名必须包含扩展名")
	}

	config = configuration{fileName: fileName, fileType: fileType, paths: []string{".", fileDir}}

	viper.SetConfigName(config.fileName)
	viper.SetConfigType(config.fileType)

	for _, path := range config.paths {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.New("file not found")
		}
		panic("Fatal error config file: " + err.Error())
	}

	return conf{}
}
