(function(e){function t(t){for(var n,s,r=t[0],c=t[1],u=t[2],l=0,d=[];l<r.length;l++)s=r[l],Object.prototype.hasOwnProperty.call(a,s)&&a[s]&&d.push(a[s][0]),a[s]=0;for(n in c)Object.prototype.hasOwnProperty.call(c,n)&&(e[n]=c[n]);f&&f(t);while(d.length)d.shift()();return i.push.apply(i,u||[]),o()}function o(){for(var e,t=0;t<i.length;t++){for(var o=i[t],n=!0,s=1;s<o.length;s++){var c=o[s];0!==a[c]&&(n=!1)}n&&(i.splice(t--,1),e=r(r.s=o[0]))}return e}var n={},a={app:0},i=[];function s(e){return r.p+"assets/js/"+({about:"about"}[e]||e)+"."+{about:"cb1df7d9"}[e]+".js"}function r(t){if(n[t])return n[t].exports;var o=n[t]={i:t,l:!1,exports:{}};return e[t].call(o.exports,o,o.exports,r),o.l=!0,o.exports}r.e=function(e){var t=[],o=a[e];if(0!==o)if(o)t.push(o[2]);else{var n=new Promise((function(t,n){o=a[e]=[t,n]}));t.push(o[2]=n);var i,c=document.createElement("script");c.charset="utf-8",c.timeout=120,r.nc&&c.setAttribute("nonce",r.nc),c.src=s(e);var u=new Error;i=function(t){c.onerror=c.onload=null,clearTimeout(l);var o=a[e];if(0!==o){if(o){var n=t&&("load"===t.type?"missing":t.type),i=t&&t.target&&t.target.src;u.message="Loading chunk "+e+" failed.\n("+n+": "+i+")",u.name="ChunkLoadError",u.type=n,u.request=i,o[1](u)}a[e]=void 0}};var l=setTimeout((function(){i({type:"timeout",target:c})}),12e4);c.onerror=c.onload=i,document.head.appendChild(c)}return Promise.all(t)},r.m=e,r.c=n,r.d=function(e,t,o){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:o})},r.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(r.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)r.d(o,n,function(t){return e[t]}.bind(null,n));return o},r.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="",r.oe=function(e){throw console.error(e),e};var c=window["webpackJsonp"]=window["webpackJsonp"]||[],u=c.push.bind(c);c.push=t,c=c.slice();for(var l=0;l<c.length;l++)t(c[l]);var f=u;i.push([0,"chunk-vendors"]),o()})({0:function(e,t,o){e.exports=o("56d7")},"034f":function(e,t,o){"use strict";o("85ec")},"08fd":function(e,t,o){"use strict";o("2414")},2414:function(e,t,o){},"4e71":function(e,t,o){"use strict";o("d7e9")},"56d7":function(e,t,o){"use strict";o.r(t);o("e260"),o("e6cf"),o("cca6"),o("a79d");var n=o("2b0e"),a=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{staticClass:"app_div",attrs:{id:"app"}},["multi"===e.defaultSetiing.default_page_mode?o("MultiPage"):e._e(),"random"===e.defaultSetiing.default_page_mode?o("RandomPage"):e._e(),"single"===e.defaultSetiing.default_page_mode?o("SinglePage"):e._e()],1)},i=[],s=(o("d3b7"),o("ddb0"),o("bc3a")),r=o.n(s),c=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"multiPage"}},[o("Header",[o("h2",[o("a",{attrs:{href:"raw/"+e.book.name}},[e._v(e._s(e.book.name))])]),o("h4",[e._v("总页数："+e._s(e.book.page_num))])]),e._l(e.book.pages,(function(t,n){return o("div",{key:t.url,staticClass:"manga"},[o("img",{directives:[{name:"lazy",rawName:"v-lazy",value:t.url,expression:"page.url"}],key:e.k,class:e._f("check_image")(t.class,t.url),attrs:{H:t.height,W:t.width}}),o("p",[e._v(e._s(n+1)+"/"+e._s(e.book.page_num))])])})),o("p"),o("v-btn",{directives:[{name:"scroll",rawName:"v-scroll",value:e.onScroll,expression:"onScroll"},{name:"show",rawName:"v-show",value:e.btnFlag,expression:"btnFlag"}],attrs:{fab:"",color:"#bbcbff",bottom:"",right:""},on:{click:e.toTop}},[e._v("▲")]),e._t("default")],2)},u=[],l=(o("25f0"),function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("header",{staticClass:"header"},[e._t("default")],2)}),f=[],d={name:"Header",data:function(){return{mybook:this.book,rangeValue:500}},methods:{changeEvent:function(){console.log("change to:"+this.rangeValue)}}},g=d,h=(o("4e71"),o("2877")),p=Object(h["a"])(g,l,f,!1,null,"c3c1e8ec",null),m=p.exports,b={components:{Header:m},data:function(){return{book:{name:"null",page_num:1,pages:[{height:2e3,width:1419,url:"/resources/favicon.ico",class:"Vertical"}]},bookshelf:{},defaultSetiing:{default_page_mode:"single"},page:1,page_mode:"multi",btnFlag:!1,duration:300,offset:0,easing:"easeInOutCubic",message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){this.initPage(),this.getBook(),this.getBookShelf(),this.hintMessage(),this.initWebSocket()},destroyed:function(){this.$socket.close()},methods:{initPage:function(){this.$cookies.keys()},getNumber:function(e){this.page=e,console.log(e)},getBook:function(){var e=this;r.a.get("/book.json").then((function(t){return e.book=t.data})),r.a.get("/setting.json").then((function(t){return e.defaultSetiing=t.data})),r.a.get("/bookshelf.json").then((function(t){return e.bookshelf=t.data})).finally()},getBookShelf:function(){},onScroll:function(e){if("undefined"!==typeof window){var t=window.pageYOffset||e.target.scrollTop||0;this.btnFlag=t>20}},toTop:function(){this.$vuetify.goTo(0)},initWebSocket:function(){this.$socket.onopen=this.websocketonopen,this.$socket.onerror=this.websocketonerror,this.$socket.onmessage=this.websocketonmessage,this.$socket.onclose=this.websocketclose,this.hint="连接建立",this.setButtonColor("green")},websocketonopen:function(e){this.hint="连接成功",this.setButtonColor("green"),console.log("连接建立",e)},websocketonerror:function(e){this.hint="连接出错",this.setButtonColor("#999999"),this.initWebSocket(),console.log("Connection Error !!!",e)},websocketonmessage:function(e){console.log(e),this.msgList.push(JSON.parse(e.data)),this.hint="接收消息",this.setButtonColor("blue",e)},onChangeBook:function(e,t){this.message.now_book_uuid=t,this.message.msg="ChangeBook",this.$socket.send(JSON.stringify(this.message)),this.getBook()},websocketsend:function(e){this.$socket.send(JSON.stringify(this.message)),console.log(this.$socket.readyState,e)},websocketclose:function(e){var t=this;this.hint="连接断开",console.log("断开连接",e),this.setButtonColor("#888888");var o=e.code,n=e.reason,a=e.wasClean;console.log(o,n,a);var i=setInterval((function(){t.$socket.onopen(),0==e.target.readyState&&clearInterval(i)}),3e3)},setButtonColor:function(e){var t=document.getElementsByClassName("hint")[0];t.style.background=e}},filters:{check_image:function(e,t){if(console.log(e),console.log(t),e=e.toString(),"Vertical"==e||"Horizontal"==e)return e;function o(e){var t=new Image;if(t.src=e,t.complete)return t.width<t.height?"Vertical":"Horizontal";t.onload=function(){return t.onload=null,t.width<t.height?"Vertical":"Horizontal"}}return""==e&&console.log("图片信息为空，开始本地JS分析"+t),e=o(t),e}}},v=b,k=(o("08fd"),o("6544")),_=o.n(k),w=o("8336"),y=o("269a"),S=o.n(y),j=o("f977"),O=Object(h["a"])(v,c,u,!1,null,null,null),x=O.exports;_()(O,{VBtn:w["a"]}),S()(O,{Scroll:j["a"]});var P=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"SinglePage"}},[o("Header",[o("h2",[o("a",{attrs:{href:"raw/"+e.book.name}},[e._v(e._s(e.book.name))])])]),o("div",{staticClass:"singlebox",staticStyle:{width:"600px",height:"800px",border:"3px #cccccc dashed",margin:"auto"}},[o("v-img",{attrs:{contain:"","lazy-src":"book.pages[page].url","max-height":"600","max-width":"800",src:"book.pages[page].url"}}),o("img",{attrs:{url:e.book.pages[0].url,height:"95%"}}),o("v-pagination",{attrs:{length:e.book.page_num},on:{input:e.getNumber},model:{value:e.page,callback:function(t){e.page=t},expression:"page"}})],1),e._t("default")],2)},$=[],C={},B=C,E=o("adda"),H=o("891e"),T=Object(h["a"])(B,P,$,!1,null,null,null),N=T.exports;_()(T,{VImg:E["a"],VPagination:H["a"]});var V={name:"app",components:{MultiPage:x,SinglePage:N},data:function(){return{book:{name:"null",page_num:1,pages:[{height:2e3,width:1419,url:"/resources/favicon.ico",class:"Vertical"}]},bookshelf:{},defaultSetiing:{default_page_mode:"single"},page:1,page_mode:"multi",duration:300,offset:0,easing:"easeInOutCubic",message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){var e=this;this.initPage(),r.a.get("/book.json").then((function(t){return e.book=t.data})),r.a.get("/setting.json").then((function(t){return e.defaultSetiing=t.data})),r.a.get("/bookshelf.json").then((function(t){return e.bookshelf=t.data})).finally()},destroyed:function(){this.$socket.close()},methods:{initPage:function(){this.$cookies.keys()},getNumber:function(e){this.page=e,console.log(e)}}},M=V,z=(o("034f"),Object(h["a"])(M,a,i,!1,null,null,null)),I=z.exports,J=o("8c4f"),F=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"Home"}})},L=[],W=(o("99af"),o("4de4"),o("2909")),A={name:"Home",components:{},data:function(){return{todos:[]}},methods:{deleteTodo:function(e){this.todos=this.todos.filter((function(t){return t.id!==e}))},addTodo:function(e){var t=this,o=e.title,n=e.completed;r.a.post("https://jsonplaceholder.typicode.com/todos",{title:o,completed:n}).then((function(e){return t.todos=[].concat(Object(W["a"])(t.todos),[e.data])})).catch((function(e){return console.log(e)}))}},created:function(){var e=this;r.a.get("https://jsonplaceholder.typicode.com/todos?_limit=4").then((function(t){return e.todos=t.data})).catch((function(e){return console.log(e)}))}},q=A,D=(o("cccb"),Object(h["a"])(q,F,L,!1,null,null,null)),R=D.exports;n["a"].use(J["a"]);var Y=[{path:"/",name:"Home",component:R},{path:"/about",name:"About",component:function(){return o.e("about").then(o.bind(null,"f820"))}}],G=new J["a"]({routes:Y}),K=G,Q=o("caf9"),U=o("b408"),X=o.n(U),Z=o("f309");n["a"].use(Z["a"]);var ee=new Z["a"]({}),te=o("2b27"),oe=o.n(te);n["a"].use(X.a,"ws://"+document.location.host+"/ws",{reconnection:!0,reconnectionAttempts:500,reconnectionDelay:1e3}),n["a"].config.productionTip=!1,n["a"].use(Q["a"],{preLoad:4.5,attempt:10}),n["a"].use(oe.a),n["a"].$cookies.config("30d"),new n["a"]({router:K,vuetify:ee,render:function(e){return e(I)}}).$mount("#app")},"5ced":function(e,t,o){},"85ec":function(e,t,o){},cccb:function(e,t,o){"use strict";o("5ced")},d7e9:function(e,t,o){}});
//# sourceMappingURL=app.cb1fd327.js.map