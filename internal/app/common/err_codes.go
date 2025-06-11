package com

type AppErrorCode int

const (
	UnknownCode        AppErrorCode = 0
	InvalidBodyCode    AppErrorCode = 1
	InvalidRequestCode AppErrorCode = 2
	NotFoundCode       AppErrorCode = 3

	TokenUnknownCode      AppErrorCode = 100
	TokenUnauthorizedCode AppErrorCode = 101
	TokenExpiredCode      AppErrorCode = 102
	NoTokenCode           AppErrorCode = 103
	InvalidTokenCode      AppErrorCode = 104
	InvalidTokenTypeCode  AppErrorCode = 105

	PermissionsUnknown AppErrorCode = 120
	PermissionsDenied  AppErrorCode = 121

	IncorrectCredentialsCode AppErrorCode = 110
	UserAlreadyExistsCode    AppErrorCode = 111

	SpotUnknownCode           AppErrorCode = 300
	UserAlreadySubscribedCode AppErrorCode = 301
)
