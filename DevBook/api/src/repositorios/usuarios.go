package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um repositorio de usuarios
func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInsertado, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInsertado), nil
}

// Buscar traz todos os usuarios que atendem um filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // usasse % para escapar os porcentagens para o LIKE
	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criado from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		)

		if erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuario do banco de dados
func (repositorio usuarios) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criado from usuarios where id = ?", id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera as informaçoes de um usuario no banco de dados
func (repositorio usuarios) Atualizar(id uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclue as informaçoes no banco de dados
func (repositorio usuarios) Deletar(id uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}
