package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func main() {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println(err)
	}

	client := armresources.NewResourceGroupsClient("b9005893-334c-45ae-b763-44f03fa20b6f", cred, nil)

	rglist, err := client.Get(context.TODO(), "example-resources", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*rglist.ID)

}
