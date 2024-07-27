package main

import (
	"fmt"
	"gin-layout/internal/data"
	"github.com/trumanwong/go-tools/robot"
	"log"
)

type AppService struct {
	data  *data.Data
	robot *robot.WorkWechatRobot
}

func newApp(
	dData *data.Data,
	robot *robot.WorkWechatRobot,
) *AppService {
	return &AppService{
		data:  dData,
		robot: robot,
	}
}

func main() {
	app := wireApp()
	err := app.data.Migrate()
	if err != nil {
		app.robot.SendText(&robot.SentTextRequest{
			Level:   robot.LevelError,
			Content: fmt.Sprintf("Migrate fail, error: %s", err.Error()),
			IsAtAll: true,
		})
		log.Fatal(err)
	}
	log.Println("Migrate success")
}
