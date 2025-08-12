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
	FCV3                = ShortGroup("fcv3")
	EIP                 = ShortGroup("eip")
	KMS                 = ShortGroup("kms")
	MessageService      = ShortGroup("messageservice")
	NATGateway          = ShortGroup("natgateway")
	OSS                 = ShortGroup("oss")
	POLARDB             = ShortGroup("polardb")
	PRIVATELINK         = ShortGroup("privatelink")
	QUOTAS              = ShortGroup("quotas")
	SLB                 = ShortGroup("slb")
	SLS                 = ShortGroup("sls")
	Tair                = ShortGroup("tair")
	VPC                 = ShortGroup("vpc")
)
