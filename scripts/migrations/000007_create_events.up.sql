CREATE TABLE IF NOT EXISTS events(
    id bigserial PRIMARY KEY,
    proposed_dates bigint[],
    confirmed_date bigint,
    location text,
    rejected_remarks text,
    created_at bigint,
    company_id bigint,
    FOREIGN KEY (company_id) references companies(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,
    vendor_id bigint,
    FOREIGN KEY (vendor_id) references vendors(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,
    event_type_id bigint NOT NULL,
    FOREIGN KEY (event_type_id) references event_types(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    event_status_id bigint NOT NULL,
    FOREIGN KEY (event_status_id) references event_status(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
