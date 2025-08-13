import { FlatCompat } from '@eslint/eslintrc'
import { dirname } from 'path'
import { fileURLToPath } from 'url'

const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)

const compat = new FlatCompat({
  baseDirectory: __dirname,
})

const eslintConfig = [
  ...compat.config({
    extends: [
      'next/core-web-vitals',
      'next/typescript',
      'plugin:prettier/recommended',
    ],
    plugins: ['simple-import-sort', 'prettier'],
    rules: {
      'simple-import-sort/imports': 'error',
      'simple-import-sort/exports': 'error',
      'react/react-in-jsx-scope': 'off',
      'prettier/prettier': [
        'error',
        {
          singleQuote: true,
          semi: false,
          plugins: ['prettier-plugin-tailwindcss'],
        },
        {
          usePrettierrc: false,
        },
      ],
    },
  }),
]

export default eslintConfig
