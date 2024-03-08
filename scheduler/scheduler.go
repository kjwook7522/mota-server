package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"mota-server/log"
)

var scheduler gocron.Scheduler

func Init() {
	var err error

	scheduler, err = gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
}

func RegisterJobs() {
	initLogSystemJob()
	// write other jobs...
}

func Start() {
	scheduler.Start()
	for _, job := range scheduler.Jobs() {
		err := job.RunNow()
		if err != nil {
			panic(err)
		}

	}
}

func Scheduler() gocron.Scheduler {
	return scheduler
}

func initLogSystemJob() {
	AtMidnight := gocron.NewAtTime(0, 0, 0)

	job, err := scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(AtMidnight)),
		gocron.NewTask(
			func() {
				if file := log.File(); file != nil {
					file.Close()
				}

				log.Init()
			},
		),
		gocron.WithName("every day log system"),
	)
	if err != nil {
		panic(err)
	}

	jobID := job.ID()
	jobName := job.Name()
	fmt.Println(fmt.Sprintf("[Job] %s(%s) register", jobName, jobID))
}
