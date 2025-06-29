create table samples(
    id VARCHAR(100) not null,
    name VARCHAR(100) not null,
    primary key (id)
) engine=InnoDB default charset=utf8mb4;

create table users(
    id INT not null auto_increment,
    name VARCHAR(100) not null,
    username VARCHAR(100) not null,
    email VARCHAR(100) not null,
    password VARCHAR(100) not null,
    status VARCHAR(100) not null,
    created_at TIMESTAMP not null default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    primary key (id)
    UNIQUE INDEX `users_username_unique`(`username` ASC) USING BTREE,
    UNIQUE INDEX `users_email_unique`(`email` ASC) USING BTREE
) engine=InnoDB default charset=utf8mb4;

create table user_logs(
    id BIGINT UNSIGNED not null auto_increment,
    user_id BIGINT UNSIGNED not null,
    action VARCHAR(100) not null,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    primary key (id),
    foreign key (user_id) references users(id) on delete cascade on update cascade
) engine=InnoDB default charset=utf8mb4;

create table todos (
    id BIGINT UNSIGNED not null auto_increment,
    user_id BIGINT UNSIGNED not null,
    title VARCHAR(100) not null,
    description TEXT,
    created_at TIMESTAMP not null default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP null,
    primary key (id),
    foreign key (user_id) references users(id) on delete cascade on update cascade
)