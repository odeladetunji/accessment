package main

import (
	"accessment.com/microservice/routers"
	commitManager "accessment.com/microservice/service/commit-manager"
)

func main() {
	//Cron Job
	cron := commitManager.CommitManager{}
	//commit runs every hour
	go cron.UpdateCommitEveryHour()

	//Router
	routers.Routers()
}
