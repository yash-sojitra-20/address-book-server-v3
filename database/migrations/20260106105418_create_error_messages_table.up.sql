CREATE TABLE IF NOT EXISTS `error_messages` (
    `id` int NOT NULL AUTO_INCREMENT,
    `code` varchar(255) NOT NULL,
    `component` varchar(255) NOT NULL,
    `response_type` varchar(255) NOT NULL,
    `one` varchar(255) NOT NULL,
    `other` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `error_messages_uk_code` (`code`)
);