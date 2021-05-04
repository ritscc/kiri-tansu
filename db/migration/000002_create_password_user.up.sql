CREATE TABLE IF NOT EXISTS `password_user` (
  `user_id` INT NOT NULL,
  `mail_address` VARCHAR(256) NOT NULL,
  `password` VARCHAR(256) NOT NULL,
  `is_lock` TINYINT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` INT NOT NULL,
  INDEX `fk_password_user_user_idx` (`user_id` ASC) VISIBLE,
  PRIMARY KEY (`user_id`),
  UNIQUE INDEX `mail_address_UNIQUE` (`mail_address` ASC) VISIBLE,
  CONSTRAINT `fk_password_user_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
