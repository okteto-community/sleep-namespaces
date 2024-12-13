package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"os/exec"

	"github.com/okteto-community/sleep-namespaces/sleeper/okteto"
)

const (
	// developmentNamespaceType defines the API type for namespaces used for development purposes (non-previews)
	developmentNamespaceType = "development"

	// namespaceSleepingStatus holds the status of a namespace that is sleeping
	namespaceSleepingStatus = "Sleeping"

	// sleepNSCommandTemplate is the command template to sleep a namespace
	sleepNSCommandTemplate = "okteto namespace sleep %s"
)

func main() {
	token := os.Getenv("OKTETO_TOKEN")
	oktetoURL := os.Getenv("OKTETO_URL")
	ctx := context.Background()

	logLevel := &slog.LevelVar{} // INFO
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	u, err := url.Parse(oktetoURL)
	if err != nil {
		logger.Error(fmt.Sprintf("Invalid OKTETO_URL %s", err))
		os.Exit(1)
	}

	client := getAPIClient(u.Host, token)

	// Retrieve all the namespaces of type "development"
	nsList, resp, err := client.NamespacesAPI.ListNamespaces(ctx).Type_(developmentNamespaceType).Execute()
	if err != nil {
		logger.Error(fmt.Sprintf("There was an error requesting the namespaces '%s'. Full HTTP response: %v", err, resp))
		os.Exit(1)
	}

	for _, ns := range nsList {
		logger.Info(fmt.Sprintf("Processing namespace '%s'", ns.Name))

		// We skip persistent namespaces and those that are already sleeping
		if ns.Persistent || ns.Status == namespaceSleepingStatus {
			logger.Info(fmt.Sprintf("Skipping namespace '%s', as it is persistent or already sleeping", ns.Name))
			logger.Info("-----------------------------------------------")
			continue
		}

		logger.Info(fmt.Sprintf("Sleeping namespace '%s'", ns.Name))

		output, err := sleepNamespace(ns.Name)
		if err != nil {
			logger.Error(fmt.Sprintf("Error sleeping namespace '%s': %s", ns.Name, err))
		} else {
			logger.Info(output)
		}

		logger.Info("-----------------------------------------------")
	}
}

// getAPIClient returns a new client for Okteto's API
func getAPIClient(apiHost, token string) *okteto.APIClient {
	config := okteto.NewConfiguration()
	config.Host = apiHost
	config.DefaultHeader["Authorization"] = fmt.Sprintf("Bearer %s", token)

	return okteto.NewAPIClient(config)
}

// sleepNamespace executes the Okteto CLI command to sleep a namespace
func sleepNamespace(nsName string) (string, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf(sleepNSCommandTemplate, nsName))

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
