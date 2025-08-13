import { SearchParams } from 'nuqs/server'

import { getContacts } from '@/api/get-contacts'
import { loadPaginationSearchParams } from '@/lib/pagination-search-params'

export default async function Home({
  searchParams,
}: {
  searchParams: Promise<SearchParams>
}) {
  const { page, per_page: perPage } =
    await loadPaginationSearchParams(searchParams)

  const { contacts } = await getContacts({ page, perPage })

  return (
    <div className="grid min-h-screen grid-rows-[20px_1fr_20px] items-center justify-items-center gap-16 p-8 pb-20 font-sans sm:p-20">
      <p className="text-3xl font-bold text-red-500">Hello World</p>
    </div>
  )
}
