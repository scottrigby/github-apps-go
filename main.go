package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte(os.Getenv("GITHUB_WEBHOOK_SECRET")))
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {
	case *github.PullRequestEvent:
		// Do something.
		fmt.Printf("Action: %s, Repository: %s\n",
			*e.Action, *e.Repo.FullName)

		// To-do: Imagine an example where we will need a client for Pull
		// Request events.
		client, err := getClient()
		if err != nil {
			log.Printf("error getting client: err=%s\n", err)
			return
		}

		// To-do: think of something better to do with GitHub API client.
		ctx := context.Background()
		owner := *e.Repo.Owner.Login
		repo := *e.Repo.Name
		number := *e.PullRequest.Number
		log.Printf("owner: %s, repo: %s, number: %v\n", owner, repo, number)

		merged, _, err := client.PullRequests.IsMerged(ctx, owner, repo, number)
		if err != nil {
			log.Printf("error getting merged status: err=%s\n", err)
			return
		}
		fmt.Printf("Merged status: %t\n", merged)
	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
		return
	}
}

// See google/go-github docs on Apps authentication:
// ref: https://github.com/google/go-github#authentication
func getClient() (*github.Client, error) {
	integrationID, err := strconv.Atoi(os.Getenv("GITHUB_APP_IDENTIFIER"))
	if err != nil {
		return nil, err
	}
	installationID, err := strconv.Atoi(os.Getenv("GITHUB_INSTALLATION_IDENTIFIER"))
	if err != nil {
		return nil, err
	}
	privateKeyFile := os.Getenv("GITHUB_PRIVATE_KEY_FILE")
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, integrationID, installationID, privateKeyFile)
	if err != nil {
		return nil, err
	}

	// Use installation transport with client.
	return github.NewClient(&http.Client{Transport: itr}), nil
}

func main() {
	fmt.Println("Listening on http://localhost:3000")
	fmt.Println("Use Ctrl-C to stop")
	http.HandleFunc("/webhooks", handleWebhook)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
