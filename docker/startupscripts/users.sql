CREATE TABLE IF NOT EXISTS `user_details` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `age` INTEGER CHECK(age>=1),
    `email` varchar(255) NOT NULL UNIQUE,
    `name` varchar(50) DEFAULT NULL,
    `gender` varchar(10) DEFAULT NULL,
    `salt` varbinary(8) NOT NULL,
    `hash` varbinary(32) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `follows` (
    id_follower int(11) NOT NULL,
    id_followed int(11) NOT NULL,
    accepted bit(1) NOT NULL,
    PRIMARY KEY (id_follower, id_followed)
) DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `likes` (
    id_route int(11) NOT NULL,
    id_user int(11) NOT NULL,
    PRIMARY KEY (id_route, id_user)
) DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `routes` (
    id int(11) NOT NULL AUTO_INCREMENT,
    id_user int(11) NOT NULL,
    description TINYTEXT NOT NULL,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8  AUTO_INCREMENT=1;
