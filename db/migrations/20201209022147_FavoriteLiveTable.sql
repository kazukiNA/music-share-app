
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `favorite_lives` (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `artist_name` VARCHAR(255),
  `event_date` DATETIME,
  `live_name` VARCHAR(255),
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  INDEX index_favorite_lives_on_user_id(`user_id`)
);	

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `favorite_lives`;