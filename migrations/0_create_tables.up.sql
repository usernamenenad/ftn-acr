CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "subjects" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    subject_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "classrooms" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    classroom_name VARCHAR(255) NOT NULL,
    classroom_building VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "programs" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    program_name VARCHAR(255) NOT NULL,
    program_short_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "time_table" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    subject_id UUID,
    classroom_id UUID,
    class VARCHAR(255),
    holding_day DATE, -- TODO: Figure out to make a Day type out of this
    time_to TIME,
    time_from TIME,
    program_id UUID,
    study_groups TEXT[],
    holding_date DATE,
    FOREIGN KEY (subject_id) REFERENCES subjects(id),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id),
    FOREIGN KEY (program_id) REFERENCES programs(id)
);
