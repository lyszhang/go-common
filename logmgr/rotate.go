/**
 * @Author: lyszhang
 * @Email: ericlyszhang@gmail.com
 * @Description:
 * @File:
 * @Version: 1.0.0
 * @Data:
 */

package logmgr

import (
	log "github.com/inconshreveable/log15"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"path"
	"time"
)

// NewRotate 日志归档
func NewRotate(dir, name string) *rotatelogs.RotateLogs {
	// 归档设置
	writer, err := rotatelogs.New(
		path.Join(dir, name),
		//// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		//rotatelogs.WithLinkName(path.Join(dir, name)),
		//
		//// WithRotationTime设置日志分割的时间，这里设置为一天分割一次
		//rotatelogs.WithRotationTime(time.Hour*24),

		// WithMaxAge设置文件清理前的最长保存时间，
		rotatelogs.WithMaxAge(time.Hour*24*30),
	)
	if err != nil {
		log.Error("config local file system for logger error", "err", err.Error())
		return nil
	}
	return writer
}
