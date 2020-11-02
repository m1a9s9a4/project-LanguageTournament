INSERT INTO player_type (id, japanese, english) VALUES (1, 'プログラミング言語', 'programming language');

INSERT INTO player (id, japanese, english, img, type_id) VALUES (1, 'php', 'php', 'https://www.php.net//images/logos/new-php-logo.svg', 1);
INSERT INTO player (id, japanese, english, img, type_id) VALUES (2, 'golang', 'golang', 'https://cdn.svgporn.com/logos/go.svg', 1);
INSERT INTO player (id, japanese, english, img, type_id) VALUES (3, 'vue.js', 'vue.js', 'https://cdn.svgporn.com/logos/vue.svg', 1);

INSERT INTO battle (player_1_id, player_2_id) VALUES (1, 2);
INSERT INTO battle (player_1_id, player_2_id) VALUES (1, 3);
INSERT INTO battle (player_1_id, player_2_id) VALUES (2, 3);

INSERT INTO question (japanese, english, player_type_id) VALUES (
    '書きやすさ', 'Easy to write', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    '書くのが楽しい', 'Fun to write', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    '初学者にとっての習得しやすい', 'Easy to study for beginners?', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    '情報が豊富（書籍、ネット含め）', 'Have many information in web or anykind', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    '処理速度が速い', 'Fast process', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    '開発環境構築のしやすい', 'Easy to create development environment', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    'コミュニティが充実している', 'Great community', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    '個人開発で使いやすい', 'Recommend for Personal Development', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    'チーム開発におすすめ', 'Recommend for Team Development?', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    'ライブラリが充実している', 'Which one is recommended to study for beginners?', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    'APIを作りやすい', 'Easy to create an API', 1
);
INSERT INTO question (japanese, english, player_type_id) VALUES (
    '公式ドキュメントが充実している', 'Better official documentation', 1
);