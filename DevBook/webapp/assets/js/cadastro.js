$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() !== $('#confirmar-senha').val()) {
        alert("As senhas nao coincidem")
        return;
    }

    $.akax({
        url: "/usuarios",
        method: 'POST',
        data: {
            nome: $('#nome').val(),
            nome: $('#email').val(),
            nome: $('#nick').val(),
            senha: $('#senha').val()
        }
    })
}