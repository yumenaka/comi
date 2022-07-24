import{_ as i,d as _,H as d,B as u,r as f,c as p,a as s,u as m,b as r,o as h,p as b,e as g,f as v}from"./index.8ad5ba17.js";const t=_({name:"AboutPage",props:["readMode"],emits:["setSome"],components:{Header:d,Bottom:u},setup(){return{model:f({interfaceColor:"#F5F5E4",backgroundColor:"#E0D9CD"})}},data(){return{book_num:0,drawerActive:!1,drawerPlacement:"right",PageTitle:""}},created(){localStorage.getItem("BackgroundColor")!=null&&(this.model.backgroundColor=localStorage.getItem("BackgroundColor")),localStorage.getItem("InterfaceColor")!=null&&(this.model.interfaceColor=localStorage.getItem("InterfaceColor"))},methods:{getUploadTitile(){return this.$store.state.server_status.SupportUploadFile===!1?this.$t("no_support_upload_file"):this.$store.state.server_status.NumberOfBooks===0?this.$t("no_book_found_hint"):this.$t("number_of_online_books")+this.$store.state.server_status.NumberOfBooks},remoteIsWindows(){return this.$store.state.server_status?(console.dir(this.$store.state.server_status.Description),this.$store.state.server_status.Description.indexOf("windows")!==-1):!1}}}),a=()=>{m(e=>({"6a3bbc1a":e.model.interfaceColor,"5ec59336":e.model.backgroundColor}))},n=t.setup;t.setup=n?(e,o)=>(a(),n(e,o)):a;const C=t,k=e=>(b("data-v-7eafb1b5"),e=e(),g(),e),I={class:"w-full h-screen flex flex-col"},$=k(()=>v("div",{class:"mian_area flex-grow"},null,-1));function S(e,o,w,B,x,D){const l=r("Header"),c=r("Bottom");return h(),p("div",I,[s(l,{class:"header flex-none h-12",bookIsFolder:!1,headerTitle:this.getUploadTitile(),showReturnIcon:!0,showSettingsIcon:!1,bookID:null,setDownLoadLink:!1},null,8,["headerTitle"]),$,s(c,{class:"bottom flex-none h-12",softVersion:this.$store.state.server_status.ServerName?this.$store.state.server_status.ServerName:"Comigo"},null,8,["softVersion"])])}const V=i(C,[["render",S],["__scopeId","data-v-7eafb1b5"]]);export{V as default};
