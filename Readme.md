In most Go apps, you'll install the following packages for authentication:

github.com/Azure/azure-sdk-for-go/sdk/azcore/to
github.com/Azure/azure-sdk-for-go/sdk/azidentity

## 1 Get CRedentials USing DEfault AZure Credentials

```
cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println(err)
	}
```

## 2 Create the client