import { createLoader, parseAsInteger } from 'nuqs/server'

export const coordinatesSearchParams = {
  page: parseAsInteger.withDefault(1),
  per_page: parseAsInteger.withDefault(10),
}

export const loadPaginationSearchParams = createLoader(coordinatesSearchParams)
