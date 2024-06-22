import withNuxt from './.nuxt/eslint.config.mjs'

export default withNuxt(
  {
    ignores: ['node_modules', 'dist', '**/dist/**', 'public', '**/public/**'],
    files: ['**/*.{vue,ts,js,json}'],
    rules: {
      'no-console': 'off',
      'brace-style': ['error', 'stroustrup'],
      'curly': ['off'],
      'vue/multi-word-component-names': 'off'
    }
  },
)


