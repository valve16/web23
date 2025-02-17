drop table post;
CREATE TABLE post
(
   `post_id`       INT NOT NULL AUTO_INCREMENT,
   `title`         VARCHAR(255) NOT NULL,
   `description`   VARCHAR(255) NOT NULL,
   `image_url`     VARCHAR(255) NOT NULL,
   `author`        VARCHAR(255) NOT NULL,
   `author_url`    VARCHAR(255) NOT NULL,
   `publish_date`  VARCHAR(255) NOT NULL,
   `image_mod`     VARCHAR(255) NOT NULL,
   `featured`      TINYINT(1),
   `content`       TEXT NOT NULL,
   PRIMARY KEY (`post_id`)
) ENGINE = InnoDB
CHARACTER SET = utf8mb4
COLLATE utf8mb4_unicode_ci
;
USE blog;
SELECT * FROM post;