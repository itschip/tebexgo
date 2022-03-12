package tebexgo

func New(secret string) (s *Session) {
	s = &Session{
		Secret: secret,
	}
	
	return s
}