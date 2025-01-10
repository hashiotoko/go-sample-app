-- Create "users" table
CREATE TABLE `users` (`id` varchar(50) NOT NULL COMMENT "ユーザーID", `name` varchar(255) NOT NULL COMMENT "ユーザー名前", `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "作成日時", `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新日時", PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
