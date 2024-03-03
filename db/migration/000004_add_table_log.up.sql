CREATE TABLE log_transaction (
    id INT AUTO_INCREMENT PRIMARY KEY,
    systemTraceAuditNumber VARCHAR(20),
    request TEXT,
    response TEXT,
    section TEXT,
    status TEXT,
    dtm_crt TIMESTAMP NOT NULL DEFAULT NOW(),
    dtm_upd TIMESTAMP NOT NULL DEFAULT NOW()
);
