(window.webpackJsonp=window.webpackJsonp||[]).push([[3],{64:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return a})),r.d(t,"metadata",(function(){return i})),r.d(t,"toc",(function(){return c})),r.d(t,"default",(function(){return p}));var n=r(3),o=r(7),s=(r(0),r(84)),a={id:"os",title:"os"},i={unversionedId:"resources/os",id:"resources/os",isDocsHomePage:!1,title:"os",description:"The os resource is used to test the operating system of the machine. There are two parameters:",source:"@site/docs\\resources\\os.md",slug:"/resources/os",permalink:"/resources/os",editUrl:"https://github.com/everettraven/hades/edit/main/docs/docs/resources/os.md",version:"current",sidebar:"someSidebar",previous:{title:"command",permalink:"/resources/command"}},c=[{value:"Examples",id:"examples",children:[]}],u={toc:c};function p(e){var t=e.components,r=Object(o.a)(e,["components"]);return Object(s.b)("wrapper",Object(n.a)({},u,r,{components:t,mdxType:"MDXLayout"}),Object(s.b)("p",null,"The os resource is used to test the operating system of the machine. There are two parameters:"),Object(s.b)("p",null,Object(s.b)("strong",{parentName:"p"},"distributionID")," - This parameter expects the distributor ID of the operating system you are expecting the machine to have. Not case sensitive."),Object(s.b)("p",null,Object(s.b)("strong",{parentName:"p"},"version ",Object(s.b)("em",{parentName:"strong"},"(optional)"))," - This parameter will be used to check the version of the operating system to this specified version. Not case sensitive."),Object(s.b)("h2",{id:"examples"},"Examples"),Object(s.b)("p",null,"Without version:"),Object(s.b)("pre",null,Object(s.b)("code",Object(n.a)({parentName:"pre"},{className:"language-hcl"}),'os {\n    distributionID = "Ubuntu"\n}\n')),Object(s.b)("p",null,"This test will pass with any operating system with the distribution ID of Ubuntu"),Object(s.b)("p",null,"With version:"),Object(s.b)("pre",null,Object(s.b)("code",Object(n.a)({parentName:"pre"},{className:"language-hcl"}),'os {\n    distributionID = "Ubuntu"\n    version = "20.04"\n}\n')),Object(s.b)("p",null,"This test will pass only with the operating system Ubuntu version 20.04."))}p.isMDXComponent=!0},84:function(e,t,r){"use strict";r.d(t,"a",(function(){return l})),r.d(t,"b",(function(){return d}));var n=r(0),o=r.n(n);function s(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){s(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var u=o.a.createContext({}),p=function(e){var t=o.a.useContext(u),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},l=function(e){var t=p(e.components);return o.a.createElement(u.Provider,{value:t},e.children)},b={inlineCode:"code",wrapper:function(e){var t=e.children;return o.a.createElement(o.a.Fragment,{},t)}},m=o.a.forwardRef((function(e,t){var r=e.components,n=e.mdxType,s=e.originalType,a=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),l=p(r),m=n,d=l["".concat(a,".").concat(m)]||l[m]||b[m]||s;return r?o.a.createElement(d,i(i({ref:t},u),{},{components:r})):o.a.createElement(d,i({ref:t},u))}));function d(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var s=r.length,a=new Array(s);a[0]=m;var i={};for(var c in t)hasOwnProperty.call(t,c)&&(i[c]=t[c]);i.originalType=e,i.mdxType="string"==typeof e?e:n,a[1]=i;for(var u=2;u<s;u++)a[u]=r[u];return o.a.createElement.apply(null,a)}return o.a.createElement.apply(null,r)}m.displayName="MDXCreateElement"}}]);