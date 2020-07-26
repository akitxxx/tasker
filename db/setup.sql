create table users (
  id integer primary key auto_increment,
  email varchar(255) unique not null,
  password varchar(100) not null
);

insert into users (email, password)
values (
  "test1@example.com",
  "password"
);
insert into users (email, password)
values (
  "test2@example.com",
  "password"
);
