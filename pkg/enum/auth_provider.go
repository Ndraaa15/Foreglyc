package enum

type AuthProvider uint8

const (
	AuthProviderUnknown AuthProvider = 0
	AuthProviderBasic   AuthProvider = 1
	AuthProviderGoogle  AuthProvider = 2
)

var (
	AuthProviderMap = map[AuthProvider]string{
		AuthProviderBasic:  "Basic",
		AuthProviderGoogle: "Google",
	}
)

func (c AuthProvider) String() string {
	return AuthProviderMap[c]
}

func ValueOfAuthProvider(value string) AuthProvider {
	for k, v := range AuthProviderMap {
		if v == value {
			return k
		}
	}
	return AuthProviderUnknown
}

func (c AuthProvider) IsValid() bool {
	_, ok := AuthProviderMap[c]
	return ok
}
