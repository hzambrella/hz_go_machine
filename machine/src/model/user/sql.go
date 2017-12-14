package user

// 格式：每五个字段换行。
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
	getUserBaseInfoByIdSql = `
SELECT
	 userid,role_code,parentid,mobile,
	 user_name,real_name,id_card,bank_card,status,
	 create_time,update_time,
    ifnull((SELECT user_name FROM machine_user_info WHERE userid=tb1.parentid),'无') as parent_name,
	ifnull((SELECT real_name FROM machine_user_info WHERE userid=tb1.parentid),'无') as parent_real_name
FROM
	machine_user_info as tb1
WHERE
	userid=?
	`

	getUserBalanceInfoByIdSql = `
SELECT
	 userid,cash,balance,withdraw,total,
	 create_time,update_time
FROM
	machine_promotor_money
WHERE
	userid=?
	`
	createUserBaseInfoSql = `
INSERT INTO 
	machine_user_info 
	(parentid,pass, mobile, user_name, real_name, 
		role_code ,id_card,bank_card,status) 
VALUES 
	(?,?,?,?,?,
	?,?,?,?);
	`

	createUserBalanceInfoSql = `
INSERT INTO 
	machine_promotor_money 
	(userid) 
VALUES 
	(?);
	`

	getParentNodeInfoSql = `
SELECT
	 userid,role_code,parentid,mobile,
	 user_name,real_name,id_card,bank_card,status,
	 create_time,update_time
FROM
	machine_user_info
WHERE
	userid=(SELECT parentid FROM machine_user_info WHERE userid=?)
	`

	updateUserBaseInfoSql = `
UPDATE
	machine_user_info
SET
	mobile=?,user_name=?,real_name=?,id_card=?,bank_card=?
WHERE
	userid=?
	`

	updateUserStatusSql = `
UPDATE
	machine_user_info
SET
	status=?
WHERE
	userid=?
	`

	getAllChildNodeInfoSql = `
SELECT
	 userid,role_code,parentid,mobile,user_name,
	 real_name,id_card,bank_card,status,create_time,
	 update_time,
	ifnull((SELECT user_name FROM machine_user_info WHERE userid=tb1.parentid),'无') as parent_name,
	ifnull((SELECT real_name FROM machine_user_info WHERE userid=tb1.parentid),'无') as parent_real_name
FROM
	machine_user_info as tb1
WHERE
	FIND_IN_SET(userid,queryAllChild(?))and userid!=?

	`

	getChildNodeWithRoleSql = `
SELECT
	 userid,role_code,parentid,mobile,user_name,
	 real_name,id_card,bank_card,status,create_time,
	 update_time,
	ifnull((SELECT user_name FROM machine_user_info WHERE userid=tb1.parentid),'无') as parent_name,
	ifnull((SELECT real_name FROM machine_user_info WHERE userid=tb1.parentid),'无') as parent_real_name
FROM
	machine_user_info as tb1
WHERE
	FIND_IN_SET(userid,queryAllChild(?)) and userid!=? and role_code=?

	`
)
