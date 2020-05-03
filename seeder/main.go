// build ignore

package main

import "fmt"

func main() {

	AutoMigrates()

	fmt.Println(fmt.Sprintf("权限填充开始！！"))
	CreatePerms()
	fmt.Println(fmt.Sprintf("权限填充完成！！"))

	fmt.Println(fmt.Sprintf("管理角色填充开始！！"))
	CreateAdminRole()
	fmt.Println(fmt.Sprintf("管理角色填充完成！！"))

	fmt.Println(fmt.Sprintf("角色填充开始！！"))
	CreateRoles()
	fmt.Println(fmt.Sprintf("角色填充完成！！"))

	fmt.Println(fmt.Sprintf("管理员填充开始！！"))
	CreateAdminUser()
	fmt.Println(fmt.Sprintf("管理员填充完成！！"))

	fmt.Println(fmt.Sprintf("商户角色账号填充开始！！"))
	CreateTenantRoleAndUser()
	fmt.Println(fmt.Sprintf("商户角色账号填充完成！！"))

	fmt.Println(fmt.Sprintf("用户填充开始！！"))
	CreateUsers()
	fmt.Println(fmt.Sprintf("用户填充完成！！"))

	//fmt.Println(fmt.Sprintf("商户填充开始！！"))
	//CreateTenants()
	//fmt.Println(fmt.Sprintf("商户填充完成！！"))

}
