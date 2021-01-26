create table users (
  id integer primary key auto_increment,
  email varchar(255) unique not null,
  password varchar(100) not null,
  created_at datetime default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
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

create table lanes (
  id integer primary key auto_increment,
  name varchar(50) not null,
  created_at datetime default current_timestamp,
  updated_at  timestamp default current_timestamp on update current_timestamp
);

insert into lanes (id, name)
values (
  1,
  "tasks1"
);

create table tasks (
  id integer primary key auto_increment,
  lane_id integer not null,
  title varchar(255) not null,
  content varchar(100),
  created_at datetime default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

insert into tasks (lane_id, title, content)
values (
  1,
  "task1",
  "task1 content"
);
