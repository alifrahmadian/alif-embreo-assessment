CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    role_id bigint,
    FOREIGN KEY (role_id) references roles(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,
    company_id bigint,
    FOREIGN KEY (company_id) references companies(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,
    vendor_id bigint,
    FOREIGN KEY (vendor_id) references vendors(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

