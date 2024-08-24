/** @type {import('tailwindcss').Config} */
// https://github.com/L-Blondy/tw-colors
const {createThemes} = require('tw-colors');
module.exports = {
    content: ['**/*.{html,templ}', './node_modules/flowbite/**/*.js'],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
        require('flowbite/plugin'),
        // colors from: https://github.com/saadeghi/daisyui/blob/master/src/theming/themes.js
        //  但是不支持oklch，需要通过这个在线工具转换成RGB颜色
        createThemes({
            aqua: {
                "color-scheme": "dark",
                "primary": "#09ecf3",
                "primary-content": "#005355",
                "secondary": "#966fb3",
                "accent": "#ffe999",
                "neutral": "#3b8ac4",
                "base-100": "#345da7",
                "info": "#2563eb",
                "success": "#16a34a",
                "warning": "#d97706",
                "error": "rgb(255, 117, 104)",
            },
            black: {
                "color-scheme": "dark",
                "primary": "#373737",
                "secondary": "#373737",
                "accent": "#373737",
                "base-100": "#000000",
                "base-200": "#141414",
                "base-300": "#262626",
                "base-content": "#d6d6d6",
                "neutral": "#373737",
                "info": "#0000ff",
                "success": "#008000",
                "warning": "#ffff00",
                "error": "#ff0000",
                // "--rounded-box": "0",
                // "--rounded-btn": "0",
                // "--rounded-badge": "0",
                // "--animation-btn": "0",
                // "--animation-input": "0",
                // "--btn-focus-scale": "1",
                // "--tab-radius": "0",
            },
            bumblebee: {
                "color-scheme": "light",
                "primary": "rgb(255, 218, 0)",
                "primary-content": "rgb(76, 69, 40)",
                "secondary": "rgb(255, 165, 0)",
                "secondary-content": "rgb(93, 63, 24)",
                "accent": "rgb(255, 166, 86)",
                "neutral": "rgb(6, 0, 35)",
                "base-100": "oklch(100% 0 0)",
            },
            cmyk: {
                "color-scheme": "light",
                "primary": "#45AEEE",
                "secondary": "#E8488A",
                "accent": "#FFF232",
                "neutral": "#1a1a1a",
                "base-100": "rgb(255, 255, 255)",
                "info": "#4AA8C0",
                "success": "#823290",
                "warning": "#EE8133",
                "error": "#E93F33",
            },
            corporate: {
                "color-scheme": "light",
                "primary": "rgb(77, 110, 255)",
                "secondary": "#7b92b2",
                "accent": "#67cba0",
                "neutral": "#181a2a",
                "neutral-content": "#edf2f7",
                "base-100": "rgb(255, 255, 255)",
                "base-content": "#181a2a",
                // "--rounded-box": "0.25rem",
                // "--rounded-btn": ".125rem",
                // "--rounded-badge": ".125rem",
                // "--tab-radius": "0.25rem",
                // "--animation-btn": "0",
                // "--animation-input": "0",
                // "--btn-focus-scale": "1",
            },
            cupcake: {
                "color-scheme": "light",
                "primary": "#65c3c8",
                "secondary": "#ef9fbc",
                "accent": "#eeaf3a",
                "neutral": "#291334",
                "base-100": "#faf7f5",
                "base-200": "#efeae6",
                "base-300": "#e7e2df",
                "base-content": "#291334",
                "--rounded-btn": "1.9rem",
                "--tab-border": "2px",
                "--tab-radius": "0.7rem",
            },
            cyberpunk: {
                "color-scheme": "light",
                "fontFamily":
                    "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,Liberation Mono,Courier New,monospace",
                "primary": "rgb(255, 111, 152)",
                "secondary": "rgb(1, 229, 249)",
                "accent": "rgb(206, 116, 255)",
                "neutral": "rgb(17, 26, 59)",
                "neutral-content": "rgb(255, 242, 72)",
                "base-100": "rgb(255, 242, 72)",
                "base-content": "#282425",
                // "--rounded-box": "0",
                // "--rounded-btn": "0",
                // "--rounded-badge": "0",
                // "--tab-radius": "0",
            },
            retro: {
                "color-scheme": "light",
                "primary": "#ef9995",
                "primary-content": "#282425",
                "secondary": "#a4cbb4",
                "secondary-content": "#282425",
                "accent": "#DC8850",
                "accent-content": "#282425",
                "neutral": "#2E282A",
                "neutral-content": "#EDE6D4",
                "base-100": "#ece3ca",
                "base-200": "#e4d8b4",
                "base-300": "#DBCA9A",
                "base-content": "#282425",
                "info": "#2563eb",
                "success": "#16a34a",
                "warning": "#d97706",
                "error": "#f35248",
                // "--rounded-box": "0.4rem",
                // "--rounded-btn": "0.4rem",
                // "--rounded-badge": "0.4rem",
                // "--tab-radius": "0.4rem",
            },
            dark: {
                "color-scheme": "dark",
                "primary": "rgb(116, 128, 255)",
                "secondary": "rgb(255, 98, 212)",
                "accent": "rgb(0, 202, 182)",
                "neutral": "#2a323c",
                "neutral-content": "#A6ADBB",
                "base-100": "#1d232a",
                "base-200": "#191e24",
                "base-300": "#15191e",
                "base-content": "#A6ADBB",
            },
            dracula: {
                "color-scheme": "dark",
                "primary": "#ff79c6",
                "secondary": "#bd93f9",
                "accent": "#ffb86c",
                "neutral": "#414558",
                "base-100": "#282a36",
                "base-content": "#f8f8f2",
                "info": "#8be9fd",
                "success": "#50fa7b",
                "warning": "#f1fa8c",
                "error": "#ff5555",
            },
            emerald: {
                "color-scheme": "light",
                "primary": "#66cc8a",
                "primary-content": "#223D30",
                "secondary": "#377cfb",
                "secondary-content": "#fff",
                "accent": "#f68067",
                "accent-content": "#000",
                "neutral": "#333c4d",
                "neutral-content": "#f9fafb",
                "base-100": "rgb(255, 255, 255)",
                "base-content": "#333c4d",
                "--animation-btn": "0",
                "--animation-input": "0",
                "--btn-focus-scale": "1",
            },
            fantasy: {
                "color-scheme": "light",
                "primary": "rgb(109, 0, 118)",
                "secondary": "rgb(0, 117, 191)",
                "accent": "rgb(255, 137, 0)",
                "neutral": "#1f2937",
                "base-100": "rgb(255, 255, 255)",
                "base-content": "#1f2937",
            },
            forest: {
                "color-scheme": "dark",
                "primary": "#1eb854",
                "primary-content": "#000000",
                "secondary": "#1DB88E",
                "accent": "#1DB8AB",
                "neutral": "#19362D",
                "base-100": "#171212",
                // "--rounded-btn": "1.9rem",
            },
            garden: {
                "color-scheme": "light",
                "primary": "rgb(251, 0, 117)",
                "primary-content": "#fff",
                "secondary": "#8E4162",
                "accent": "#5c7f67",
                "neutral": "#291E00",
                "neutral-content": "#e9e7e7",
                "base-100": "#e9e7e7",
                "base-content": "#100f0f",
            },
            halloween: {
                "color-scheme": "dark",
                "primary": "rgb(255, 147, 0)",
                "primary-content": "#131616",
                "secondary": "rgb(122, 0, 194)",
                "accent": "rgb(68, 169, 0)",
                "accent-content": "#000000",
                "neutral": "#2F1B05",
                "base-100": "#212121",
                "base-content": "#ff7557",
                "info": "#2563eb",
                "success": "#16a34a",
                "warning": "#d97706",
                "error": "rgb(243, 82, 72)",
            },
            light: {
                "color-scheme": "light",
                "primary": "rgb(74, 0, 255)",
                "secondary": "rgb(255, 30, 204)",
                "secondary-content": "rgb(255, 248, 253)",
                "accent": "rgb(0, 211, 190)",
                "neutral": "#2B3440",
                "neutral-content": "#D7DDE4",
                "base-100": "rgb(255, 255, 255)",
                "base-200": "#F2F2F2",
                "base-300": "#E5E6E6",
                "base-content": "#1f2937",
            },
            lofi: {
                "color-scheme": "light",
                "primary": "#0D0D0D",
                "primary-content": "rgb(255, 255, 255)",
                "secondary": "#1A1919",
                "secondary-content": "rgb(255, 255, 255)",
                "accent": "#262626",
                "accent-content": "oklch(100% 0 0)",
                "neutral": "#000000",
                "neutral-content": "rgb(255, 255, 255)",
                "base-100": "rgb(255, 255, 255)",
                "base-200": "#F2F2F2",
                "base-300": "#E6E5E5",
                "base-content": "#000000",
                "info": "rgb(95, 207, 221)",
                "success": "rgb(105, 254, 195)",
                "warning": "rgb(255, 206, 105)",
                "error": "rgb(255, 146, 130)",
                // "--rounded-box": "0.25rem",
                // "--rounded-btn": "0.125rem",
                // "--rounded-badge": "0.125rem",
                // "--tab-radius": "0.125rem",
                // "--animation-btn": "0",
                // "--animation-input": "0",
                // "--btn-focus-scale": "1",
            },
            luxury: {
                "color-scheme": "dark",
                "primary": "rgb(255, 255, 255)",
                "secondary": "#152747",
                "accent": "#513448",
                "neutral": "#331800",
                "neutral-content": "#FFE7A3",
                "base-100": "#09090b",
                "base-200": "#171618",
                "base-300": "#2e2d2f",
                "base-content": "#dca54c",
                "info": "#66c6ff",
                "success": "#87d039",
                "warning": "#e2d562",
                "error": "#ff6f6f",
            },
            pastel: {
                "color-scheme": "light",
                "primary": "#d1c1d7",
                "secondary": "#f6cbd1",
                "accent": "#b4e9d6",
                "neutral": "#70acc7",
                "base-100": "rgb(255, 255, 255)",
                "base-200": "#f9fafb",
                "base-300": "#d1d5db",
                // "--rounded-btn": "1.9rem",
                // "--tab-radius": "0.7rem",
            },
            synthwave: {
                "color-scheme": "dark",
                "primary": "#e779c1",
                "secondary": "#58c7f3",
                "accent": "rgb(255, 210, 0)",
                "neutral": "#221551",
                "neutral-content": "#f9f7fd",
                "base-100": "#1a103d",
                "base-content": "#f9f7fd",
                "info": "#53c0f3",
                "info-content": "#201047",
                "success": "#71ead2",
                "success-content": "#201047",
                "warning": "#eace6c",
                "warning-content": "#201047",
                "error": "#ec8c78",
                "error-content": "#201047",
            },
            valentine: {
                "color-scheme": "light",
                "primary": "#e96d7b",
                "secondary": "#a991f7",
                "accent": "#66b1b3",
                "neutral": "#af4670",
                "neutral-content": "#f0d6e8",
                "base-100": "#fae7f4",
                "base-content": "#632c3b",
                "info": "#2563eb",
                "success": "#16a34a",
                "warning": "#d97706",
                "error": "rgb(255, 111, 98)",
                "--rounded-btn": "1.9rem",
                "--tab-radius": "0.7rem",
            },
            wireframe: {
                "color-scheme": "light",
                "fontFamily": "Chalkboard,comic sans ms,'sans-serif'",
                "primary": "#b8b8b8",
                "secondary": "#b8b8b8",
                "accent": "#b8b8b8",
                "neutral": "#ebebeb",
                "base-100": "rgb(255, 255, 255)",
                "base-200": "#eeeeee",
                "base-300": "#dddddd",
                "info": "#0000ff",
                "success": "#008000",
                "warning": "#a6a659",
                "error": "#ff0000",
                // "--rounded-box": "0.2rem",
                // "--rounded-btn": "0.2rem",
                // "--rounded-badge": "0.2rem",
                // "--tab-radius": "0.2rem",
            },
            autumn: {
                "color-scheme": "light",
                "primary": "#8C0327",
                "secondary": "#D85251",
                "accent": "#D59B6A",
                "neutral": "#826A5C",
                "base-100": "#f1f1f1",
                "info": "#42ADBB",
                "success": "#499380",
                "warning": "#E97F14",
                "error": "oklch(53.07% 0.241 24.16)",
            },
            business: {
                "color-scheme": "dark",
                "primary": "#1C4E80",
                "secondary": "#7C909A",
                "accent": "#EA6947",
                "neutral": "#23282E",
                "base-100": "#202020",
                "info": "#0091D5",
                "success": "#6BB187",
                "warning": "#DBAE59",
                "error": "#AC3E31",
                // "--rounded-box": "0.25rem",
                // "--rounded-btn": ".125rem",
                // "--rounded-badge": ".125rem",
            },
            acid: {
                "color-scheme": "light",
                "primary": "oklch(71.9% 0.357 330.7595734057481)",
                "secondary": "oklch(73.37% 0.224 48.25087840015526)",
                "accent": "oklch(92.78% 0.264 122.96295065960891)",
                "neutral": "oklch(21.31% 0.128 278.68)",
                "base-100": "#fafafa",
                "info": "oklch(60.72% 0.227 252.05)",
                "success": "oklch(85.72% 0.266 158.53)",
                "warning": "oklch(91.01% 0.212 100.5)",
                "error": "oklch(64.84% 0.293 29.34918758658804)",
                // "--rounded-box": "1.25rem",
                // "--rounded-btn": "1rem",
                // "--rounded-badge": "1rem",
                // "--tab-radius": "0.7rem",
            },
            lemonade: {
                "color-scheme": "light",
                "primary": "oklch(58.92% 0.199 134.6)",
                "secondary": "oklch(77.75% 0.196 111.09)",
                "accent": "oklch(85.39% 0.201 100.73)",
                "neutral": "oklch(30.98% 0.075 108.6)",
                "base-100": "oklch(98.71% 0.02 123.72)",
                "info": "oklch(86.19% 0.047 224.14)",
                "success": "oklch(86.19% 0.047 157.85)",
                "warning": "oklch(86.19% 0.047 102.15)",
                "error": "oklch(86.19% 0.047 25.85)",
            },
            night: {
                "color-scheme": "dark",
                "primary": "#38bdf8",
                "secondary": "#818CF8",
                "accent": "#F471B5",
                "neutral": "#1E293B",
                "base-100": "#0F172A",
                "info": "#0CA5E9",
                "info-content": "#000000",
                "success": "#2DD4BF",
                "warning": "#F4BF50",
                "error": "#FB7085",
            },
            coffee: {
                "color-scheme": "dark",
                "primary": "#DB924B",
                "secondary": "#263E3F",
                "accent": "#10576D",
                "neutral": "#120C12",
                "base-100": "#20161F",
                "base-content": "#c59f60",
                "info": "#8DCAC1",
                "success": "#9DB787",
                "warning": "#FFD25F",
                "error": "#FC9581",
            },
            winter: {
                "color-scheme": "light",
                "primary": "rgb(0, 105, 255)",
                "secondary": "#463AA2",
                "accent": "#C148AC",
                "neutral": "#021431",
                "base-100": "rgb(255, 255, 255)",
                "base-200": "#F2F7FF",
                "base-300": "#E3E9F4",
                "base-content": "#394E6A",
                "info": "#93E7FB",
                "success": "#81CFD1",
                "warning": "#EFD7BB",
                "error": "#E58B8B",
            },
            dim: {
                "color-scheme": "dark",
                "primary": "#9FE88D",
                "secondary": "#FF7D5C",
                "accent": "#C792E9",
                "neutral": "#1c212b",
                "neutral-content": "#B2CCD6",
                "base-100": "#2A303C",
                "base-200": "#242933",
                "base-300": "#20252E",
                "base-content": "#B2CCD6",
                "info": "#28ebff",
                "success": "#62efbd",
                "warning": "#efd057",
                "error": "#ffae9b",
            },
            nord: {
                "color-scheme": "light",
                "primary": "#5E81AC",
                "secondary": "#81A1C1",
                "accent": "#88C0D0",
                "neutral": "#4C566A",
                "neutral-content": "#D8DEE9",
                "base-100": "#ECEFF4",
                "base-200": "#E5E9F0",
                "base-300": "#D8DEE9",
                "base-content": "#2E3440",
                "info": "#B48EAD",
                "success": "#A3BE8C",
                "warning": "#EBCB8B",
                "error": "#BF616A",
                // "--rounded-box": "0.4rem",
                // "--rounded-btn": "0.2rem",
                // "--rounded-badge": "0.4rem",
                // "--tab-radius": "0.2rem",
            },
            sunset: {
                "color-scheme": "dark",
                "primary": "#FF865B",
                "secondary": "#FD6F9C",
                "accent": "#B387FA",
                "neutral": "rgb(27, 38, 44)",
                "neutral-content": "rgb(148, 160, 169)",
                "base-100": "rgb(18, 28, 34)",
                "base-200": "rgb(14, 23, 30)",
                "base-300": "rgb(9, 19, 25)",
                "base-content": "#9fb9d0",
                "info": "#89e0eb",
                "success": "#addfad",
                "warning": "#f1c891",
                "error": "#ffbbbd",
                // "--rounded-box": "1.2rem",
                // "--rounded-btn": "0.8rem",
                // "--rounded-badge": "0.4rem",
                // "--tab-radius": "0.7rem",
            },
        })
    ],
}


