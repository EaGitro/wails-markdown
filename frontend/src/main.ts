import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
// import { Greet } from '../wailsjs/go/main/App';

import { GetContent } from "../wailsjs/go/main/OperateFile"

import markdownit from "markdown-it";
import hljs from 'highlight.js';
import { riscvasm } from './highlightjs-riscvasm';
import type * as CSS from 'csstype';


function css2str(css: CSS.PropertiesHyphen){
    let propSets = []
    for(const prop in css){
        const val = css[prop as keyof typeof css];
        propSets.push(`${prop}: ${val}`)
    }
    return propSets.join("; ")
}




//
// start: markdown-it 
// 

hljs.registerLanguage("riscv", riscvasm);
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



let defaultRenderLinkOpen = md.renderer.rules.link_open || function (tokens, idx, options, _env, self) {
    return self.renderToken(tokens, idx, options);
};
md.renderer.rules.link_open = function (tokens, idx, options, env, self) {
    // Below is an official example from https://github.com/markdown-it/markdown-it/blob/master/docs/architecture.md
    // Add a new `target` attribute, or replace the value of the existing one.
    // tokens[idx].attrSet('target', '_blank');

    tokens[idx].attrSet("onclick",                          // Opens the URL in the system browser. 
        `window.runtime.BrowserOpenURL("${tokens[idx].attrGet("href")}"); return false`    // see: https://wails.io/docs/reference/runtime/browser/
    )


    console.log({tokens, idx, options, env, self});
    // Pass the token to the default renderer.
    return defaultRenderLinkOpen(tokens, idx, options, env, self);
};



let defaultRenderBlockquoteOpen = md.renderer.rules.blockquote_open || function (tokens, idx, options, _env, self) {
    return self.renderToken(tokens, idx, options);
};
md.renderer.rules.blockquote_open = function (tokens, idx, options, env, self) {
    // console.log("blockquote_open")
    // console.log({tokens, idx, options, env, self});

    tokens[idx].attrSet("style", css2str({
        "border-left": "thick double #32a1ce",
        "padding": "1em"
    }))

    return defaultRenderBlockquoteOpen(tokens, idx, options, env, self);

}


window._md = md;

// 
// end : markdown-it
// 

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
            // console.log(res)
            if (!res.Err && res.IsModified)
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
        setMdContent: (content: string) => void;
        runtime: any
        _md: any;
    }
}
