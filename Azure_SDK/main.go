package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func main() {
	// Create a default Azure credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Define the service URL
	accountName := "rggolangstorage"
	containerName := "test"
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	// Create a new service client
	serviceClient, err := azblob.NewClient(serviceURL, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new container client
	containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

	// List blobs in the container
	pager := containerClient.NewListBlobsFlatPager(&azblob.ListBlobsFlatOptions{})

	fmt.Println("All blobs in the container:")
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, blob := range resp.Segment.BlobItems {
			fmt.Println("Blob name:", *blob.Name)
		}
	}

	// Example of listing blobs with a prefix
	fmt.Println("\nBlobs with prefix 'test':")
	listBlobsFlatOptions(serviceClient, containerName, "test")
}

func listBlobsFlatOptions(client *azblob.Client, containerName string, prefix string) {
	// List the blobs in the container with a prefix
	pager := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Prefix: &prefix,
	})

	fmt.Println("List blobs with prefix:")
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		for _, blob := range resp.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}
