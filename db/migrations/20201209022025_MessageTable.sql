
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `messages` (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `connection_id` INT NOT NULL,
  `content` TEXT,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  INDEX index_messages_on_user_id(`user_id`),
  INDEX index_messages_on_connection_id(`connection_id`)
);	

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `messages`;
