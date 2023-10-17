import axios from 'axios'

axios.defaults.baseURL = process.env.API_URL

// Endpoint responsável por pegar os dados dos pacientes
const getPacientes = async params => {
  const res = await axios.get('/pacientes', { params })
  return res.data
}

// Endpoint responsável por adicionar um novo paciente
const createNewPaciente = async params => {
  const res = await axios.post('/pacientes', params)
  return res.data
}

// Endpoint responsável por retornar o histórico de sinais vitais do paciente
const getSinaisByPaciente = async (idPaciente, params) => {
  const res = await axios.get(`/sinais_vitais/${idPaciente}`, { params })
  return res.data
}

export { getPacientes, createNewPaciente, getSinaisByPaciente }
