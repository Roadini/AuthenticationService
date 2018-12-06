USE db;

CREATE TABLE IF NOT EXISTS user_details (
    id          BIGINT          NOT NULL AUTO_INCREMENT,
    age         INTEGER         CHECK(age>=1),
    email       varchar(255)    NOT NULL UNIQUE,
    name        varchar(50)     NOT NULL,
    gender      varchar(10)     DEFAULT "Male",
    hash        varbinary(32)   DEFAULT NULL,
    salt        varbinary(8)    DEFAULT NULL,
    description varchar(255)    DEFAULT NULL,

    accessToken varchar(300)    DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS follows (
    id_follower int(11) NOT NULL,
    id_followed int(11) NOT NULL,
    accepted    bit(1)  NOT NULL,
    PRIMARY KEY (id_follower, id_followed)
) DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS routes (
    id          BIGINT     NOT NULL AUTO_INCREMENT,
    id_user     BIGINT     NOT NULL,
    description TINYTEXT    NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id_user) REFERENCES user_details(id) ON DELETE CASCADE
) DEFAULT CHARSET=utf8  AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS likes (
    id_route    BIGINT     NOT NULL,
    id_user     BIGINT     NOT NULL,
    PRIMARY KEY (id_route, id_user),
    FOREIGN KEY (id_user) REFERENCES user_details(id) ON DELETE CASCADE, 
    FOREIGN KEY (id_route) REFERENCES routes(id) ON DELETE CASCADE
) DEFAULT CHARSET=utf8;
