package common

type ShortGroup string

const (
	CDN            = ShortGroup("cdn")
	ECS            = ShortGroup("ecs")
	MessageService = ShortGroup("messageservice")
	POLARDB        = ShortGroup("polardb")
	SLS            = ShortGroup("sls")
	VPC            = ShortGroup("vpc")
)
