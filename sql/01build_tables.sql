/*
    Table: workstream
    Description: contains workstream details
    version 1.0.0
    rename table to express more than just the workstream name.
*/
CREATE TABLE IF NOT EXISTS workstream(
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(128) NOT NULL UNIQUE,
    PRIMARY KEY(id)
    ) ENGINE=InnoDB;

/*
    Table: engineer_details
    Description: holds details about the developer.
    version 0.0.2
    increase the size of the key int 
*/
CREATE TABLE IF NOT EXISTS engineer_details(
    id INT(11) NOT NULL AUTO_INCREMENT,
    first_name TEXT,
    last_name TEXT,
    email VARCHAR(128) NOT NULL UNIQUE,
    velocity INT,
    PRIMARY KEY(id)
) ENGINE=InnoDB;

/*
    Table: sprint
    Description: contains sprint details
    version 1.0.0
    rename table to express more than just sprint name.
*/
CREATE TABLE IF NOT EXISTS sprint(
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(128) NOT NULL UNIQUE,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB;

/*
    Table: sprint_line_item
    Description: a single line item for a sprint
    version 0.0.2 - increase the size of the primary key
*/
CREATE TABLE IF NOT EXISTS sprint_line_item(
    id INT(11) NOT NULL AUTO_INCREMENT,
    current_availability INT,
    previous_availability INT,
    capacity INT,
    target_points INT,
    committed_points_this_sprint INT,
    completed_points_this_sprint INT,
    completed_points_last_sprint INT,
    PRIMARY KEY (id)
) ENGINE=InnoDB;

/*
    Table: workstream_sprint_engineer_sprint_line_item_map
    Description: maps workstream/sprint/engineer/sprint line item
    version 1.0.0
    rename columns to match with renamed tables and preferred column names.
*/
CREATE TABLE IF NOT EXISTS workstream_sprint_engineer_sprint_line_item_map(
    workstream_id INTEGER NOT NULL,
    sprint_id INTEGER NOT NULL,
    sprint_summary_id INTEGER NOT NULL,
    PRIMARY KEY (
        workstream_id, sprint_id, sprint_summary_id
    )
) ENGINE=InnoDB;