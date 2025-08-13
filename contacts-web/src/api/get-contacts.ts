import { api } from '.'

type Request = {
  page: number
  perPage: number
}

type Response = {
  contacts: {
    id: number
    name: string
    email: string
    cpf_cnpj: string
    phone: string
    created_at: string
    updated_at: string
  }[]
  page: number
  perPage: number
  total: number
}

export async function getContacts({ page, perPage }: Request) {
  const response = await api
    .get(`contacts?page=${page}&per_page=${perPage}`, {})
    .json<Response>()

  return response
}
