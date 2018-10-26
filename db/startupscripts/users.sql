CREATE TABLE IF NOT EXISTS `user_details` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT,
    `age` INTEGER CHECK(age>=1),
    `email` varchar(255) NOT NULL UNIQUE,
    `name` varchar(50) DEFAULT NULL,
    `gender` varchar(10) DEFAULT NULL,
    `salt` varbinary(8) NOT NULL,
    `hash` varbinary(32) NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;
