(function(e){function t(t){for(var n,i,r=t[0],c=t[1],u=t[2],l=0,f=[];l<r.length;l++)i=r[l],Object.prototype.hasOwnProperty.call(s,i)&&s[i]&&f.push(s[i][0]),s[i]=0;for(n in c)Object.prototype.hasOwnProperty.call(c,n)&&(e[n]=c[n]);d&&d(t);while(f.length)f.shift()();return a.push.apply(a,u||[]),o()}function o(){for(var e,t=0;t<a.length;t++){for(var o=a[t],n=!0,i=1;i<o.length;i++){var c=o[i];0!==s[c]&&(n=!1)}n&&(a.splice(t--,1),e=r(r.s=o[0]))}return e}var n={},s={app:0},a=[];function i(e){return r.p+"assets/js/"+({about:"about"}[e]||e)+"."+{about:"d3306ce8"}[e]+".js"}function r(t){if(n[t])return n[t].exports;var o=n[t]={i:t,l:!1,exports:{}};return e[t].call(o.exports,o,o.exports,r),o.l=!0,o.exports}r.e=function(e){var t=[],o=s[e];if(0!==o)if(o)t.push(o[2]);else{var n=new Promise((function(t,n){o=s[e]=[t,n]}));t.push(o[2]=n);var a,c=document.createElement("script");c.charset="utf-8",c.timeout=120,r.nc&&c.setAttribute("nonce",r.nc),c.src=i(e);var u=new Error;a=function(t){c.onerror=c.onload=null,clearTimeout(l);var o=s[e];if(0!==o){if(o){var n=t&&("load"===t.type?"missing":t.type),a=t&&t.target&&t.target.src;u.message="Loading chunk "+e+" failed.\n("+n+": "+a+")",u.name="ChunkLoadError",u.type=n,u.request=a,o[1](u)}s[e]=void 0}};var l=setTimeout((function(){a({type:"timeout",target:c})}),12e4);c.onerror=c.onload=a,document.head.appendChild(c)}return Promise.all(t)},r.m=e,r.c=n,r.d=function(e,t,o){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:o})},r.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(r.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)r.d(o,n,function(t){return e[t]}.bind(null,n));return o},r.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="",r.oe=function(e){throw console.error(e),e};var c=window["webpackJsonp"]=window["webpackJsonp"]||[],u=c.push.bind(c);c.push=t,c=c.slice();for(var l=0;l<c.length;l++)t(c[l]);var d=u;a.push([0,"chunk-vendors"]),o()})({0:function(e,t,o){e.exports=o("56d7")},"034f":function(e,t,o){"use strict";o("85ec")},"08fd":function(e,t,o){"use strict";o("2414")},2414:function(e,t,o){},"4e71":function(e,t,o){"use strict";o("d7e9")},"56d7":function(e,t,o){"use strict";o.r(t);o("4de4"),o("96cf");var n=o("1da1"),s=(o("e260"),o("e6cf"),o("cca6"),o("a79d"),o("2b0e")),a=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"app"}},[e.defaultSetting?o("div",["multi"===e.defaultSetting.template?o("MultiPage"):e._e(),"single"===e.defaultSetting.template?o("SinglePage"):e._e(),"sketch"===e.defaultSetting.template?o("SketchPage"):e._e()],1):o("p",[e._v("loading.....")])])},i=[],r=(o("d3b7"),o("ddb0"),o("bc3a")),c=o.n(r),u=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"multiPage"}},[o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】")]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name))]):e._e()]),o("h4",[e._v("总页数："+e._s(this.$store.state.book.page_num))])]),e._l(this.$store.state.book.pages,(function(t,n){return o("div",{key:t.url,staticClass:"manga"},[o("img",{directives:[{name:"lazy",rawName:"v-lazy",value:t.url,expression:"page.url"}],key:n,class:e._f("check_image")(t.class,t.url),attrs:{H:t.height,W:t.width}}),o("p",[e._v(e._s(n+1)+"/"+e._s(e.AllPageNum))])])})),o("p"),o("v-btn",{directives:[{name:"scroll",rawName:"v-scroll",value:e.onScroll,expression:"onScroll"},{name:"show",rawName:"v-show",value:e.btnFlag,expression:"btnFlag"}],attrs:{fab:"",color:"#bbcbff",bottom:"",right:""},on:{click:e.toTop}},[e._v("▲")]),e._t("default")],2)},l=[],d=(o("25f0"),function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("header",{staticClass:"header"},[e._t("default")],2)}),f=[],h={name:"Header",data:function(){return{mybook:this.book,rangeValue:500}},methods:{changeEvent:function(){console.log("change to:"+this.rangeValue)}}},g=h,m=(o("4e71"),o("2877")),p=Object(m["a"])(g,d,f,!1,null,"c3c1e8ec",null),b=p.exports,k={components:{Header:b},data:function(){return{page_mode:"multi",btnFlag:!1,duration:300,offset:0,easing:"easeInOutCubic",AllPageNum:this.$store.state.book.page_num,message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){this.initPage()},destroyed:function(){},methods:{initPage:function(){this.$cookies.keys()},getBook:function(){return this.$store.state.book},getNumber:function(e){this.page=e,console.log(e)},onScroll:function(e){if("undefined"!==typeof window){var t=window.pageYOffset||e.target.scrollTop||0;this.btnFlag=t>20}},toTop:function(){this.$vuetify.goTo(0)},initWebSocket:function(){this.$socket.onopen=this.websocketonopen,this.$socket.onerror=this.websocketonerror,this.$socket.onmessage=this.websocketonmessage,this.$socket.onclose=this.websocketclose,this.hint="连接建立"},websocketonopen:function(e){this.hint="连接成功",console.log("连接建立",e)},websocketonerror:function(e){this.hint="连接出错",this.initWebSocket(),console.log("Connection Error !!!",e)},websocketonmessage:function(e){console.log(e),this.msgList.push(JSON.parse(e.data)),this.hint="接收消息"},onChangeBook:function(e,t){this.message.now_book_uuid=t,this.message.msg="ChangeBook",this.$socket.send(JSON.stringify(this.message)),this.getBook()},websocketsend:function(e){this.$socket.send(JSON.stringify(this.message)),console.log(this.$socket.readyState,e)},websocketclose:function(e){var t=this;this.hint="连接断开",console.log("断开连接",e);var o=e.code,n=e.reason,s=e.wasClean;console.log(o,n,s);var a=setInterval((function(){t.$socket.onopen(),0==e.target.readyState&&clearInterval(a)}),3e3)}},filters:{check_image:function(e,t){if(e=e.toString(),"Vertical"==e||"Horizontal"==e)return e;function o(e){var t=new Image;if(t.src=e,t.complete)return t.width<t.height?"Vertical":"Horizontal";t.onload=function(){return t.onload=null,t.width<t.height?"Vertical":"Horizontal"}}return""==e&&console.log("图片信息为空，开始本地JS分析"+t),e=o(t),e}}},v=k,_=(o("08fd"),o("6544")),w=o.n(_),y=o("8336"),$=o("269a"),S=o.n($),P=o("f977"),O=Object(m["a"])(v,u,l,!1,null,null,null),j=O.exports;w()(O,{VBtn:y["a"]}),S()(O,{Scroll:P["a"]});var x=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"SinglePage"}},[o("Header",[o("h2",[e.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+e.book.name}},[e._v(e._s(e.book.name)+"【Download】")]),e.book.IsFolder?o("a",{attrs:{href:"raw/"+e.book.name}},[e._v(e._s(e.book.name))]):e._e()])]),o("div",{staticClass:"singe_page",on:{click:e.nextPage}},[o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:e.book.pages[e.page-1].url}}),o("img")]),e._t("default")],2)},C=[],H={components:{Header:b},data:function(){return{page:1,alert:!1,easing:"easeInOutCubic",book:null,bookshelf:null,defaultSetting:null}},mounted:function(){this.book=this.$store.state.book,this.bookshelf=this.$store.state.bookshelf,this.defaultSetting=this.$store.state.defaultSetting},destroyed:function(){},methods:{initPage:function(){},nextPage:function(e){this.page<this.book.page_num?this.page=this.page+1:alert("Last Page!"),console.log(e)},toPage:function(e){this.page=e,console.log(e)},moveSomething:function(e){switch(e.keyCode){case 37:break;case 32:window.innerHeight,window.scrollY,document.body.offsetHeight;break;case 39:break;case 17:break}},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t){var o=t.key,n=t.keyCode;console.log(n),console.log(o)}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}}},T=H,E=(o("6320"),Object(m["a"])(T,x,C,!1,null,null,null)),I=E.exports,F=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{staticClass:"SketchPage",attrs:{id:"SketchPage"}},[o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】")]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(e.book.name))]):e._e()])]),o("div",{staticClass:"sketch_div",on:{click:function(t){return e.nextPage(2)}}},[o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.page-1].url}}),o("img")]),o("v-pagination",{attrs:{circle:"",length:this.$store.state.book.page_num,"total-visible":10},on:{input:e.toPage},model:{value:e.page,callback:function(t){e.page=t},expression:"page"}}),e._t("default")],2)},N=[],B={components:{Header:b},data:function(){return{book:null,bookshelf:null,defaultSetting:null,page:1,time_cont:0,alert:!1,easing:"easeInOutCubic"}},mounted:function(){this.time_cont=0,this.$cookies.keys(),this.book=this.$store.state.book,this.bookshelf=this.$store.state.bookshelf,this.defaultSetting=this.$store.state.defaultSetting},destroyed:function(){},methods:{initPage:function(){},nextPage:function(e){this.page<this.book.page_num?this.page=this.page+e:this.alert=!0,console.log(e)},toPage:function(e){this.page=e,console.log(e)},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t){var o=t.key,n=t.keyCode;console.log(n),console.log(o)}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}}},V=B,z=(o("dda7"),o("891e")),A=Object(m["a"])(V,F,N,!1,null,null,null),D=A.exports;w()(A,{VPagination:z["a"]});var M={name:"app",components:{MultiPage:j,SinglePage:I,SketchPage:D},data:function(){return{book:null,bookshelf:{},defaultSetting:{},page:1,duration:300,offset:0,easing:"easeInOutCubic",message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){this.initPage(),this.$cookies.keys()},destroyed:function(){},methods:{initPage:function(){var e=this;this.book=this.$store.book,this.defaultSetting=this.$store.defaultSetting,this.bookshelf=this.$store.bookshelf,c.a.get("/book.json").then((function(t){return e.$store.state.book=t.data})),c.a.get("/setting.json").then((function(t){return e.defaultSetting=t.data})),c.a.get("/bookshelf.json").then((function(t){return e.$store.state.bookshelf=t.data})).finally()},getNumber:function(e){this.page=e,console.log(e)}}},J=M,L=(o("034f"),Object(m["a"])(J,a,i,!1,null,null,null)),W=L.exports,K=o("8c4f"),R=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"Home"}})},Y=[],q=(o("99af"),o("2909")),G={name:"Home",components:{},data:function(){return{todos:[]}},methods:{deleteTodo:function(e){this.todos=this.todos.filter((function(t){return t.id!==e}))},addTodo:function(e){var t=this,o=e.title,n=e.completed;c.a.post("https://jsonplaceholder.typicode.com/todos",{title:o,completed:n}).then((function(e){return t.todos=[].concat(Object(q["a"])(t.todos),[e.data])})).catch((function(e){return console.log(e)}))}},created:function(){var e=this;c.a.get("https://jsonplaceholder.typicode.com/todos?_limit=4").then((function(t){return e.todos=t.data})).catch((function(e){return console.log(e)}))}},Q=G,U=(o("cccb"),Object(m["a"])(Q,R,Y,!1,null,null,null)),X=U.exports;s["a"].use(K["a"]);var Z=[{path:"/",name:"Home",component:X},{path:"/about",name:"About",component:function(){return o.e("about").then(o.bind(null,"f820"))}}],ee=new K["a"]({routes:Z}),te=ee,oe=o("caf9"),ne=o("b408"),se=o.n(ne),ae=o("f309");s["a"].use(ae["a"]);var ie=new ae["a"]({}),re=o("2b27"),ce=o.n(re),ue=o("2f62"),le=void 0;s["a"].use(se.a,"ws://"+document.location.host+"/ws",{reconnection:!0,reconnectionAttempts:500,reconnectionDelay:1e3}),s["a"].config.productionTip=!1,s["a"].use(oe["a"],{preLoad:4.5,attempt:10}),s["a"].use(ce.a),s["a"].$cookies.config("30d"),s["a"].use(ue["a"]);var de=new ue["a"].Store({state:{count:0,todos:[{id:1,text:"...",done:!0},{id:2,text:"...",done:!1}],now_page:1,book:{name:"loading",page_num:1,pages:[{height:500,width:449,url:"/resources/favicon.ico",class:"Vertical"}]},bookshelf:{},defaultSetting:{default_page_template:"multi"},message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}},getters:{doneTodos:function(e){return e.todos.filter((function(e){return e.done}))},now_page:function(e){return e.now_page},book:function(e){return console.log(le.state.book),e.book},bookshelf:function(e){return e.bookshelf},defaultSetting:function(e){return e.defaultSetting},message:function(e){return e.message}},mutations:{increment:function(e){e.count++},syncBookDate:function(e,t){e.book=t.msg,console.log(e.book),console.log("syncBookDate run")}},actions:{incrementAction:function(e){e.commit("increment")},getMessageAction:function(e){return Object(n["a"])(regeneratorRuntime.mark((function t(){var o,n;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,c.a.get("/book.json").then((function(e){return e.data}),(function(){return""}));case 2:o=t.sent,n={message:o},e.commit("syncBookDate",n);case 5:case"end":return t.stop()}}),t)})))()}}});new s["a"]({router:te,vuetify:ie,store:de,render:function(e){return e(W)}}).$mount("#app")},"5ced":function(e,t,o){},6320:function(e,t,o){"use strict";o("bcdd")},"85ec":function(e,t,o){},bcdd:function(e,t,o){},cccb:function(e,t,o){"use strict";o("5ced")},d7e9:function(e,t,o){},dd61:function(e,t,o){},dda7:function(e,t,o){"use strict";o("dd61")}});
//# sourceMappingURL=app.f074f8cb.js.map