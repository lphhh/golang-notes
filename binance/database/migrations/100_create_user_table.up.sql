CREATE TABLE IF NOT EXISTS `users`
(
    id         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    account    varchar(191)        NOT NULL,
    email      varchar(191) DEFAULT NULL,
    password   varchar(191)        NOT NULL,
    name       varchar(191)        NOT NULL,
    created_at datetime     DEFAULT NULL,
    updated_at datetime     DEFAULT NULL,
    PRIMARY KEY (`id`)
)
