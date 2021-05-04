CREATE TABLE IF NOT EXISTS `lending` (
  `user_id` INT NOT NULL,
  `item_id` INT NOT NULL,
  `returned_at` DATE NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` INT NOT NULL,
  INDEX `fk_lending_user1_idx` (`user_id` ASC) VISIBLE,
  INDEX `fk_lending_item1_idx` (`item_id` ASC) VISIBLE,
  PRIMARY KEY (`user_id`, `item_id`),
  CONSTRAINT `fk_lending_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_lending_item1`
    FOREIGN KEY (`item_id`)
    REFERENCES `item` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
