package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"svc-template/util"
)

var databaseConnection *firestore.Client

func createClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClientWithDatabase(ctx, util.Config.ProjectId, util.Config.FireStoreDatabase)
	if err != nil {
		util.SLogger.Errorf("failed to create firestore client: %v", err)
	}

	databaseConnection = client

	// it is the responsibility of the calling function to ensure that the connection is closed
	// defer client.Close()

	return client
}
