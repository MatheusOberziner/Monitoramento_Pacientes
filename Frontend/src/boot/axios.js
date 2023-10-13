import axios from 'axios'

axios.defaults.baseURL = process.env.API_URL

const getPacientes = async params => {
  const res = await axios.get('/pacientes', { params })
  return res.data
}

const createNewPaciente = async params => {
  const res = await axios.post('/pacientes', params)
  return res.data
}

export { getPacientes, createNewPaciente }
