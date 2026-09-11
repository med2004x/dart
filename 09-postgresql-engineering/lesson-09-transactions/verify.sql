SELECT id, owner_name, balance
FROM app.accounts
ORDER BY id;

SELECT SUM(balance) AS total_money
FROM app.accounts;
