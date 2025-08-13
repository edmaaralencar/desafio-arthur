import ky from 'ky'

function delay(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export const api = ky.create({
  prefixUrl: 'http://localhost:8080',
  hooks: {
    beforeRequest: [
      async () => {
        if (process.env.NODE_ENV !== 'production') {
          await delay(2000)
        }
      },
    ],
  },
})
