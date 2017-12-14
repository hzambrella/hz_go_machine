-- 机器码分销项目

-- 所有金额的单位为厘(小数点后四位)

Drop table if exists machine_user_info;
-- 机器码分销用户表
CREATE TABLE machine_user_info (
    -- 用户id
    userid BIGINT(11) NOT NULL AUTO_INCREMENT,
    -- 用户角色
    role_code VARCHAR(32) NOT NULL,
     -- 用户父节点，0代表根。
    parentid BIGINT(20) NOT NULL DEFAULT '0',
    -- 密码
    pass VARCHAR(128) NOT NULL,
    -- 手机号
    mobile VARCHAR(32) NOT NULL,
    -- 用户名
    user_name VARCHAR(25)NOT NULL,
    -- 真实名字
    real_name VARCHAR(32)NOT NULL,
    -- 身份证号
    id_card VARCHAR(32) NOT NULL,
    -- 预设置的银行卡号
    bank_card VARCHAR(32) NOT NULL DEFAULT '0',
    -- 状态
    -- 0:正常  1:冻结或删除  2：审核中
    status VARCHAR(45) NOT NULL DEFAULT '0',
    -- 创建时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 更新时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (userid)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=100000 PARTITION BY HASH(userid DIV 1000000) PARTITIONS 10;
CREATE INDEX machine_user_info_idx0 ON machine_user_info(role_code);
CREATE INDEX machine_user_info_idx1 ON machine_user_info(parentid);
INSERT INTO `machine_user_info` (`userid`,`pass`, `mobile`, `user_name`, `real_name`, `role_code`, `parentid`,`id_card`) VALUES (1,'haozhao', '15671682392', 'haozhao', '开发者', 'ROOT', '0','0');
INSERT INTO `machine_user_info` (`userid`,`pass`, `mobile`, `user_name`, `real_name`, `role_code`, `parentid`,`id_card`) VALUES (2,'test_ckzx', '15671682392', 'test_ckzx', '测试者-创客中心', 'CKZX', '1','0');
INSERT INTO `machine_user_info` (`userid`,`pass`, `mobile`, `user_name`, `real_name`, `role_code`, `parentid`,`id_card`) VALUES (3,'test_consumer1', '15671682392', 'test_consumer1', '测试者-消费者1', 'CONSUMER', '2','0');
INSERT INTO `machine_user_info` (`userid`,`pass`, `mobile`, `user_name`, `real_name`, `role_code`, `parentid`,`id_card`) VALUES (4,'test_consumer2', '15671682392', 'test_consumer2', '测试者-消费者2', 'CONSUMER', '1','0');
ALTER TABLE machine_user_info auto_increment=100000;


Drop table if exists machine_role;
-- 用户角色表
CREATE TABLE machine_role (
    -- 角色信息码。系统管理员填写这个code,尽量简单明了，一目了然。
    role_code VARCHAR(8) NOT NULL,
    -- 名字
    role_name VARCHAR(48) NOT NULL,
    -- 描述
    detail VARCHAR(256) NOT NULL  DEFAULT '',
    -- 创建时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 更新时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 状态
    -- 0:正常  1:冻结或删除 
    status INT(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (role_code)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;
INSERT INTO `machine_role` (`role_code`, `role_name`, `detail`) VALUES ('ADMIN','平台管理员','系统管理员，可以查看所有角色，无需审核地创建大多数角色，查看所有信息');
INSERT INTO `machine_role` (`role_code`, `role_name`, `detail`) VALUES ('CB','消费商','消费商');
INSERT INTO `machine_role` (`role_code`, `role_name`, `detail`) VALUES ('CKZX','创客中心','创客中心');
INSERT INTO `machine_role` (`role_code`, `role_name`, `detail`) VALUES ('CONSUMER','消费者','消费者');
INSERT INTO `machine_role` (`role_code`, `role_name`, `detail`) VALUES ('ROOT','系统管理员','系统管理员，用来创建角色，行为，维护系统');

/*
Drop table if exists machine_action;
-- 行为表
CREATE TABLE machine_action (
    -- 行为码。系统管理员填写这个code,尽量简单明了，一目了然。 
    action_code VARCHAR(16) NOT NULL,
    -- 名字
    acrtion_name VARCHAR(48) NOT NULL,
    -- 描述
    detail VARCHAR(256) NOT NULL DEFAULT '',
    -- 创建时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 更新时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 状态
    -- 0:正常  1:冻结或删除 
    status INT(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (action_code)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;
INSERT INTO `machine_action` (`action_code`, `acrtion_name`, `detail`) VALUES ('ROOT', 'root', 'root行为');
INSERT INTO `machine_action` (`action_code`, `acrtion_name`, `detail`) VALUES ('ADMIN', '平台管理', '默认的平台管理行为');
INSERT INTO `machine_action` (`action_code`, `acrtion_name`, `detail`) VALUES ('S_CONSUME', '查看消费者', '查看消费者行为');
INSERT INTO `machine_action` (`action_code`, `acrtion_name`, `detail`) VALUES ('C_CONSUME', '创建消费者', '创建消费者行为');
INSERT INTO `machine_action` (`action_code`, `acrtion_name`, `detail`) VALUES ('U_CONSUME', '更改消费者', '更改消费者行为');
INSERT INTO `machine_action` (`action_code`, `acrtion_name`, `detail`) VALUES ('D_CONSUME', '删除消费者', '删除消费者行为');
*/

Drop table if exists machine_role_act;
-- 角色-行为表
CREATE TABLE machine_role_act (
    -- 角色码
    role_code VARCHAR(32) NOT NULL,
    -- 行为码 暂定为: （查：S，改：U，增：A，删：D）+目标角色码。如创客中心可以创建消费者和查询消费者，SA_CONSUME。到时候用字符串处理来判断。
    -- 特殊权限：ADMIN，ROOT
    action_code VARCHAR(32) NOT NULL, 
    -- 创建时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 更新时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 状态
    -- 0:正常  1:冻结或删除
    status INT(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (role_code,action_code)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;
-- 平台管理员，行为是系统默认的平台管理行为
INSERT INTO `machine_role_act` (`role_code`, `action_code`) VALUES ('ADMIN', 'ADMIN');
-- 系统管理员，行为是系统默认的系统管理行为
INSERT INTO `machine_role_act` (`role_code`, `action_code`) VALUES ('ROOT', 'ROOT');
-- 创客中心，查询和删除role_code为CONSUME的行为
INSERT INTO `machine_role_act` (`role_code`, `action_code`) VALUES ('CKZX', 'SA_CONSUME');


Drop table if exists machine_bank_info;
-- 用户银行卡信息
CREATE TABLE machine_bank_info (
    -- 用户ID
    userid BIGINT(20) NOT NULL,
    -- 银行卡号
    card_sn VARCHAR(64) NOT NULL,
    -- 绑定卡的身份证姓名
    idcard_name VARCHAR(64) NOT NULL,
     -- 绑定卡的身份证号码
    idcard_num VARCHAR(18) NOT NULL,
    -- 开户银行
    head_bank VARCHAR(64) NOT NULL,
    -- 开户支行
    sub_bank VARCHAR(128) NOT NULL,
    -- 银行预留的手机号码
    mobile VARCHAR(32) NOT NULL,
    -- 创建时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 更新时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 状态
    -- 0:正常  1:冻结或删除
    status INT(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (userid)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;

Drop table if exists machine_addr_info;
-- 用户地址
CREATE TABLE machine_addr_info (
    -- 地址id
    addr_id INT(11) NOT NULL,
    -- 用户id
    userid BIGINT(20) NOT NULL,
    -- 用户地址
    addr VARCHAR(128) DEFAULT NULL,
    -- 创建时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 更新时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 状态
    -- 0:正常  1:冻结或删除
    status INT(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (addr_id)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;

Drop table if exists machine_promotor_money;
-- 用户佣金表，当用户审核通过后创键，状态随用户状态而改变。
CREATE TABLE machine_promotor_money (
    -- 用户ID
    userid BIGINT(20) NOT NULL AUTO_INCREMENT,
    -- 可提现佣金
    cash INT(11) NOT NULL DEFAULT '0',
    -- 拥金余额(压款+可提现金额，单位:元后小数点4位)
    balance INT NOT NULL DEFAULT '0',
    -- 已提现佣金
    withdraw INT(11) NOT NULL DEFAULT '0',
    -- 佣金总额
    total INT(11) DEFAULT '0',
     -- 创建时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      -- 更新时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (userid)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;

INSERT INTO `machine_promotor_money` (`userid`) VALUES ('1');
INSERT INTO `machine_promotor_money` (`userid`) VALUES ('2');
INSERT INTO `machine_promotor_money` (`userid`) VALUES ('3');
INSERT INTO `machine_promotor_money` (`userid`) VALUES ('4');


Drop table if exists machine_promoter_withdraw;
-- 佣金提现表
CREATE TABLE machine_promoter_withdraw (
    -- 订单表
    order_id VARCHAR(64) NOT NULL DEFAULT '',
    -- 创建时间
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 更新时间
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 用户ID
    userid BIGINT(20) NOT NULL,
    -- 订单金额(单位:元后小数点4位), 正数为收入
    amount INT(11) NOT NULL DEFAULT '0',
    -- 提现渠道，0微信号, 1银行卡提现
    pay_channel INT(11) DEFAULT NULL,
    -- 提现，0时，此值为微信的appid; 1时为银行类别
    appid_item VARCHAR(256) NOT NULL DEFAULT '',
    -- 提现，0时，此值为微信的appid; 1时为银行类别
    openid_card VARCHAR(256) NOT NULL DEFAULT '',
    -- 个税百分之几, 只存整数部分, 如:20%存储为20
    tax INT(11) NOT NULL DEFAULT '0',
    -- 税款
    cash_tax BIGINT(20) NOT NULL DEFAULT '0',
    -- 实收现金
    cash_pay BIGINT(20) NOT NULL DEFAULT '0',
    -- 订单状态
    -- 0, 初始化(未审核);1，已通过(已支付)，2，已取消, 填写备注值
    status INT(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (order_id),
    KEY machine_promoter_withdraw_idx0 (userid),
    KEY machine_promoter_withdraw_idx1 (create_time)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8 PARTITION BY KEY (order_id) PARTITIONS 10;

Drop table if exists machine_distribute_record;
-- 机器码分配记录表，兼机器码归属表
CREATE TABLE machine_distribute_record (
    -- 发放者 （0代表根,机器码的最初拥有者）
    from_userid BIGINT(20) NOT NULL DEFAULT '0',
    -- 分配者（机器码的拥有者）
    to_userid BIGINT(20) NOT NULL,
    -- 机器码
    machine_code VARCHAR(45) NOT NULL,
    -- 创建时间，若需机器码查询走向，可以用这个排序
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 是否是机器码的当前拥有者
    -- 0是,1否
    is_owner INT(11)NOT NULL DEFAULT '0',
    -- 0：正常 1：售出状态，机器码售出具体状态可查看订单 2:待结算，表示这条记录“分配者”佣金正在结算
    -- 3：已结算，表示这条记录的“分配者”的佣金已经结算 :4：被冻结，失效
    status INT(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (from_userid,to_userid,machine_code)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;
CREATE INDEX machine_distribute_record_idx0 ON machine_distribute_record(from_userid);
CREATE INDEX machine_distribute_record_idx1 ON machine_distribute_record(to_userid);
CREATE INDEX machine_distribute_record_idx2 ON machine_distribute_record(machine_code);

Drop table if exists  machine_sale_record;
-- 机器订单和交易记录表
CREATE TABLE machine_sale_record (
    order_id VARCHAR(64) NOT NULL,
    -- 卖家
    seller_id BIGINT(20) NOT NULL,
    -- 买家
    purchase_id BIGINT(20) NOT NULL,
    -- 机器码
    machine_code VARCHAR(45) DEFAULT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 0：正常 1：已下单 2：已发件 3：已到货 4：已收货 5：已付款 6：退货中 7：已退货 8:被冻结
    status INT(11) NOT NULL DEFAULT '0',
    -- 地址信息
    addr VARCHAR(128) NOT NULL,
    -- 电话信息
    mobile VARCHAR(32) NOT NULL,
    -- 备注
    memo VARCHAR(128) NOT NULL,
    PRIMARY KEY (order_id)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8 PARTITION BY KEY (order_id) PARTITIONS 10;
CREATE INDEX machine_sale_record_idx0 ON machine_sale_record(create_time);
CREATE INDEX machine_sale_record_idx1 ON machine_sale_record(seller_id);
CREATE INDEX machine_sale_record_idx2 ON machine_sale_record(purchase_id);


-- 查询所有子节点
-- SELECT * FROM machine_user_info 
-- WHERE FIND_IN_SET(userid,queryAllChild(?))and userid!=? and role_code=?
DROP function IF EXISTS `queryAllChild`;

DELIMITER $$
USE `hz_test`$$
CREATE FUNCTION queryAllChild(areaId INT)
RETURNS VARCHAR(4000)
BEGIN
DECLARE sTemp VARCHAR(4000);
DECLARE sTempChd VARCHAR(4000);

SET sTemp='$';
SET sTempChd = CAST(areaId AS CHAR);

WHILE sTempChd IS NOT NULL DO

SET sTemp= CONCAT(sTemp,',',sTempChd);
SELECT 
    GROUP_CONCAT(userid)
INTO sTempChd FROM
    machine_user_info
WHERE
    FIND_IN_SET(parentid, sTempChd) > 0 ;


END WHILE;
RETURN sTemp;
END$$

DELIMITER ;

