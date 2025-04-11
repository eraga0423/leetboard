INSERT INTO posts (title, post_content, post_image, post_time) VALUES
                                                                   ('Привет, мир!', 'Это мой первый пост на 1337b04rd. Всем хай!', 'https://example.com/images/hello.png', NOW()),
                                                                   ('Вопрос по Go', 'Кто-нибудь может объяснить, что такое интерфейсы в Go?', 'https://example.com/images/golang.png', NOW() - INTERVAL '5 minutes'),
                                                                   ('Rick and Morty', 'Какой ваш любимый персонаж из Rick and Morty?', 'https://example.com/images/rickmorty.jpg', NOW() - INTERVAL '10 minutes');
