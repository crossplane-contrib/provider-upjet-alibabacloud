package common

type ShortGroup string

const (
	ALIDNS              = ShortGroup("alidns")
	CDN                 = ShortGroup("cdn")
	CLOUDMONITORSERVICE = ShortGroup("cloudmonitorservice")
	ECS                 = ShortGroup("ecs")
	MessageService      = ShortGroup("messageservice")
	POLARDB             = ShortGroup("polardb")
	SLS                 = ShortGroup("sls")
	Tair                = ShortGroup("tair")
	VPC                 = ShortGroup("vpc")
)
