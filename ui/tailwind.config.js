/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./app/**/*.{js,ts,jsx,tsx}', './pages/**/*.{js,ts,jsx,tsx}'],
  theme: {
    container: {
      center: true
    },
    fontFamily: {
      sans: [
        'Hiragino Kaku Gothic ProN',
        'Hiragino Sans',
        'Segoe UI',
        'Roboto',
        'Noto Sans CJK JP',
        'sans-serif',
        'Apple Color Emoji',
        'Segoe UI Emoji',
        'Segoe UI Symbol',
        'Noto Sans Emoji'
      ]
    },
    extend: {}
  },
  plugins: []
}
