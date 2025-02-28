package auth

import (
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

const (
	defaultValidDuration = 6 * time.Hour
)

// AccessToken produces token signed with API key and secret
type AccessToken struct {
	apiKey   string
	secret   string
	grant    ClaimGrants
	validFor time.Duration
}

func NewAccessToken(key string, secret string) *AccessToken {
	return &AccessToken{
		apiKey: key,
		secret: secret,
	}
}

func (t *AccessToken) SetIdentity(identity string) *AccessToken {
	t.grant.Identity = identity
	return t
}

func (t *AccessToken) SetValidFor(duration time.Duration) *AccessToken {
	t.validFor = duration
	return t
}

func (t *AccessToken) SetName(name string) *AccessToken {
	t.grant.Name = name
	return t
}

func (t *AccessToken) AddGrant(grant *VideoGrant) *AccessToken {
	t.grant.Video = grant
	return t
}

func (t *AccessToken) SetMetadata(md string) *AccessToken {
	t.grant.Metadata = md
	return t
}

func (t *AccessToken) SetSha256(sha string) *AccessToken {
	t.grant.Sha256 = sha
	return t
}

func (t *AccessToken) ToJWT() (string, error) {
	if t.apiKey == "" || t.secret == "" {
		return "", ErrKeysMissing
	}

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(t.secret)},
		(&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	validFor := defaultValidDuration
	if t.validFor > 0 {
		validFor = t.validFor
	}

	cl := jwt.Claims{
		Issuer:    t.apiKey,
		NotBefore: jwt.NewNumericDate(time.Now()),
		Expiry:    jwt.NewNumericDate(time.Now().Add(validFor)),
		Subject:   t.grant.Identity,
	}
	return jwt.Signed(sig).Claims(cl).Claims(&t.grant).CompactSerialize()
}
