package common

type ShortGroup string

const (
	ECS            = ShortGroup("ecs")
	MessageService = ShortGroup("messageservice")
	POLARDB        = ShortGroup("polardb")
	VPC            = ShortGroup("vpc")
)
