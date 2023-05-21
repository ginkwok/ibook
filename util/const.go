package util

import "io/fs"

type key string

const (
	LOGGER_KEY key = "logger"
	MYSQL_KEY  key = "mysql"
)

const (
	PERM_OF_DIR  fs.FileMode = 0775
	PERM_OF_FILE fs.FileMode = 0644
)

const (
	ResvStatusUnsignin  = "unsignin"
	ResvStatusSignined  = "signined"
	ResvStatusSignouted = "signouted"
	ResvStatusCancelled = "cancelled"
)

var (
	NoticeBeforeStart = `Dear user, Your seat reservation will begin within 15 minutes. Please signin promptly.`
	NoticeAfterStart  = `Dear user, Your seat reservation has arrived at the start time. Please signin promptly.`
	NoticeCancel      = "Dear user, your seat reservation has exceeded the start time by 30 minutes and has not been signed in. The system has automatically cancelled the reservation."
)
