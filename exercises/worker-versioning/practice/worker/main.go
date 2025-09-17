package main

import (
	"log"
	loanprocess "worker-versioning/exercises/worker-versioning/solution"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, loanprocess.TaskQueueName, worker.Options{
		DeploymentOptions: worker.DeploymentOptions{
			UseVersioning: true,
			Version: worker.WorkerDeploymentVersion{
				DeploymentName: "worker_versioning_demo",
				// TODO Part A: set BuildId and a DefaultVersioningBehavior
				BuildId: "",
			},
			DefaultVersioningBehavior: workflow.VersioningBehaviorUnspecified,
		},
	})

	w.RegisterWorkflow(loanprocess.LoanProcessingWorkflow)
	w.RegisterActivity(loanprocess.ChargeCustomer)
	w.RegisterActivity(loanprocess.SendThankYouToCustomer)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
