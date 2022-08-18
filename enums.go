package mooncake

type MooncakeTypes string

const (
	ANY MooncakeTypes = "any"
)

type MooncakeLifetime int

const (
	LT_ANY_TIME MooncakeLifetime = iota
	LT_ONE_CALL
	LT_REPEAT
)
