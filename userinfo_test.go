package goiguserid

import (
	"fmt"
)

func ExampleGetUserInfo() {
	user, err := GetUserInfo("instagram")
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Id)
	fmt.Println(user.Biography)
	// Output:
	// 25025320
	// Discovering — and telling — stories from around the world. Curated by Instagram’s community team.
}

func ExampleGetUserId() {
	fmt.Println(GetUserId("instagram"))
	// Output: 25025320 <nil>
}
