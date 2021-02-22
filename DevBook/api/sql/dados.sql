
INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$Mx9BWJl8IqKeZ02w2iNZSORNFXGwFhWmFeg3CofmESOx2KXj1iN0q"), -- usuario1
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$Mx9BWJl8IqKeZ02w2iNZSORNFXGwFhWmFeg3CofmESOx2KXj1iN0q"), -- usuario2
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$Mx9BWJl8IqKeZ02w2iNZSORNFXGwFhWmFeg3CofmESOx2KXj1iN0q"); -- usuario3

INSERT INTO seguidores (usuario_id, seguidor_id)
values
(1, 2), -- Usuario 2 segue ao usuario 1
(3, 1), -- Usuario 1 segue ao usuario 3
(1, 3); -- Usuario 3 segue ao usuario 1