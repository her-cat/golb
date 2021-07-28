package global

import (
	"github.com/her-cat/golb/pkg/logger"
	"github.com/her-cat/golb/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
