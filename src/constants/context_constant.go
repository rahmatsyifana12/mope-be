package constants

type ContextKey string

const (
	AuthClaimsKey   = "auth-claims"
	AccessToken     = "access-token"
	RequestIDKey    = ContextKey("requestid")
)