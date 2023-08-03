CREATE TABLE appointments (
    id serial PRIMARY KEY,
    patient_id INT NOT NULL REFERENCES accounts,
    doctor_id INT NOT NULL REFERENCES accounts,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    time timestamp(0) with time zone NOT NULL
);
