-- if exist douyin database, drop it
drop database if exists douyin;

create database douyin;
use douyin;

CREATE TABLE user (
                      id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                      created_at TIMESTAMP NOT NULL,
                      updated_at TIMESTAMP NOT NULL,
                      deleted_at TIMESTAMP,
                      name VARCHAR(255) NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      avatar VARCHAR(255) NOT NULL,
                      background_image VARCHAR(255) NOT NULL,
                      signature VARCHAR(255) NOT NULL,
                      follow_count INT UNSIGNED NOT NULL,
                      follower_count INT UNSIGNED NOT NULL,
                      total_favorited INT UNSIGNED NOT NULL,
                      work_count INT UNSIGNED NOT NULL,
                      favorite_count INT UNSIGNED NOT NULL
);

CREATE TABLE video (
                       id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       deleted_at TIMESTAMP,
                       author_id INT UNSIGNED NOT NULL,
                       play_url VARCHAR(255) NOT NULL,
                       cover_url VARCHAR(255) NOT NULL,
                       favorite_count INT UNSIGNED NOT NULL,
                       comment_count INT UNSIGNED NOT NULL,
                       title VARCHAR(255) NOT NULL,
                       FOREIGN KEY (author_id) REFERENCES user(id)
);


CREATE TABLE comment (
                         id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                         created_at TIMESTAMP NOT NULL,
                         updated_at TIMESTAMP NOT NULL,
                         deleted_at TIMESTAMP,
                         video_id INT UNSIGNED NOT NULL,
                         user_id INT UNSIGNED NOT NULL,
                         content TEXT NOT NULL,
                         FOREIGN KEY (video_id) REFERENCES video(id),
                         FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE following (
                           id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                           created_at TIMESTAMP NOT NULL,
                           updated_at TIMESTAMP NOT NULL,
                           deleted_at TIMESTAMP,
                           host_id INT UNSIGNED NOT NULL,
                           guest_id INT UNSIGNED NOT NULL,
                           FOREIGN KEY (host_id) REFERENCES user(id),
                           FOREIGN KEY (guest_id) REFERENCES user(id)
);

CREATE TABLE favorite (
                          id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                          created_at TIMESTAMP NOT NULL,
                          updated_at TIMESTAMP NOT NULL,
                          deleted_at TIMESTAMP,
                          user_id INT UNSIGNED NOT NULL,
                          video_id INT UNSIGNED NOT NULL,
                          status INT UNSIGNED NOT NULL,
                          FOREIGN KEY (user_id) REFERENCES user(id),
                          FOREIGN KEY (video_id) REFERENCES video(id)
);

CREATE TABLE message (
                         id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                         created_at TIMESTAMP NOT NULL,
                         updated_at TIMESTAMP NOT NULL,
                         deleted_at TIMESTAMP,
                         host_id INT UNSIGNED NOT NULL,
                         guest_id INT UNSIGNED NOT NULL,
                         content TEXT NOT NULL,
                         FOREIGN KEY (host_id) REFERENCES user(id),
                         FOREIGN KEY (guest_id) REFERENCES user(id)
);
