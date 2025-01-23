# Azure Blob Storage Example

This demonstrates how to list blobs in an Azure Blob Storage container using the Azure SDK for Go.

## Setup

In the folder, you will need to do the following setup:

1. Authenticate to Azure using Azure CLI and ensure that the user has the Storage Account Blob Contributor role.

2. Initialize a Go module:

    ```sh
    go mod init github.com/yourusername/azure-blob-storage-example
    ```

3. Install the required dependencies:

    ```sh
    go get github.com/Azure/azure-sdk-for-go/sdk/azidentity
    go get github.com/Azure/azure-sdk-for-go/sdk/storage/azblob
    ```