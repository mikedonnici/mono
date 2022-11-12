create table attribute
(
    id   int auto_increment primary key,
    type varchar(255) not null,
    name varchar(255) not null,
    constraint attribute_name_uindex unique (name)
);
INSERT INTO attribute (id, type, name) VALUES (1, 'visual', 'colour');
INSERT INTO attribute (id, type, name) VALUES (2, 'measurement', 'height');
INSERT INTO attribute (id, type, name) VALUES (3, 'measurement', 'weight');

create table user
(
    id   int auto_increment primary key,
    firstname varchar(255) not null,
    lastname varchar(255) not null,
    email varchar(255) not null,
    constraint user_email_uindex unique (email)
);
INSERT INTO user (id, firstname, lastname, email) VALUES (1, 'Mike', 'Donnici', 'mike@mono.com');
INSERT INTO user (id, firstname, lastname, email) VALUES (2, 'Christie', 'Wood', 'christie@mono.com');
INSERT INTO user (id, firstname, lastname, email) VALUES (3, 'Maia', 'Donnici', 'maia@mono.com');
INSERT INTO user (id, firstname, lastname, email) VALUES (4, 'Leo', 'Donnici', 'leo@mono.com');
