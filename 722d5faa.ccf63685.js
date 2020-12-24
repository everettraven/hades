(window.webpackJsonp=window.webpackJsonp||[]).push([[5],{75:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return s})),n.d(t,"metadata",(function(){return l})),n.d(t,"toc",(function(){return c})),n.d(t,"default",(function(){return u}));var r=n(3),o=n(7),a=(n(0),n(84)),s={id:"multiple_tests",title:"Multiple Tests"},l={unversionedId:"guides/multiple_tests",id:"guides/multiple_tests",isDocsHomePage:!1,title:"Multiple Tests",description:"This guide shows how to use hades with a simple test file that runs a very simple command. All source code for this example can be found here",source:"@site/docs\\guides\\multiple_tests.md",slug:"/guides/multiple_tests",permalink:"/guides/multiple_tests",editUrl:"https://github.com/everettraven/hades/edit/main/docs/docs/guides/multiple_tests.md",version:"current",sidebar:"someSidebar",previous:{title:"Host File",permalink:"/guides/host_file"},next:{title:"command",permalink:"/resources/command"}},c=[{value:"Docker Setup",id:"docker-setup",children:[]},{value:"Create the Test Files",id:"create-the-test-files",children:[]},{value:"Run hades on the Remote System",id:"run-hades-on-the-remote-system",children:[]},{value:"Docker Cleanup",id:"docker-cleanup",children:[]}],i={toc:c};function u(e){var t=e.components,n=Object(o.a)(e,["components"]);return Object(a.b)("wrapper",Object(r.a)({},i,n,{components:t,mdxType:"MDXLayout"}),Object(a.b)("p",null,"This guide shows how to use hades with a simple test file that runs a very simple command. All source code for this example can be found ",Object(a.b)("a",Object(r.a)({parentName:"p"},{href:"https://github.com/everettraven/hades/tree/main/examples/multiple_tests"}),"here")),Object(a.b)("p",null,"This guide uses Docker to create a container that would simulate a remote machine running on your local machine. If you have an existing remote machine feel free to skip the Docker setup steps."),Object(a.b)("h2",{id:"docker-setup"},"Docker Setup"),Object(a.b)("ol",null,Object(a.b)("li",{parentName:"ol"},"Make sure that you have docker installed by running:")),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"docker -v\n")),Object(a.b)("p",null,"If you do not have Docker installed you can install it by following the offical Docker installation instructions: ",Object(a.b)("a",Object(r.a)({parentName:"p"},{href:"https://docs.docker.com/get-docker/"}),"https://docs.docker.com/get-docker/")),Object(a.b)("ol",{start:2},Object(a.b)("li",{parentName:"ol"},"Pull the Docker Image we are going to use by running:")),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"docker pull bpalmer/ssh_test\n")),Object(a.b)("ol",{start:3},Object(a.b)("li",{parentName:"ol"},"Run the Docker Container by running:")),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"docker run --name multiple_tests -d -p 5000:22 bpalmer/ssh_test\n")),Object(a.b)("p",null,"This will run the Docker Container in the background with your localhost port 5000 mapped to port 22 (the standard SSH port) on the Docker Container. It also names the container 'multiple_tests' for easy cleanup when we are done with it. Feel free to play with these values as you see fit, but make sure to adjust where these values are used in the future pieces of the guide."),Object(a.b)("p",null,"For reference:"),Object(a.b)("p",null,"Both the username and password for this Docker Container is ",Object(a.b)("strong",{parentName:"p"},"root")),Object(a.b)("h2",{id:"create-the-test-files"},"Create the Test Files"),Object(a.b)("p",null,"Make sure you are in the directory you would like to store the source code of this guide in and make a directory named ",Object(a.b)("strong",{parentName:"p"},"tests"),"."),Object(a.b)("p",null,"hades will by default look for the ",Object(a.b)("strong",{parentName:"p"},"tests")," folder and run all the tests within that folder but you can name this directory whatever you would like and pass the ",Object(a.b)("inlineCode",{parentName:"p"},"--test-dir")," flag followed by the name of your test directory and hades will use that folder."),Object(a.b)("p",null,"In the ",Object(a.b)("strong",{parentName:"p"},"tests")," folder we are going to create 2 test files named ",Object(a.b)("strong",{parentName:"p"},"test1.hcl")," and ",Object(a.b)("strong",{parentName:"p"},"test2.hcl"),"."),Object(a.b)("p",null,"In ",Object(a.b)("strong",{parentName:"p"},"test1.hcl")," put:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'title = "Multiple Tests - Test #1"\n\ncommand {\n    cmd = "echo"\n    args = ["Multiple"]\n    expectedOutput = "Multiple"\n}\n\ncommand {\n    cmd = "ls"\n    args = ["/usr"]\n    expectedOutput = "bin\\ngames\\ninclude\\nlib\\nlib32\\nlib64\\nlibx32\\nlocal\\nsbin\\nshare\\nsrc"\n}\n\nos {\n    distributionID = "ubuntu"\n}\n')),Object(a.b)("p",null,"You can also do the following to get rid of the optional args parameter in the command block if you would like:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'title = "Multiple Tests - Test #1"\n\ncommand {\n    cmd = "echo Multiple"\n    expectedOutput = "Multiple"\n}\n\ncommand {\n    cmd = "ls /usr"\n    expectedOutput = "bin\\ngames\\ninclude\\nlib\\nlib32\\nlib64\\nlibx32\\nlocal\\nsbin\\nshare\\nsrc"\n}\n\nos {\n    distributionID = "ubuntu"\n}\n')),Object(a.b)("p",null,"In ",Object(a.b)("strong",{parentName:"p"},"test2.hcl")," put:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'title = "Multiple Tests - Test #2"\n\ncommand {\n    cmd = "echo"\n    args = ["hades is working!"]\n    expectedOutput = "hades is working!"\n}\n')),Object(a.b)("p",null,"You can also do the following to get rid of the optional args parameter in the command block if you would like:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'title = "Multiple Tests - Test #2"\n\ncommand {\n    cmd = "echo hades is working!"\n    expectedOutput = "hades is working!"\n}\n')),Object(a.b)("p",null,"We recommend that you create a hosts folder for better organization, but it is not necessary when creating the ",Object(a.b)("strong",{parentName:"p"},"hosts.hcl")," file as hades will look for it as long as it is in the current directory or a sub-directory of the current directory. In our case we are going to create the hosts folder and then within that folder create the ",Object(a.b)("strong",{parentName:"p"},"hosts.hcl")," file. In the ",Object(a.b)("strong",{parentName:"p"},"hosts.hcl")," file we are going to put the following:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-hcl"}),'host {\n    ip = "127.0.0.1"\n    port = "5000"\n    user = "root"\n}\n')),Object(a.b)("p",null,"If you wanted to run the tests on multiple hosts you can place multiple host blocks in the hosts file. hades will run all the tests on each of the hosts."),Object(a.b)("h2",{id:"run-hades-on-the-remote-system"},"Run hades on the Remote System"),Object(a.b)("p",null,"Now that we have a simple test file created we can run hades to test the remote system. In this case we will run it on our Docker Container we created for this guide by running:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"hades --pass root\n")),Object(a.b)("p",null,"If you did not name the test directory ",Object(a.b)("strong",{parentName:"p"},"tests")," you can let hades know what directory to use for getting the tests by running:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"hades --test-dir [directory] --pass root\n")),Object(a.b)("p",null,"It is not recommended to run hades by passing the ",Object(a.b)("inlineCode",{parentName:"p"},"--pass")," flag to the command as this puts the password for the remote machine in your command line history as plaintext. We recommend running hades using an SSH key. If you are interested in doing so we first need to do some setup:"),Object(a.b)("ol",null,Object(a.b)("li",{parentName:"ol"},"Create an SSH key by running:")),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"ssh-keygen\n")),Object(a.b)("ol",{start:2},Object(a.b)("li",{parentName:"ol"},"Send the SSH key to the remote system (or Docker Container in this case):")),Object(a.b)("p",null,"Windows (Powershell):"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-powershell"}),'type $env:USERPROFILE\\.ssh\\id_rsa.pub | ssh root@127.0.0.1 -p 5000 "cat >> .ssh/authorized_keys"\n')),Object(a.b)("p",null,"Unix:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"ssh-copy-id -i ~/.ssh/id_rsa root@127.0.0.1 -p 5000\n")),Object(a.b)("p",null,"Now that the SSH Key has been sent to the remote system we can run hades with the SSH key. By default hades attempts to get the SSH key from ",Object(a.b)("inlineCode",{parentName:"p"},"~/.ssh/id_rsa")," so if you created the key to a different directory you can use the flag ",Object(a.b)("inlineCode",{parentName:"p"},"--key-file")," followed by the path to the SSH key."),Object(a.b)("p",null,"Running hades with the SSH Key:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"hades\n")),Object(a.b)("p",null,"If you did not name the test directory ",Object(a.b)("strong",{parentName:"p"},"tests")," you can let hades know what directory to use for getting the tests by running:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"hades --test-dir [directory]\n")),Object(a.b)("p",null,"Running hades with a non-default SSH Key path:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"hades --key-file [key path]\n")),Object(a.b)("p",null,"If you did not name the test directory ",Object(a.b)("strong",{parentName:"p"},"tests")," you can let hades know what directory to use for getting the tests by running:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"hades --test-dir [directory] --key-file [key path]\n")),Object(a.b)("h2",{id:"docker-cleanup"},"Docker Cleanup"),Object(a.b)("p",null,"All the guides use Docker Containers so if you plan to continue with the rest of the guides you can return to this step when you are finished with the guides you would like to go through."),Object(a.b)("p",null,"If you would like to cleanup the Docker Container now you can run the following commands to stop and remove the Docker Container:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"docker container stop multiple_tests\n")),Object(a.b)("p",null,"and"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"docker container rm multiple_tests\n")),Object(a.b)("p",null,"All the guides will use the same Docker Image for their Containers, but if you would like to remove the image as well you can run:"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{}),"docker image rm bpalmer/ssh_test\n")))}u.isMDXComponent=!0},84:function(e,t,n){"use strict";n.d(t,"a",(function(){return b})),n.d(t,"b",(function(){return h}));var r=n(0),o=n.n(r);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var i=o.a.createContext({}),u=function(e){var t=o.a.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},b=function(e){var t=u(e.components);return o.a.createElement(i.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return o.a.createElement(o.a.Fragment,{},t)}},d=o.a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,a=e.originalType,s=e.parentName,i=c(e,["components","mdxType","originalType","parentName"]),b=u(n),d=r,h=b["".concat(s,".").concat(d)]||b[d]||p[d]||a;return n?o.a.createElement(h,l(l({ref:t},i),{},{components:n})):o.a.createElement(h,l({ref:t},i))}));function h(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var a=n.length,s=new Array(a);s[0]=d;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:r,s[1]=l;for(var i=2;i<a;i++)s[i]=n[i];return o.a.createElement.apply(null,s)}return o.a.createElement.apply(null,n)}d.displayName="MDXCreateElement"}}]);