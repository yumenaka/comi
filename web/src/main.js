// import "es6-promise/auto";
import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import VueLazyload from "vue-lazyload";
// import websocket from 'vue-native-websocket'
import vuetify from "./plugins/vuetify";
import VueCookies from "vue-cookies";
import axios from "axios";
import Vuex from "vuex"; //[1]引入vuex  参考：https://my.oschina.net/u/4395108/blog/3317345

// Vue.use(websocket, "ws://" + document.location.host + "/ws", {//服务器的地址
//   reconnection: true, // (Boolean)是否自动重连，默认false
//   reconnectionAttempts: 500, // 重连次数
//   reconnectionDelay: 1000, // 再次重连等待时间间隔(1000)
// })

Vue.config.productionTip = false;
Vue.use(VueLazyload, {
  preLoad: 4.5,
  attempt: 10,
});

// https://github.com/cmp-cc/vue-cookies
Vue.use(VueCookies);
Vue.$cookies.config("30d");

//https://vuex.vuejs.org/zh/guide/
Vue.use(Vuex); //[2]使用vuex

const store = new Vuex.Store({
  //[3]创建一个store实例
  state: {
    //[4]所有组件共用数据存放处
    count: 0,
    todos: [
      { id: 1, text: "...", done: true },
      { id: 2, text: "...", done: false },
    ],
    now_page: 1,
    book: {
      name: "loading",
      page_num: 1,
      pages: [
        {
          height: 500,
          width: 449,
          url: "/resources/favicon.ico",
          class: "Vertical",
        },
      ],
    },
    bookshelf: {},
    setting: {
      template: "scroll",
      sketch_count_seconds: 90,
    },
    message: {
      user_uuid: "",
      server_status: "",
      now_book_uuid: "",
      read_percent: 0.0,
      msg: "",
    },
  },

  // mutaitions内只能执行同步操作
  mutations: {
    change_template_to_scroll(state) {
      state.setting.template = "scroll";
      console.log("change_template_to_scroll:" + state.setting.template);
    },
    change_template_to_double(state) {
      state.setting.template = "double";
      console.log("change_template_to_double:" + state.setting.template);
    },
    change_template_to_single(state) {
      state.setting.template = "single";
      console.log("change_template_to_single:" + state.setting.template);
    },
    change_template_to_sketch(state) {
      state.setting.template = "sketch";
      console.log("change_template_to_sketch:" + state.setting.template);
    },
    increment(state) {
      state.count++;
    },
    syncSettingData(state, payload) {
      state.setting = payload.message;
    },
    syncBookData(state, payload) {
      state.book = payload.message;
    },
    syncBookShelfData(state, payload) {
      state.bookshelf = payload.message;
    },
  },
  // Action 可以包含任意异步操作，通过 store.dispatch 方法触发
  actions: {
    incrementAction(context) {
      context.commit("increment");
    },
    //拉取远程设定数据
    async syncSettingDataAction(context) {
      const msg = await axios.get("/setting.json").then(
        (res) => res.data,
        () => ""
      );
      const payload = {
        message: msg,
      };
      context.commit("syncSettingData", payload);
    },
    //拉取当前阅读书籍数据
    async syncBookDataAction(context) {
      const msg = await axios.get("/book.json").then(
        (res) => res.data,
        () => ""
      );
      const payload = {
        message: msg,
      };
      context.commit("syncBookData", payload);
    },
    //拉取书架数据
    async syncBookShelfDataAction(context) {
      const msg = await axios.get("/bookshelf.json").then(
        (res) => res.data,
        () => ""
      );
      const payload = {
        message: msg,
      };
      context.commit("syncBookShelfData", payload);
    },
  },
  getters: {
    doneTodos: (state) => {
      return state.todos.filter((todo) => todo.done);
    },
    now_page: (state) => {
      return state.now_page;
    },
    book: (state) => {
      return state.book;
    },
    bookshelf: (state) => {
      return state.bookshelf;
    },
    setting: (state) => {
      return state.setting;
    },
    message: (state) => {
      return state.message;
    },
  },
});

new Vue({
  router,
  vuetify,
  store, //[5]注入store
  render: (h) => h(App),
}).$mount("#app");
