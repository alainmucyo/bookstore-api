package environment

type Environment struct {
	Port      string
	DBURL     string
	JWTSecret string
}

func New(
	port string,
	DBURL string,
	JWTSecret string,
) *Environment {
	return &Environment{
		Port:      port,
		DBURL:     DBURL,
		JWTSecret: JWTSecret,
	}
}
