insert into funds (nickname, currency, amount)
values
("ben", "btc", 0.2),
("ben", "doge", 1.8),
("alex", "btc", 12.2),
("alex", "doge", 1.123),
("alex", "ada", 1.8),
("ben", "eth", 0.2),
("ben", "ada", 1.8),
("pedro", "btc", 1.2),
("pedro", "doge", 1.044),
("pedro", "ada", 5.2),
("pedro", "eth", 1.864) ON DUPLICATE KEY UPDATE amount = amount + amount;




