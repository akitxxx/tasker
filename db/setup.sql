create table users (
  id integer primary key auto_increment,
  email varchar(255) unique not null,
  password varchar(100) not null,
  created_at datetime default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

insert into users (id, email, password)
values (
  1,
  "test1@example.com",
  "password"
);

insert into users (id ,email, password)
values (
  2,
  "test2@example.com",
  "password"
);

create table lanes (
  id integer primary key auto_increment,
  user_id integer not null,
  name varchar(50) not null,
  created_at datetime default current_timestamp,
  updated_at  timestamp default current_timestamp on update current_timestamp
);

insert into lanes (id, user_id, name)
values (
  1,
  1,
  "tasks1"
);

create table tasks (
  id integer primary key auto_increment,
  user_id integer not null,
  lane_id integer not null,
  title varchar(255) not null,
  content varchar(100),
  created_at datetime default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

insert into tasks (lane_id, user_id, title, content)
values (
  1,
  1,
  "task1",
  "task1 content"
);
