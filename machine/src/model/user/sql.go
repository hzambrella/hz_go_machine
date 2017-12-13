package user

const (
	checkUserByNameSql = `
SELECT
    status
FROM
	machine_user_info
WHERE
	user_name=?
`

	checkLoginSql = `
SELECT
    userid,pass,status
FROM
	machine_user_info
WHERE
	user_name=?
	`
	getUserByIdSql = `
SELECT
	 userid,role_code,parentid,pass,mobile,
	 user_name,real_name,id_card,bank_card,status,
	 create_time,update_time
FROM
	machine_user_info
WHERE
	userid=?
	`
)
