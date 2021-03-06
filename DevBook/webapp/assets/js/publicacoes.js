function criarPublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo:  $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao criar a publicaçao");
    })
}

function curtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "GET",
    }).done(function() {
        let contadorDeCurtidas = elementoClicado.next('span');
        if (contadorDeCurtidas.length === 0)
            contadorDeCurtidas = elementoClicado.parent().next('span');

        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);
    }).fail(function() {
        alert("Erro ao curtir publicaçao");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}

$(document).ready(function() {
    $('#nova-publicacao').on('submit', criarPublicacao);
    $('.curtir-publicacao').on('click', curtirPublicacao);
});
