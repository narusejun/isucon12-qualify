CREATE INDEX `nj_1` ON `player_score` (`competition_id`, `player_id`, `row_num` DESC);
CREATE INDEX `nj_2` ON `competition` (`created_at`);
CREATE INDEX `nj_3` ON `player` (`created_at`);
CREATE INDEX `nj_4` ON `competition` (`created_at` DESC);
CREATE INDEX `nj_5` ON `player` (`created_at` DESC);
