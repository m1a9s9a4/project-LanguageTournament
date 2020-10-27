CREATE DATABASE comparision_box;

use comparision_box;

CREATE TABLE player (
    id INT NOT NULL, 
    name varchar(100) NOT NULL,
    imgsrc text NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id),
    INDEX name_index (name)
);

CREATE TABLE player_type (
    id INT NOT NULL, 
    name varchar(100) NOT NULL,
    parent_id INT DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id),
    INDEX name_parent_id_index (name, parent_id), 
    FOREIGN KEY (parent_id) REFERENCES player_type (id)
);

CREATE TABLE player_types (
    id INT NOT NULL,
    player_id INT NOT NULL,
    type_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (player_id) REFERENCES player (id),
    FOREIGN KEY (type_id) REFERENCES player_type (id),
    PRIMARY KEY (id),
    INDEX player_id_type_id_index (player_id, type_id)
);

CREATE TABLE battle (
    id INT NOT NULL,
    player_1_id INT NOT NULL,
    player_2_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (player_1_id) REFERENCES player (id),
    FOREIGN KEY (player_2_id) REFERENCES player (id),
    PRIMARY KEY (id),
    INDEX player_id_index (player_1_id, player_2_id)
);

CREATE TABLE question (
    id INT NOT NULL,
    content text NOT NULL,
    player_type_id INT NOT NULL,
    FOREIGN KEY (player_type_id) REFERENCES player_type (id),
    PRIMARY KEY (id),
    INDEX player_type_id_index (player_type_id),
);

CREATE TABLE answer (
    id INT NOT NULL,
    battle_id INT NOT NULL,
    question_id INT NOT NULL,
    selected_player_id INT NOT NULL,
    FOREIGN KEY (battle_id) REFERENCES battle (id),
    FOREIGN KEY (question_id) REFERENCES question (id),
    FOREIGN KEY (selected_player_id) REFERENCES player (id),
    PRIMARY KEY (id),
    INDEX battle_question_player_ids_index (battle_id, question_id, selected_player_id),
);
