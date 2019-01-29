package main

import (
	"fmt"
	"golab/DB/dblab"
)

func main() {
	// -----------------------Insert--------------------------------------
	// err := dblab.InsertUserInfo()
	// if err != nil {
	// 	fmt.Printf("mysql-lab[InsertUserInfo] err: %+v", err)
	// }

	// -----------------------Select------------------------------------
	// userInfos, err := dblab.SelectUserInfo()
	// if err != nil {
	// 	fmt.Printf("mysql-lab[SelectUserInfo] err: %+v", err)
	// }
	// for i, v := range userInfos {
	// 	fmt.Printf("UserInfos[%d]: %v \n", i, *v)
	// }

	// ------------------------InsertIncludeThings----------------
	if err := dblab.InsertIncludeThings(); err != nil {
		fmt.Printf("mysql-lab[InsertUserInfo] err: %+v", err)
	}
}
