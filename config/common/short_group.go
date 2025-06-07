package common

type ShortGroup string

const (
	ACK                 = ShortGroup("ack")
	ACKONE              = ShortGroup("ackone")
	ALB                 = ShortGroup("alb")
	ALIDNS              = ShortGroup("alidns")
	CDN                 = ShortGroup("cdn")
	CLOUDMONITORSERVICE = ShortGroup("cloudmonitorservice")
	ECS                 = ShortGroup("ecs")
	KMS                 = ShortGroup("kms")
	MessageService      = ShortGroup("messageservice")
	OSS                 = ShortGroup("oss")
	POLARDB             = ShortGroup("polardb")
	PRIVATELINK         = ShortGroup("privatelink")
	SLS                 = ShortGroup("sls")
	Tair                = ShortGroup("tair")
	VPC                 = ShortGroup("vpc")
)
