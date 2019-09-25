-- Sample migration
-- +migrate Up
CREATE TABLE user (
id bigint unsigned PRIMARY KEY AUTO_INCREMENT,
name varchar(255) NOT NULL,
email varchar(255) NOT NULL,
created_at datetime,
updated_at datetime)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC
;
-- +migrate Down
DROP TABLE IF EXISTS user;
