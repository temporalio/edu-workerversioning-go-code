package loanprocess

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func LoanProcessingWorkflow(ctx workflow.Context, input CustomerInfo) (string, error) {
	logger := workflow.GetLogger(ctx)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 60,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var totalPaid int
	var err error

	// for workflow executions started before the change, send thank you before the loop
	// TODO Part B: Comment this out and uncomment the identical block below the loop
	var notifyConfirmation string
	err = workflow.ExecuteActivity(ctx, SendThankYouToCustomer, input).Get(ctx, &notifyConfirmation)
	if err != nil {
		return "", err
	}

	for period := 1; period <= input.NumberOfPeriods; period++ {

		chargeInput := ChargeInput{
			CustomerID:      input.CustomerID,
			Amount:          input.Amount,
			PeriodNumber:    period,
			NumberOfPeriods: input.NumberOfPeriods,
		}

		var chargeConfirmation string
		err = workflow.ExecuteActivity(ctx, ChargeCustomer, chargeInput).Get(ctx, &chargeConfirmation)
		if err != nil {
			return "", err
		}

		totalPaid += chargeInput.Amount
		logger.Info("Payment complete", "Period", period, "Total Paid", totalPaid)

		workflow.Sleep(ctx, time.Minute*1)
	}

	// for workflow executions started after the change, send thank you after the loop
	// TODO Part B: Uncomment this and comment out the identical block above the loop
	// var notifyConfirmation string
	// err = workflow.ExecuteActivity(ctx, SendThankYouToCustomer, input).Get(ctx, &notifyConfirmation)
	// if err != nil {
	// 	return "", err
	// }

	result := fmt.Sprintf("Loan for customer '%s' has been fully paid (total=%d)", input.CustomerID, totalPaid)
	return result, nil
}
