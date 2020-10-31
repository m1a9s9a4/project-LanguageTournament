CREATE DATABASE comparision_box;

use comparision_box;

CREATE TABLE player_type (
    id INT, 
    japanese varchar(100) NOT NULL,
    english varchar(100) NOT NULL,
    parent_id INT DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id),
    INDEX name_parent_id_index (id, japanese, english, parent_id), 
    FOREIGN KEY (parent_id) REFERENCES player_type (id)
);

INSERT INTO player_type (id, japanese, english) VALUES (1, 'プログラミング言語', 'programming language');

CREATE TABLE player (
    id INT AUTO_INCREMENT, 
    japanese varchar(100) NOT NULL,
    english varchar(100) NOT NULL,
    type_id INT NOT NULL,
    img text NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (type_id) REFERENCES player_type (id),
    INDEX name_type_id_index (japanese, english, type_id) 
);

insert into player (id, japanese, english, img, type_id) VALUES (1, 'php', 'php', 'https://www.php.net//images/logos/new-php-logo.svg', 1);
INSERT INTO player (id, japanese, english, img, type_id) VALUES (2, 'golang', 'golang', 'https://cdn.svgporn.com/logos/go.svg', 1);
INSERT INTO player (id, japanese, english, img, type_id) VALUES (3, 'vue.js', 'vue.js', 'https://cdn.svgporn.com/logos/vue.svg', 1);

CREATE TABLE battle (
    id INT AUTO_INCREMENT,
    player_1_id INT NOT NULL,
    player_2_id INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (player_1_id) REFERENCES player (id),
    FOREIGN KEY (player_2_id) REFERENCES player (id),
    PRIMARY KEY (id),
    INDEX player_id_index (player_1_id, player_2_id)
);

CREATE TABLE question (
    id INT AUTO_INCREMENT,
    japanese_content text NOT NULL,
    english_content text NOT NULL,
    player_type_id INT NOT NULL,
    FOREIGN KEY (player_type_id) REFERENCES player_type (id),
    PRIMARY KEY (id),
    INDEX player_type_id_index (player_type_id)
);

CREATE TABLE answer (
    id INT AUTO_INCREMENT,
    battle_id INT NOT NULL,
    question_id INT NOT NULL,
    selected_player_id INT NOT NULL,
    FOREIGN KEY (battle_id) REFERENCES battle (id),
    FOREIGN KEY (question_id) REFERENCES question (id),
    FOREIGN KEY (selected_player_id) REFERENCES player (id),
    PRIMARY KEY (id),
    INDEX battle_question_player_ids_index (battle_id, question_id, selected_player_id)
);


