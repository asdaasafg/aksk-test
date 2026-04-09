package main

import (
	"fmt"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

func main() {
	credential := tos.NewStaticCredentials("AKLTOGRlNjNmNDQ5ZTU3NDM5MGFhY2Q4MTI3YTJmOWMxMzc", "WmpCbU1HTm1Oamt3WVRFMU5HRmpZVGd5WlRVM056YzBaak01TTJOaU9UVQ==")
	client, err := tos.NewClientV2(
		"https://tos-cn-beijing.volces.com",
		tos.WithCredentials(credential),
		tos.WithRegion("cn-beijing"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("TOS client initialized:", client)
}
