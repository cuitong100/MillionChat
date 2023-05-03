package user

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		timeout.Duration(2) * time.Second,
	}		
}

func (s *service) CreatedUser(c context.Context, req *CreatedUserReq) (*CreatedUserRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	u, err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}

