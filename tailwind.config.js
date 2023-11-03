/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      'components/**/*.templ',
    ],
    darkMode: 'class',
    theme: {
      extend: {
        fontFamily: {
          mono: ['Courier Prime', 'monospace'],
        }
      },
    },
    corePlugins: {
      preflight: true,
    }
  }