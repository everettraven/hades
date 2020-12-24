(window.webpackJsonp=window.webpackJsonp||[]).push([[2],{63:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return o})),n.d(t,"metadata",(function(){return s})),n.d(t,"toc",(function(){return i})),n.d(t,"default",(function(){return p}));var r=n(3),a=n(7),c=(n(0),n(84)),o={id:"command",title:"command"},s={unversionedId:"resources/command",id:"resources/command",isDocsHomePage:!1,title:"command",description:"The command testing resource allows you to run any command on the system being tested. The command resource has three different parameters:",source:"@site/docs\\resources\\command.md",slug:"/resources/command",permalink:"/hades/resources/command",editUrl:"https://github.com/everettraven/hades/edit/main/docs/docs/resources/command.md",version:"current",sidebar:"someSidebar",previous:{title:"Multiple Tests",permalink:"/hades/guides/multiple_tests"},next:{title:"os",permalink:"/hades/resources/os"}},i=[{value:"Examples",id:"examples",children:[]}],u={toc:i};function p(e){var t=e.components,n=Object(a.a)(e,["components"]);return Object(c.b)("wrapper",Object(r.a)({},u,n,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"The command testing resource allows you to run any command on the system being tested. The command resource has three different parameters:"),Object(c.b)("p",null,Object(c.b)("strong",{parentName:"p"},"cmd")," - This parameter specifies the actual command you would like to run. Case sensitive."),Object(c.b)("p",null,Object(c.b)("strong",{parentName:"p"},"args ",Object(c.b)("em",{parentName:"strong"},"(optional)"))," - This parameter is used to pass arguments to the command specified in the cmd parameter, but is optional as you can pass arguments in the cmd parameter as well."),Object(c.b)("p",null,Object(c.b)("strong",{parentName:"p"},"expectedOutput")," - This parameter is used to test the output from running the command specified in the cmd parameter to ensure it matches the expected output. The text in this is case sensitive and accepts standard string special characters such as ",Object(c.b)("inlineCode",{parentName:"p"},"\\n"),"."),Object(c.b)("h2",{id:"examples"},"Examples"),Object(c.b)("p",null,"Simple echo command with args parameter:"),Object(c.b)("pre",null,Object(c.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'command {\n    cmd = "echo"\n    args = ["hello world"]\n    expectedOutput = "hello world"\n}\n')),Object(c.b)("p",null,"Multiple arguments in args parameter:"),Object(c.b)("pre",null,Object(c.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'command {\n    cmd = "head"\n    args = ["-1", "/etc/os-release"]\n    expectedOutput = "NAME=\\"Ubuntu\\""\n}\n')),Object(c.b)("p",null,"Without args parameter:"),Object(c.b)("pre",null,Object(c.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'command {\n    cmd = "lsb_release -a | grep -i Release"\n    expectedOutput = "Release:\\t20.04"\n}\n')))}p.isMDXComponent=!0},84:function(e,t,n){"use strict";n.d(t,"a",(function(){return m})),n.d(t,"b",(function(){return b}));var r=n(0),a=n.n(r);function c(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){c(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},c=Object.keys(e);for(r=0;r<c.length;r++)n=c[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(r=0;r<c.length;r++)n=c[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var u=a.a.createContext({}),p=function(e){var t=a.a.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},m=function(e){var t=p(e.components);return a.a.createElement(u.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},d=a.a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,c=e.originalType,o=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),m=p(n),d=r,b=m["".concat(o,".").concat(d)]||m[d]||l[d]||c;return n?a.a.createElement(b,s(s({ref:t},u),{},{components:n})):a.a.createElement(b,s({ref:t},u))}));function b(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var c=n.length,o=new Array(c);o[0]=d;var s={};for(var i in t)hasOwnProperty.call(t,i)&&(s[i]=t[i]);s.originalType=e,s.mdxType="string"==typeof e?e:r,o[1]=s;for(var u=2;u<c;u++)o[u]=n[u];return a.a.createElement.apply(null,o)}return a.a.createElement.apply(null,n)}d.displayName="MDXCreateElement"}}]);