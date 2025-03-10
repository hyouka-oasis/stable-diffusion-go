module.exports = {
    extends: [
        'eslint:recommended',
        'plugin:react/recommended',
        'plugin:@typescript-eslint/recommended',
    ],
    rules: {
        'react/jsx-first-prop-new-line': [ 2, 'multiline' ],
        'react/jsx-indent': [ 'error', 4 ],
        'react/jsx-indent-props': [ 'error', 4 ],
        'import/no-extraneous-dependencies': 'off',
        'react/react-in-jsx-scope': 'off',
        'react/jsx-filename-extension': 'off',
        'import/extensions': 'off',
        'import/no-unresolved': 'off',
        'import/no-import-module-exports': 'off',
        'import/no-duplicates': 'off',
        'import/no-self-import': 'off',
        'import/no-useless-path-segments': 'off',
        'import/order': 'off',
        'import/no-cycle': 'off',
        'import/no-relative-packages': 'off',
        'import/no-named-as-default': 'off',
        'import/no-named-as-default-member': 'off',
        'global-require': 'off',
        'promise/always-return': 'off',
        'no-unused-vars': 'off',
        'no-useless-catch': 'off',
        'jsx-a11y/click-events-have-key-events': 'off',
        'jsx-a11y/no-static-element-interactions': 'off',
        'max-classes-per-file': 'off',
        'class-methods-use-this': 'off',
        'no-plusplus': 'off',
        'no-continue': 'off',
        'no-shadow': 'off',
        'no-undef': 'off',
        'no-await-in-loop': 'off',
        'import/prefer-default-export': 'off',
        'no-nested-ternary': 'off',
        'consistent-return': 'off',
        'no-use-before-define': 'off',
        'react/function-component-definition': 'off',
        'jsx-a11y/media-has-caption': 'off',
        'react/no-unknown-property': 'off',
        'react/button-has-type': 'off',
        'no-console': 'off',
        semi: [ 'error', 'always' ],
        indent: [ 'error', 4, { SwitchCase: 1 } ],
        '@typescript-eslint/ban-ts-comment': 'off',
        '@typescript-eslint/no-empty-interface': 'off',
        '@typescript-eslint/no-unused-vars': 'off',
        '@typescript-eslint/no-extra-non-null-assertion': 'off',
        '@typescript-eslint/no-inferrable-types': 'off',
        '@typescript-eslint/explicit-module-boundary-types': 'off',
        'no-useless-escape': 'off',
        '@typescript-eslint/no-var-requires': 'off',
        'array-bracket-spacing': [ 'error', 'always' ],
        'react/jsx-max-props-per-line': 'off',
        'react/jsx-closing-bracket-location': [
            'error',
            {
                nonEmpty: false,
                selfClosing: 'line-aligned',
            },
        ],
        'object-curly-newline': [
            'error',
            {
                multiline: true,
                consistent: true,
            },
        ],
        "@typescript-eslint/no-explicit-any": "off",
        "@typescript-eslint/no-non-null-assertion": "off",
        "@typescript-eslint/no-empty-function": "off"
    },
    parserOptions: {
        ecmaVersion: 2020,
        sourceType: 'module',
        project: './tsconfig.json',
        tsconfigRootDir: __dirname,
        createDefaultProgram: true,
    },
    settings: {
        'import/resolver': {
            node: {},
            webpack: {
                config: require.resolve('./.erb/configs/webpack.config.eslint.ts'),
            },
            typescript: {},
        },
        'import/parsers': {
            '@typescript-eslint/parser': [ '.ts', '.tsx' ],
        },
    },
};
