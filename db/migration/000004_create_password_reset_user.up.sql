CREATE TABLE IF NOT EXISTS `password_reset_user` (
  `user_id` INT NOT NULL,
  `reset_hash` VARCHAR(256) NOT NULL,
  `expired_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX `fk_password_reset_user_user1_idx` (`user_id` ASC) VISIBLE,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_password_reset_user_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
