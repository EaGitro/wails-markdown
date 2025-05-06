import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
// import { Greet } from '../wailsjs/go/main/App';

import { GetContent } from "../wailsjs/go/main/OperateFile"

import markdownit from "markdown-it";
import hljs from 'highlight.js';

const md = markdownit({
    highlight: function (str, lang): string {
        if (lang && hljs.getLanguage(lang)) {
            try {
                return '<pre>' + lang + '<code class="hljs">' +
                    hljs.highlight(str, { language: lang, ignoreIllegals: true }).value +
                    '</code></pre>';
            } catch (__) { }
        }

        return '<pre><code class="hljs">' + md.utils.escapeHtml(str) + '</code></pre>';
    }
});

document.addEventListener('keydown', e => {
    if (e.ctrlKey && e.key === 's') {
        // Prevent the Save dialog to open
        e.preventDefault();
        // Place your code here
        console.log('CTRL + S');
    }
});

document.addEventListener('keydown', e => {
    if (e.ctrlKey && e.key === 'r') {
        // Prevent the Save dialog to open
        e.preventDefault();
        // Place your code here
        console.log('CTRL + R: js captured');
        // reloadMd()
    }
});

// User's reload operation
window.reloadMd = function () {
    try {
        GetContent().then((res) => {
            document.getElementById("app")!.innerHTML = md.render(res.FileContent);
        }).catch((err) => {
            console.error(err);
        });
    } catch (err) {
        console.error(err);
    }
};


window.reloadMd();

/**
 * TODO: Hot reload
// Hot reload 
window.setMdContent = function (content: string) {
try {
    document.getElementById("app")!.innerHTML = md.render(content)
} catch (error) {
    console.log(error)
}
}
**/

// // Setup the greet function
// window.greet = function () {
//     // Get name
//     let name = nameElement!.value;

//     // Check if the input is empty
//     if (name === "") return;

//     // Call App.Greet(name)
//     try {
//         Greet(name)
//             .then((result) => {
//                 // Update result with data back from App.Greet()
//                 resultElement!.innerText = result;
//             })
//             .catch((err) => {
//                 console.error(err);
//             });
//     } catch (err) {
//         console.error(err);
//     }
// };

// document.querySelector('#app')!.innerHTML = `
//     <img id="logo" class="logo">
//       <div class="result" id="result">Please enter your name below 👇</div>
//       <div class="input-box" id="input">
//         <input class="input" id="name" type="text" autocomplete="off" />
//         <button class="btn" onclick="greet()">Greet</button>
//       </div>
//     </div>
// `;
{ (document.getElementById('logo') as HTMLImageElement).src = logo; }

// let nameElement = (document.getElementById("name") as HTMLInputElement);
// nameElement.focus();
// let resultElement = document.getElementById("result");

declare global {
    interface Window {
        reloadMd: () => void;
        setMdContent: (content:string) => void;
    }
}
