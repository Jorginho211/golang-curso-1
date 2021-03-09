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
        Swal.fire("Ops...", "Erro ao criar a publicaçao!", "error");
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
        Swal.fire("Ops...", "Erro ao curtir publicaçao!", "error");
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
        Swal.fire("Ops...", "Erro ao descurtir publicaçao!", "error");
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
        Swal.fire('Sucesso', 'Publicaçao editada com sucesso!', 'success')
            .then(function() {
                window.location = "/home";
            });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao editar a publicaçao!", "error");
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false);
    });
}

function deletarPublicacao(evento) {
    evento.preventDefault();

    Swal.fire({
        title: "Atençao!",
        text: "Tem certeza que deseja excluir essa publicaçao? Essa açao é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

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
            Swal.fire("Ops...", "Erro ao excluir a publicaçao!", "error");
        }).always(function() {
            elementoClicado.prop('disabled', false);
        });
    });
}

$(document).ready(function() {
    $('#nova-publicacao').on('submit', criarPublicacao);
    $(document).on('click', '.curtir-publicacao', curtirPublicacao);
    $(document).on('click', '.descurtir-publicacao', descurtirPublicacao);
    $('#atualizar-publicacao').on('click', atualizarPublicacao);
    $('.deletar-publicacao').on('click', deletarPublicacao);
});
