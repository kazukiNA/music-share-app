
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `favorite_genres` (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `name` VARCHAR(255),
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  INDEX index_favorite_albums_on_user_id(`user_id`)
);	

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `favorite_genres`;