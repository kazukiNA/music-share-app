
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `favorite_artists` (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `first_place` VARCHAR(255),
  `socond_place` VARCHAR(255),
  `third_place` VARCHAR(255),
  `fourth_place` VARCHAR(255),
  `fifth_place` VARCHAR(255),
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  INDEX index_favorite_artists_on_user_id(`user_id`)
);	

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `favorite_artists`;