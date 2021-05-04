CREATE TABLE IF NOT EXISTS `item_tagging` (
  `item_id` INT NOT NULL,
  `tag_id` INT NOT NULL,
  INDEX `fk_item_tag_item1_idx` (`item_id` ASC) VISIBLE,
  INDEX `fk_item_tag_tag1_idx` (`tag_id` ASC) VISIBLE,
  CONSTRAINT `fk_item_tag_item1`
    FOREIGN KEY (`item_id`)
    REFERENCES `item` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_item_tag_tag1`
    FOREIGN KEY (`tag_id`)
    REFERENCES `tag` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
