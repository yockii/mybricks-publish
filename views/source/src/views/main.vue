<script setup lang="ts">
import {onMounted, ref} from "vue";
import {useUserStore} from "../store/user";
import router from "../router";
import {loadMicroApp, MicroApp} from "qiankun";
import {useRoute} from "vue-router";

const route = useRoute()
const userStore = useUserStore()

const menus = ref<{
  name: string,
  displayName: string,
  activeRule: string,
  entry: string,
}[]>([])

const appContainer = ref<HTMLElement|null>(null)
const currentApp = ref<MicroApp|null>(null)

const switchApp = (menu:any) => {
  if (appContainer.value) {
    if (currentApp.value) {
      currentApp.value.unmount()
    }
    router.push(menu.activeRule)
    currentApp.value = loadMicroApp({
      name: menu.name,
      entry: menu.entry,
      container: appContainer.value,
      props: {
        token: userStore.token,
      }
    })
  }
}

onMounted(() => {
  // 获取route信息，注册微前端
  fetch("/api/v1/route/list?offset=-1&limit=-1", {
    headers: {
      "Authorization": "Bearer " + userStore.token
    }
  }).then(response => {
    if (!response.ok) {
      throw new Error(response.statusText)
    }
    return response.json()
  }).then(response => {
    menus.value = response.data.items.map((item: any) => {
      const microApp = {
        name: item.code,
        displayName: item.displayName,
        activeRule: "/main" + item.activeRule,
        entry: item.entry,
      }
      if (!currentApp.value && route.path === microApp.activeRule) {
        switchApp(microApp)
      }
      return microApp
    })
  }).catch(reason => {
    if (reason.message === "Unauthorized") {
      userStore.$reset()
      router.push("/login")
    }
  })
})
</script>

<template>
  <div class="header">
    <span class="ml-16px">页面发布部署系统</span>
  </div>
  <div class="main">
    <div class="h-full w-200px bg-#FFFFFF3A">
      <template v-for="menu in menus">
        <div class="menu" :class="{'active': menu.activeRule === route.path}" @click="switchApp(menu)">
          {{menu.displayName}}
        </div>
      </template>
    </div>
    <div ref="appContainer" class="app-container"></div>
  </div>
</template>

<style scoped>
.header {
  width: 100%;
  height: 40px;
  color: #FFFFFF;
  font-size: 22px;
  display: flex;
  align-items: center;
  background:
      radial-gradient(rgba(255,255,255,0) 0, rgba(255,255,255,.15) 30%, rgba(255,255,255,.3) 32%, rgba(255,255,255,0) 33%) 0 0,
      radial-gradient(rgba(255,255,255,0) 0, rgba(255,255,255,.1) 11%, rgba(255,255,255,.3) 13%, rgba(255,255,255,0) 14%) 0 0,
      radial-gradient(rgba(255,255,255,0) 0, rgba(255,255,255,.2) 17%, rgba(255,255,255,.43) 19%, rgba(255,255,255,0) 20%) 0 110px,
      radial-gradient(rgba(255,255,255,0) 0, rgba(255,255,255,.2) 11%, rgba(255,255,255,.4) 13%, rgba(255,255,255,0) 14%) -130px -170px,
      radial-gradient(rgba(255,255,255,0) 0, rgba(255,255,255,.2) 11%, rgba(255,255,255,.4) 13%, rgba(255,255,255,0) 14%) 130px 370px,
      radial-gradient(rgba(255,255,255,0) 0, rgba(255,255,255,.1) 11%, rgba(255,255,255,.2) 13%, rgba(255,255,255,0) 14%) 0 0,
      linear-gradient(45deg, #343702 0%, #184500 20%, #187546 30%, #006782 40%, #0b1284 50%, #760ea1 60%, #83096e 70%, #840b2a 80%, #b13e12 90%, #e27412 100%);
  background-size: 470px 470px, 970px 970px, 410px 410px, 610px 610px, 530px 530px, 730px 730px, 100% 100%;
  background-color: #840b2a;
}
.main{
  display: flex;
  height: calc(100vh - 40px);
  width: 100%;
  color: #FFFFFF;
  background-color: hsl(2, 57%, 40%);
  background-image: repeating-linear-gradient(transparent, transparent 50px, rgba(0,0,0,.4) 50px, rgba(0,0,0,.4) 53px,
  transparent 53px, transparent 63px, rgba(0,0,0,.4) 63px, rgba(0,0,0,.4) 66px, transparent 66px, transparent 116px,
  rgba(0,0,0,.5) 116px, rgba(0,0,0,.5) 166px, rgba(255,255,255,.2) 166px, rgba(255,255,255,.2) 169px,
  rgba(0,0,0,.5) 169px, rgba(0,0,0,.5) 179px, rgba(255,255,255,.2) 179px, rgba(255,255,255,.2) 182px,
  rgba(0,0,0,.5) 182px, rgba(0,0,0,.5) 232px, transparent 232px),
  repeating-linear-gradient(270deg, transparent, transparent 50px, rgba(0,0,0,.4) 50px, rgba(0,0,0,.4) 53px,
      transparent 53px, transparent 63px, rgba(0,0,0,.4) 63px, rgba(0,0,0,.4) 66px, transparent 66px, transparent 116px,
      rgba(0,0,0,.5) 116px, rgba(0,0,0,.5) 166px, rgba(255,255,255,.2) 166px, rgba(255,255,255,.2) 169px,
      rgba(0,0,0,.5) 169px, rgba(0,0,0,.5) 179px, rgba(255,255,255,.2) 179px, rgba(255,255,255,.2) 182px,
      rgba(0,0,0,.5) 182px, rgba(0,0,0,.5) 232px, transparent 232px),
  repeating-linear-gradient(125deg, transparent, transparent 2px, rgba(0,0,0,.2) 2px, rgba(0,0,0,.2) 3px,
      transparent 3px, transparent 5px, rgba(0,0,0,.2) 5px);
}
.menu {
  font-size: 18px;
  line-height: 32px;
  margin: 8px 0;
  padding: 8px;
  cursor: pointer;
}
.menu:hover {
  background-color: #69c0ff4f;
}
.menu.active {
  background-color: #69c0ff8f;
}
.app-container {
  padding: 8px;
  width: calc(100vw - 200px);
  opacity: 0.8;
}
</style>