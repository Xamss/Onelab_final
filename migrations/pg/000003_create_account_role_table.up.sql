CREATE TABLE account_roles (
    account_id INT NOT NULL,
    role_id INT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (role_id)
       REFERENCES roles (id),
    FOREIGN KEY (user_id)
       REFERENCES accounts (id)
);