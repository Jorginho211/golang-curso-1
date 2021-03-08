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

    let elementoClicado = $(evento.target);
    if (!evento.target.classList.contains('curtir-publicacao')) {
        elementoClicado = elementoClicado.parent()
    }
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "GET",
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');
    }).fail(function() {
        alert("Erro ao curtir publicaçao");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}

function descurtirPublicacao(evento) {
    evento.preventDefault();

    let elementoClicado = $(evento.target);
    if (!evento.target.classList.contains('descurtir-publicacao')) {
        elementoClicado = elementoClicado.parent()
    }
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "GET",
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.addClass('curtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.removeClass('descurtir-publicacao');
    }).fail(function() {
        alert("Erro ao descurtir publicaçao");
    }).always(function () {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao() {
    $(this).prop('disabled', true);

    const publicacaoId = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        alert("Publicaçao editada com sucesso!");
    }).fail(function() {
        alert("Erro ao editar a publicaçao!");
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false);
    });
}

function deletarPublicacao(evento) {
    evento.preventDefault();

    let elementoClicado = $(evento.target);
    if (!evento.target.classList.contains('deletar-publicacao')) {
        elementoClicado = elementoClicado.parent()
    }
    const publicacao = elementoClicado.closest('div');
    const publicacaoId = publicacao.data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "DELETE"
    }).done(function() {
        publicacao.fadeOut("slow", function() {
            $(this).remove();
        });
    }).fail(function() {
        alert("Erro ao excluir a publicaçao");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

$(document).ready(function() {
    $('#nova-publicacao').on('submit', criarPublicacao);
    $(document).on('click', '.curtir-publicacao', curtirPublicacao);
    $(document).on('click', '.descurtir-publicacao', descurtirPublicacao);
    $('#atualizar-publicacao').on('click', atualizarPublicacao);
    $('.deletar-publicacao').on('click', deletarPublicacao);
});
