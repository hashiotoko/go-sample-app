-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `email_address` varchar(255) NOT NULL COMMENT "ユーザーのメールアドレス", ADD INDEX `email_address_idx` (`email_address`) COMMENT "メールアドレスインデックス";
