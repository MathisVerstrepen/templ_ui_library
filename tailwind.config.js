/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./components/**/*.{html,js,templ}", "./ui/**/*.{html,js,templ}"],
    theme: {
        extend: {
            colors: {
                "c-grey": "#222222",
                "c-cyan": "#40BCF4",
                "c-red": "#F44336",
            },
        },
    },
    plugins: [],
};
