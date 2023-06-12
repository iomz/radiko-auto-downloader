package main

const (
	// DatetimeLayout for time strings from radiko
	DatetimeLayout = "20060102150405"
	// DefaultArea for radiko are
	DefaultArea = "JP13"
	// RetryDelaySecond for initial delay
	DefaultInitialDelaySeconds = 60
	// DefaultInterval to fetch the programs
	DefaultInterval = "168h"
	// Language for ID3v2 tags
	ID3v2LangJPN = "jpn"
	// DefaultMaxConcurrents
	MaxConcurrency = 64
	// MaxRetryAttempts for BackOffDelay
	MaxRetryAttempts = 8
	// The API for region full
	RegionFullAPI = "http://radiko.jp/v3/station/region/full.xml"
	// OutputDatetimeLayout for downloaded files
	OutputDatetimeLayout = "200601021504"
	// TZTokyo for time location
	TZTokyo = "Asia/Tokyo"

	// HTTP Headers
	// auth1 req
	UserAgentHeader        = "User-Agent"
	RadikoAreaIDHeader     = "X-Radiko-AreaId"
	RadikoAppHeader        = "X-Radiko-App"
	RadikoAppVersionHeader = "X-Radiko-App-Version"
	RadikoDeviceHeader     = "X-Radiko-Device"
	RadikoUserHeader       = "X-Radiko-User"
	// auth1 res
	RadikoAuthTokenHeader = "X-Radiko-AuthToken"
	RadikoKeyLentghHeader = "X-Radiko-KeyLength"
	RadikoKeyOffsetHeader = "X-Radiko-KeyOffset"
	// auth2 req
	RadikoConnectionHeader = "X-Radiko-Connection"
	RadikoLocationHeader   = "X-Radiko-Location"
	RadikoPartialKeyHeader = "X-Radiko-Partialkey"
)
