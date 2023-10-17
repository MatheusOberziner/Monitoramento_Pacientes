export const convertDate = {
  convertDateToBr (data) {
    // Separa a data e o horário
    const partes = data.split('T')
    const dataPartes = partes[0].split('-')
    const horarioPartes = partes[1].split(':')
    // Obtém as partes da data
    const ano = dataPartes[0]
    const mes = dataPartes[1]
    const dia = dataPartes[2]
    // Obtém as partes do horário
    const hora = horarioPartes[0]
    const minuto = horarioPartes[1]
    const segundo = horarioPartes[2]
    // Monta a data no formato brasileiro
    const dataBrasileira = dia + '/' + mes + '/' + ano + ' ' + hora + ':' + minuto + ':' + segundo
    return dataBrasileira.replace('Z', '')
  }
}
