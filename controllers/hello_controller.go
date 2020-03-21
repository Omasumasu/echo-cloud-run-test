package controllers

import (
	"example.com/go-echo-cloud-run/usecases"
)

type HelloController struct{
	UseCase usecases.HelloUseCase
}

func New() HelloController {
	controller := HelloController{
		UseCase: usecases.New(),
	}
	return controller
}

func (controller *HelloController) Get_main() string{
	return controller.UseCase.Return_main_page()
}

func (controller *HelloController) Get_with_param(param string) string{
	return controller.UseCase.Return_param_page(param)
}

func (controller *HelloController) Get_goodbye() string {
	return controller.UseCase.Return_goodbye()
}