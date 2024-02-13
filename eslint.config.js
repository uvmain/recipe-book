import antfu from '@antfu/eslint-config'

export default antfu({
  vue: true,
  unocss: true,
  typescript: {
    tsconfigPath: 'tsconfig.json',
  },
  ignores: ['node_modules', 'dist', '**/dist/**', 'public', '**/public/**'],
  files: ['**/*.{vue,ts,js,json}'],
}, {
  rules: {
    'no-console': 'off',
    'brace-style': ['error', 'stroustrup'],
    'curly': ['off'],
  },
})
