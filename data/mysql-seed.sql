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
