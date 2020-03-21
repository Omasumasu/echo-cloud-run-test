package usecases

type HelloUseCase struct{
	test string
}

func New() HelloUseCase {
	return HelloUseCase{
		test: "Masuta",
	}
}


func (u *HelloUseCase) Return_main_page() string {
	return "Hello world " + u.test
}

func (u *HelloUseCase) Return_param_page(param string) string {
	return "Hello " + param + " " + u.test
}

func (u *HelloUseCase) Return_goodbye() string {
	return "Good Bye " + u.test
}

