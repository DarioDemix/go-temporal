package workflow

import (
	"app/activity"
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	
	var spanishGreeting string
	err := workflow.ExecuteActivity(ctx, activity.GreetInSpanish, name).Get(ctx, &spanishGreeting)
	if err != nil {
		return "", err
	}

	return spanishGreeting, nil
}