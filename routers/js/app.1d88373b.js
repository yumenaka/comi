(function(e){function t(t){for(var n,i,r=t[0],l=t[1],c=t[2],g=0,h=[];g<r.length;g++)i=r[g],Object.prototype.hasOwnProperty.call(a,i)&&a[i]&&h.push(a[i][0]),a[i]=0;for(n in l)Object.prototype.hasOwnProperty.call(l,n)&&(e[n]=l[n]);u&&u(t);while(h.length)h.shift()();return s.push.apply(s,c||[]),o()}function o(){for(var e,t=0;t<s.length;t++){for(var o=s[t],n=!0,r=1;r<o.length;r++){var l=o[r];0!==a[l]&&(n=!1)}n&&(s.splice(t--,1),e=i(i.s=o[0]))}return e}var n={},a={app:0},s=[];function i(t){if(n[t])return n[t].exports;var o=n[t]={i:t,l:!1,exports:{}};return e[t].call(o.exports,o,o.exports,i),o.l=!0,o.exports}i.m=e,i.c=n,i.d=function(e,t,o){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:o})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(i.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)i.d(o,n,function(t){return e[t]}.bind(null,n));return o},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="";var r=window["webpackJsonp"]=window["webpackJsonp"]||[],l=r.push.bind(r);r.push=t,r=r.slice();for(var c=0;c<r.length;c++)t(r[c]);var u=l;s.push([0,"chunk-vendors"]),o()})({0:function(e,t,o){e.exports=o("56d7")},"034f":function(e,t,o){"use strict";o("85ec")},"0a90":function(e,t,o){},"3bc4":function(e,t,o){"use strict";o("d1e9")},"4e71":function(e,t,o){"use strict";o("d7e9")},"56d7":function(e,t,o){"use strict";o.r(t);o("4de4"),o("96cf");var n=o("1da1"),a=(o("e260"),o("e6cf"),o("cca6"),o("a79d"),o("2b0e")),s=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"app"}},[this.$store.state.defaultSetting?o("div",["scroll"===this.$store.state.defaultSetting.template?o("ScrollTemplate"):e._e(),"sketch"===this.$store.state.defaultSetting.template?o("SketchTemplate"):e._e(),"single"===this.$store.state.defaultSetting.template?o("SinglePageTemplate"):e._e(),"double"===this.$store.state.defaultSetting.template?o("DoublePageTemplate"):e._e()],1):o("p",[e._v("loading.....")])])},i=[],r=(o("d3b7"),o("ddb0"),o("bc3a")),l=o.n(r),c=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"ScrollPage"}},[o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】")]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name))]):e._e()]),o("h4",[e._v("总页数："+e._s(this.$store.state.book.all_page_num))])]),e._l(this.$store.state.book.pages,(function(t,n){return o("div",{key:t.url,staticClass:"manga"},[o("img",{directives:[{name:"lazy",rawName:"v-lazy",value:t.url,expression:"page.url"}],key:n,class:e._f("check_image")(t.image_type,t.url),attrs:{H:t.height,W:t.width}}),e.showPageNum?o("p",[e._v(e._s(n+1)+"/"+e._s(e.AllPageNum))]):e._e()])})),o("p"),o("v-btn",{directives:[{name:"scroll",rawName:"v-scroll",value:e.onScroll,expression:"onScroll"},{name:"show",rawName:"v-show",value:e.btnFlag,expression:"btnFlag"}],attrs:{fab:"",color:"#bbcbff",bottom:"",right:""},on:{click:e.toTop}},[e._v("▲")]),e._t("default")],2)},u=[],g=(o("25f0"),o("93d3")),h={components:{Header:g["a"]},data:function(){return{page_mode:"multi",btnFlag:!1,showPageNum:!1,duration:300,offset:0,easing:"easeInOutCubic",AllPageNum:this.$store.state.book.all_page_num,message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){this.initPage()},destroyed:function(){},methods:{initPage:function(){this.$cookies.keys()},getBook:function(){return this.$store.state.book},getNumber:function(e){this.page=e,console.log(e)},onScroll:function(e){if("undefined"!==typeof window){var t=window.pageYOffset||e.target.scrollTop||0;this.btnFlag=t>20}},toTop:function(){this.$vuetify.goTo(0)},initWebSocket:function(){this.$socket.onopen=this.websocketonopen,this.$socket.onerror=this.websocketonerror,this.$socket.onmessage=this.websocketonmessage,this.$socket.onclose=this.websocketclose,this.hint="连接建立"},websocketonopen:function(e){this.hint="连接成功",console.log("连接建立",e)},websocketonerror:function(e){this.hint="连接出错",this.initWebSocket(),console.log("Connection Error !!!",e)},websocketonmessage:function(e){console.log(e),this.msgList.push(JSON.parse(e.data)),this.hint="接收消息"},onChangeBook:function(e,t){this.message.now_book_uuid=t,this.message.msg="ChangeBook",this.$socket.send(JSON.stringify(this.message)),this.getBook()},websocketsend:function(e){this.$socket.send(JSON.stringify(this.message)),console.log(this.$socket.readyState,e)},websocketclose:function(e){var t=this;this.hint="连接断开",console.log("断开连接",e);var o=e.code,n=e.reason,a=e.wasClean;console.log(o,n,a);var s=setInterval((function(){t.$socket.onopen(),0==e.target.readyState&&clearInterval(s)}),3e3)}},filters:{check_image:function(e,t){if(e=e.toString(),"SinglePage"==e||"DoublePage"==e)return e;function o(e){var t=new Image;if(t.src=e,t.complete)return t.width<t.height?"SinglePage":"DoublePage";t.onload=function(){return t.onload=null,t.width<t.height?"SinglePage":"DoublePage"}}return""==e&&console.log("图片信息为空，开始本地JS分析"+t),e=o(t),e}}},p=h,d=(o("9803"),o("2877")),_=o("6544"),f=o.n(_),m=o("8336"),w=o("269a"),b=o.n(w),k=o("f977"),v=Object(d["a"])(p,c,u,!1,null,null,null),P=v.exports;f()(v,{VBtn:m["a"]}),b()(v,{Scroll:k["a"]});var $=o("d47c"),y=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"SinglePageTemplate"}},[e.showHeader?o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】")]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name))]):e._e()])]):e._e(),o("div",{staticClass:"single_page_main"},[e.now_page<=this.$store.state.book.all_page_num&&e.now_page>=1?o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.now_page-1].url},on:{click:function(t){return e.addPage(1)}}}):e._e(),o("img")]),o("v-pagination",{attrs:{length:this.$store.state.book.all_page_num,"total-visible":10},on:{input:e.toPage},model:{value:e.now_page,callback:function(t){e.now_page=t},expression:"now_page"}}),e._t("default")],2)},S=[],T={components:{Header:g["a"]},data:function(){return{now_page:1,showHeader:!1,showPagination:!0,alert:!1,easing:"easeInOutCubic",book:null,bookshelf:null,defaultSetting:null}},mounted:function(){this.book=this.$store.state.book,this.bookshelf=this.$store.state.bookshelf,this.defaultSetting=this.$store.state.defaultSetting,window.addEventListener("keyup",this.handleKeyup)},destroyed:function(){window.removeEventListener("keyup",this.handleKeyup)},methods:{initPage:function(){},addPage:function(e){this.now_page+e<this.book.all_page_num&&this.now_page+e>=1&&(this.now_page=this.now_page+e)},toPage:function(e){e<=this.book.all_page_num&&e>=1&&(this.now_page=e)},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t)switch(t.key){case"ArrowUp":case"PageUp":case"ArrowLeft":this.addPage(-1);break;case"Space":case"ArrowDown":case"PageDown":case"ArrowRight":this.addPage(1);break;case"Home":this.toPage(1);break;case"End":this.toPage(this.book.all_page_num);break;case"Ctrl":break}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}}},x=T,D=(o("9f5f"),o("891e")),O=Object(d["a"])(x,y,S,!1,null,null,null),H=O.exports;f()(O,{VPagination:D["a"]});var A=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"DoublePageTemplate"}},[e.showHeader?o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】现在时刻："+e._s(e.currentTime))]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"现在时刻："+e._s(e.currentTime))]):e._e()])]):e._e(),o("div",{staticClass:"double_page_main"},[e.now_page<this.$store.state.book.all_page_num&&"SinglePage"==this.$store.state.book.pages[e.now_page-1].image_type&&"SinglePage"==this.$store.state.book.pages[e.now_page].image_type?o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.now_page].url},on:{click:e.nextPageSinglePage}}):e._e(),o("img"),e.now_page-1>=0&&"SinglePage"==this.$store.state.book.pages[e.now_page-1].image_type?o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.now_page-1].url},on:{click:e.previousPageSinglePage}}):e._e(),o("img"),e.now_page-1>=0&&"DoublePage"==this.$store.state.book.pages[e.now_page-1].image_type?o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.now_page-1].url},on:{click:e.nextPageDoublePage}}):e._e(),o("img")]),e.showPagination?o("v-pagination",{attrs:{length:this.$store.state.book.all_page_num,"total-visible":15},on:{input:e.toPage},model:{value:e.now_page,callback:function(t){e.now_page=t},expression:"now_page"}}):e._e(),e._t("default")],2)},j=[],E={methods:{initPage:function(){},toPage:function(e){e>this.$store.state.book.all_page_num&&e<0&&console.log("now_page error",e),this.now_page=e,console.log(e)},nextPageDoublePage:function(){this.now_page=this.now_page+1,console.log(this.now_page)},nextPageSinglePage:function(){var e=this;if(e.now_page+1>e.AllPageNum)console.log(e.now_page);else{if(e.now_page+1!=e.AllPageNum)return e.now_page+1<e.AllPageNum&&"SinglePage"==this.$store.state.book.pages[e.now_page].image_type&&"SinglePage"==this.$store.state.book.pages[e.now_page-1].image_type?(e.now_page=e.now_page+2,void console.log(e.now_page)):(e.now_page=e.now_page+1,void console.log(e.now_page));e.now_page=e.now_page+1}},previousPageSinglePage:function(){if(0!=this.now_page)return 1==this.now_page?(this.now_page=1,void console.log(this.now_page)):2==this.now_page?(this.now_page=this.now_page-1,void console.log(this.now_page)):"SinglePage"==this.$store.state.book.pages[this.now_page-2].image_type&&"SinglePage"==this.$store.state.book.pages[this.now_page-1].image_type?(this.now_page=this.now_page-2,void console.log(this.now_page)):(this.now_page=this.now_page-1,void console.log(this.now_page));console.log(this.now_page)},nextPage:function(){this.now_page>this.$store.state.book.all_page_num?console.log(this.now_page):this.now_page!=this.$store.state.book.all_page_num?"SinglePage"!=this.$store.state.book.pages[this.now_page].image_type?"DoublePage"!=this.$store.state.book.pages[this.now_page].image_type||this.nextPageDoublePage():this.nextPageSinglePage():console.log(this.now_page)},previousPage:function(){this.now_page>this.$store.state.book.all_page_num?console.log(this.now_page):this.now_page!=this.$store.state.book.all_page_num?"SinglePage"==this.$store.state.book.pages[this.now_page].image_type?this.previousPageSinglePage():"DoublePage"==this.$store.state.book.pages[this.now_page].image_type&&this.now_page-1>=0&&(this.now_page=this.now_page-1,console.log(this.now_page)):this.now_page=this.now_page-1},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t)switch(t.key){case"PageUp":case"ArrowUp":case"ArrowLeft":this.previousPage();break;case"Space":case"ArrowDown":case"PageDown":case"ArrowRight":this.nextPage();break;case"Home":this.toPage(1);break;case"End":this.toPage(this.$store.state.book.all_page_num);break;case"Ctrl":break}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}},components:{Header:g["a"]},data:function(){return{book:null,bookshelf:null,defaultSetting:null,now_page:1,showHeader:!1,showPagination:!0,AllPageNum:this.$store.state.book.all_page_num,time_cont:0,alert:!1,easing:"easeInOutCubic",timer:"",currentTime:new Date}},created:function(){var e=this;this.timer=setInterval((function(){var t=new Date,o=t.getFullYear(),n=t.getMonth()+1,a=t.getDate();n>=1&&n<=9&&(n="0"+n),a>=0&&a<=9&&(a="0"+a);var s=o+" 年 "+n+" 月 "+a+" 日 ",i=t.getHours();i>=0&&i<=9&&(i="0"+i);var r=t.getMinutes();r>=0&&r<=9&&(r="0"+r);var l=t.getSeconds();l>=0&&l<=9&&(l="0"+l),e.currentTime=s+" "+i+":"+r+":"+l}),1e3)},mounted:function(){this.time_cont=0,this.$cookies.keys(),window.addEventListener("keyup",this.handleKeyup)},destroyed:function(){window.removeEventListener("keyup",this.handleKeyup),this.timer&&clearInterval(this.timer)}},C=E,I=(o("810a"),Object(d["a"])(C,A,j,!1,null,null,null)),N=I.exports;f()(I,{VPagination:D["a"]});var F={name:"app",components:{ScrollTemplate:P,SketchTemplate:$["default"],SinglePageTemplate:H,DoublePageTemplate:N},data:function(){return{book:null,bookshelf:{},defaultSetting:{},now_page:1,duration:300,offset:0,easing:"easeInOutCubic",message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){this.initPage(),this.$cookies.keys()},destroyed:function(){},methods:{initPage:function(){var e=this;l.a.get("/book.json").then((function(t){return e.$store.state.book=t.data})).finally(this.book=this.$store.book),l.a.get("/setting.json").then((function(t){return e.$store.state.defaultSetting=t.data})).finally(this.defaultSetting=this.$store.defaultSetting),l.a.get("/bookshelf.json").then((function(t){return e.$store.state.bookshelf=t.data})).finally(this.bookshelf=this.$store.bookshelf)},getNumber:function(e){this.page=e,console.log(e)}}},L=F,K=(o("034f"),Object(d["a"])(L,s,i,!1,null,null,null)),M=K.exports,B=o("8c4f");a["a"].use(B["a"]);var z=[{path:"/",name:"Scrool",component:P},{path:"/sketch",name:"Sketch",component:function(){return Promise.resolve().then(o.bind(null,"d47c"))}}],V=new B["a"]({routes:z}),J=V,U=o("caf9"),W=o("b408"),R=o.n(W),Y=o("f309");a["a"].use(Y["a"]);var q=new Y["a"]({}),G=o("2b27"),Q=o.n(G),X=o("2f62"),Z=void 0;a["a"].use(R.a,"ws://"+document.location.host+"/ws",{reconnection:!0,reconnectionAttempts:500,reconnectionDelay:1e3}),a["a"].config.productionTip=!1,a["a"].use(U["a"],{preLoad:4.5,attempt:10}),a["a"].use(Q.a),a["a"].$cookies.config("30d"),a["a"].use(X["a"]);var ee=new X["a"].Store({state:{count:0,todos:[{id:1,text:"...",done:!0},{id:2,text:"...",done:!1}],now_page:1,book:{name:"loading",page_num:1,pages:[{height:500,width:449,url:"/resources/favicon.ico",class:"Vertical"}]},bookshelf:{},defaultSetting:{default_page_template:"scroll",sketch_count_seconds:90},message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}},getters:{doneTodos:function(e){return e.todos.filter((function(e){return e.done}))},now_page:function(e){return e.now_page},book:function(e){return console.log(Z.state.book),e.book},bookshelf:function(e){return e.bookshelf},defaultSetting:function(e){return e.defaultSetting},message:function(e){return e.message}},mutations:{increment:function(e){e.count++},syncBookDate:function(e,t){e.book=t.msg,console.log(e.book),console.log("syncBookDate run")}},actions:{incrementAction:function(e){e.commit("increment")},getMessageAction:function(e){return Object(n["a"])(regeneratorRuntime.mark((function t(){var o,n;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,l.a.get("/book.json").then((function(e){return e.data}),(function(){return""}));case 2:o=t.sent,n={message:o},e.commit("syncBookDate",n);case 5:case"end":return t.stop()}}),t)})))()}}});new a["a"]({router:J,vuetify:q,store:ee,render:function(e){return e(M)}}).$mount("#app")},6528:function(e,t,o){},"810a":function(e,t,o){"use strict";o("0a90")},"85ec":function(e,t,o){},"93d3":function(e,t,o){"use strict";var n=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("header",{staticClass:"header"},[e._t("default")],2)},a=[],s={name:"Header",data:function(){return{mybook:this.book,rangeValue:500}},methods:{changeEvent:function(){console.log("change to:"+this.rangeValue)}}},i=s,r=(o("4e71"),o("2877")),l=Object(r["a"])(i,n,a,!1,null,"c3c1e8ec",null);t["a"]=l.exports},9803:function(e,t,o){"use strict";o("6528")},"9f5f":function(e,t,o){"use strict";o("f2b5")},d1e9:function(e,t,o){},d47c:function(e,t,o){"use strict";o.r(t);var n=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"SketchPage"}},[e.showHeader?o("Header",[o("h2",[o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"现在时刻："+e._s(e.currentTime))])])]):e._e(),o("div",{staticClass:"sketch_main"},[o("div",{attrs:{id:"SketchHint"}},[o("p",[e._v(" "+e._s(this.$store.state.defaultSetting.sketch_count_seconds)+"秒翻页,"+e._s(e.getNowCount())+"⏳ ")])]),o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.now_page-1].url},on:{click:function(t){return e.addPage(1)}}}),o("img"),o("div",{attrs:{id:"SketchHint"}},[o("p",[e._v("🕒"+e._s(e.currentTime))])])]),e.showPagination?o("v-pagination",{attrs:{length:this.$store.state.book.all_page_num,"total-visible":10},on:{input:e.toPage},model:{value:e.now_page,callback:function(t){e.now_page=t},expression:"now_page"}}):e._e(),e._t("default")],2)},a=[],s=(o("d3b7"),o("ddb0"),o("93d3")),i={components:{Header:s["a"]},data:function(){return{time_cont:1,WaitSeconds:this.$store.state.defaultSetting.sketch_count_seconds,book:null,bookshelf:null,defaultSetting:null,showHeader:!1,showPagination:!0,now_page:1,alert:!1,easing:"easeInOutCubic",timer:"",currentTime:null}},created:function(){var e=this;this.timer=setInterval((function(){var t=new Date,o=t.getFullYear(),n=t.getMonth()+1,a=t.getDate();n>=1&&n<=9&&(n="0"+n),a>=0&&a<=9&&(a="0"+a);var s=o+"年"+n+"月"+a+"日",i=t.getHours();i>=0&&i<=9&&(i="0"+i);var r=t.getMinutes();r>=0&&r<=9&&(r="0"+r);var l=t.getSeconds();l>=0&&l<=9&&(l="0"+l),e.currentTime=i+":"+r+":"+l,console.log(s+"time_cont："+e.time_cont),e.time_cont<e.WaitSeconds?e.time_cont++:(e.time_cont=1,console.log("时间到，翻页："+e.currentTime+"秒"),e.now_page<e.$store.state.book.all_page_num?e.now_page+=1:e.now_page=1)}),1e3)},mounted:function(){this.time_cont=0,this.$cookies.keys(),window.addEventListener("keyup",this.handleKeyup)},destroyed:function(){window.removeEventListener("keyup",this.handleKeyup),this.timer&&clearInterval(this.timer)},methods:{initPage:function(){},getWaitSeconds:function(){return this.$store.state.defaultSetting.sketch_count_seconds},getNowCount:function(){var e=this.time_cont;return e>=0&&e<=9&&(e="0"+e),e},addPage:function(e){this.now_page+e<this.$store.state.book.all_page_num&&this.now_page+e>=1&&(this.now_page=this.now_page+e)},toPage:function(e){this.now_page=e,console.log(e)},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t)switch(t.key){case"PageUp":case"ArrowUp":case"ArrowLeft":this.addPage(-1);break;case"Space":case"ArrowDown":case"PageDown":case"ArrowRight":this.addPage(1);break;case"Home":this.toPage(1);break;case"End":this.toPage(this.$store.state.book.all_page_num-1);break;case"Ctrl":break}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}}},r=i,l=(o("3bc4"),o("2877")),c=o("6544"),u=o.n(c),g=o("891e"),h=Object(l["a"])(r,n,a,!1,null,null,null);t["default"]=h.exports;u()(h,{VPagination:g["a"]})},d7e9:function(e,t,o){},f2b5:function(e,t,o){}});
//# sourceMappingURL=app.1d88373b.js.map