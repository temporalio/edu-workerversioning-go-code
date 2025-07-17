# Exercise 1: Worker Versioning

During this exercise, you will

- Configure a Worker for Versioning
- Mark a Workflow as Pinned
- Move a Pinned Workflow to a new Version
- Cut over your running Workers
- Sunset an old deployment Version

Make your changes to the code in the `practice` subdirectory (look for
`TODO` comments that will guide you to where you should make changes to
the code). If you need a hint or want to verify your changes, look at
the complete version in the `solution` subdirectory.

## Part A: Configure a Worker for Versioning

1. Run `go run worker/main.go` in a terminal to start a Worker.

## Part B: Mark a Workflow as Pinned

1. This Workflow uses the `SendThankYouToCustomer` Activity to
   send a thank you message to the customer before charging
   them with the first loan payment, but this was a mistake.
   This Activity should run after the last payment. To fix this,
   edit the `workflow.go` file and move the five lines of code
   (which begin with the `var notifyConfirmation string` statement)
   related to that Activity from just before the loop to just
   after it.

## Part C: Move a Pinned Workflow to a new Version

1. Edit the `workflow_test.go` file and uncomment the two import
   statements near the top of the file, then implement the following
   in the `TestReplayWorkflowHistoryFromFile` function:
   - Create the Workflow Replayer

## Part D: Cut over your Running Workers

Just above the loop, where the `ExecuteActivity` call was prior to
the change, add a call to `GetVersion`:

```go
version := workflow.GetVersion(ctx, "MovedThankYouAfterLoop", workflow.DefaultVersion, 1)
```

## Part E: Sunset an old Deployment Version

Test test test

### This is the end of the exercise.
