(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[312],{656:(e,t,n)=>{"use strict";n.r(t),n.d(t,{default:()=>s});n(1504);var o=n(3664),i=n(7624);function s(e){let{children:t,fallback:n}=e;return(0,o.c)()?(0,i.jsx)(i.Fragment,{children:t?.()}):n??null}},304:(e,t,n)=>{"use strict";n.d(t,{c:()=>E});var o=n(1504),i=n(7624);function s(e){const{mdxAdmonitionTitle:t,rest:n}=function(e){const t=o.Children.toArray(e),n=t.find((e=>o.isValidElement(e)&&"mdxAdmonitionTitle"===e.type)),s=t.filter((e=>e!==n)),c=n?.props.children;return{mdxAdmonitionTitle:c,rest:s.length>0?(0,i.jsx)(i.Fragment,{children:s}):null}}(e.children),s=e.title??t;return{...e,...s&&{title:s},children:n}}var c=n(5456),r=n(4357),a=n(5864);const l={admonition:"admonition_xJq3",admonitionHeading:"admonitionHeading_Gvgb",admonitionIcon:"admonitionIcon_Rf37",admonitionContent:"admonitionContent_BuS1"};function u(e){let{type:t,className:n,children:o}=e;return(0,i.jsx)("div",{className:(0,c.c)(a.W.common.admonition,a.W.common.admonitionType(t),l.admonition,n),children:o})}function d(e){let{icon:t,title:n}=e;return(0,i.jsxs)("div",{className:l.admonitionHeading,children:[(0,i.jsx)("span",{className:l.admonitionIcon,children:t}),n]})}function m(e){let{children:t}=e;return t?(0,i.jsx)("div",{className:l.admonitionContent,children:t}):null}function f(e){const{type:t,icon:n,title:o,children:s,className:c}=e;return(0,i.jsxs)(u,{type:t,className:c,children:[(0,i.jsx)(d,{title:o,icon:n}),(0,i.jsx)(m,{children:s})]})}function h(e){return(0,i.jsx)("svg",{viewBox:"0 0 14 16",...e,children:(0,i.jsx)("path",{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"})})}const p={icon:(0,i.jsx)(h,{}),title:(0,i.jsx)(r.c,{id:"theme.admonition.note",description:"The default label used for the Note admonition (:::note)",children:"note"})};function g(e){return(0,i.jsx)(f,{...p,...e,className:(0,c.c)("alert alert--secondary",e.className),children:e.children})}function b(e){return(0,i.jsx)("svg",{viewBox:"0 0 12 16",...e,children:(0,i.jsx)("path",{fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"})})}const v={icon:(0,i.jsx)(b,{}),title:(0,i.jsx)(r.c,{id:"theme.admonition.tip",description:"The default label used for the Tip admonition (:::tip)",children:"tip"})};function x(e){return(0,i.jsx)(f,{...v,...e,className:(0,c.c)("alert alert--success",e.className),children:e.children})}function j(e){return(0,i.jsx)("svg",{viewBox:"0 0 14 16",...e,children:(0,i.jsx)("path",{fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"})})}const k={icon:(0,i.jsx)(j,{}),title:(0,i.jsx)(r.c,{id:"theme.admonition.info",description:"The default label used for the Info admonition (:::info)",children:"info"})};function y(e){return(0,i.jsx)(f,{...k,...e,className:(0,c.c)("alert alert--info",e.className),children:e.children})}function w(e){return(0,i.jsx)("svg",{viewBox:"0 0 16 16",...e,children:(0,i.jsx)("path",{fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"})})}const N={icon:(0,i.jsx)(w,{}),title:(0,i.jsx)(r.c,{id:"theme.admonition.warning",description:"The default label used for the Warning admonition (:::warning)",children:"warning"})};function C(e){return(0,i.jsx)("svg",{viewBox:"0 0 12 16",...e,children:(0,i.jsx)("path",{fillRule:"evenodd",d:"M5.05.31c.81 2.17.41 3.38-.52 4.31C3.55 5.67 1.98 6.45.9 7.98c-1.45 2.05-1.7 6.53 3.53 7.7-2.2-1.16-2.67-4.52-.3-6.61-.61 2.03.53 3.33 1.94 2.86 1.39-.47 2.3.53 2.27 1.67-.02.78-.31 1.44-1.13 1.81 3.42-.59 4.78-3.42 4.78-5.56 0-2.84-2.53-3.22-1.25-5.61-1.52.13-2.03 1.13-1.89 2.75.09 1.08-1.02 1.8-1.86 1.33-.67-.41-.66-1.19-.06-1.78C8.18 5.31 8.68 2.45 5.05.32L5.03.3l.02.01z"})})}const L={icon:(0,i.jsx)(C,{}),title:(0,i.jsx)(r.c,{id:"theme.admonition.danger",description:"The default label used for the Danger admonition (:::danger)",children:"danger"})};const _={icon:(0,i.jsx)(w,{}),title:(0,i.jsx)(r.c,{id:"theme.admonition.caution",description:"The default label used for the Caution admonition (:::caution)",children:"caution"})};const B={...{note:g,tip:x,info:y,warning:function(e){return(0,i.jsx)(f,{...N,...e,className:(0,c.c)("alert alert--warning",e.className),children:e.children})},danger:function(e){return(0,i.jsx)(f,{...L,...e,className:(0,c.c)("alert alert--danger",e.className),children:e.children})}},...{secondary:e=>(0,i.jsx)(g,{title:"secondary",...e}),important:e=>(0,i.jsx)(y,{title:"important",...e}),success:e=>(0,i.jsx)(x,{title:"success",...e}),caution:function(e){return(0,i.jsx)(f,{..._,...e,className:(0,c.c)("alert alert--warning",e.className),children:e.children})}}};function E(e){const t=s(e),n=(o=t.type,B[o]||(console.warn(`No admonition component found for admonition type "${o}". Using Info as fallback.`),B.info));var o;return(0,i.jsx)(n,{...t})}},1608:(e,t,n)=>{"use strict";n.r(t),n.d(t,{default:()=>$});var o=n(1504),i=n(3664),s=n(5456),c=n(6528),r=n(1824);function a(){const{prism:e}=(0,r.y)(),{colorMode:t}=(0,c.U)(),n=e.theme,o=e.darkTheme||n;return"dark"===t?o:n}var l=n(5864),u=n(6504),d=n.n(u);const m=/title=(?<quote>["'])(?<title>.*?)\1/,f=/\{(?<range>[\d,-]+)\}/,h={js:{start:"\\/\\/",end:""},jsBlock:{start:"\\/\\*",end:"\\*\\/"},jsx:{start:"\\{\\s*\\/\\*",end:"\\*\\/\\s*\\}"},bash:{start:"#",end:""},html:{start:"\x3c!--",end:"--\x3e"}},p={...h,lua:{start:"--",end:""},wasm:{start:"\\;\\;",end:""},tex:{start:"%",end:""},vb:{start:"['\u2018\u2019]",end:""},vbnet:{start:"(?:_\\s*)?['\u2018\u2019]",end:""},rem:{start:"[Rr][Ee][Mm]\\b",end:""},f90:{start:"!",end:""},ml:{start:"\\(\\*",end:"\\*\\)"},cobol:{start:"\\*>",end:""}},g=Object.keys(h);function b(e,t){const n=e.map((e=>{const{start:n,end:o}=p[e];return`(?:${n}\\s*(${t.flatMap((e=>[e.line,e.block?.start,e.block?.end].filter(Boolean))).join("|")})\\s*${o})`})).join("|");return new RegExp(`^\\s*(?:${n})\\s*$`)}function v(e,t){let n=e.replace(/\n$/,"");const{language:o,magicComments:i,metastring:s}=t;if(s&&f.test(s)){const e=s.match(f).groups.range;if(0===i.length)throw new Error(`A highlight range has been given in code block's metastring (\`\`\` ${s}), but no magic comment config is available. Docusaurus applies the first magic comment entry's className for metastring ranges.`);const t=i[0].className,o=d()(e).filter((e=>e>0)).map((e=>[e-1,[t]]));return{lineClassNames:Object.fromEntries(o),code:n}}if(void 0===o)return{lineClassNames:{},code:n};const c=function(e,t){switch(e){case"js":case"javascript":case"ts":case"typescript":return b(["js","jsBlock"],t);case"jsx":case"tsx":return b(["js","jsBlock","jsx"],t);case"html":return b(["js","jsBlock","html"],t);case"python":case"py":case"bash":return b(["bash"],t);case"markdown":case"md":return b(["html","jsx","bash"],t);case"tex":case"latex":case"matlab":return b(["tex"],t);case"lua":case"haskell":case"sql":return b(["lua"],t);case"wasm":return b(["wasm"],t);case"vb":case"vba":case"visual-basic":return b(["vb","rem"],t);case"vbnet":return b(["vbnet","rem"],t);case"batch":return b(["rem"],t);case"basic":return b(["rem","f90"],t);case"fsharp":return b(["js","ml"],t);case"ocaml":case"sml":return b(["ml"],t);case"fortran":return b(["f90"],t);case"cobol":return b(["cobol"],t);default:return b(g,t)}}(o,i),r=n.split("\n"),a=Object.fromEntries(i.map((e=>[e.className,{start:0,range:""}]))),l=Object.fromEntries(i.filter((e=>e.line)).map((e=>{let{className:t,line:n}=e;return[n,t]}))),u=Object.fromEntries(i.filter((e=>e.block)).map((e=>{let{className:t,block:n}=e;return[n.start,t]}))),m=Object.fromEntries(i.filter((e=>e.block)).map((e=>{let{className:t,block:n}=e;return[n.end,t]})));for(let d=0;d<r.length;){const e=r[d].match(c);if(!e){d+=1;continue}const t=e.slice(1).find((e=>void 0!==e));l[t]?a[l[t]].range+=`${d},`:u[t]?a[u[t]].start=d:m[t]&&(a[m[t]].range+=`${a[m[t]].start}-${d-1},`),r.splice(d,1)}n=r.join("\n");const h={};return Object.entries(a).forEach((e=>{let[t,{range:n}]=e;d()(n).forEach((e=>{h[e]??=[],h[e].push(t)}))})),{lineClassNames:h,code:n}}const x={codeBlockContainer:"codeBlockContainer_Ckt0"};var j=n(7624);function k(e){let{as:t,...n}=e;const o=function(e){const t={color:"--prism-color",backgroundColor:"--prism-background-color"},n={};return Object.entries(e.plain).forEach((e=>{let[o,i]=e;const s=t[o];s&&"string"==typeof i&&(n[s]=i)})),n}(a());return(0,j.jsx)(t,{...n,style:o,className:(0,s.c)(n.className,x.codeBlockContainer,l.W.common.codeBlock)})}const y={codeBlockContent:"codeBlockContent_biex",codeBlockTitle:"codeBlockTitle_Ktv7",codeBlock:"codeBlock_bY9V",codeBlockStandalone:"codeBlockStandalone_MEMb",codeBlockLines:"codeBlockLines_e6Vv",codeBlockLinesWithNumbering:"codeBlockLinesWithNumbering_o6Pm",buttonGroup:"buttonGroup__atx"};function w(e){let{children:t,className:n}=e;return(0,j.jsx)(k,{as:"pre",tabIndex:0,className:(0,s.c)(y.codeBlockStandalone,"thin-scrollbar",n),children:(0,j.jsx)("code",{className:y.codeBlockLines,children:t})})}var N=n(1100);const C={attributes:!0,characterData:!0,childList:!0,subtree:!0};function L(e,t){const[n,i]=(0,o.useState)(),s=(0,o.useCallback)((()=>{i(e.current?.closest("[role=tabpanel][hidden]"))}),[e,i]);(0,o.useEffect)((()=>{s()}),[s]),function(e,t,n){void 0===n&&(n=C);const i=(0,N.yA)(t),s=(0,N.Mh)(n);(0,o.useEffect)((()=>{const t=new MutationObserver(i);return e&&t.observe(e,s),()=>t.disconnect()}),[e,i,s])}(n,(e=>{e.forEach((e=>{"attributes"===e.type&&"hidden"===e.attributeName&&(t(),s())}))}),{attributes:!0,characterData:!1,childList:!1,subtree:!1})}var _=n(5720);const B={codeLine:"codeLine_lJS_",codeLineNumber:"codeLineNumber_Tfdd",codeLineContent:"codeLineContent_feaV"};function E(e){let{line:t,classNames:n,showLineNumbers:o,getLineProps:i,getTokenProps:c}=e;1===t.length&&"\n"===t[0].content&&(t[0].content="");const r=i({line:t,className:(0,s.c)(n,o&&B.codeLine)}),a=t.map(((e,t)=>(0,j.jsx)("span",{...c({token:e,key:t})},t)));return(0,j.jsxs)("span",{...r,children:[o?(0,j.jsxs)(j.Fragment,{children:[(0,j.jsx)("span",{className:B.codeLineNumber}),(0,j.jsx)("span",{className:B.codeLineContent,children:a})]}):a,(0,j.jsx)("br",{})]})}var T=n(4357);function H(e){return(0,j.jsx)("svg",{viewBox:"0 0 24 24",...e,children:(0,j.jsx)("path",{fill:"currentColor",d:"M19,21H8V7H19M19,5H8A2,2 0 0,0 6,7V21A2,2 0 0,0 8,23H19A2,2 0 0,0 21,21V7A2,2 0 0,0 19,5M16,1H4A2,2 0 0,0 2,3V17H4V3H16V1Z"})})}function S(e){return(0,j.jsx)("svg",{viewBox:"0 0 24 24",...e,children:(0,j.jsx)("path",{fill:"currentColor",d:"M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"})})}const R={copyButtonCopied:"copyButtonCopied_obH4",copyButtonIcons:"copyButtonIcons_eSgA",copyButtonIcon:"copyButtonIcon_y97N",copyButtonSuccessIcon:"copyButtonSuccessIcon_LjdS"};function M(e){let{code:t,className:n}=e;const[i,c]=(0,o.useState)(!1),r=(0,o.useRef)(void 0),a=(0,o.useCallback)((()=>{!function(e,t){let{target:n=document.body}=void 0===t?{}:t;if("string"!=typeof e)throw new TypeError(`Expected parameter \`text\` to be a \`string\`, got \`${typeof e}\`.`);const o=document.createElement("textarea"),i=document.activeElement;o.value=e,o.setAttribute("readonly",""),o.style.contain="strict",o.style.position="absolute",o.style.left="-9999px",o.style.fontSize="12pt";const s=document.getSelection(),c=s.rangeCount>0&&s.getRangeAt(0);n.append(o),o.select(),o.selectionStart=0,o.selectionEnd=e.length;let r=!1;try{r=document.execCommand("copy")}catch{}o.remove(),c&&(s.removeAllRanges(),s.addRange(c)),i&&i.focus()}(t),c(!0),r.current=window.setTimeout((()=>{c(!1)}),1e3)}),[t]);return(0,o.useEffect)((()=>()=>window.clearTimeout(r.current)),[]),(0,j.jsx)("button",{type:"button","aria-label":i?(0,T.G)({id:"theme.CodeBlock.copied",message:"Copied",description:"The copied button label on code blocks"}):(0,T.G)({id:"theme.CodeBlock.copyButtonAriaLabel",message:"Copy code to clipboard",description:"The ARIA label for copy code blocks button"}),title:(0,T.G)({id:"theme.CodeBlock.copy",message:"Copy",description:"The copy button label on code blocks"}),className:(0,s.c)("clean-btn",n,R.copyButton,i&&R.copyButtonCopied),onClick:a,children:(0,j.jsxs)("span",{className:R.copyButtonIcons,"aria-hidden":"true",children:[(0,j.jsx)(H,{className:R.copyButtonIcon}),(0,j.jsx)(S,{className:R.copyButtonSuccessIcon})]})})}function O(e){return(0,j.jsx)("svg",{viewBox:"0 0 24 24",...e,children:(0,j.jsx)("path",{fill:"currentColor",d:"M4 19h6v-2H4v2zM20 5H4v2h16V5zm-3 6H4v2h13.25c1.1 0 2 .9 2 2s-.9 2-2 2H15v-2l-3 3l3 3v-2h2c2.21 0 4-1.79 4-4s-1.79-4-4-4z"})})}const A={wordWrapButtonIcon:"wordWrapButtonIcon_Bwma",wordWrapButtonEnabled:"wordWrapButtonEnabled_EoeP"};function z(e){let{className:t,onClick:n,isEnabled:o}=e;const i=(0,T.G)({id:"theme.CodeBlock.wordWrapToggle",message:"Toggle word wrap",description:"The title attribute for toggle word wrapping button of code block lines"});return(0,j.jsx)("button",{type:"button",onClick:n,className:(0,s.c)("clean-btn",t,o&&A.wordWrapButtonEnabled),"aria-label":i,title:i,children:(0,j.jsx)(O,{className:A.wordWrapButtonIcon,"aria-hidden":"true"})})}function I(e){let{children:t,className:n="",metastring:i,title:c,showLineNumbers:l,language:u}=e;const{prism:{defaultLanguage:d,magicComments:f}}=(0,r.y)(),h=function(e){return e?.toLowerCase()}(u??function(e){const t=e.split(" ").find((e=>e.startsWith("language-")));return t?.replace(/language-/,"")}(n)??d),p=a(),g=function(){const[e,t]=(0,o.useState)(!1),[n,i]=(0,o.useState)(!1),s=(0,o.useRef)(null),c=(0,o.useCallback)((()=>{const n=s.current.querySelector("code");e?n.removeAttribute("style"):(n.style.whiteSpace="pre-wrap",n.style.overflowWrap="anywhere"),t((e=>!e))}),[s,e]),r=(0,o.useCallback)((()=>{const{scrollWidth:e,clientWidth:t}=s.current,n=e>t||s.current.querySelector("code").hasAttribute("style");i(n)}),[s]);return L(s,r),(0,o.useEffect)((()=>{r()}),[e,r]),(0,o.useEffect)((()=>(window.addEventListener("resize",r,{passive:!0}),()=>{window.removeEventListener("resize",r)})),[r]),{codeBlockRef:s,isEnabled:e,isCodeScrollable:n,toggle:c}}(),b=function(e){return e?.match(m)?.groups.title??""}(i)||c,{lineClassNames:x,code:w}=v(t,{metastring:i,language:h,magicComments:f}),N=l??function(e){return Boolean(e?.includes("showLineNumbers"))}(i);return(0,j.jsxs)(k,{as:"div",className:(0,s.c)(n,h&&!n.includes(`language-${h}`)&&`language-${h}`),children:[b&&(0,j.jsx)("div",{className:y.codeBlockTitle,children:b}),(0,j.jsxs)("div",{className:y.codeBlockContent,children:[(0,j.jsx)(_.gl,{theme:p,code:w,language:h??"text",children:e=>{let{className:t,style:n,tokens:o,getLineProps:i,getTokenProps:c}=e;return(0,j.jsx)("pre",{tabIndex:0,ref:g.codeBlockRef,className:(0,s.c)(t,y.codeBlock,"thin-scrollbar"),style:n,children:(0,j.jsx)("code",{className:(0,s.c)(y.codeBlockLines,N&&y.codeBlockLinesWithNumbering),children:o.map(((e,t)=>(0,j.jsx)(E,{line:e,getLineProps:i,getTokenProps:c,classNames:x[t],showLineNumbers:N},t)))})})}}),(0,j.jsxs)("div",{className:y.buttonGroup,children:[(g.isEnabled||g.isCodeScrollable)&&(0,j.jsx)(z,{className:y.codeButton,onClick:()=>g.toggle(),isEnabled:g.isEnabled}),(0,j.jsx)(M,{className:y.codeButton,code:w})]})]})]})}function $(e){let{children:t,...n}=e;const s=(0,i.c)(),c=function(e){return o.Children.toArray(e).some((e=>(0,o.isValidElement)(e)))?e:Array.isArray(e)?e.join(""):e}(t),r="string"==typeof c?I:w;return(0,j.jsx)(r,{...n,children:c},String(s))}},6656:(e,t,n)=>{"use strict";n.d(t,{c:()=>H});var o=n(1504),i=n(2172),s=n(6952),c=n(6752),r=n.n(c),a=n(7624);function l(e){return(0,a.jsx)("code",{...e})}var u=n(867);var d=n(5456),m=n(5976),f=n(3664),h=n(8448);const p={details:"details_lb9f",isBrowser:"isBrowser_bmU9",collapsibleContent:"collapsibleContent_i85q"};function g(e){return!!e&&("SUMMARY"===e.tagName||g(e.parentElement))}function b(e,t){return!!e&&(e===t||b(e.parentElement,t))}function v(e){let{summary:t,children:n,...i}=e;(0,m.c)().collectAnchor(i.id);const s=(0,f.c)(),c=(0,o.useRef)(null),{collapsed:r,setCollapsed:l}=(0,h.a)({initialState:!i.open}),[u,v]=(0,o.useState)(i.open),x=o.isValidElement(t)?t:(0,a.jsx)("summary",{children:t??"Details"});return(0,a.jsxs)("details",{...i,ref:c,open:u,"data-collapsed":r,className:(0,d.c)(p.details,s&&p.isBrowser,i.className),onMouseDown:e=>{g(e.target)&&e.detail>1&&e.preventDefault()},onClick:e=>{e.stopPropagation();const t=e.target;g(t)&&b(t,c.current)&&(e.preventDefault(),r?(l(!1),v(!0)):l(!0))},children:[x,(0,a.jsx)(h.U,{lazy:!1,collapsed:r,disableSSRStyle:!0,onCollapseTransitionEnd:e=>{l(e),v(!e)},children:(0,a.jsx)("div",{className:p.collapsibleContent,children:n})})]})}const x={details:"details_b_Ee"},j="alert alert--info";function k(e){let{...t}=e;return(0,a.jsx)(v,{...t,className:(0,d.c)(j,x.details,t.className)})}function y(e){const t=o.Children.toArray(e.children),n=t.find((e=>o.isValidElement(e)&&"summary"===e.type)),i=(0,a.jsx)(a.Fragment,{children:t.filter((e=>e!==n))});return(0,a.jsx)(k,{...e,summary:n,children:i})}var w=n(6448);function N(e){return(0,a.jsx)(w.c,{...e})}const C={containsTaskList:"containsTaskList_mC6p"};function L(e){if(void 0!==e)return(0,d.c)(e,e?.includes("contains-task-list")&&C.containsTaskList)}const _={img:"img_ev3q"};var B=n(304),E=n(2944);const T={Head:s.c,details:y,Details:y,code:function(e){return function(e){return void 0!==e.children&&o.Children.toArray(e.children).every((e=>"string"==typeof e&&!e.includes("\n")))}(e)?(0,a.jsx)(l,{...e}):(0,a.jsx)(r(),{...e})},a:function(e){return(0,a.jsx)(u.c,{...e})},pre:function(e){return(0,a.jsx)(a.Fragment,{children:e.children})},ul:function(e){return(0,a.jsx)("ul",{...e,className:L(e.className)})},li:function(e){return(0,m.c)().collectAnchor(e.id),(0,a.jsx)("li",{...e})},img:function(e){return(0,a.jsx)("img",{decoding:"async",loading:"lazy",...e,className:(t=e.className,(0,d.c)(t,_.img))});var t},h1:e=>(0,a.jsx)(N,{as:"h1",...e}),h2:e=>(0,a.jsx)(N,{as:"h2",...e}),h3:e=>(0,a.jsx)(N,{as:"h3",...e}),h4:e=>(0,a.jsx)(N,{as:"h4",...e}),h5:e=>(0,a.jsx)(N,{as:"h5",...e}),h6:e=>(0,a.jsx)(N,{as:"h6",...e}),admonition:B.c,mermaid:E.c};function H(e){let{children:t}=e;return(0,a.jsx)(i.I,{components:T,children:t})}},6592:(e,t,n)=>{"use strict";n.d(t,{c:()=>r});n(1504);var o=n(5456),i=n(4357),s=n(6448),c=n(7624);function r(e){let{className:t}=e;return(0,c.jsx)("main",{className:(0,o.c)("container margin-vert--xl",t),children:(0,c.jsx)("div",{className:"row",children:(0,c.jsxs)("div",{className:"col col--6 col--offset-3",children:[(0,c.jsx)(s.c,{as:"h1",className:"hero__title",children:(0,c.jsx)(i.c,{id:"theme.NotFound.title",description:"The title of the 404 page",children:"Page Not Found"})}),(0,c.jsx)("p",{children:(0,c.jsx)(i.c,{id:"theme.NotFound.p1",description:"The first paragraph of the 404 page",children:"We could not find what you were looking for."})}),(0,c.jsx)("p",{children:(0,c.jsx)(i.c,{id:"theme.NotFound.p2",description:"The 2nd paragraph of the 404 page",children:"Please contact the owner of the site that linked you to the original URL and let them know their link is broken."})})]})})})}},5124:(e,t,n)=>{"use strict";n.d(t,{c:()=>l});n(1504);var o=n(5456),i=n(3088);const s={tableOfContents:"tableOfContents_bqdL",docItemContainer:"docItemContainer_F8PC"};var c=n(7624);const r="table-of-contents__link toc-highlight",a="table-of-contents__link--active";function l(e){let{className:t,...n}=e;return(0,c.jsx)("div",{className:(0,o.c)(s.tableOfContents,"thin-scrollbar",t),children:(0,c.jsx)(i.c,{...n,linkClassName:r,linkActiveClassName:a})})}},3088:(e,t,n)=>{"use strict";n.d(t,{c:()=>p});var o=n(1504),i=n(1824);function s(e){const t=e.map((e=>({...e,parentIndex:-1,children:[]}))),n=Array(7).fill(-1);t.forEach(((e,t)=>{const o=n.slice(2,e.level);e.parentIndex=Math.max(...o),n[e.level]=t}));const o=[];return t.forEach((e=>{const{parentIndex:n,...i}=e;n>=0?t[n].children.push(i):o.push(i)})),o}function c(e){let{toc:t,minHeadingLevel:n,maxHeadingLevel:o}=e;return t.flatMap((e=>{const t=c({toc:e.children,minHeadingLevel:n,maxHeadingLevel:o});return function(e){return e.level>=n&&e.level<=o}(e)?[{...e,children:t}]:t}))}function r(e){const t=e.getBoundingClientRect();return t.top===t.bottom?r(e.parentNode):t}function a(e,t){let{anchorTopOffset:n}=t;const o=e.find((e=>r(e).top>=n));if(o){return function(e){return e.top>0&&e.bottom<window.innerHeight/2}(r(o))?o:e[e.indexOf(o)-1]??null}return e[e.length-1]??null}function l(){const e=(0,o.useRef)(0),{navbar:{hideOnScroll:t}}=(0,i.y)();return(0,o.useEffect)((()=>{e.current=t?0:document.querySelector(".navbar").clientHeight}),[t]),e}function u(e){const t=(0,o.useRef)(void 0),n=l();(0,o.useEffect)((()=>{if(!e)return()=>{};const{linkClassName:o,linkActiveClassName:i,minHeadingLevel:s,maxHeadingLevel:c}=e;function r(){const e=function(e){return Array.from(document.getElementsByClassName(e))}(o),r=function(e){let{minHeadingLevel:t,maxHeadingLevel:n}=e;const o=[];for(let i=t;i<=n;i+=1)o.push(`h${i}.anchor`);return Array.from(document.querySelectorAll(o.join()))}({minHeadingLevel:s,maxHeadingLevel:c}),l=a(r,{anchorTopOffset:n.current}),u=e.find((e=>l&&l.id===function(e){return decodeURIComponent(e.href.substring(e.href.indexOf("#")+1))}(e)));e.forEach((e=>{!function(e,n){n?(t.current&&t.current!==e&&t.current.classList.remove(i),e.classList.add(i),t.current=e):e.classList.remove(i)}(e,e===u)}))}return document.addEventListener("scroll",r),document.addEventListener("resize",r),r(),()=>{document.removeEventListener("scroll",r),document.removeEventListener("resize",r)}}),[e,n])}var d=n(867),m=n(7624);function f(e){let{toc:t,className:n,linkClassName:o,isChild:i}=e;return t.length?(0,m.jsx)("ul",{className:i?void 0:n,children:t.map((e=>(0,m.jsxs)("li",{children:[(0,m.jsx)(d.c,{to:`#${e.id}`,className:o??void 0,dangerouslySetInnerHTML:{__html:e.value}}),(0,m.jsx)(f,{isChild:!0,toc:e.children,className:n,linkClassName:o})]},e.id)))}):null}const h=o.memo(f);function p(e){let{toc:t,className:n="table-of-contents table-of-contents__left-border",linkClassName:r="table-of-contents__link",linkActiveClassName:a,minHeadingLevel:l,maxHeadingLevel:d,...f}=e;const p=(0,i.y)(),g=l??p.tableOfContents.minHeadingLevel,b=d??p.tableOfContents.maxHeadingLevel,v=function(e){let{toc:t,minHeadingLevel:n,maxHeadingLevel:i}=e;return(0,o.useMemo)((()=>c({toc:s(t),minHeadingLevel:n,maxHeadingLevel:i})),[t,n,i])}({toc:t,minHeadingLevel:g,maxHeadingLevel:b});return u((0,o.useMemo)((()=>{if(r&&a)return{linkClassName:r,linkActiveClassName:a,minHeadingLevel:g,maxHeadingLevel:b}}),[r,a,g,b])),(0,m.jsx)(h,{toc:v,className:n,linkClassName:r,...f})}},1528:(e,t,n)=>{"use strict";n.d(t,{c:()=>f});n(1504);var o=n(5456),i=n(4357),s=n(6952),c=n(7624);function r(){return(0,c.jsx)(i.c,{id:"theme.unlistedContent.title",description:"The unlisted content banner title",children:"Unlisted page"})}function a(){return(0,c.jsx)(i.c,{id:"theme.unlistedContent.message",description:"The unlisted content banner message",children:"This page is unlisted. Search engines will not index it, and only users having a direct link can access it."})}function l(){return(0,c.jsx)(s.c,{children:(0,c.jsx)("meta",{name:"robots",content:"noindex, nofollow"})})}var u=n(5864),d=n(304);function m(e){let{className:t}=e;return(0,c.jsx)(d.c,{type:"caution",title:(0,c.jsx)(r,{}),className:(0,o.c)(t,u.W.common.unlistedBanner),children:(0,c.jsx)(a,{})})}function f(e){return(0,c.jsxs)(c.Fragment,{children:[(0,c.jsx)(l,{}),(0,c.jsx)(m,{...e})]})}},6504:(e,t)=>{function n(e){let t,n=[];for(let o of e.split(",").map((e=>e.trim())))if(/^-?\d+$/.test(o))n.push(parseInt(o,10));else if(t=o.match(/^(-?\d+)(-|\.\.\.?|\u2025|\u2026|\u22EF)(-?\d+)$/)){let[e,o,i,s]=t;if(o&&s){o=parseInt(o),s=parseInt(s);const e=o<s?1:-1;"-"!==i&&".."!==i&&"\u2025"!==i||(s+=e);for(let t=o;t!==s;t+=e)n.push(t)}}return n}t.default=n,e.exports=n},3416:(e,t,n)=>{"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.useCodeblockThemeConfig=void 0;var o=n(56);Object.defineProperty(t,"useCodeblockThemeConfig",{enumerable:!0,get:function(){return o.useCodeblockThemeConfig}})},56:function(e,t,n){"use strict";var o=this&&this.__importDefault||function(e){return e&&e.__esModule?e:{default:e}};Object.defineProperty(t,"__esModule",{value:!0}),t.useCodeblockThemeConfig=void 0;const i=o(n(8264)),s={showGithubLink:!0,githubLinkLabel:"View on GitHub",showRunmeLink:!1,runmeLinkLabel:"Checkout via Runme"};t.useCodeblockThemeConfig=function(){const{siteConfig:{themeConfig:e}}=(0,i.default)();return Object.assign(s,e.codeblock||{})}},6752:function(e,t,n){"use strict";var o=this&&this.__importDefault||function(e){return e&&e.__esModule?e:{default:e}};Object.defineProperty(t,"__esModule",{value:!0});const i=o(n(1504)),s=o(n(8720)),c=o(n(1608));t.default=(c.default,e=>function(e){return e.reference||e.metastring?.split(" ").includes("reference")}(e)?i.default.createElement(s.default,{...e}):i.default.createElement(c.default,{...e}))},7096:function(e,t,n){"use strict";var o=this&&this.__importDefault||function(e){return e&&e.__esModule?e:{default:e}};Object.defineProperty(t,"__esModule",{value:!0}),t.getRunmeLink=t.RunmeLink=void 0;const i=o(n(1504)),s=o(n(656)),c=n(3416),r=n(1756),a=n(6428),l=/(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino/i,u=/1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-/i;function d(e){let{reference:t,metastring:n}=e;const o=navigator.userAgent||navigator.vendor||window.opera,s=(0,c.useCodeblockThemeConfig)(),r=!function(e){return l.test(e)||u.test(e.substr(0,4))}(o)&&s.showRunmeLink&&n;return r?i.default.createElement("a",{href:`vscode://stateful.runme?${m(t,n)}`,className:"runmeLink",target:"_blank",style:a.buttonStyles},s.runmeLinkLabel):null}function m(e,t){const n=new URLSearchParams({command:"setup"}),o=t.match(/runmeRepository="(?<repository>[^"]*)"/),i=t.match(/runmeFileToOpen="(?<fileToOpen>[^"]*)"/),s=t.match(/useHTTPS|useHTTPS=(false|true)/);if(e.endsWith(".md"))return n.set("fileToOpen",(0,r.parseReference)(e).url),n.toString();if(i?.groups?.fileToOpen)return n.set("fileToOpen",i.groups.fileToOpen),o?.groups?.repository&&n.set("repository",o.groups.repository),n.toString();const{org:c,repo:a,title:l}=(0,r.parseReference)(e),u=s&&s.input&&(s.input.includes("useHTTPS=true")||s.input.includes(" useHTTPS ")||s.input.endsWith("useHTTPS"))?"https://github.com/":"git@github.com:";return n.set("repository",`${u}${c}/${a}.git`),n.set("fileToOpen",l.split("/").slice(0,-1).join("/")+"/README.md"),n.toString()}t.RunmeLink=function(e){return i.default.createElement(s.default,null,(()=>i.default.createElement(d,{...e})))},t.getRunmeLink=m},6428:(e,t)=>{"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.buttonStyles=t.initialFetchResultState=void 0,t.initialFetchResultState={code:"loading...",error:null,loading:null},t.buttonStyles={margin:"0 10px"}},8720:function(e,t,n){"use strict";var o=this&&this.__createBinding||(Object.create?function(e,t,n,o){void 0===o&&(o=n);var i=Object.getOwnPropertyDescriptor(t,n);i&&!("get"in i?!t.__esModule:i.writable||i.configurable)||(i={enumerable:!0,get:function(){return t[n]}}),Object.defineProperty(e,o,i)}:function(e,t,n,o){void 0===o&&(o=n),e[o]=t[n]}),i=this&&this.__setModuleDefault||(Object.create?function(e,t){Object.defineProperty(e,"default",{enumerable:!0,value:t})}:function(e,t){e.default=t}),s=this&&this.__importStar||function(e){if(e&&e.__esModule)return e;var t={};if(null!=e)for(var n in e)"default"!==n&&Object.prototype.hasOwnProperty.call(e,n)&&o(t,e,n);return i(t,e),t},c=this&&this.__importDefault||function(e){return e&&e.__esModule?e:{default:e}};Object.defineProperty(t,"__esModule",{value:!0});const r=s(n(1504)),a=c(n(1608)),l=n(3416),u=n(7096),d=n(1756),m=n(6428),f={fontSize:".9em",fontWeight:600,color:"#0E75DD",textAlign:"right",paddingBottom:"13px",textDecoration:"underline"};t.default=function(e){const[t,n]=(0,r.useReducer)(d.codeReducer,m.initialFetchResultState),o=(0,d.parseReference)(e.children);!1!==t.loading&&(0,d.fetchCode)(o,n);const i=e.metastring?.match(/title="(?<title>[^"]*)"/),s={...e,metastring:i?.groups?.title?`${e.metastring} title="${i?.groups?.title}"`:`${e.metastring} title="${o.title}"`,children:m.initialFetchResultState.code},c=(0,l.useCodeblockThemeConfig)(),h=c.showGithubLink||c.showRunmeLink;return r.default.createElement("div",{className:"docusaurus-theme-github-codeblock"},r.default.createElement(a.default,{...s},t.code),h&&r.default.createElement("div",{style:f},r.default.createElement(u.RunmeLink,{reference:e.children,metastring:e.metastring}),c.showGithubLink&&r.default.createElement("a",{href:e.children,className:"githubLink",style:m.buttonStyles,target:"_blank"},c.githubLinkLabel)))}},1756:(e,t,n)=>{"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.codeReducer=t.fetchCode=t.parseReference=void 0;const o=n(6428);t.parseReference=function(e){const t=e.slice(e.indexOf("https")),[o,i]=t.split("#"),[s,c,r,a,...l]=new n.g.URL(o).pathname.split("/").slice(1),[u,d]=i?i.split("-").map((e=>parseInt(e.slice(1),10)-1)):[0,1/0];return{url:`https://raw.githubusercontent.com/${s}/${c}/${a}/${l.join("/")}`,fromLine:u,toLine:d,title:l.join("/"),org:s,repo:c}},t.fetchCode=async function(e,t){let n,{url:o,fromLine:i,toLine:s}=e;try{n=await fetch(o)}catch(a){return t({type:"error",value:a})}if(200!==n.status){return t({type:"error",value:await n.text()})}const c=(await n.text()).split("\n").slice(i,(s||i)+1),r=c.reduce(((e,t)=>{if(0===t.length)return e;const n=t.match(/^\s+/);return n?Math.min(e,n[0].length):0}),1/0);return 0===c.length?t({type:"error",value:`Error: No code found at ${o} from line ${i} to line ${s}`}):t({type:"loaded",value:c.map((e=>e.slice(r))).join("\n")})},t.codeReducer=function(e,t){let{type:n,value:i}=t;switch(n){case"reset":return o.initialFetchResultState;case"loading":return{...e,loading:!0};case"loaded":case"error":return{...e,code:i,loading:!1};default:return e}}},2172:(e,t,n)=>{"use strict";n.d(t,{I:()=>r,M:()=>c});var o=n(1504);const i={},s=o.createContext(i);function c(e){const t=o.useContext(s);return o.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function r(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:c(e.components),o.createElement(s.Provider,{value:t},e.children)}}}]);