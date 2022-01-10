
CREATE TABLE IF NOT EXISTS `users` (
    `id` int(11) PRIMARY KEY AUTO_INCREMENT COMMENT 'User ID',
    `name` varchar(32) NOT NULL COMMENT 'User Name',
    `created_at` datetime NOT NULL COMMENT 'Created Time',
    updated_at datetime COMMENT 'Updated Time'
) COMMENT='User';
