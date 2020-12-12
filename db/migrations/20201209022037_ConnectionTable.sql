
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `connections` (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `previous_user_id` INT NOT NULL,
  `following_user_id` INT NOT NULL,
  `status` VARCHAR(255),
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  INDEX index_connections_on_previous_user_id(`previous_user_id`),
  INDEX index_connections_on_following_user_id(`following_user_id`)
);	

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `connections`;