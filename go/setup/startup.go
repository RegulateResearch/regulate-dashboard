package setup

import (
	"context"
	"frascati/comp/background"
	"frascati/config"
	"frascati/constants"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/typing"
)

type startupTask struct {
	taskMap   map[string]func() (any, exception.Exception)
	processor background.Processor
	services  services
}

func initStartupTask(services services, processor background.Processor) startupTask {
	return startupTask{
		services:  services,
		processor: processor,
		taskMap: map[string]func() (any, exception.Exception){
			"register admin": func() (any, exception.Exception) {
				ctx := typing.NewDictionaryContext(context.Background())
				adminConfig := config.GetAdminConfig()
				res, exc := services.auth.Register(ctx, entity.User{
					Email:    adminConfig.Email,
					Username: adminConfig.Username,
					Password: adminConfig.Password,
					Role:     constants.ROLE_ADMIN,
				})

				return res, exc
			},
		},
	}
}

func (st startupTask) execute() {
	for key := range st.taskMap {
		function := st.taskMap[key]
		st.processor.AddTask(key, function)
	}
}
