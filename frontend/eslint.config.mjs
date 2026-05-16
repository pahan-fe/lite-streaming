import prettierConfig from 'eslint-config-prettier/flat'
import perfectionist from 'eslint-plugin-perfectionist'

// @ts-check
import withNuxt from './.nuxt/eslint.config.mjs'

export default withNuxt(
  {
    name: 'app/ignores',
    ignores: ['.nuxt/**', '.output/**', 'node_modules/**', 'public/**', 'dist/**', '**/*.d.ts'],
  },

  {
    name: 'app/plugins',
    plugins: {
      perfectionist,
    },
  },

  {
    name: 'app/rules',
    rules: {
      // ------- imports & ordering -------
      'perfectionist/sort-imports': [
        'error',
        {
          type: 'natural',
          order: 'asc',
          newlinesBetween: 1,
          groups: [
            'builtin',
            'external',
            'internal',
            ['parent', 'sibling', 'index'],
            'type',
            'style',
          ],
        },
      ],
      'perfectionist/sort-named-imports': ['error', { type: 'natural', order: 'asc' }],
      'perfectionist/sort-exports': ['error', { type: 'natural', order: 'asc' }],

      // ------- Vue specifics -------
      'vue/multi-word-component-names': 'off',
      'vue/block-order': ['error', { order: ['script', 'template', 'style'] }],
      'vue/define-macros-order': [
        'error',
        {
          order: ['defineProps', 'defineEmits', 'defineSlots', 'defineExpose', 'defineOptions'],
        },
      ],
      'vue/component-name-in-template-casing': [
        'error',
        'PascalCase',
        { registeredComponentsOnly: false },
      ],
      'vue/require-default-prop': 'off',

      // ------- TypeScript -------
      '@typescript-eslint/no-unused-vars': [
        'error',
        { argsIgnorePattern: '^_', varsIgnorePattern: '^_' },
      ],
      '@typescript-eslint/no-explicit-any': 'error',

      // ------- general -------
      'no-console': ['warn', { allow: ['warn', 'error'] }],
      'no-debugger': 'error',
      'object-shorthand': ['error', 'always'],
      'prefer-const': 'error',
    },
  },

  prettierConfig,
)
